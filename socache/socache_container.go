package socache

import (
	"math/rand"
	"sync"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/paralin/go-dota2/cso"
	gcsdkm "github.com/paralin/go-dota2/protocol"
)

// SOCacheContainer contains a type of object in the cache.
type SOCacheContainer struct {
	le            logrus.FieldLogger
	typeID        uint32
	objects       sync.Map // map[uint64]proto.Message
	subscriptions sync.Map // map[uint32]chan *CacheEvent
}

// NewSOCacheContainer builds a new container for a type id.
func NewSOCacheContainer(le logrus.FieldLogger, typeID uint32) (*SOCacheContainer, error) {
	typ := cso.CSOType(typeID)
	if _, err := cso.NewSharedObject(typ); err != nil {
		return nil, err
	}

	return &SOCacheContainer{typeID: typeID, le: le.WithField("object-type", typ.String())}, nil
}

// GetTypeID returns the type id the container contains.
func (c *SOCacheContainer) GetTypeID() uint32 {
	return c.typeID
}

// parseObject parses an object.
func (c *SOCacheContainer) parseObject(obj *gcsdkm.CMsgSOCacheSubscribed_SubscribedType) (proto.Message, error) {
	objData := obj.GetObjectData()[0]

	so, err := cso.NewSharedObject(cso.CSOType(c.GetTypeID()))
	if err != nil {
		c.le.Debugf("unknown: %v", obj.String())
		return nil, err
	}

	if err := proto.Unmarshal(objData, so); err != nil {
		return nil, err
	}

	return so, nil
}

// stringifyObject converts an object to a json string.
func (c *SOCacheContainer) stringifyObject(obj proto.Message) string {
	m := &jsonpb.Marshaler{}
	str, _ := m.MarshalToString(obj)
	return str
}

// emitEvent emits an event to all listeners.
func (c *SOCacheContainer) emitEvent(event *CacheEvent) {
	c.subscriptions.Range(func(key interface{}, value interface{}) bool {
		ch := value.(chan *CacheEvent)
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
		return true
	})
}

// addUpdateObject handles an added / updated object.
func (c *SOCacheContainer) addUpdateObject(soid *gcsdkm.CMsgSOIDOwner, obj *gcsdkm.CMsgSOCacheSubscribed_SubscribedType) error {
	soID := soid.GetId()
	if soID == 0 {
		return errors.New("object has empty shared object id")
	}

	so, err := c.parseObject(obj)
	if err != nil {
		return err
	}

	// c.le.Debugf("received object from GC: %s", c.stringifyObject(so))

	eventType := EventTypeCreate
	_, existed := c.objects.Load(soID)
	if existed {
		eventType = EventTypeUpdate
	}

	c.objects.Store(soID, so)
	c.emitEvent(&CacheEvent{
		EventType: eventType,
		Object:    so,
	})

	return nil
}

// removeObject attempts to remove an object.
func (c *SOCacheContainer) removeObject(soid *gcsdkm.CMsgSOIDOwner) error {
	soID := soid.GetId()
	if soID == 0 {
		return errors.New("object has empty shared object id")
	}

	soInter, ok := c.objects.Load(soID)
	if !ok {
		return nil
	}
	so := soInter.(proto.Message)

	c.objects.Delete(soID)
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

// Subscribe subscribes to events for the object type.
func (c *SOCacheContainer) Subscribe() (<-chan *CacheEvent, CacheUnsubscribeFunc, error) {
	subID := uint32(rand.Int31())
	ch := make(chan *CacheEvent, 10)
	c.subscriptions.Store(subID, (chan *CacheEvent)(ch))
	return ch, func() {
		c.subscriptions.Delete(subID)
	}, nil
}

// Get tries to look up an object by ID.
func (c *SOCacheContainer) Get(id uint64) (proto.Message, bool) {
	obj, ok := c.objects.Load(id)
	if !ok {
		return nil, false
	}

	return obj.(proto.Message), true
}

// GetOne returns an object assuming there is only one.
func (c *SOCacheContainer) GetOne() (obj proto.Message) {
	c.objects.Range(func(key interface{}, value interface{}) bool {
		obj = value.(proto.Message)
		return false
	})

	return
}

// Range ranges over the objects in the container.
func (c *SOCacheContainer) Range(cb func(id uint64, obj proto.Message) error) error {
	var retErr error
	c.objects.Range(func(key interface{}, value interface{}) bool {
		id := key.(uint64)
		val := value.(proto.Message)

		if err := cb(id, val); err != nil {
			retErr = err
			return false
		}

		return true
	})

	return retErr
}
