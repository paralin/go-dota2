package events

import (
	"github.com/golang/protobuf/proto"
	gcmcc "github.com/paralin/go-dota2/protocol"
)

// ChatMessage is emitted when a chat message is observed.
type ChatMessage struct {
	gcmcc.CMsgDOTAChatMessage
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatMessage) GetDotaEventMsgID() gcmcc.EDOTAGCMsg {
	return gcmcc.EDOTAGCMsg_k_EMsgGCChatMessage
}

// GetEventBody returns the event body.
func (e *ChatMessage) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatMessage
}

// GetEventName returns the chat message event name.
func (e *ChatMessage) GetEventName() string {
	return "ChatMessage"
}

// JoinedChatChannel is emitted when we join a chat channel.
type JoinedChatChannel struct {
	gcmcc.CMsgDOTAJoinChatChannelResponse
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *JoinedChatChannel) GetDotaEventMsgID() gcmcc.EDOTAGCMsg {
	return gcmcc.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse
}

// GetEventBody returns the event body.
func (e *JoinedChatChannel) GetEventBody() proto.Message {
	return &e.CMsgDOTAJoinChatChannelResponse
}

// GetEventName returns the chat message event name.
func (e *JoinedChatChannel) GetEventName() string {
	return "JoinedChatChannel"
}
