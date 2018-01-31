package dota2

import (
	gcccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	"github.com/paralin/go-dota2/protocol/dota_shared_enums"
)

// JoinLobbyTeam switches team in a lobby.
func (d *Dota2) JoinLobbyTeam(team dota_shared_enums.DOTA_GC_TEAM, slot uint32) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team: &team,
		Slot: &slot,
	})
}

// SetLobbySlotBotDifficulty sets the difficulty of a slot to a given bot difficulty.
func (d *Dota2) SetLobbySlotBotDifficulty(team dota_shared_enums.DOTA_GC_TEAM, slot uint32, botDifficulty dota_shared_enums.DOTABotDifficulty) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team:          &team,
		Slot:          &slot,
		BotDifficulty: &botDifficulty,
	})
}
