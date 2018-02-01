package events

import (
	"github.com/golang/protobuf/proto"
	gcmcc "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// ChatMessage is emitted when a chat message is observed.
type ChatMessage struct {
	gcmcc.CMsgDOTAChatMessage
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatMessage) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCChatMessage
}

// GetEventBody returns the event body.
func (e *ChatMessage) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatMessage
}

// GetEventName returns the chat message event name.
func (e *ChatMessage) GetEventName() string {
	return "ChatMessage"
}
