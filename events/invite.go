package events

import (
	bgcm "github.com/paralin/go-dota2/protocol/base_gcmessages"
)

// InvitationCreated is fired when the
type InvitationCreated struct {
	bgcm.CMsgInvitationCreated
}
