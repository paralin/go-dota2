package dota2

import (
	"github.com/faceit/go-steam/steamid"
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	"github.com/paralin/go-dota2/protocol/dota_shared_enums"
)

// LobbyJoinTeam switches team in a lobby.
func (d *Dota2) LobbyJoinTeam(team dota_shared_enums.DOTA_GC_TEAM, slot uint32) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team: &team,
		Slot: &slot,
	})
}

// LobbySetSlotBotDifficulty sets the difficulty of a slot to a given bot difficulty.
func (d *Dota2) LobbySetSlotBotDifficulty(team dota_shared_enums.DOTA_GC_TEAM, slot uint32, botDifficulty dota_shared_enums.DOTABotDifficulty) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team:          &team,
		Slot:          &slot,
		BotDifficulty: &botDifficulty,
	})
}

// LobbyKickPlayerFromTeam kicks a player from whatever team they are on.
func (d *Dota2) LobbyKickPlayerFromTeam(player steamid.SteamId) {
	accountID := player.GetAccountId()
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam), &gcccm.CMsgPracticeLobbyKickFromTeam{
		AccountId: &accountID,
	})
}
