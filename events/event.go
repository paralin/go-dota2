package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

// Event is a DOTA event.
type Event interface {
	// GetDotaEventMsgID returns the DOTA event message ID.
	GetDotaEventMsgID() protocol.EDOTAGCMsg
	// GetEventBody event body.
	GetEventBody() proto.Message
	// GetEventName returns the event name.
	GetEventName() string
}
