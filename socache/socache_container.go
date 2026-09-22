package socache

import (
	"maps"
	"sync"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/paralin/go-dota2/cso"
	gcsdkm "github.com/paralin/go-dota2/protocol"
)

// SOCacheContainer contains a type of object in the cache.
type SOCacheContainer struct {
	// le records cache diagnostics without retaining message payloads.
	le logrus.FieldLogger
	// typeID identifies the shared-object schema.
	typeID uint32
	// mtx guards objects, subscriptions, and their notification order.
	mtx sync.Mutex
	// objects retains decoded messages, immutable after publication.
	objects map[uint64]protobuf.Message
	// subscriptions receives bounded, best-effort change notifications.
	subscriptions map[uint64]chan *CacheEvent
	// subIDCounter allocates subscription identities under mtx.
	subIDCounter uint64
}

// NewSOCacheContainer builds a new container for a type id.
func NewSOCacheContainer(le logrus.FieldLogger, typeID uint32) (*SOCacheContainer, error) {
	// Reject unsupported schemas before publishing a container.
	typ := cso.CSOType(typeID) //nolint:gosec
	if _, err := cso.NewSharedObject(typ); err != nil {
		return nil, err
	}

	// Allocate both maps before the container accepts concurrent use.
	return &SOCacheContainer{
		typeID: typeID, le: le.WithField("object-type", typ.String()),
		objects:       make(map[uint64]protobuf.Message),
		subscriptions: make(map[uint64]chan *CacheEvent),
	}, nil
}

// GetTypeID returns the type id the container contains.
func (c *SOCacheContainer) GetTypeID() uint32 {
	return c.typeID
}

// parseObject parses an object.
func (c *SOCacheContainer) parseObject(obj *gcsdkm.CMsgSOCacheSubscribed_SubscribedType) (protobuf.Message, error) {
	// Construct the schema selected by this container.
	so, err := cso.NewSharedObject(cso.CSOType(c.GetTypeID())) //nolint:gosec
	if err != nil {
		c.le.Debugf("unknown: %v", obj.String())
		return nil, err
	}

	// Decode into a fresh message before sharing it with readers.
	if err := so.UnmarshalVT(obj.GetObjectData()[0]); err != nil {
		return nil, err
	}

	return so, nil
}

// emitEvent emits a bounded notification while mtx is held.
// Slow subscribers lose older events; subscription channels are never closed.
func (c *SOCacheContainer) emitEvent(event *CacheEvent) {
	for _, ch := range c.subscriptions {
		tries := 0
	RetryLoop:
		for tries < 3 {
			select {
			case ch <- event:
				break RetryLoop
			default:
				select {
				case <-ch:
					c.le.Warn("dropping event due to channel overflow")
				default:
				}
			}

			tries++
		}
	}
}

// addUpdateObject handles an added / updated object.
func (c *SOCacheContainer) addUpdateObject(soid *gcsdkm.CMsgSOIDOwner, obj *gcsdkm.CMsgSOCacheSubscribed_SubscribedType) error {
	// Require the coordinator identity before changing the cache.
	soID := soid.GetId()
	if soID == 0 {
		return errors.New("object has empty shared object id")
	}

	// Decode before acquiring the publication lock.
	so, err := c.parseObject(obj)
	if err != nil {
		return err
	}

	// Publish the object and notification in the same mutation order.
	c.mtx.Lock()
	defer c.mtx.Unlock()
	eventType := EventTypeCreate
	_, existed := c.objects[soID]
	if existed {
		eventType = EventTypeUpdate
	}

	c.objects[soID] = so
	c.emitEvent(&CacheEvent{
		EventType: eventType,
		Object:    so,
	})

	return nil
}

// removeObject attempts to remove an object.
func (c *SOCacheContainer) removeObject(soid *gcsdkm.CMsgSOIDOwner) error {
	// Require the coordinator identity before changing the cache.
	soID := soid.GetId()
	if soID == 0 {
		return errors.New("object has empty shared object id")
	}

	// Remove and notify atomically relative to other cache mutations.
	c.mtx.Lock()
	defer c.mtx.Unlock()
	so, ok := c.objects[soID]
	if !ok {
		return nil
	}

	delete(c.objects, soID)
	c.emitEvent(&CacheEvent{
		EventType: EventTypeDestroy,
		Object:    so,
	})
	return nil
}

// HandleSubscribed handles an incoming object from a Subscribed event.
func (c *SOCacheContainer) HandleSubscribed(msg *gcsdkm.CMsgSOCacheSubscribed, obj *gcsdkm.CMsgSOCacheSubscribed_SubscribedType) error {
	if len(obj.GetObjectData()) == 0 {
		return errors.Errorf("expected object data for cache type %d", c.GetTypeID())
	}

	return c.addUpdateObject(msg.GetOwnerSoid(), obj)
}

// HandleUnsubscribed handles a cache unsubscribe packet.
func (c *SOCacheContainer) HandleUnsubscribed(msg *gcsdkm.CMsgSOCacheUnsubscribed) error {
	return c.removeObject(msg.GetOwnerSoid())
}

// HandleDestroy handles a cache object destroy packet.
func (c *SOCacheContainer) HandleDestroy(msg *gcsdkm.CMsgSOSingleObject) error {
	return c.removeObject(msg.GetOwnerSoid())
}

// Subscribe registers a best-effort change stream; cancellation never closes it.
func (c *SOCacheContainer) Subscribe() (<-chan *CacheEvent, CacheUnsubscribeFunc, error) {
	// Register before publishing the stream to its caller.
	c.mtx.Lock()
	c.subIDCounter++
	subID := c.subIDCounter
	ch := make(chan *CacheEvent, 10)
	c.subscriptions[subID] = ch
	c.mtx.Unlock()

	// Cancellation is idempotent and excludes future notification sends.
	return ch, func() {
		c.mtx.Lock()
		delete(c.subscriptions, subID)
		c.mtx.Unlock()
	}, nil
}

// Get looks up an immutable shared object by coordinator identity.
func (c *SOCacheContainer) Get(id uint64) (protobuf.Message, bool) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	obj, ok := c.objects[id]
	return obj, ok
}

// GetOne returns any retained object, or nil when the container is empty.
func (c *SOCacheContainer) GetOne() protobuf.Message {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	for _, obj := range c.objects {
		return obj
	}
	return nil
}

// Range visits a snapshot; callbacks may call back into the container.
func (c *SOCacheContainer) Range(cb func(id uint64, obj protobuf.Message) error) error {
	// Release the cache lock before invoking caller code.
	c.mtx.Lock()
	objects := maps.Clone(c.objects)
	c.mtx.Unlock()

	// Stop at the first callback failure without changing the retained objects.
	for id, obj := range objects {
		if err := cb(id, obj); err != nil {
			return err
		}
	}
	return nil
}
