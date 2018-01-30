package dota2

import (
	"github.com/faceit/go-steam/steamid"
	bgcm "github.com/paralin/go-dota2/protocol/base_gcmessages"
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// LobbyKickPlayer kicks a player from the lobby by account ID.
func (d *Dota2) LobbyKickPlayer(playerID steamid.SteamId) {
	accountID := playerID.GetAccountId()
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick), &gcccm.CMsgPracticeLobbyKick{
		AccountId: &accountID,
	})
}

// LobbyInvitePlayer attempts to invite a player to the current lobby.
func (d *Dota2) LobbyInvitePlayer(playerID steamid.SteamId) {
	steamID := playerID.ToUint64()
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToLobby), &bgcm.CMsgInviteToLobby{
		SteamId: &steamID,
	})
}
