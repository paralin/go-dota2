package events

import (
	gcmcc "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
)

// ChatMessage is emitted when a chat message is observed.
type ChatMessage struct {
	gcmcc.CMsgDOTAChatMessage
}

// LeftChatChannel is emitted when we left a channel.
type LeftChatChannel struct {
	gcmcc.CMsgDOTAOtherLeftChatChannel
}

// JoinedChatChannel is emitted when we joined a channel.
type JoinedChatChannel struct {
	gcmcc.CMsgDOTAOtherJoinedChatChannel
}
