package events

import (
	bgcm "github.com/paralin/go-dota2/protocol"
	"google.golang.org/protobuf/proto"
)

// InvitationCreated confirms that an invitation has been created.
type InvitationCreated struct {
	bgcm.CMsgInvitationCreated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *InvitationCreated) GetDotaEventMsgID() bgcm.EDOTAGCMsg {
	return bgcm.EDOTAGCMsg(bgcm.EGCBaseMsg_k_EMsgGCInvitationCreated)
}

// GetEventBody returns the event body.
func (e *InvitationCreated) GetEventBody() proto.Message {
	return &e.CMsgInvitationCreated
}

// GetEventName returns the event name.
func (e *InvitationCreated) GetEventName() string {
	return "InvitationCreated"
}
