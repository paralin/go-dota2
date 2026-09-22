package events

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	"github.com/paralin/go-dota2/protocol"
)

// Event is a DOTA event.
type Event interface {
	// GetDotaEventMsgID returns the DOTA event message ID.
	GetDotaEventMsgID() protocol.EDOTAGCMsg
	// GetEventBody returns the generated lite message.
	GetEventBody() protobuf.Message
	// GetEventName returns the event name.
	GetEventName() string
}
