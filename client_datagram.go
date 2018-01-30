package dota2

import (
	"github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/faceit/go-steam/steamid"
	"github.com/paralin/go-dota2/events"
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// RequestSteamDatagramTicket requests a steam datagram ticket.
func (d *Dota2) RequestSteamDatagramTicket(serverID steamid.SteamId) {
	steamID := serverID.ToUint64()
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicket), &gcccm.CMsgClientToGCRequestSteamDatagramTicket{
		ServerSteamId: &steamID,
	})
}

// handleSteamDatagramTicket handles an incoming steam datagram ticket.
func (d *Dota2) handleSteamDatagramTicket(packet *gamecoordinator.GCPacket) error {
	ev := &events.SteamDatagramTicket{}
	if err := d.unmarshalBody(packet, &ev.CMsgGCToClientSteamDatagramTicket); err != nil {
		return err
	}

	d.emit(ev)
	return nil
}
