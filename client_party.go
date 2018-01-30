package dota2

import (
	"github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-dota2/events"
	bgcm "github.com/paralin/go-dota2/protocol/base_gcmessages"
)

// LeaveParty attempts to leave the current party.
func (d *Dota2) LeaveParty() {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLeaveParty), &bgcm.CMsgLeaveParty{})
}

// handleGcInvitationCreated handles game coordinator invitation created.
func (d *Dota2) handleGcInvitationCreated(packet *gamecoordinator.GCPacket) error {
	eve := &events.InvitationCreated{}
	if err := d.unmarshalBody(packet, &eve.CMsgInvitationCreated); err != nil {
		return err
	}

	d.emit(eve)
	return nil
}
