package dota2

import (
	bgcm "github.com/paralin/go-dota2/protocol/base_gcmessages"
)

// LeaveParty attempts to leave the current party.
func (d *Dota2) LeaveParty() {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLeaveParty), &bgcm.CMsgLeaveParty{})
}
