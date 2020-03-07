package socache

import (
	"sync"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	gcsdkm "github.com/paralin/go-dota2/protocol"
)

// SOCache implements the shared-object cache from DOTA.
type SOCache struct {
	containerMap sync.Map
	le           logrus.FieldLogger
}

// NewSOCache builds a new SOCache.
func NewSOCache(le logrus.FieldLogger) *SOCache {
	return &SOCache{le: le}
}

// GetContainerForTypeID returns the container for the type id.
func (c *SOCache) GetContainerForTypeID(id uint32) (*SOCacheContainer, error) {
	if valInter, ok := c.containerMap.Load(id); ok {
		return valInter.(*SOCacheContainer), nil
	}

	ctr, err := NewSOCacheContainer(c.le, id)
	if err != nil {
		return nil, err
	}

	valInter, loaded := c.containerMap.LoadOrStore(id, ctr)
	if loaded {
		ctr = valInter.(*SOCacheContainer)
	}
	return ctr, nil
}

// HandleSubscribed handles a subscribed message.
func (c *SOCache) HandleSubscribed(msg *gcsdkm.CMsgSOCacheSubscribed) error {
	var retErr error
	for _, obj := range msg.GetObjects() {
		if obj.GetTypeId() == 0 {
			continue
		}

		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId()))
		if err == nil {
			err = ctr.HandleSubscribed(msg, obj)
		}

		if retErr == nil && err != nil {
			retErr = err
		}
	}

	return retErr
}

// HandleUnsubscribed handles a subscribed message.
func (c *SOCache) HandleUnsubscribed(msg *gcsdkm.CMsgSOCacheUnsubscribed) error {
	if msg.GetOwnerSoid().GetId() == 0 {
		return errors.Errorf("object id was empty in unsubscribed")
	}

	var retErr error
	c.containerMap.Range(func(key interface{}, value interface{}) bool {
		val := value.(*SOCacheContainer)
		_, hasObj := val.objects.Load(msg.GetOwnerSoid().GetId())
		if !hasObj {
			return true
		}

		retErr = val.HandleUnsubscribed(msg)
		return false
	})

	return retErr
}

// HandleUpdateMultiple handles a update multiple message.
func (c *SOCache) HandleUpdateMultiple(msg *gcsdkm.CMsgSOMultipleObjects) error {
	addUpdateObj := func(obj *gcsdkm.CMsgSOMultipleObjects_SingleObject) error {
		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId()))
		if err != nil {
			return nil
		}

		m := &gcsdkm.CMsgSOCacheSubscribed_SubscribedType{
			TypeId:     obj.TypeId,
			ObjectData: [][]byte{obj.ObjectData},
		}

		if err := ctr.addUpdateObject(msg.GetOwnerSoid(), m); err != nil {
			return err
		}

		return nil
	}

	for _, obj := range msg.GetObjectsAdded() {
		if err := addUpdateObj(obj); err != nil {
			return err
		}
	}

	for _, obj := range msg.GetObjectsModified() {
		if err := addUpdateObj(obj); err != nil {
			return err
		}
	}

	for _, obj := range msg.GetObjectsRemoved() {
		ctr, err := c.GetContainerForTypeID(uint32(obj.GetTypeId()))
		if err != nil {
			continue
		}

		if err := ctr.removeObject(msg.GetOwnerSoid()); err != nil {
			return err
		}
	}

	return nil
}

// HandleDestroy handles a object destroy message.
func (c *SOCache) HandleDestroy(msg *gcsdkm.CMsgSOSingleObject) error {
	ctr, err := c.GetContainerForTypeID(uint32(msg.GetTypeId()))
	if err != nil {
		return err
	}

	return ctr.HandleDestroy(msg)
}
