package dota2

import (
	"context"

	"github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-dota2/events"
	gcmcc "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	"github.com/paralin/go-dota2/protocol/dota_shared_enums"
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

// JoinChatChannel attempts to join a chat channel by name and type.
func (d *Dota2) JoinChatChannel(ctx context.Context, name string, channelType dota_shared_enums.DOTAChatChannelTypeT) (*gcmcc.CMsgDOTAJoinChatChannelResponse, error) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCJoinChatChannel), &gcmcc.CMsgDOTAJoinChatChannel{
		ChannelName: &name,
		ChannelType: &channelType,
	})

	defer d.joinChatChannelHandlers.Delete(name)
	ch := make(chan *gcmcc.CMsgDOTAJoinChatChannelResponse, 1)
	d.joinChatChannelHandlers.Store(name, ch)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case v := <-ch:
		return v, nil
	}
}

// handleJoinChatChannelResponse handles the response to the join chat channel request
func (d *Dota2) handleJoinChatChannelResponse(packet *gamecoordinator.GCPacket) error {
	resp := &gcmcc.CMsgDOTAJoinChatChannelResponse{}
	if err := d.unmarshalBody(packet, resp); err != nil {
		return err
	}

	chanName := resp.GetChannelName()
	reqCh, ok := d.joinChatChannelHandlers.Load(chanName)
	if ok {
		reqCh.(chan *gcmcc.CMsgDOTAJoinChatChannelResponse) <- resp
	}

	return nil
}

// handleChatMessage handles an incoming chat message.
func (d *Dota2) handleChatMessage(packet *gamecoordinator.GCPacket) error {
	eve := &events.ChatMessage{}
	resp := &eve.CMsgDOTAChatMessage
	if err := d.unmarshalBody(packet, resp); err != nil {
		return err
	}

	d.emit(eve)
	return nil
}

// handleLeftChannel handles a channel leave event.
func (d *Dota2) handleLeftChannel(packet *gamecoordinator.GCPacket) error {
	eve := &events.LeftChatChannel{}
	resp := &eve.CMsgDOTAOtherLeftChatChannel
	if err := d.unmarshalBody(packet, resp); err != nil {
		return err
	}

	d.emit(eve)
	return nil
}

// handleJoinedChannel handles a channel join event.
func (d *Dota2) handleJoinedChannel(packet *gamecoordinator.GCPacket) error {
	eve := &events.JoinedChatChannel{}
	resp := &eve.CMsgDOTAOtherJoinedChatChannel
	if err := d.unmarshalBody(packet, resp); err != nil {
		return err
	}

	d.emit(eve)
	return nil
}
