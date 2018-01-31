package dota2

import (
	gcmcc "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// SendChannelMessage attempts to send a message in a channel.
func (d *Dota2) SendChannelMessage(channelID uint64, message string) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCChatMessage), &gcmcc.CMsgDOTAChatMessage{
		ChannelId: &channelID,
		Text:      &message,
	})
}

// SendChannelMessageEx attempts to send a message in a channel with explicit details.
// At minimum ChannelId and Text must be set.
func (d *Dota2) SendChannelMessageEx(channelID uint64, msg *gcmcc.CMsgDOTAChatMessage) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCChatMessage), msg)
}
