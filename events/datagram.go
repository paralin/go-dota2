package events

import (
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
)

// SteamDatagramTicket contains an incoming steam datagram ticket.
type SteamDatagramTicket struct {
	gcccm.CMsgGCToClientSteamDatagramTicket
}
