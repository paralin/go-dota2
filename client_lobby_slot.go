package dota2

import (
	gcccm "github.com/paralin/go-dota2/protocol"
	gcm "github.com/paralin/go-dota2/protocol"
)

// JoinLobbyTeam switches team in a lobby.
func (d *Dota2) JoinLobbyTeam(team gcm.DOTA_GC_TEAM, slot uint32) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team: &team,
		Slot: &slot,
	})
}

// SetLobbySlotBotDifficulty sets the difficulty of a slot to a given bot difficulty.
func (d *Dota2) SetLobbySlotBotDifficulty(team gcm.DOTA_GC_TEAM, slot uint32, botDifficulty gcm.DOTABotDifficulty) {
	d.write(uint32(gcm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot), &gcccm.CMsgPracticeLobbySetTeamSlot{
		Team:          &team,
		Slot:          &slot,
		BotDifficulty: &botDifficulty,
	})
}
