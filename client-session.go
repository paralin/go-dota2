package dota2

import (
	"context"

	"github.com/golang/protobuf/proto"
	devents "github.com/paralin/go-dota2/events"
	gcsm "github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
)

// SetPlaying announces whether this Steam connection is running Dota2.
// Stopping immediately interrupts requests and rejects late welcomes until resumed.
func (d *Dota2) SetPlaying(playing bool) {
	// Record playback intent before Steam can return another welcome.
	d.mtx.Lock()
	if d.closed {
		d.mtx.Unlock()
		return
	}
	d.stopped = !playing
	d.mtx.Unlock()

	// Release session admission before sending the stop announcement.
	if !playing {
		d.setConnectionStatus(gcsm.GCConnectionStatus_GCConnectionStatus_NO_SESSION, nil)
		d.coordinator.SetGamesPlayed()
		return
	}
	d.coordinator.SetGamesPlayed(AppID)
}

// SayHello requests a GC session with the supplied shared-object cache versions.
func (d *Dota2) SayHello(haveCacheVersions ...*gcsm.CMsgSOCacheHaveVersion) {
	d.write(uint32(gcsm.EGCBaseClientMsg_k_EMsgGCClientHello), &gcsm.CMsgClientHello{
		ClientLauncher:      gcsm.PartnerAccountType_PARTNER_NONE.Enum(),
		Engine:              gcsm.ESourceEngine_k_ESE_Source2.Enum(),
		ClientSessionNeed:   proto.Uint32(104),
		SocacheHaveVersions: haveCacheVersions,
	})
}

// handleClientWelcome admits requests after applying the coordinator's caches.
func (d *Dota2) handleClientWelcome(packet *gamecoordinator.GCPacket) error {
	// Decode the welcome and synchronize its shared-object subscriptions.
	welcome := &gcsm.CMsgClientWelcome{}
	if err := d.unmarshalBody(packet, welcome); err != nil {
		return err
	}
	for _, cache := range welcome.GetUptodateSubscribedCaches() {
		d.RequestCacheSubscriptionRefresh(cache.GetOwnerSoid())
	}
	for _, cache := range welcome.GetOutofdateSubscribedCaches() {
		if err := d.cache.HandleSubscribed(cache); err != nil {
			d.le.WithError(err).Warn("unable to handle welcome cache")
		}
	}

	// Publish a welcome only when the current handler still accepts sessions.
	if d.setConnectionStatus(gcsm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil) {
		d.emit(&devents.ClientWelcomed{Welcome: welcome})
	}
	return nil
}

// handleConnectionStatus applies a coordinator-reported session transition.
func (d *Dota2) handleConnectionStatus(packet *gamecoordinator.GCPacket) error {
	status := &gcsm.CMsgConnectionStatus{}
	if err := d.unmarshalBody(packet, status); err != nil {
		return err
	}
	if status.Status != nil {
		d.setConnectionStatus(*status.Status, status)
	}
	return nil
}

// setConnectionStatus changes admission and cancels requests before publishing events.
// It returns false when the requested transition does not establish a new state.
func (d *Dota2) setConnectionStatus(status gcsm.GCConnectionStatus, update *gcsm.CMsgConnectionStatus) bool {
	// Replace session state and its request lifetime under the same lock.
	d.mtx.Lock()
	if d.state.ConnectionStatus == status ||
		(status == gcsm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION && (d.closed || d.stopped)) {
		d.mtx.Unlock()
		return false
	}
	previous := d.state
	d.state.ClearState()
	d.state.ConnectionStatus = status
	d.state.LastConnectionStatusUpdate = update
	if d.connectionCtxCancel != nil {
		d.connectionCtxCancel()
		d.connectionCtxCancel = nil
		d.connectionCtx = nil
	}
	if status == gcsm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION {
		d.connectionCtx, d.connectionCtxCancel = context.WithCancel(context.Background())
	}
	next := d.state
	d.mtx.Unlock()

	// Event consumers may call back into Dota2 without holding its state lock.
	d.emit(&devents.GCConnectionStatusChanged{
		OldState: previous.ConnectionStatus,
		NewState: status,
		Update:   update,
	})
	d.emit(devents.ClientStateChanged{OldState: previous, NewState: next})
	return true
}
