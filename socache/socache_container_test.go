package socache

import (
	"io"
	"testing"

	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-dota2/cso"
	gcsm "github.com/paralin/go-dota2/protocol"
	"github.com/sirupsen/logrus"
)

// TestCacheNotifications checks lite decoding and callback reentry together.
func TestCacheNotifications(t *testing.T) {
	// Subscribe before receiving a generated shared-object packet.
	logger := logrus.New()
	logger.SetOutput(io.Discard)
	cache, err := NewSOCacheContainer(logger, uint32(cso.EconItem))
	if err != nil {
		t.Fatal(err)
	}
	events, cancel, err := cache.Subscribe()
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()
	owner := &gcsm.CMsgSOIDOwner{Id: new(uint64(73))}
	encoded, err := (&gcsm.CSOEconItem{Id: new(uint64(9007199254740993))}).MarshalVT()
	if err != nil {
		t.Fatal(err)
	}
	if err := cache.HandleSubscribed(&gcsm.CMsgSOCacheSubscribed{OwnerSoid: owner}, &gcsm.CMsgSOCacheSubscribed_SubscribedType{ObjectData: [][]byte{encoded}}); err != nil {
		t.Fatal(err)
	}

	// A subscriber sees the exact decoded value retained by the cache.
	event := <-events
	item := event.Object.(*gcsm.CSOEconItem)
	if event.EventType != EventTypeCreate || item.GetId() != 9007199254740993 {
		t.Fatal("shared object lost on publication")
	}
	if err := cache.Range(func(id uint64, object protobuf.Message) error {
		stored, ok := cache.Get(id)
		if !ok || stored != object {
			t.Fatal("range callback could not read its retained object")
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	// Unsubscribe is idempotent and stops later mutation notifications.
	cancel()
	cancel()
	if err := cache.HandleDestroy(&gcsm.CMsgSOSingleObject{OwnerSoid: owner}); err != nil {
		t.Fatal(err)
	}
	if cache.GetOne() != nil {
		t.Fatal("destroy left the object in the cache")
	}
	select {
	case <-events:
		t.Fatal("received a notification after unsubscribe")
	default:
	}
}
