package socache

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/cso"
)

// EventType marks the type of event this is.
type EventType int

//go:generate stringer -type=EventType
const (
	// EventTypeCreate is emitted when an object is inserted into the cache.
	EventTypeCreate EventType = iota
	// EventTypeUpdate is emitted when an object is updated in the cache.
	EventTypeUpdate
	// EventTypeDestroy is emitted when an object is removed from the cache.
	EventTypeDestroy
)

// CacheEvent is a cache event.
type CacheEvent struct {
	// EventType marks the type of event.
	EventType EventType
	// Object contains the affected object.
	Object proto.Message
}

// CacheUnsubscribeFunc is called to unsubscribe a subscription.
type CacheUnsubscribeFunc func()

// SubscribeType subscribes to events for an object type.
func (c *SOCache) SubscribeType(id cso.CSOType) (<-chan *CacheEvent, CacheUnsubscribeFunc, error) {
	ctr, err := c.GetContainerForTypeID(uint32(id))
	if err != nil {
		return nil, nil, err
	}

	return ctr.Subscribe()
}
