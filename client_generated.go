package dota2

import (
	"context"

	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol"
	"github.com/paralin/go-steam/steamid"
)

// AbandonLobby abandons the current practice lobby or live game.
func (d *Dota2) AbandonLobby() {
	req := &protocol.CMsgAbandonCurrentGame{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame), req)
}

// AckPartyReadyCheck acknowledges a party-wide ready check on behalf of your
// account.
func (d *Dota2) AckPartyReadyCheck(
	readyStatus protocol.EReadyCheckStatus,
) {
	req := &protocol.CMsgPartyReadyCheckAcknowledge{
		ReadyStatus: &readyStatus,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge), req)
}

// ApplyGemCombiner applys a gem combiner.
//
// Sends the GC message k_EMsgClientToGCApplyGemCombiner (CMsgClientToGCApplyGemCombiner). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ApplyGemCombiner(
	itemID1 uint64,
	itemID2 uint64,
) {
	req := &protocol.CMsgClientToGCApplyGemCombiner{
		ItemId_1: &itemID1,
		ItemId_2: &itemID2,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCApplyGemCombiner), req)
}

// ApplyModerationShowcaseModeration applys a moderation showcase moderation.
//
// Sends the GC message k_EMsgClientToGCShowcaseModerationApplyModeration (CMsgClientToGCShowcaseModerationApplyModeration) and awaits the response k_EMsgClientToGCShowcaseModerationApplyModerationResponse,
// delivered as *CMsgClientToGCShowcaseModerationApplyModerationResponse.
func (d *Dota2) ApplyModerationShowcaseModeration(
	ctx context.Context,
	accountID uint32,
	showcaseType protocol.EShowcaseType,
	showcaseTimestamp uint32,
	approve bool,
) (*protocol.CMsgClientToGCShowcaseModerationApplyModerationResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseModerationApplyModeration{
		AccountId:         &accountID,
		ShowcaseType:      &showcaseType,
		ShowcaseTimestamp: &showcaseTimestamp,
		Approve:           &approve,
	}
	resp := &protocol.CMsgClientToGCShowcaseModerationApplyModerationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseModerationApplyModeration),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseModerationApplyModerationResponse),
		resp,
	)
}

// ApplyTeamToLobby applys a team to lobby.
//
// Sends the GC message k_EMsgGCApplyTeamToPracticeLobby (CMsgApplyTeamToPracticeLobby). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ApplyTeamToLobby(
	teamID uint32,
) {
	req := &protocol.CMsgApplyTeamToPracticeLobby{
		TeamId: &teamID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCApplyTeamToPracticeLobby), req)
}

// AutographReward autographs a reward.
//
// Sends the GC message k_EMsgGameAutographReward (CMsgDOTAGameAutographReward) and awaits the response k_EMsgGameAutographRewardResponse,
// delivered as *CMsgDOTAGameAutographRewardResponse.
func (d *Dota2) AutographReward(
	ctx context.Context,
	badgeID string,
) (*protocol.CMsgDOTAGameAutographRewardResponse, error) {
	req := &protocol.CMsgDOTAGameAutographReward{
		BadgeId: &badgeID,
	}
	resp := &protocol.CMsgDOTAGameAutographRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGameAutographReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGameAutographRewardResponse),
		resp,
	)
}

// CancelGameFightingChallengeFriend cancels a game fighting challenge friend.
//
// Sends the GC message k_EMsgClientToGCFightingGameCancelChallengeFriend (CMsgClientToGCFightingGameCancelChallengeFriend). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CancelGameFightingChallengeFriend(
	friendAccountID uint32,
) {
	req := &protocol.CMsgClientToGCFightingGameCancelChallengeFriend{
		FriendAccountId: &friendAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFightingGameCancelChallengeFriend), req)
}

// CancelInviteToGuild cancels a invite to guild.
//
// Sends the GC message k_EMsgClientToGCCancelInviteToGuild (CMsgClientToGCCancelInviteToGuild) and awaits the response k_EMsgClientToGCCancelInviteToGuildResponse,
// delivered as *CMsgClientToGCCancelInviteToGuildResponse.
func (d *Dota2) CancelInviteToGuild(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCCancelInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCCancelInviteToGuild{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCCancelInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelInviteToGuildResponse),
		resp,
	)
}

// CancelPartyInvites cancels all outstanding party invites sent by your
// account.
func (d *Dota2) CancelPartyInvites(
	invitedSteamids []uint64,
	invitedGroupids []uint64,
) {
	req := &protocol.CMsgDOTACancelGroupInvites{
		InvitedSteamids: invitedSteamids,
		InvitedGroupids: invitedGroupids,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites), req)
}

// CancelWatchGame cancels a watch game.
//
// Sends the GC message k_EMsgGCCancelWatchGame (CMsgCancelWatchGame). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CancelWatchGame() {
	req := &protocol.CMsgCancelWatchGame{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCCancelWatchGame), req)
}

// ClaimBingoRow claims a bingo row.
//
// Sends the GC message k_EMsgClientToGCBingoClaimRow (CMsgClientToGCBingoClaimRow) and awaits the response k_EMsgClientToGCBingoClaimRowResponse,
// delivered as *CMsgClientToGCBingoClaimRowResponse.
func (d *Dota2) ClaimBingoRow(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
	rowIndex uint32,
) (*protocol.CMsgClientToGCBingoClaimRowResponse, error) {
	req := &protocol.CMsgClientToGCBingoClaimRow{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
		RowIndex:    &rowIndex,
	}
	resp := &protocol.CMsgClientToGCBingoClaimRowResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoClaimRow),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoClaimRowResponse),
		resp,
	)
}

// ClaimCrawlCavernRoom claims a crawl cavern room.
//
// Sends the GC message k_EMsgClientToGCCavernCrawlClaimRoom (CMsgClientToGCCavernCrawlClaimRoom) and awaits the response k_EMsgClientToGCCavernCrawlClaimRoomResponse,
// delivered as *CMsgClientToGCCavernCrawlClaimRoomResponse.
func (d *Dota2) ClaimCrawlCavernRoom(
	ctx context.Context,
	eventID uint32,
	roomID uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlClaimRoomResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlClaimRoom{
		EventId:    &eventID,
		RoomId:     &roomID,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlClaimRoomResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlClaimRoom),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlClaimRoomResponse),
		resp,
	)
}

// ClaimEventAction claims a event action.
//
// Sends the GC message k_EMsgDOTAClaimEventAction (CMsgDOTAClaimEventAction) and awaits the response k_EMsgDOTAClaimEventActionResponse,
// delivered as *CMsgDOTAClaimEventActionResponse.
func (d *Dota2) ClaimEventAction(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
	quantity uint32,
	data protocol.CMsgDOTAClaimEventActionData,
	scoreMode protocol.EEventActionScoreMode,
	suppressRewards bool,
) (*protocol.CMsgDOTAClaimEventActionResponse, error) {
	req := &protocol.CMsgDOTAClaimEventAction{
		EventId:         &eventID,
		ActionId:        &actionID,
		Quantity:        &quantity,
		Data:            &data,
		ScoreMode:       &scoreMode,
		SuppressRewards: &suppressRewards,
	}
	resp := &protocol.CMsgDOTAClaimEventActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAClaimEventAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAClaimEventActionResponse),
		resp,
	)
}

// ClaimEventActionUsingItem claims a event action using item.
//
// Sends the GC message k_EMsgClientToGCClaimEventActionUsingItem (CMsgClientToGCClaimEventActionUsingItem) and awaits the response k_EMsgClientToGCClaimEventActionUsingItemResponse,
// delivered as *CMsgClientToGCClaimEventActionUsingItemResponse.
func (d *Dota2) ClaimEventActionUsingItem(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
	itemID uint64,
	quantity uint32,
	suppressRewards bool,
) (*protocol.CMsgClientToGCClaimEventActionUsingItemResponse, error) {
	req := &protocol.CMsgClientToGCClaimEventActionUsingItem{
		EventId:         &eventID,
		ActionId:        &actionID,
		ItemId:          &itemID,
		Quantity:        &quantity,
		SuppressRewards: &suppressRewards,
	}
	resp := &protocol.CMsgClientToGCClaimEventActionUsingItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimEventActionUsingItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimEventActionUsingItemResponse),
		resp,
	)
}

// ClaimGatedEvent claims a gated event.
//
// Sends the GC message k_EMsgClientToGCClaimGatedEvent (CMsgDOTAClaimGatedEvent) and awaits the response k_EMsgClientToGCClaimGatedEventResponse,
// delivered as *CMsgDOTAClaimGatedEventResponse.
func (d *Dota2) ClaimGatedEvent(
	ctx context.Context,
	eventID protocol.EEvent,
) (*protocol.CMsgDOTAClaimGatedEventResponse, error) {
	req := &protocol.CMsgDOTAClaimGatedEvent{
		EventId: &eventID,
	}
	resp := &protocol.CMsgDOTAClaimGatedEventResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimGatedEvent),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimGatedEventResponse),
		resp,
	)
}

// ClaimHunterDevMonsterInvestigationRewards claims hunter dev monster investigation rewards.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterDevClaimInvestigationRewards (CMsgClientToGCMonsterHunterDevClaimInvestigationRewards) and awaits the response k_EMsgClientToGCMonsterHunterDevClaimInvestigationRewardsResponse,
// delivered as *CMsgClientToGCMonsterHunterDevClaimInvestigationRewardsResponse.
func (d *Dota2) ClaimHunterDevMonsterInvestigationRewards(
	ctx context.Context,
	investigationGameState protocol.CMsgMonsterHunterInvestigationGameState,
	win bool,
) (*protocol.CMsgClientToGCMonsterHunterDevClaimInvestigationRewardsResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterDevClaimInvestigationRewards{
		InvestigationGameState: &investigationGameState,
		Win:                    &win,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterDevClaimInvestigationRewardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevClaimInvestigationRewards),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevClaimInvestigationRewardsResponse),
		resp,
	)
}

// ClaimHunterMonsterCodexReward claims a hunter monster codex reward.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterClaimCodexReward (CMsgClientToGCMonsterHunterClaimCodexReward) and awaits the response k_EMsgClientToGCMonsterHunterClaimCodexRewardResponse,
// delivered as *CMsgClientToGCMonsterHunterClaimCodexRewardResponse.
func (d *Dota2) ClaimHunterMonsterCodexReward(
	ctx context.Context,
	codexID uint32,
	reward uint32,
) (*protocol.CMsgClientToGCMonsterHunterClaimCodexRewardResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterClaimCodexReward{
		CodexId: &codexID,
		Reward:  &reward,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterClaimCodexRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimCodexReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimCodexRewardResponse),
		resp,
	)
}

// ClaimHunterMonsterReward claims a hunter monster reward.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterClaimReward (CMsgClientToGCMonsterHunterClaimReward) and awaits the response k_EMsgClientToGCMonsterHunterClaimRewardResponse,
// delivered as *CMsgClientToGCMonsterHunterClaimRewardResponse.
func (d *Dota2) ClaimHunterMonsterReward(
	ctx context.Context,
	req *protocol.CMsgClientToGCMonsterHunterClaimReward,
) (*protocol.CMsgClientToGCMonsterHunterClaimRewardResponse, error) {
	resp := &protocol.CMsgClientToGCMonsterHunterClaimRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimRewardResponse),
		resp,
	)
}

// ClaimHunterMonsterSetReward claims a hunter monster set reward.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterClaimSetReward (CMsgClientToGCMonsterHunterClaimSetReward) and awaits the response k_EMsgClientToGCMonsterHunterClaimSetRewardResponse,
// delivered as *CMsgClientToGCMonsterHunterClaimSetRewardResponse.
func (d *Dota2) ClaimHunterMonsterSetReward(
	ctx context.Context,
	itemSets []*protocol.CMsgMonsterHunterItemSet,
) (*protocol.CMsgClientToGCMonsterHunterClaimSetRewardResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterClaimSetReward{
		ItemSets: itemSets,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterClaimSetRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimSetReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimSetRewardResponse),
		resp,
	)
}

// ClaimLeaderboardRewards claims leaderboard rewards.
//
// Sends the GC message k_EMsgClientToGCClaimLeaderboardRewards (CMsgClientToGCClaimLeaderboardRewards) and awaits the response k_EMsgClientToGCClaimLeaderboardRewardsResponse,
// delivered as *CMsgClientToGCClaimLeaderboardRewardsResponse.
func (d *Dota2) ClaimLeaderboardRewards(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCClaimLeaderboardRewardsResponse, error) {
	req := &protocol.CMsgClientToGCClaimLeaderboardRewards{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCClaimLeaderboardRewardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimLeaderboardRewards),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimLeaderboardRewardsResponse),
		resp,
	)
}

// ClaimOverworldEncounterReward claims a overworld encounter reward.
//
// Sends the GC message k_EMsgClientToGCOverworldClaimEncounterReward (CMsgClientToGCOverworldClaimEncounterReward) and awaits the response k_EMsgClientToGCOverworldClaimEncounterRewardResponse,
// delivered as *CMsgClientToGCOverworldClaimEncounterRewardResponse.
func (d *Dota2) ClaimOverworldEncounterReward(
	ctx context.Context,
	req *protocol.CMsgClientToGCOverworldClaimEncounterReward,
) (*protocol.CMsgClientToGCOverworldClaimEncounterRewardResponse, error) {
	resp := &protocol.CMsgClientToGCOverworldClaimEncounterRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimEncounterReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimEncounterRewardResponse),
		resp,
	)
}

// ClaimOverworldFortunePermanentReward claims a overworld fortune permanent reward.
//
// Sends the GC message k_EMsgClientToGCOverworldClaimFortunePermanentReward (CMsgClientToGCOverworldClaimFortunePermanentReward) and awaits the response k_EMsgClientToGCOverworldClaimFortunePermanentRewardResponse,
// delivered as *CMsgClientToGCOverworldClaimFortunePermanentRewardResponse.
func (d *Dota2) ClaimOverworldFortunePermanentReward(
	ctx context.Context,
	overworldID uint32,
	fortuneID uint32,
) (*protocol.CMsgClientToGCOverworldClaimFortunePermanentRewardResponse, error) {
	req := &protocol.CMsgClientToGCOverworldClaimFortunePermanentReward{
		OverworldId: &overworldID,
		FortuneId:   &fortuneID,
	}
	resp := &protocol.CMsgClientToGCOverworldClaimFortunePermanentRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortunePermanentReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortunePermanentRewardResponse),
		resp,
	)
}

// ClaimOverworldFortuneReward claims a overworld fortune reward.
//
// Sends the GC message k_EMsgClientToGCOverworldClaimFortuneReward (CMsgClientToGCOverworldClaimFortuneReward) and awaits the response k_EMsgClientToGCOverworldClaimFortuneRewardResponse,
// delivered as *CMsgClientToGCOverworldClaimFortuneRewardResponse.
func (d *Dota2) ClaimOverworldFortuneReward(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldClaimFortuneRewardResponse, error) {
	req := &protocol.CMsgClientToGCOverworldClaimFortuneReward{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldClaimFortuneRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortuneReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortuneRewardResponse),
		resp,
	)
}

// ClaimOverworldFortuneTellerStoryNode claims a overworld fortune teller story node.
//
// Sends the GC message k_EMsgClientToGCOverworldClaimFortuneTellerStoryNode (CMsgClientToGCOverworldClaimFortuneTellerStoryNode) and awaits the response k_EMsgClientToGCOverworldClaimFortuneTellerStoryNodeResponse,
// delivered as *CMsgClientToGCOverworldClaimFortuneTellerStoryNodeResponse.
func (d *Dota2) ClaimOverworldFortuneTellerStoryNode(
	ctx context.Context,
	overworldID uint32,
	storyNodeID uint32,
) (*protocol.CMsgClientToGCOverworldClaimFortuneTellerStoryNodeResponse, error) {
	req := &protocol.CMsgClientToGCOverworldClaimFortuneTellerStoryNode{
		OverworldId: &overworldID,
		StoryNodeId: &storyNodeID,
	}
	resp := &protocol.CMsgClientToGCOverworldClaimFortuneTellerStoryNodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortuneTellerStoryNode),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldClaimFortuneTellerStoryNodeResponse),
		resp,
	)
}

// ClaimSwag claims a swag.
//
// Sends the GC message k_EMsgClientToGCClaimSwag (CMsgClientToGCClaimSwag) and awaits the response k_EMsgGCToClientClaimSwagResponse,
// delivered as *CMsgClientToGCClaimSwagResponse.
func (d *Dota2) ClaimSwag(
	ctx context.Context,
	eventID protocol.EEvent,
	actionID uint32,
	data uint32,
) (*protocol.CMsgClientToGCClaimSwagResponse, error) {
	req := &protocol.CMsgClientToGCClaimSwag{
		EventId:  &eventID,
		ActionId: &actionID,
		Data:     &data,
	}
	resp := &protocol.CMsgClientToGCClaimSwagResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCClaimSwag),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientClaimSwagResponse),
		resp,
	)
}

// CloseLobbyBroadcastChannel closes a lobby broadcast channel.
//
// Sends the GC message k_EMsgGCPracticeLobbyCloseBroadcastChannel (CMsgPracticeLobbyCloseBroadcastChannel). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CloseLobbyBroadcastChannel(
	channel uint32,
) {
	req := &protocol.CMsgPracticeLobbyCloseBroadcastChannel{
		Channel: &channel,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyCloseBroadcastChannel), req)
}

// CreateBotGame creates a bot game.
//
// Sends the GC message k_EMsgGCBotGameCreate (CMsgBotGameCreate). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CreateBotGame(
	searchKey string,
	difficultyRadiant protocol.DOTABotDifficulty,
	team protocol.DOTA_GC_TEAM,
	gameMode uint32,
	difficultyDire protocol.DOTABotDifficulty,
) {
	req := &protocol.CMsgBotGameCreate{
		SearchKey:         &searchKey,
		DifficultyRadiant: &difficultyRadiant,
		Team:              &team,
		GameMode:          &gameMode,
		DifficultyDire:    &difficultyDire,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCBotGameCreate), req)
}

// CreateGuild creates a guild.
//
// Sends the GC message k_EMsgClientToGCCreateGuild (CMsgClientToGCCreateGuild) and awaits the response k_EMsgClientToGCCreateGuildResponse,
// delivered as *CMsgClientToGCCreateGuildResponse.
func (d *Dota2) CreateGuild(
	ctx context.Context,
	guildInfo protocol.CMsgGuildInfo,
	guildChatType protocol.EGuildChatType,
) (*protocol.CMsgClientToGCCreateGuildResponse, error) {
	req := &protocol.CMsgClientToGCCreateGuild{
		GuildInfo:     &guildInfo,
		GuildChatType: &guildChatType,
	}
	resp := &protocol.CMsgClientToGCCreateGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateGuildResponse),
		resp,
	)
}

// CreateHeroStatue creates a hero statue.
//
// Sends the GC message k_EMsgClientToGCCreateHeroStatue (CMsgClientToGCCreateHeroStatue). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CreateHeroStatue(
	req *protocol.CMsgClientToGCCreateHeroStatue,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateHeroStatue), req)
}

// CreatePlayerCardPack creates a player card pack.
//
// Sends the GC message k_EMsgClientToGCCreatePlayerCardPack (CMsgClientToGCCreatePlayerCardPack) and awaits the response k_EMsgClientToGCCreatePlayerCardPackResponse,
// delivered as *CMsgClientToGCCreatePlayerCardPackResponse.
func (d *Dota2) CreatePlayerCardPack(
	ctx context.Context,
	cardDustItemID uint64,
	eventID uint32,
	premiumPack bool,
) (*protocol.CMsgClientToGCCreatePlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCCreatePlayerCardPack{
		CardDustItemId: &cardDustItemID,
		EventId:        &eventID,
		PremiumPack:    &premiumPack,
	}
	resp := &protocol.CMsgClientToGCCreatePlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse),
		resp,
	)
}

// CreateSpectatorLobby creates a spectator lobby.
//
// Sends the GC message k_EMsgClientToGCCreateSpectatorLobby (CMsgCreateSpectatorLobby). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) CreateSpectatorLobby(
	details protocol.CMsgSetSpectatorLobbyDetails,
) {
	req := &protocol.CMsgCreateSpectatorLobby{
		Details: &details,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobby), req)
}

// CreateTeam creates a team.
//
// Sends the GC message k_EMsgGCCreateTeam (CMsgDOTACreateTeam) and awaits the response k_EMsgGCCreateTeamResponse,
// delivered as *CMsgDOTACreateTeamResponse.
func (d *Dota2) CreateTeam(
	ctx context.Context,
	req *protocol.CMsgDOTACreateTeam,
) (*protocol.CMsgDOTACreateTeamResponse, error) {
	resp := &protocol.CMsgDOTACreateTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCreateTeam),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCreateTeamResponse),
		resp,
	)
}

// CreateTeamPlayerCardPack creates a team player card pack.
//
// Sends the GC message k_EMsgClientToGCCreateTeamPlayerCardPack (CMsgClientToGCCreateTeamPlayerCardPack) and awaits the response k_EMsgClientToGCCreateTeamPlayerCardPackResponse,
// delivered as *CMsgClientToGCCreateTeamPlayerCardPackResponse.
func (d *Dota2) CreateTeamPlayerCardPack(
	ctx context.Context,
	cardDustItemID uint64,
	eventID uint32,
	premiumPack bool,
	teamID uint32,
) (*protocol.CMsgClientToGCCreateTeamPlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCCreateTeamPlayerCardPack{
		CardDustItemId: &cardDustItemID,
		EventId:        &eventID,
		PremiumPack:    &premiumPack,
		TeamId:         &teamID,
	}
	resp := &protocol.CMsgClientToGCCreateTeamPlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateTeamPlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateTeamPlayerCardPackResponse),
		resp,
	)
}

// DemotePrivateChatMember demotes a private chat member.
//
// Sends the GC message k_EMsgClientToGCPrivateChatDemote (CMsgClientToGCPrivateChatDemote). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) DemotePrivateChatMember(
	privateChatChannelName string,
	demoteAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatDemote{
		PrivateChatChannelName: &privateChatChannelName,
		DemoteAccountId:        &demoteAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote), req)
}

// DestroyLobby destroys the practice lobby you lead. Only the lobby leader
// can destroy it.
func (d *Dota2) DestroyLobby(
	ctx context.Context,
) (*protocol.CMsgDOTADestroyLobbyResponse, error) {
	req := &protocol.CMsgDOTADestroyLobbyRequest{}
	resp := &protocol.CMsgDOTADestroyLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDestroyLobbyRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDestroyLobbyResponse),
		resp,
	)
}

// EditTeamDetails edits team details.
//
// Sends the GC message k_EMsgGCEditTeamDetails (CMsgDOTAEditTeamDetails) and awaits the response k_EMsgGCEditTeamDetailsResponse,
// delivered as *CMsgDOTAEditTeamDetailsResponse.
func (d *Dota2) EditTeamDetails(
	ctx context.Context,
	req *protocol.CMsgDOTAEditTeamDetails,
) (*protocol.CMsgDOTAEditTeamDetailsResponse, error) {
	resp := &protocol.CMsgDOTAEditTeamDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCEditTeamDetails),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCEditTeamDetailsResponse),
		resp,
	)
}

// FindTopSourceTVGames finds top source tv games.
//
// Sends the GC message k_EMsgClientToGCFindTopSourceTVGames (CMsgClientToGCFindTopSourceTVGames) and awaits the response k_EMsgGCToClientFindTopSourceTVGamesResponse,
// delivered as *CMsgGCToClientFindTopSourceTVGamesResponse.
func (d *Dota2) FindTopSourceTVGames(
	ctx context.Context,
	searchKey string,
	leagueID uint32,
	heroID int32,
	startGame uint32,
	gameListIndex uint32,
	lobbyIDs []uint64,
) (*protocol.CMsgGCToClientFindTopSourceTVGamesResponse, error) {
	req := &protocol.CMsgClientToGCFindTopSourceTVGames{
		SearchKey:     &searchKey,
		LeagueId:      &leagueID,
		HeroId:        &heroID,
		StartGame:     &startGame,
		GameListIndex: &gameListIndex,
		LobbyIds:      lobbyIDs,
	}
	resp := &protocol.CMsgGCToClientFindTopSourceTVGamesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse),
		resp,
	)
}

// FlipLobbyTeams swaps every member of the lobby between the Radiant and
// Dire teams.
func (d *Dota2) FlipLobbyTeams() {
	req := &protocol.CMsgFlipLobbyTeams{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCFlipLobbyTeams), req)
}

// GetAdminShowcaseReportsRollup gets a admin showcase reports rollup.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminGetReportsRollup (CMsgClientToGCShowcaseAdminGetReportsRollup) and awaits the response k_EMsgClientToGCShowcaseAdminGetReportsRollupResponse,
// delivered as *CMsgClientToGCShowcaseAdminGetReportsRollupResponse.
func (d *Dota2) GetAdminShowcaseReportsRollup(
	ctx context.Context,
	rollupID uint32,
) (*protocol.CMsgClientToGCShowcaseAdminGetReportsRollupResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminGetReportsRollup{
		RollupId: &rollupID,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminGetReportsRollupResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetReportsRollup),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetReportsRollupResponse),
		resp,
	)
}

// GetAdminShowcaseReportsRollupList gets a admin showcase reports rollup list.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminGetReportsRollupList (CMsgClientToGCShowcaseAdminGetReportsRollupList) and awaits the response k_EMsgClientToGCShowcaseAdminGetReportsRollupListResponse,
// delivered as *CMsgClientToGCShowcaseAdminGetReportsRollupListResponse.
func (d *Dota2) GetAdminShowcaseReportsRollupList(
	ctx context.Context,
) (*protocol.CMsgClientToGCShowcaseAdminGetReportsRollupListResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminGetReportsRollupList{}
	resp := &protocol.CMsgClientToGCShowcaseAdminGetReportsRollupListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetReportsRollupList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetReportsRollupListResponse),
		resp,
	)
}

// GetAdminShowcaseUserDetails gets admin showcase user details.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminGetUserDetails (CMsgClientToGCShowcaseAdminGetUserDetails) and awaits the response k_EMsgClientToGCShowcaseAdminGetUserDetailsResponse,
// delivered as *CMsgClientToGCShowcaseAdminGetUserDetailsResponse.
func (d *Dota2) GetAdminShowcaseUserDetails(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCShowcaseAdminGetUserDetailsResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminGetUserDetails{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminGetUserDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetUserDetails),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminGetUserDetailsResponse),
		resp,
	)
}

// GetAllHeroOrder gets all hero order.
//
// Sends the GC message k_EMsgClientToGCGetAllHeroOrder (CMsgClientToGCGetAllHeroOrder) and awaits the response k_EMsgClientToGCGetAllHeroOrderResponse,
// delivered as *CMsgClientToGCGetAllHeroOrderResponse.
func (d *Dota2) GetAllHeroOrder(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetAllHeroOrderResponse, error) {
	req := &protocol.CMsgClientToGCGetAllHeroOrder{}
	resp := &protocol.CMsgClientToGCGetAllHeroOrderResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrder),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrderResponse),
		resp,
	)
}

// GetAllHeroProgress gets all hero progress.
//
// Sends the GC message k_EMsgClientToGCGetAllHeroProgress (CMsgClientToGCGetAllHeroProgress) and awaits the response k_EMsgClientToGCGetAllHeroProgressResponse,
// delivered as *CMsgClientToGCGetAllHeroProgressResponse.
func (d *Dota2) GetAllHeroProgress(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetAllHeroProgressResponse, error) {
	req := &protocol.CMsgClientToGCGetAllHeroProgress{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetAllHeroProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgress),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgressResponse),
		resp,
	)
}

// GetAvailablePrivateCoachingSessions gets available private coaching sessions.
//
// Sends the GC message k_EMsgClientToGCGetAvailablePrivateCoachingSessions (CMsgClientToGCGetAvailablePrivateCoachingSessions) and awaits the response k_EMsgClientToGCGetAvailablePrivateCoachingSessionsResponse,
// delivered as *CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse.
func (d *Dota2) GetAvailablePrivateCoachingSessions(
	ctx context.Context,
	language uint32,
) (*protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse, error) {
	req := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessions{
		Language: &language,
	}
	resp := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsResponse),
		resp,
	)
}

// GetAvailablePrivateCoachingSessionsSummary gets a available private coaching sessions summary.
//
// Sends the GC message k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummary (CMsgClientToGCGetAvailablePrivateCoachingSessionsSummary) and awaits the response k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse,
// delivered as *CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse.
func (d *Dota2) GetAvailablePrivateCoachingSessionsSummary(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse, error) {
	req := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummary{}
	resp := &protocol.CMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummary),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetAvailablePrivateCoachingSessionsSummaryResponse),
		resp,
	)
}

// GetBattleReport gets a battle report.
//
// Sends the GC message k_EMsgClientToGCGetBattleReport (CMsgClientToGCGetBattleReport) and awaits the response k_EMsgClientToGCGetBattleReportResponse,
// delivered as *CMsgClientToGCGetBattleReportResponse.
func (d *Dota2) GetBattleReport(
	ctx context.Context,
	accountID uint32,
	timestamp uint32,
	duration uint32,
) (*protocol.CMsgClientToGCGetBattleReportResponse, error) {
	req := &protocol.CMsgClientToGCGetBattleReport{
		AccountId: &accountID,
		Timestamp: &timestamp,
		Duration:  &duration,
	}
	resp := &protocol.CMsgClientToGCGetBattleReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReport),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportResponse),
		resp,
	)
}

// GetBattleReportAggregateStats gets battle report aggregate stats.
//
// Sends the GC message k_EMsgClientToGCGetBattleReportAggregateStats (CMsgClientToGCGetBattleReportAggregateStats) and awaits the response k_EMsgClientToGCGetBattleReportAggregateStatsResponse,
// delivered as *CMsgClientToGCGetBattleReportAggregateStatsResponse.
func (d *Dota2) GetBattleReportAggregateStats(
	ctx context.Context,
	aggregateKeys []*protocol.CMsgClientToGCGetBattleReportAggregateStats_CMsgBattleReportAggregateKey,
	timestamp uint32,
	duration uint32,
	rank uint32,
) (*protocol.CMsgClientToGCGetBattleReportAggregateStatsResponse, error) {
	req := &protocol.CMsgClientToGCGetBattleReportAggregateStats{
		AggregateKeys: aggregateKeys,
		Timestamp:     &timestamp,
		Duration:      &duration,
		Rank:          &rank,
	}
	resp := &protocol.CMsgClientToGCGetBattleReportAggregateStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportAggregateStats),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportAggregateStatsResponse),
		resp,
	)
}

// GetBattleReportInfo gets a battle report info.
//
// Sends the GC message k_EMsgClientToGCGetBattleReportInfo (CMsgClientToGCGetBattleReportInfo) and awaits the response k_EMsgClientToGCGetBattleReportInfoResponse,
// delivered as *CMsgClientToGCGetBattleReportInfoResponse.
func (d *Dota2) GetBattleReportInfo(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetBattleReportInfoResponse, error) {
	req := &protocol.CMsgClientToGCGetBattleReportInfo{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetBattleReportInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportInfoResponse),
		resp,
	)
}

// GetBattleReportMatchHistory gets a battle report match history.
//
// Sends the GC message k_EMsgClientToGCGetBattleReportMatchHistory (CMsgClientToGCGetBattleReportMatchHistory) and awaits the response k_EMsgClientToGCGetBattleReportMatchHistoryResponse,
// delivered as *CMsgClientToGCGetBattleReportMatchHistoryResponse.
func (d *Dota2) GetBattleReportMatchHistory(
	ctx context.Context,
	accountID uint32,
	timestamp uint32,
	duration uint32,
) (*protocol.CMsgClientToGCGetBattleReportMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCGetBattleReportMatchHistory{
		AccountId: &accountID,
		Timestamp: &timestamp,
		Duration:  &duration,
	}
	resp := &protocol.CMsgClientToGCGetBattleReportMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportMatchHistory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetBattleReportMatchHistoryResponse),
		resp,
	)
}

// GetBattlerItemUserData gets a battler item user data.
//
// Sends the GC message k_EMsgClientToGCItemBattlerGetUserData (CMsgClientToGCItemBattlerGetUserData) and awaits the response k_EMsgClientToGCItemBattlerGetUserDataResponse,
// delivered as *CMsgClientToGCItemBattlerGetUserDataResponse.
func (d *Dota2) GetBattlerItemUserData(
	ctx context.Context,
) (*protocol.CMsgClientToGCItemBattlerGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCItemBattlerGetUserData{}
	resp := &protocol.CMsgClientToGCItemBattlerGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerGetUserDataResponse),
		resp,
	)
}

// GetBingoStatsData gets a bingo stats data.
//
// Sends the GC message k_EMsgClientToGCBingoGetStatsData (CMsgClientToGCBingoGetStatsData) and awaits the response k_EMsgClientToGCBingoGetStatsDataResponse,
// delivered as *CMsgClientToGCBingoGetStatsDataResponse.
func (d *Dota2) GetBingoStatsData(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
) (*protocol.CMsgClientToGCBingoGetStatsDataResponse, error) {
	req := &protocol.CMsgClientToGCBingoGetStatsData{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
	}
	resp := &protocol.CMsgClientToGCBingoGetStatsDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoGetStatsData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoGetStatsDataResponse),
		resp,
	)
}

// GetBingoUserData gets a bingo user data.
//
// Sends the GC message k_EMsgClientToGCBingoGetUserData (CMsgClientToGCBingoGetUserData) and awaits the response k_EMsgClientToGCBingoGetUserDataResponse,
// delivered as *CMsgClientToGCBingoGetUserDataResponse.
func (d *Dota2) GetBingoUserData(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgClientToGCBingoGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCBingoGetUserData{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgClientToGCBingoGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoGetUserDataResponse),
		resp,
	)
}

// GetChatMemberCount gets a chat member count.
//
// Sends the GC message k_EMsgDOTAChatGetMemberCount (CMsgDOTAChatGetMemberCount) and awaits the response k_EMsgDOTAChatGetMemberCountResponse,
// delivered as *CMsgDOTAChatGetMemberCountResponse.
func (d *Dota2) GetChatMemberCount(
	ctx context.Context,
	channelName string,
	channelType protocol.DOTAChatChannelTypeT,
) (*protocol.CMsgDOTAChatGetMemberCountResponse, error) {
	req := &protocol.CMsgDOTAChatGetMemberCount{
		ChannelName: &channelName,
		ChannelType: &channelType,
	}
	resp := &protocol.CMsgDOTAChatGetMemberCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCount),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCountResponse),
		resp,
	)
}

// GetCraftingFantasyData gets a crafting fantasy data.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingGetData (CMsgClientToGCFantasyCraftingGetData) and awaits the response k_EMsgClientToGCFantasyCraftingGetDataResponse,
// delivered as *CMsgClientToGCFantasyCraftingGetDataResponse.
func (d *Dota2) GetCraftingFantasyData(
	ctx context.Context,
	fantasyLeague uint32,
	accountID uint32,
) (*protocol.CMsgClientToGCFantasyCraftingGetDataResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingGetData{
		FantasyLeague: &fantasyLeague,
		AccountId:     &accountID,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingGetDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingGetData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingGetDataResponse),
		resp,
	)
}

// GetCraftworksUserData gets a craftworks user data.
//
// Sends the GC message k_EMsgClientToGCCraftworksGetUserData (CMsgClientToGCCraftworksGetUserData) and awaits the response k_EMsgClientToGCCraftworksGetUserDataResponse,
// delivered as *CMsgClientToGCCraftworksGetUserDataResponse.
func (d *Dota2) GetCraftworksUserData(
	ctx context.Context,
	craftworksID uint32,
) (*protocol.CMsgClientToGCCraftworksGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCCraftworksGetUserData{
		CraftworksId: &craftworksID,
	}
	resp := &protocol.CMsgClientToGCCraftworksGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksGetUserDataResponse),
		resp,
	)
}

// GetCrawlCavernClaimedRoomCount gets a crawl cavern claimed room count.
//
// Sends the GC message k_EMsgClientToGCCavernCrawlGetClaimedRoomCount (CMsgClientToGCCavernCrawlGetClaimedRoomCount) and awaits the response k_EMsgClientToGCCavernCrawlGetClaimedRoomCountResponse,
// delivered as *CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse.
func (d *Dota2) GetCrawlCavernClaimedRoomCount(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCount{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlGetClaimedRoomCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlGetClaimedRoomCount),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlGetClaimedRoomCountResponse),
		resp,
	)
}

// GetCurrentPrivateCoachingSession gets a current private coaching session.
//
// Sends the GC message k_EMsgClientToGCGetCurrentPrivateCoachingSession (CMsgClientToGCGetCurrentPrivateCoachingSession) and awaits the response k_EMsgClientToGCGetCurrentPrivateCoachingSessionResponse,
// delivered as *CMsgClientToGCGetCurrentPrivateCoachingSessionResponse.
func (d *Dota2) GetCurrentPrivateCoachingSession(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetCurrentPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCGetCurrentPrivateCoachingSession{}
	resp := &protocol.CMsgClientToGCGetCurrentPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetCurrentPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetCurrentPrivateCoachingSessionResponse),
		resp,
	)
}

// GetDPCFavorites gets dpc favorites.
//
// Sends the GC message k_EMsgClientToGCGetDPCFavorites (CMsgClientToGCGetDPCFavorites) and awaits the response k_EMsgClientToGCGetDPCFavoritesResponse,
// delivered as *CMsgClientToGCGetDPCFavoritesResponse.
func (d *Dota2) GetDPCFavorites(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetDPCFavoritesResponse, error) {
	req := &protocol.CMsgClientToGCGetDPCFavorites{}
	resp := &protocol.CMsgClientToGCGetDPCFavoritesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetDPCFavorites),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetDPCFavoritesResponse),
		resp,
	)
}

// GetEventCoupon gets a event coupon.
//
// Sends the GC message k_EMsgClientToGCGetEventCoupon (CMsgClientToGCGetEventCoupon) and awaits the response k_EMsgClientToGCGetEventCouponResponse,
// delivered as *CMsgClientToGCGetEventCouponResponse.
func (d *Dota2) GetEventCoupon(
	ctx context.Context,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCGetEventCouponResponse, error) {
	req := &protocol.CMsgClientToGCGetEventCoupon{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCGetEventCouponResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetEventCoupon),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetEventCouponResponse),
		resp,
	)
}

// GetEventPoints gets event points.
//
// Sends the GC message k_EMsgDOTAGetEventPoints (CMsgDOTAGetEventPoints) and awaits the response k_EMsgDOTAGetEventPointsResponse,
// delivered as *CMsgDOTAGetEventPointsResponse.
func (d *Dota2) GetEventPoints(
	ctx context.Context,
	eventID uint32,
	accountID uint32,
) (*protocol.CMsgDOTAGetEventPointsResponse, error) {
	req := &protocol.CMsgDOTAGetEventPoints{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAGetEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetEventPointsResponse),
		resp,
	)
}

// GetEventRanking gets a event ranking.
//
// Sends the GC message k_EMsgClientToGCGetEventRanking (CMsgClientToGCGetEventRanking) and awaits the response k_EMsgClientToGCGetEventRankingResponse,
// delivered as *CMsgClientToGCGetEventRankingResponse.
func (d *Dota2) GetEventRanking(
	ctx context.Context,
	eventID protocol.EEvent,
	accountID uint32,
) (*protocol.CMsgClientToGCGetEventRankingResponse, error) {
	req := &protocol.CMsgClientToGCGetEventRanking{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetEventRankingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetEventRanking),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetEventRankingResponse),
		resp,
	)
}

// GetFavoritePlayers gets favorite players.
//
// Sends the GC message k_EMsgClientToGCGetFavoritePlayers (CMsgClientToGCGetFavoritePlayers) and awaits the response k_EMsgGCToClientGetFavoritePlayersResponse,
// delivered as *CMsgGCToClientGetFavoritePlayersResponse.
func (d *Dota2) GetFavoritePlayers(
	ctx context.Context,
	paginationKey uint64,
	paginationCount int32,
) (*protocol.CMsgGCToClientGetFavoritePlayersResponse, error) {
	req := &protocol.CMsgClientToGCGetFavoritePlayers{
		PaginationKey:   &paginationKey,
		PaginationCount: &paginationCount,
	}
	resp := &protocol.CMsgGCToClientGetFavoritePlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetFavoritePlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGetFavoritePlayersResponse),
		resp,
	)
}

// GetFilteredPlayers gets filtered players.
//
// Sends the GC message k_EMsgClientToGCGetFilteredPlayers (CMsgClientToGCGetFilteredPlayers) and awaits the response k_EMsgGCToClientGetFilteredPlayersResponse,
// delivered as *CMsgGCToClientGetFilteredPlayersResponse.
func (d *Dota2) GetFilteredPlayers(
	ctx context.Context,
) (*protocol.CMsgGCToClientGetFilteredPlayersResponse, error) {
	req := &protocol.CMsgClientToGCGetFilteredPlayers{}
	resp := &protocol.CMsgGCToClientGetFilteredPlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetFilteredPlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGetFilteredPlayersResponse),
		resp,
	)
}

// GetGiftPermissions gets gift permissions.
//
// Sends the GC message k_EMsgClientToGCGetGiftPermissions (CMsgClientToGCGetGiftPermissions) and awaits the response k_EMsgClientToGCGetGiftPermissionsResponse,
// delivered as *CMsgClientToGCGetGiftPermissionsResponse.
func (d *Dota2) GetGiftPermissions(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetGiftPermissionsResponse, error) {
	req := &protocol.CMsgClientToGCGetGiftPermissions{}
	resp := &protocol.CMsgClientToGCGetGiftPermissionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissionsResponse),
		resp,
	)
}

// GetHeroStandings gets hero standings.
//
// Sends the GC message k_EMsgGCGetHeroStandings (CMsgGCGetHeroStandings) and awaits the response k_EMsgGCGetHeroStandingsResponse,
// delivered as *CMsgGCGetHeroStandingsResponse.
func (d *Dota2) GetHeroStandings(
	ctx context.Context,
) (*protocol.CMsgGCGetHeroStandingsResponse, error) {
	req := &protocol.CMsgGCGetHeroStandings{}
	resp := &protocol.CMsgGCGetHeroStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStandings),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStandingsResponse),
		resp,
	)
}

// GetHeroStatsHistory gets a hero stats history.
//
// Sends the GC message k_EMsgGCGetHeroStatsHistory (CMsgGCGetHeroStatsHistory) and awaits the response k_EMsgGCGetHeroStatsHistoryResponse,
// delivered as *CMsgGCGetHeroStatsHistoryResponse.
func (d *Dota2) GetHeroStatsHistory(
	ctx context.Context,
	heroID int32,
) (*protocol.CMsgGCGetHeroStatsHistoryResponse, error) {
	req := &protocol.CMsgGCGetHeroStatsHistory{
		HeroId: &heroID,
	}
	resp := &protocol.CMsgGCGetHeroStatsHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistoryResponse),
		resp,
	)
}

// GetHeroStickers gets hero stickers.
//
// Sends the GC message k_EMsgClientToGCGetHeroStickers (CMsgClientToGCGetHeroStickers) and awaits the response k_EMsgClientToGCGetHeroStickersResponse,
// delivered as *CMsgClientToGCGetHeroStickersResponse.
func (d *Dota2) GetHeroStickers(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetHeroStickersResponse, error) {
	req := &protocol.CMsgClientToGCGetHeroStickers{}
	resp := &protocol.CMsgClientToGCGetHeroStickersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetHeroStickers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetHeroStickersResponse),
		resp,
	)
}

// GetHunterMonsterUserData gets a hunter monster user data.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterGetUserData (CMsgClientToGCMonsterHunterGetUserData) and awaits the response k_EMsgClientToGCMonsterHunterGetUserDataResponse,
// delivered as *CMsgClientToGCMonsterHunterGetUserDataResponse.
func (d *Dota2) GetHunterMonsterUserData(
	ctx context.Context,
) (*protocol.CMsgClientToGCMonsterHunterGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterGetUserData{}
	resp := &protocol.CMsgClientToGCMonsterHunterGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterGetUserDataResponse),
		resp,
	)
}

// GetModerationShowcaseQueue gets a moderation showcase queue.
//
// Sends the GC message k_EMsgClientToGCShowcaseModerationGetQueue (CMsgClientToGCShowcaseModerationGetQueue) and awaits the response k_EMsgClientToGCShowcaseModerationGetQueueResponse,
// delivered as *CMsgClientToGCShowcaseModerationGetQueueResponse.
func (d *Dota2) GetModerationShowcaseQueue(
	ctx context.Context,
	startTimestamp uint32,
	resultCount uint32,
) (*protocol.CMsgClientToGCShowcaseModerationGetQueueResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseModerationGetQueue{
		StartTimestamp: &startTimestamp,
		ResultCount:    &resultCount,
	}
	resp := &protocol.CMsgClientToGCShowcaseModerationGetQueueResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseModerationGetQueue),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseModerationGetQueueResponse),
		resp,
	)
}

// GetOWMatchDetails gets ow match details.
//
// Sends the GC message k_EMsgClientToGCGetOWMatchDetails (CMsgClientToGCGetOWMatchDetails) and awaits the response k_EMsgClientToGCGetOWMatchDetailsResponse,
// delivered as *CMsgClientToGCGetOWMatchDetailsResponse.
func (d *Dota2) GetOWMatchDetails(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetOWMatchDetailsResponse, error) {
	req := &protocol.CMsgClientToGCGetOWMatchDetails{}
	resp := &protocol.CMsgClientToGCGetOWMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetOWMatchDetails),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetOWMatchDetailsResponse),
		resp,
	)
}

// GetOverworldDynamicImage gets a overworld dynamic image.
//
// Sends the GC message k_EMsgClientToGCOverworldGetDynamicImage (CMsgClientToGCOverworldGetDynamicImage) and awaits the response k_EMsgClientToGCOverworldGetDynamicImageResponse,
// delivered as *CMsgClientToGCOverworldGetDynamicImageResponse.
func (d *Dota2) GetOverworldDynamicImage(
	ctx context.Context,
	magic uint32,
	imageID uint32,
	language uint32,
) (*protocol.CMsgClientToGCOverworldGetDynamicImageResponse, error) {
	req := &protocol.CMsgClientToGCOverworldGetDynamicImage{
		Magic:    &magic,
		ImageId:  &imageID,
		Language: &language,
	}
	resp := &protocol.CMsgClientToGCOverworldGetDynamicImageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGetDynamicImage),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGetDynamicImageResponse),
		resp,
	)
}

// GetOverworldUserData gets a overworld user data.
//
// Sends the GC message k_EMsgClientToGCOverworldGetUserData (CMsgClientToGCOverworldGetUserData) and awaits the response k_EMsgClientToGCOverworldGetUserDataResponse,
// delivered as *CMsgClientToGCOverworldGetUserDataResponse.
func (d *Dota2) GetOverworldUserData(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCOverworldGetUserData{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGetUserDataResponse),
		resp,
	)
}

// GetPeriodicResource gets a periodic resource.
//
// Sends the GC message k_EMsgDOTAGetPeriodicResource (CMsgDOTAGetPeriodicResource) and awaits the response k_EMsgDOTAGetPeriodicResourceResponse,
// delivered as *CMsgDOTAGetPeriodicResourceResponse.
func (d *Dota2) GetPeriodicResource(
	ctx context.Context,
	accountID uint32,
	periodicResourceID uint32,
	timestamp uint32,
) (*protocol.CMsgDOTAGetPeriodicResourceResponse, error) {
	req := &protocol.CMsgDOTAGetPeriodicResource{
		AccountId:          &accountID,
		PeriodicResourceId: &periodicResourceID,
		Timestamp:          &timestamp,
	}
	resp := &protocol.CMsgDOTAGetPeriodicResourceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResource),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResourceResponse),
		resp,
	)
}

// GetPlayerCardItemInfo gets a player card item info.
//
// Sends the GC message k_EMsgGCGetPlayerCardItemInfo (CMsgGCGetPlayerCardItemInfo) and awaits the response k_EMsgGCGetPlayerCardItemInfoResponse,
// delivered as *CMsgGCGetPlayerCardItemInfoResponse.
func (d *Dota2) GetPlayerCardItemInfo(
	ctx context.Context,
	accountID uint32,
	playerCardItemIDs []uint64,
	allForEvent uint32,
) (*protocol.CMsgGCGetPlayerCardItemInfoResponse, error) {
	req := &protocol.CMsgGCGetPlayerCardItemInfo{
		AccountId:         &accountID,
		PlayerCardItemIds: playerCardItemIDs,
		AllForEvent:       &allForEvent,
	}
	resp := &protocol.CMsgGCGetPlayerCardItemInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfoResponse),
		resp,
	)
}

// GetPlayerMatchHistory returns a page of an account's match history with
// per-match results. Pass start_at_match_id from the previous response to page
// through older matches.
func (d *Dota2) GetPlayerMatchHistory(
	ctx context.Context,
	req *protocol.CMsgDOTAGetPlayerMatchHistory,
) (*protocol.CMsgDOTAGetPlayerMatchHistoryResponse, error) {
	resp := &protocol.CMsgDOTAGetPlayerMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistoryResponse),
		resp,
	)
}

// GetProfileCard requests the profile card of an account.
//
// The card is not returned directly; watch for a ProfileCardUpdated event after
// sending this request.
func (d *Dota2) GetProfileCard(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgDOTAProfileCard, error) {
	req := &protocol.CMsgClientToGCGetProfileCard{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAProfileCard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse),
		resp,
	)
}

// GetProfileTickets gets profile tickets.
//
// Sends the GC message k_EMsgClientToGCGetProfileTickets (CMsgClientToGCGetProfileTickets) and awaits the response k_EMsgClientToGCGetProfileTicketsResponse,
// delivered as *CMsgDOTAProfileTickets.
func (d *Dota2) GetProfileTickets(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgDOTAProfileTickets, error) {
	req := &protocol.CMsgClientToGCGetProfileTickets{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgDOTAProfileTickets{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileTickets),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse),
		resp,
	)
}

// GetQuestProgress gets quest progress.
//
// Sends the GC message k_EMsgClientToGCGetQuestProgress (CMsgClientToGCGetQuestProgress) and awaits the response k_EMsgClientToGCGetQuestProgressResponse,
// delivered as *CMsgClientToGCGetQuestProgressResponse.
func (d *Dota2) GetQuestProgress(
	ctx context.Context,
	questIDs []uint32,
) (*protocol.CMsgClientToGCGetQuestProgressResponse, error) {
	req := &protocol.CMsgClientToGCGetQuestProgress{
		QuestIds: questIDs,
	}
	resp := &protocol.CMsgClientToGCGetQuestProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgress),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgressResponse),
		resp,
	)
}

// GetShopCandyUserData gets a shop candy user data.
//
// Sends the GC message k_EMsgClientToGCCandyShopGetUserData (CMsgClientToGCCandyShopGetUserData) and awaits the response k_EMsgClientToGCCandyShopGetUserDataResponse,
// delivered as *CMsgClientToGCCandyShopGetUserDataResponse.
func (d *Dota2) GetShopCandyUserData(
	ctx context.Context,
	candyShopID uint32,
) (*protocol.CMsgClientToGCCandyShopGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopGetUserData{
		CandyShopId: &candyShopID,
	}
	resp := &protocol.CMsgClientToGCCandyShopGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopGetUserDataResponse),
		resp,
	)
}

// GetShowcaseUserData gets a showcase user data.
//
// Sends the GC message k_EMsgClientToGCShowcaseGetUserData (CMsgClientToGCShowcaseGetUserData) and awaits the response k_EMsgClientToGCShowcaseGetUserDataResponse,
// delivered as *CMsgClientToGCShowcaseGetUserDataResponse.
func (d *Dota2) GetShowcaseUserData(
	ctx context.Context,
	accountID uint32,
	showcaseType protocol.EShowcaseType,
) (*protocol.CMsgClientToGCShowcaseGetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseGetUserData{
		AccountId:    &accountID,
		ShowcaseType: &showcaseType,
	}
	resp := &protocol.CMsgClientToGCShowcaseGetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseGetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseGetUserDataResponse),
		resp,
	)
}

// GetToTIRoadActiveQuest gets a to ti road active quest.
//
// Sends the GC message k_EMsgClientToGCRoadToTIGetActiveQuest (CMsgClientToGCRoadToTIGetActiveQuest) and awaits the response k_EMsgClientToGCRoadToTIGetActiveQuestResponse,
// delivered as *CMsgClientToGCRoadToTIGetActiveQuestResponse.
func (d *Dota2) GetToTIRoadActiveQuest(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCRoadToTIGetActiveQuestResponse, error) {
	req := &protocol.CMsgClientToGCRoadToTIGetActiveQuest{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRoadToTIGetActiveQuestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIGetActiveQuest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIGetActiveQuestResponse),
		resp,
	)
}

// GetToTIRoadQuests gets to ti road quests.
//
// Sends the GC message k_EMsgClientToGCRoadToTIGetQuests (CMsgClientToGCRoadToTIGetQuests) and awaits the response k_EMsgClientToGCRoadToTIGetQuestsResponse,
// delivered as *CMsgClientToGCRoadToTIGetQuestsResponse.
func (d *Dota2) GetToTIRoadQuests(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCRoadToTIGetQuestsResponse, error) {
	req := &protocol.CMsgClientToGCRoadToTIGetQuests{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRoadToTIGetQuestsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIGetQuests),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIGetQuestsResponse),
		resp,
	)
}

// GetTourneyWeekendPlayerStats gets tourney weekend player stats.
//
// Sends the GC message k_EMsgClientToGCWeekendTourneyGetPlayerStats (CMsgDOTAWeekendTourneyPlayerStatsRequest) and awaits the response k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse,
// delivered as *CMsgDOTAWeekendTourneyPlayerStats.
func (d *Dota2) GetTourneyWeekendPlayerStats(
	ctx context.Context,
	accountID uint32,
	seasonTrophyID uint32,
) (*protocol.CMsgDOTAWeekendTourneyPlayerStats, error) {
	req := &protocol.CMsgDOTAWeekendTourneyPlayerStatsRequest{
		AccountId:      &accountID,
		SeasonTrophyId: &seasonTrophyID,
	}
	resp := &protocol.CMsgDOTAWeekendTourneyPlayerStats{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse),
		resp,
	)
}

// GetWeekendTourneySchedule gets a weekend tourney schedule.
//
// Sends the GC message k_EMsgDOTAGetWeekendTourneySchedule (CMsgRequestWeekendTourneySchedule). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) GetWeekendTourneySchedule() {
	req := &protocol.CMsgRequestWeekendTourneySchedule{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule), req)
}

// GrantBattlerDevItemItem grants a battler dev item item.
//
// Sends the GC message k_EMsgClientToGCItemBattlerDevGrantItem (CMsgClientToGCItemBattlerDevGrantItem) and awaits the response k_EMsgClientToGCItemBattlerDevGrantItemResponse,
// delivered as *CMsgClientToGCItemBattlerDevGrantItemResponse.
func (d *Dota2) GrantBattlerDevItemItem(
	ctx context.Context,
	itemDefinitionID uint32,
) (*protocol.CMsgClientToGCItemBattlerDevGrantItemResponse, error) {
	req := &protocol.CMsgClientToGCItemBattlerDevGrantItem{
		ItemDefinitionId: &itemDefinitionID,
	}
	resp := &protocol.CMsgClientToGCItemBattlerDevGrantItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerDevGrantItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerDevGrantItemResponse),
		resp,
	)
}

// GrantDevEventAction grants a dev event action.
//
// Sends the GC message k_EMsgDevGrantEventAction (CMsgDevGrantEventAction) and awaits the response k_EMsgDevGrantEventActionResponse,
// delivered as *CMsgDevGrantEventActionResponse.
func (d *Dota2) GrantDevEventAction(
	ctx context.Context,
	eventID protocol.EEvent,
	actionID uint32,
	actionScore uint32,
) (*protocol.CMsgDevGrantEventActionResponse, error) {
	req := &protocol.CMsgDevGrantEventAction{
		EventId:     &eventID,
		ActionId:    &actionID,
		ActionScore: &actionScore,
	}
	resp := &protocol.CMsgDevGrantEventActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventActionResponse),
		resp,
	)
}

// GrantDevEventPoints grants dev event points.
//
// Sends the GC message k_EMsgDevGrantEventPoints (CMsgDevGrantEventPoints) and awaits the response k_EMsgDevGrantEventPointsResponse,
// delivered as *CMsgDevGrantEventPointsResponse.
func (d *Dota2) GrantDevEventPoints(
	ctx context.Context,
	eventID protocol.EEvent,
	eventPoints uint32,
	premiumPoints uint32,
) (*protocol.CMsgDevGrantEventPointsResponse, error) {
	req := &protocol.CMsgDevGrantEventPoints{
		EventId:       &eventID,
		EventPoints:   &eventPoints,
		PremiumPoints: &premiumPoints,
	}
	resp := &protocol.CMsgDevGrantEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevGrantEventPointsResponse),
		resp,
	)
}

// GrantDevOverworldFortuneTellerCoin grants a dev overworld fortune teller coin.
//
// Sends the GC message k_EMsgClientToGCOverworldDevGrantFortuneTellerCoin (CMsgClientToGCOverworldDevGrantFortuneTellerCoin) and awaits the response k_EMsgClientToGCOverworldDevGrantFortuneTellerCoinResponse,
// delivered as *CMsgClientToGCOverworldDevGrantFortuneTellerCoinResponse.
func (d *Dota2) GrantDevOverworldFortuneTellerCoin(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldDevGrantFortuneTellerCoinResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevGrantFortuneTellerCoin{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevGrantFortuneTellerCoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevGrantFortuneTellerCoin),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevGrantFortuneTellerCoinResponse),
		resp,
	)
}

// GrantDevOverworldTokens grants dev overworld tokens.
//
// Sends the GC message k_EMsgClientToGCOverworldDevGrantTokens (CMsgClientToGCOverworldDevGrantTokens) and awaits the response k_EMsgClientToGCOverworldDevGrantTokensResponse,
// delivered as *CMsgClientToGCOverworldDevGrantTokensResponse.
func (d *Dota2) GrantDevOverworldTokens(
	ctx context.Context,
	overworldID uint32,
	tokenQuantity protocol.CMsgOverworldTokenQuantity,
) (*protocol.CMsgClientToGCOverworldDevGrantTokensResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevGrantTokens{
		OverworldId:   &overworldID,
		TokenQuantity: &tokenQuantity,
	}
	resp := &protocol.CMsgClientToGCOverworldDevGrantTokensResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevGrantTokens),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevGrantTokensResponse),
		resp,
	)
}

// GrantEventSupportConsumeItem grants a event support consume item.
//
// Sends the GC message k_EMsgConsumeEventSupportGrantItem (CMsgConsumeEventSupportGrantItem) and awaits the response k_EMsgConsumeEventSupportGrantItemResponse,
// delivered as *CMsgConsumeEventSupportGrantItemResponse.
func (d *Dota2) GrantEventSupportConsumeItem(
	ctx context.Context,
	itemID uint64,
) (*protocol.CMsgConsumeEventSupportGrantItemResponse, error) {
	req := &protocol.CMsgConsumeEventSupportGrantItem{
		ItemId: &itemID,
	}
	resp := &protocol.CMsgConsumeEventSupportGrantItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgConsumeEventSupportGrantItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgConsumeEventSupportGrantItemResponse),
		resp,
	)
}

// GrantHunterDevMonsterMaterials grants hunter dev monster materials.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterDevGrantMaterials (CMsgClientToGCMonsterHunterDevGrantMaterials) and awaits the response k_EMsgClientToGCMonsterHunterDevGrantMaterialsResponse,
// delivered as *CMsgClientToGCMonsterHunterDevGrantMaterialsResponse.
func (d *Dota2) GrantHunterDevMonsterMaterials(
	ctx context.Context,
	materialQuantity protocol.CMsgMonsterHunterMaterialQuantity,
) (*protocol.CMsgClientToGCMonsterHunterDevGrantMaterialsResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterDevGrantMaterials{
		MaterialQuantity: &materialQuantity,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterDevGrantMaterialsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevGrantMaterials),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevGrantMaterialsResponse),
		resp,
	)
}

// GrantShopDevCandyCandy grants a shop dev candy candy.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevGrantCandy (CMsgClientToGCCandyShopDevGrantCandy) and awaits the response k_EMsgClientToGCCandyShopDevGrantCandyResponse,
// delivered as *CMsgClientToGCCandyShopDevGrantCandyResponse.
func (d *Dota2) GrantShopDevCandyCandy(
	ctx context.Context,
	candyShopID uint32,
	candyQuantity protocol.CMsgCandyShopCandyQuantity,
) (*protocol.CMsgClientToGCCandyShopDevGrantCandyResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevGrantCandy{
		CandyShopId:   &candyShopID,
		CandyQuantity: &candyQuantity,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevGrantCandyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantCandy),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantCandyResponse),
		resp,
	)
}

// GrantShopDevCandyCandyBags grants shop dev candy candy bags.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevGrantCandyBags (CMsgClientToGCCandyShopDevGrantCandyBags) and awaits the response k_EMsgClientToGCCandyShopDevGrantCandyBagsResponse,
// delivered as *CMsgClientToGCCandyShopDevGrantCandyBagsResponse.
func (d *Dota2) GrantShopDevCandyCandyBags(
	ctx context.Context,
	candyShopID uint32,
	quantity uint32,
) (*protocol.CMsgClientToGCCandyShopDevGrantCandyBagsResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevGrantCandyBags{
		CandyShopId: &candyShopID,
		Quantity:    &quantity,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevGrantCandyBagsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantCandyBags),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantCandyBagsResponse),
		resp,
	)
}

// GrantShopDevCandyRerollCharges grants shop dev candy reroll charges.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevGrantRerollCharges (CMsgClientToGCCandyShopDevGrantRerollCharges) and awaits the response k_EMsgClientToGCCandyShopDevGrantRerollChargesResponse,
// delivered as *CMsgClientToGCCandyShopDevGrantRerollChargesResponse.
func (d *Dota2) GrantShopDevCandyRerollCharges(
	ctx context.Context,
	candyShopID uint32,
	rerollCharges uint32,
) (*protocol.CMsgClientToGCCandyShopDevGrantRerollChargesResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevGrantRerollCharges{
		CandyShopId:   &candyShopID,
		RerollCharges: &rerollCharges,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevGrantRerollChargesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantRerollCharges),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevGrantRerollChargesResponse),
		resp,
	)
}

// InvitePlayerToTeam is undocumented.
//
// Sends the GC message k_EMsgGCTeamInvite_InviterToGC (CMsgDOTATeamInvite_InviterToGC) and awaits the response k_EMsgGCTeamInvite_GCImmediateResponseToInviter,
// delivered as *CMsgDOTATeamInvite_GCImmediateResponseToInviter.
func (d *Dota2) InvitePlayerToTeam(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTATeamInvite_GCImmediateResponseToInviter, error) {
	req := &protocol.CMsgDOTATeamInvite_InviterToGC{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &protocol.CMsgDOTATeamInvite_GCImmediateResponseToInviter{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter),
		resp,
	)
}

// InvitePrivateChatMember is undocumented.
//
// Sends the GC message k_EMsgClientToGCPrivateChatInvite (CMsgClientToGCPrivateChatInvite). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) InvitePrivateChatMember(
	privateChatChannelName string,
	invitedAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatInvite{
		PrivateChatChannelName: &privateChatChannelName,
		InvitedAccountId:       &invitedAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite), req)
}

// JoinChatChannel joins a chat channel by name and type. The response lists
// the channel members; other joins and leaves arrive as events.
func (d *Dota2) JoinChatChannel(
	ctx context.Context,
	channelName string,
	channelType protocol.DOTAChatChannelTypeT,
	silentRejection bool,
) (*protocol.CMsgDOTAJoinChatChannelResponse, error) {
	req := &protocol.CMsgDOTAJoinChatChannel{
		ChannelName:     &channelName,
		ChannelType:     &channelType,
		SilentRejection: &silentRejection,
	}
	resp := &protocol.CMsgDOTAJoinChatChannelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinChatChannel),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse),
		resp,
	)
}

// JoinGuild joins a guild.
//
// Sends the GC message k_EMsgClientToGCJoinGuild (CMsgClientToGCJoinGuild) and awaits the response k_EMsgClientToGCJoinGuildResponse,
// delivered as *CMsgClientToGCJoinGuildResponse.
func (d *Dota2) JoinGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCJoinGuildResponse, error) {
	req := &protocol.CMsgClientToGCJoinGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCJoinGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinGuildResponse),
		resp,
	)
}

// JoinLobby joins a practice lobby by ID, optionally with a pass key.
func (d *Dota2) JoinLobby(
	ctx context.Context,
	lobbyID uint64,
	passKey string,
	customGameCrc uint64,
	customGameTimestamp uint32,
) (*protocol.CMsgPracticeLobbyJoinResponse, error) {
	req := &protocol.CMsgPracticeLobbyJoin{
		LobbyId:             &lobbyID,
		PassKey:             &passKey,
		CustomGameCrc:       &customGameCrc,
		CustomGameTimestamp: &customGameTimestamp,
	}
	resp := &protocol.CMsgPracticeLobbyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoin),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinResponse),
		resp,
	)
}

// JoinLobbyBroadcastChannel joins a lobby broadcast channel.
//
// Sends the GC message k_EMsgGCPracticeLobbyJoinBroadcastChannel (CMsgPracticeLobbyJoinBroadcastChannel). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) JoinLobbyBroadcastChannel(
	channel uint32,
	preferredDescription string,
	preferredCountryCode string,
	preferredLanguageCode string,
) {
	req := &protocol.CMsgPracticeLobbyJoinBroadcastChannel{
		Channel:               &channel,
		PreferredDescription:  &preferredDescription,
		PreferredCountryCode:  &preferredCountryCode,
		PreferredLanguageCode: &preferredLanguageCode,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinBroadcastChannel), req)
}

// JoinPartyFromBeacon joins a party from beacon.
//
// Sends the GC message k_EMsgClientToGCJoinPartyFromBeacon (CMsgClientToGCJoinPartyFromBeacon) and awaits the response k_EMsgGCToClientJoinPartyFromBeaconResponse,
// delivered as *CMsgGCToClientJoinPartyFromBeaconResponse.
func (d *Dota2) JoinPartyFromBeacon(
	ctx context.Context,
	partyID uint64,
	accountID uint32,
	beaconType int32,
) (*protocol.CMsgGCToClientJoinPartyFromBeaconResponse, error) {
	req := &protocol.CMsgClientToGCJoinPartyFromBeacon{
		PartyId:    &partyID,
		AccountId:  &accountID,
		BeaconType: &beaconType,
	}
	resp := &protocol.CMsgGCToClientJoinPartyFromBeaconResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPartyFromBeacon),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientJoinPartyFromBeaconResponse),
		resp,
	)
}

// JoinPlaytest joins a playtest.
//
// Sends the GC message k_EMsgClientToGCJoinPlaytest (CMsgClientToGCJoinPlaytest) and awaits the response k_EMsgClientToGCJoinPlaytestResponse,
// delivered as *CMsgClientToGCJoinPlaytestResponse.
func (d *Dota2) JoinPlaytest(
	ctx context.Context,
) (*protocol.CMsgClientToGCJoinPlaytestResponse, error) {
	req := &protocol.CMsgClientToGCJoinPlaytest{}
	resp := &protocol.CMsgClientToGCJoinPlaytestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytestResponse),
		resp,
	)
}

// JoinPrivateCoachingSessionLobby joins a private coaching session lobby.
//
// Sends the GC message k_EMsgClientToGCJoinPrivateCoachingSessionLobby (CMsgClientToGCJoinPrivateCoachingSessionLobby) and awaits the response k_EMsgClientToGCJoinPrivateCoachingSessionLobbyResponse,
// delivered as *CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse.
func (d *Dota2) JoinPrivateCoachingSessionLobby(
	ctx context.Context,
) (*protocol.CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse, error) {
	req := &protocol.CMsgClientToGCJoinPrivateCoachingSessionLobby{}
	resp := &protocol.CMsgClientToGCJoinPrivateCoachingSessionLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPrivateCoachingSessionLobby),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCJoinPrivateCoachingSessionLobbyResponse),
		resp,
	)
}

// JoinQuickCustomLobby joins a quick custom lobby.
//
// Sends the GC message k_EMsgGCQuickJoinCustomLobby (CMsgQuickJoinCustomLobby) and awaits the response k_EMsgGCQuickJoinCustomLobbyResponse,
// delivered as *CMsgQuickJoinCustomLobbyResponse.
func (d *Dota2) JoinQuickCustomLobby(
	ctx context.Context,
	legacyServerRegion uint32,
	customGameID uint64,
	createLobbyDetails protocol.CMsgPracticeLobbySetDetails,
	allowAnyMap bool,
	legacyRegionPings []*protocol.CMsgQuickJoinCustomLobby_LegacyRegionPing,
	pingData protocol.CMsgClientPingData,
) (*protocol.CMsgQuickJoinCustomLobbyResponse, error) {
	req := &protocol.CMsgQuickJoinCustomLobby{
		LegacyServerRegion: &legacyServerRegion,
		CustomGameId:       &customGameID,
		CreateLobbyDetails: &createLobbyDetails,
		AllowAnyMap:        &allowAnyMap,
		LegacyRegionPings:  legacyRegionPings,
		PingData:           &pingData,
	}
	resp := &protocol.CMsgQuickJoinCustomLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobby),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobbyResponse),
		resp,
	)
}

// KickGuildMember kicks a guild member.
//
// Sends the GC message k_EMsgClientToGCKickGuildMember (CMsgClientToGCKickGuildMember) and awaits the response k_EMsgClientToGCKickGuildMemberResponse,
// delivered as *CMsgClientToGCKickGuildMemberResponse.
func (d *Dota2) KickGuildMember(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCKickGuildMemberResponse, error) {
	req := &protocol.CMsgClientToGCKickGuildMember{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCKickGuildMemberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCKickGuildMember),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCKickGuildMemberResponse),
		resp,
	)
}

// KickLobbyMember kicks a member from your practice lobby. Lobby leaders
// only.
func (d *Dota2) KickLobbyMember(
	accountID uint32,
) {
	req := &protocol.CMsgPracticeLobbyKick{
		AccountId: &accountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick), req)
}

// KickLobbyMemberFromTeam kicks a lobby member from team.
//
// Sends the GC message k_EMsgGCPracticeLobbyKickFromTeam (CMsgPracticeLobbyKickFromTeam). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) KickLobbyMemberFromTeam(
	accountID uint32,
) {
	req := &protocol.CMsgPracticeLobbyKickFromTeam{
		AccountId: &accountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam), req)
}

// KickPrivateChatMember kicks a private chat member.
//
// Sends the GC message k_EMsgClientToGCPrivateChatKick (CMsgClientToGCPrivateChatKick). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) KickPrivateChatMember(
	privateChatChannelName string,
	kickAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatKick{
		PrivateChatChannelName: &privateChatChannelName,
		KickAccountId:          &kickAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick), req)
}

// KickTeamMember kicks a team member.
//
// Sends the GC message k_EMsgGCKickTeamMember (CMsgDOTAKickTeamMember) and awaits the response k_EMsgGCKickTeamMemberResponse,
// delivered as *CMsgDOTAKickTeamMemberResponse.
func (d *Dota2) KickTeamMember(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTAKickTeamMemberResponse, error) {
	req := &protocol.CMsgDOTAKickTeamMember{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &protocol.CMsgDOTAKickTeamMemberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCKickTeamMember),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCKickTeamMemberResponse),
		resp,
	)
}

// LaunchLobby starts the match for the practice lobby you lead.
func (d *Dota2) LaunchLobby() {
	req := &protocol.CMsgPracticeLobbyLaunch{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyLaunch), req)
}

// LeaveChatChannel leaves a joined chat channel.
func (d *Dota2) LeaveChatChannel(
	channelID uint64,
) {
	req := &protocol.CMsgDOTALeaveChatChannel{
		ChannelId: &channelID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveChatChannel), req)
}

// LeaveGuild leaves a guild.
//
// Sends the GC message k_EMsgClientToGCLeaveGuild (CMsgClientToGCLeaveGuild) and awaits the response k_EMsgClientToGCLeaveGuildResponse,
// delivered as *CMsgClientToGCLeaveGuildResponse.
func (d *Dota2) LeaveGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCLeaveGuildResponse, error) {
	req := &protocol.CMsgClientToGCLeaveGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCLeaveGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeaveGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeaveGuildResponse),
		resp,
	)
}

// LeaveLobby leaves the current practice lobby without destroying it.
func (d *Dota2) LeaveLobby() {
	req := &protocol.CMsgPracticeLobbyLeave{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyLeave), req)
}

// LeavePrivateCoachingSession leaves a private coaching session.
//
// Sends the GC message k_EMsgClientToGCLeavePrivateCoachingSession (CMsgClientToGCLeavePrivateCoachingSession) and awaits the response k_EMsgClientToGCLeavePrivateCoachingSessionResponse,
// delivered as *CMsgClientToGCLeavePrivateCoachingSessionResponse.
func (d *Dota2) LeavePrivateCoachingSession(
	ctx context.Context,
) (*protocol.CMsgClientToGCLeavePrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCLeavePrivateCoachingSession{}
	resp := &protocol.CMsgClientToGCLeavePrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeavePrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLeavePrivateCoachingSessionResponse),
		resp,
	)
}

// LeaveTeam leaves a team.
//
// Sends the GC message k_EMsgGCLeaveTeam (CMsgDOTALeaveTeam) and awaits the response k_EMsgGCLeaveTeamResponse,
// delivered as *CMsgDOTALeaveTeamResponse.
func (d *Dota2) LeaveTeam(
	ctx context.Context,
	teamID uint32,
) (*protocol.CMsgDOTALeaveTeamResponse, error) {
	req := &protocol.CMsgDOTALeaveTeam{
		TeamId: &teamID,
	}
	resp := &protocol.CMsgDOTALeaveTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveTeam),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCLeaveTeamResponse),
		resp,
	)
}

// LeaveTourneyWeekend leaves a tourney weekend.
//
// Sends the GC message k_EMsgClientToGCWeekendTourneyLeave (CMsgWeekendTourneyLeave). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) LeaveTourneyWeekend() {
	req := &protocol.CMsgWeekendTourneyLeave{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave), req)
}

// ListChatChannel lists a chat channel.
//
// Sends the GC message k_EMsgGCRequestChatChannelList (CMsgDOTARequestChatChannelList) and awaits the response k_EMsgGCRequestChatChannelListResponse,
// delivered as *CMsgDOTARequestChatChannelListResponse.
func (d *Dota2) ListChatChannel(
	ctx context.Context,
) (*protocol.CMsgDOTARequestChatChannelListResponse, error) {
	req := &protocol.CMsgDOTARequestChatChannelList{}
	resp := &protocol.CMsgDOTARequestChatChannelListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRequestChatChannelList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRequestChatChannelListResponse),
		resp,
	)
}

// ListCustomGamesTop lists a custom games top.
//
// Sends the GC message k_EMsgGCTopCustomGamesList (CMsgGCTopCustomGamesList). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ListCustomGamesTop(
	topCustomGames []uint64,
	gameOfTheDay uint64,
) {
	req := &protocol.CMsgGCTopCustomGamesList{
		TopCustomGames: topCustomGames,
		GameOfTheDay:   &gameOfTheDay,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTopCustomGamesList), req)
}

// ListFriendLobby lists a friend lobby.
//
// Sends the GC message k_EMsgGCFriendPracticeLobbyListRequest (CMsgFriendPracticeLobbyListRequest) and awaits the response k_EMsgGCFriendPracticeLobbyListResponse,
// delivered as *CMsgFriendPracticeLobbyListResponse.
func (d *Dota2) ListFriendLobby(
	ctx context.Context,
	friends []uint32,
) (*protocol.CMsgFriendPracticeLobbyListResponse, error) {
	req := &protocol.CMsgFriendPracticeLobbyListRequest{
		Friends: friends,
	}
	resp := &protocol.CMsgFriendPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListResponse),
		resp,
	)
}

// ListLobbies lists lobbies.
//
// Sends the GC message k_EMsgGCPracticeLobbyList (CMsgPracticeLobbyList) and awaits the response k_EMsgGCPracticeLobbyListResponse,
// delivered as *CMsgPracticeLobbyListResponse.
func (d *Dota2) ListLobbies(
	ctx context.Context,
	passKey string,
	region uint32,
	gameMode protocol.DOTA_GameMode,
) (*protocol.CMsgPracticeLobbyListResponse, error) {
	req := &protocol.CMsgPracticeLobbyList{
		PassKey:  &passKey,
		Region:   &region,
		GameMode: &gameMode,
	}
	resp := &protocol.CMsgPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyListResponse),
		resp,
	)
}

// ListLobbySpectator lists a lobby spectator.
//
// Sends the GC message k_EMsgClientToGCSpectatorLobbyList (CMsgSpectatorLobbyList) and awaits the response k_EMsgClientToGCSpectatorLobbyListResponse,
// delivered as *CMsgSpectatorLobbyListResponse.
func (d *Dota2) ListLobbySpectator(
	ctx context.Context,
) (*protocol.CMsgSpectatorLobbyListResponse, error) {
	req := &protocol.CMsgSpectatorLobbyList{}
	resp := &protocol.CMsgSpectatorLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyListResponse),
		resp,
	)
}

// ListOverworldFortuneLobby lists a overworld fortune lobby.
//
// Sends the GC message k_EMsgLobbyOverworldFortuneList (CMsgLobbyOverworldFortuneList). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ListOverworldFortuneLobby(
	accountID []uint32,
	fortune []*protocol.CMsgOverworldFortune,
) {
	req := &protocol.CMsgLobbyOverworldFortuneList{
		AccountId: accountID,
		Fortune:   fortune,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyOverworldFortuneList), req)
}

// ListTrophies lists trophies.
//
// Sends the GC message k_EMsgClientToGCGetTrophyList (CMsgClientToGCGetTrophyList) and awaits the response k_EMsgClientToGCGetTrophyListResponse,
// delivered as *CMsgClientToGCGetTrophyListResponse.
func (d *Dota2) ListTrophies(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetTrophyListResponse, error) {
	req := &protocol.CMsgClientToGCGetTrophyList{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetTrophyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetTrophyListResponse),
		resp,
	)
}

// OpenPlayerCardPack opens a player card pack.
//
// Sends the GC message k_EMsgClientToGCOpenPlayerCardPack (CMsgClientToGCOpenPlayerCardPack) and awaits the response k_EMsgClientToGCOpenPlayerCardPackResponse,
// delivered as *CMsgClientToGCOpenPlayerCardPackResponse.
func (d *Dota2) OpenPlayerCardPack(
	ctx context.Context,
	playerCardPackItemID uint64,
	teamID uint32,
	deprecatedLeagueID uint32,
	region protocol.ELeagueRegion,
) (*protocol.CMsgClientToGCOpenPlayerCardPackResponse, error) {
	req := &protocol.CMsgClientToGCOpenPlayerCardPack{
		PlayerCardPackItemId: &playerCardPackItemID,
		TeamId:               &teamID,
		DeprecatedLeagueId:   &deprecatedLeagueID,
		Region:               &region,
	}
	resp := &protocol.CMsgClientToGCOpenPlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPack),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPackResponse),
		resp,
	)
}

// OpenShopCandyBags opens shop candy bags.
//
// Sends the GC message k_EMsgClientToGCCandyShopOpenBags (CMsgClientToGCCandyShopOpenBags) and awaits the response k_EMsgClientToGCCandyShopOpenBagsResponse,
// delivered as *CMsgClientToGCCandyShopOpenBagsResponse.
func (d *Dota2) OpenShopCandyBags(
	ctx context.Context,
	candyShopID uint32,
	bagCount uint32,
) (*protocol.CMsgClientToGCCandyShopOpenBagsResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopOpenBags{
		CandyShopId: &candyShopID,
		BagCount:    &bagCount,
	}
	resp := &protocol.CMsgClientToGCCandyShopOpenBagsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopOpenBags),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopOpenBagsResponse),
		resp,
	)
}

// PromotePrivateChatMember promotes a private chat member.
//
// Sends the GC message k_EMsgClientToGCPrivateChatPromote (CMsgClientToGCPrivateChatPromote). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) PromotePrivateChatMember(
	privateChatChannelName string,
	promoteAccountID uint32,
) {
	req := &protocol.CMsgClientToGCPrivateChatPromote{
		PrivateChatChannelName: &privateChatChannelName,
		PromoteAccountId:       &promoteAccountID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote), req)
}

// PublishUserStat publishs a user stat.
//
// Sends the GC message k_EMsgClientToGCPublishUserStat (CMsgClientToGCPublishUserStat). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) PublishUserStat(
	userStatsEvent uint32,
	referenceData uint64,
) {
	req := &protocol.CMsgClientToGCPublishUserStat{
		UserStatsEvent: &userStatsEvent,
		ReferenceData:  &referenceData,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPublishUserStat), req)
}

// PurchaseFilteredPlayerSlot purchases a filtered player slot.
//
// Sends the GC message k_EMsgClientToGCPurchaseFilteredPlayerSlot (CMsgClientToGCPurchaseFilteredPlayerSlot) and awaits the response k_EMsgGCToClientPurchaseFilteredPlayerSlotResponse,
// delivered as *CMsgGCToClientPurchaseFilteredPlayerSlotResponse.
func (d *Dota2) PurchaseFilteredPlayerSlot(
	ctx context.Context,
	additionalSlotsCurrent int32,
) (*protocol.CMsgGCToClientPurchaseFilteredPlayerSlotResponse, error) {
	req := &protocol.CMsgClientToGCPurchaseFilteredPlayerSlot{
		AdditionalSlotsCurrent: &additionalSlotsCurrent,
	}
	resp := &protocol.CMsgGCToClientPurchaseFilteredPlayerSlotResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPurchaseFilteredPlayerSlot),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPurchaseFilteredPlayerSlotResponse),
		resp,
	)
}

// PurchaseHeroRandomRelic purchases a hero random relic.
//
// Sends the GC message k_EMsgPurchaseHeroRandomRelic (CMsgPurchaseHeroRandomRelic) and awaits the response k_EMsgPurchaseHeroRandomRelicResponse,
// delivered as *CMsgPurchaseHeroRandomRelicResponse.
func (d *Dota2) PurchaseHeroRandomRelic(
	ctx context.Context,
	heroID int32,
	relicRarity protocol.EHeroRelicRarity,
) (*protocol.CMsgPurchaseHeroRandomRelicResponse, error) {
	req := &protocol.CMsgPurchaseHeroRandomRelic{
		HeroId:      &heroID,
		RelicRarity: &relicRarity,
	}
	resp := &protocol.CMsgPurchaseHeroRandomRelicResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseHeroRandomRelic),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseHeroRandomRelicResponse),
		resp,
	)
}

// PurchaseItemWithEventPoints purchases item with event points.
//
// Sends the GC message k_EMsgPurchaseItemWithEventPoints (CMsgPurchaseItemWithEventPoints) and awaits the response k_EMsgPurchaseItemWithEventPointsResponse,
// delivered as *CMsgPurchaseItemWithEventPointsResponse.
func (d *Dota2) PurchaseItemWithEventPoints(
	ctx context.Context,
	itemDef uint32,
	quantity uint32,
	eventID protocol.EEvent,
	usePremiumPoints bool,
) (*protocol.CMsgPurchaseItemWithEventPointsResponse, error) {
	req := &protocol.CMsgPurchaseItemWithEventPoints{
		ItemDef:          &itemDef,
		Quantity:         &quantity,
		EventId:          &eventID,
		UsePremiumPoints: &usePremiumPoints,
	}
	resp := &protocol.CMsgPurchaseItemWithEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPointsResponse),
		resp,
	)
}

// PurchaseLabyrinthBlessings purchases labyrinth blessings.
//
// Sends the GC message k_EMsgClientToGCPurchaseLabyrinthBlessings (CMsgClientToGCPurchaseLabyrinthBlessings) and awaits the response k_EMsgClientToGCPurchaseLabyrinthBlessingsResponse,
// delivered as *CMsgClientToGCPurchaseLabyrinthBlessingsResponse.
func (d *Dota2) PurchaseLabyrinthBlessings(
	ctx context.Context,
	eventID protocol.EEvent,
	blessingIDs []int32,
	debug bool,
	debugRemove bool,
) (*protocol.CMsgClientToGCPurchaseLabyrinthBlessingsResponse, error) {
	req := &protocol.CMsgClientToGCPurchaseLabyrinthBlessings{
		EventId:     &eventID,
		BlessingIds: blessingIDs,
		Debug:       &debug,
		DebugRemove: &debugRemove,
	}
	resp := &protocol.CMsgClientToGCPurchaseLabyrinthBlessingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPurchaseLabyrinthBlessings),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPurchaseLabyrinthBlessingsResponse),
		resp,
	)
}

// PurchasePlayerCardSpecific purchases a player card specific.
//
// Sends the GC message k_EMsgClientToGCPlayerCardSpecificPurchaseRequest (CMsgClientToGCPlayerCardSpecificPurchaseRequest) and awaits the response k_EMsgClientToGCPlayerCardSpecificPurchaseResponse,
// delivered as *CMsgClientToGCPlayerCardSpecificPurchaseResponse.
func (d *Dota2) PurchasePlayerCardSpecific(
	ctx context.Context,
	playerAccountID uint32,
	eventID uint32,
	cardDustItemID uint64,
) (*protocol.CMsgClientToGCPlayerCardSpecificPurchaseResponse, error) {
	req := &protocol.CMsgClientToGCPlayerCardSpecificPurchaseRequest{
		PlayerAccountId: &playerAccountID,
		EventId:         &eventID,
		CardDustItemId:  &cardDustItemID,
	}
	resp := &protocol.CMsgClientToGCPlayerCardSpecificPurchaseResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerCardSpecificPurchaseRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerCardSpecificPurchaseResponse),
		resp,
	)
}

// PurchaseShopCandyReward purchases a shop candy reward.
//
// Sends the GC message k_EMsgClientToGCCandyShopPurchaseReward (CMsgClientToGCCandyShopPurchaseReward) and awaits the response k_EMsgClientToGCCandyShopPurchaseRewardResponse,
// delivered as *CMsgClientToGCCandyShopPurchaseRewardResponse.
func (d *Dota2) PurchaseShopCandyReward(
	ctx context.Context,
	candyShopID uint32,
	rewardID uint64,
) (*protocol.CMsgClientToGCCandyShopPurchaseRewardResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopPurchaseReward{
		CandyShopId: &candyShopID,
		RewardId:    &rewardID,
	}
	resp := &protocol.CMsgClientToGCCandyShopPurchaseRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopPurchaseReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopPurchaseRewardResponse),
		resp,
	)
}

// QueryHasItem queries to check if the target has item.
//
// Sends the GC message k_EMsgGCHasItemQuery (CMsgDOTAHasItemQuery). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) QueryHasItem(
	accountID uint32,
	itemID uint64,
) {
	req := &protocol.CMsgDOTAHasItemQuery{
		AccountId: &accountID,
		ItemId:    &itemID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCHasItemQuery), req)
}

// RecordContestVote records a contest vote.
//
// Sends the GC message k_EMsgClientToGCRecordContestVote (CMsgClientToGCRecordContestVote) and awaits the response k_EMsgGCToClientRecordContestVoteResponse,
// delivered as *CMsgGCToClientRecordContestVoteResponse.
func (d *Dota2) RecordContestVote(
	ctx context.Context,
	contestID uint32,
	contestItemID uint64,
	vote int32,
) (*protocol.CMsgGCToClientRecordContestVoteResponse, error) {
	req := &protocol.CMsgClientToGCRecordContestVote{
		ContestId:     &contestID,
		ContestItemId: &contestItemID,
		Vote:          &vote,
	}
	resp := &protocol.CMsgGCToClientRecordContestVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecordContestVote),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRecordContestVoteResponse),
		resp,
	)
}

// RecyclePlayerCard recycles a player card.
//
// Sends the GC message k_EMsgClientToGCRecyclePlayerCard (CMsgClientToGCRecyclePlayerCard) and awaits the response k_EMsgClientToGCRecyclePlayerCardResponse,
// delivered as *CMsgClientToGCRecyclePlayerCardResponse.
func (d *Dota2) RecyclePlayerCard(
	ctx context.Context,
	playerCardItemIDs []uint64,
	eventID uint32,
) (*protocol.CMsgClientToGCRecyclePlayerCardResponse, error) {
	req := &protocol.CMsgClientToGCRecyclePlayerCard{
		PlayerCardItemIds: playerCardItemIDs,
		EventId:           &eventID,
	}
	resp := &protocol.CMsgClientToGCRecyclePlayerCardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCardResponse),
		resp,
	)
}

// RedeemDraftUnderReward redeems a draft under reward.
//
// Sends the GC message k_EMsgClientToGCUnderDraftRedeemReward (CMsgClientToGCUnderDraftRedeemReward) and awaits the response k_EMsgClientToGCUnderDraftRedeemRewardResponse,
// delivered as *CMsgClientToGCUnderDraftRedeemRewardResponse.
func (d *Dota2) RedeemDraftUnderReward(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
) (*protocol.CMsgClientToGCUnderDraftRedeemRewardResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRedeemReward{
		EventId:  &eventID,
		ActionId: &actionID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRedeemRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRedeemReward),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRedeemRewardResponse),
		resp,
	)
}

// RedeemItem redeems a item.
//
// Sends the GC message k_EMsgDOTARedeemItem (CMsgDOTARedeemItem) and awaits the response k_EMsgDOTARedeemItemResponse,
// delivered as *CMsgDOTARedeemItemResponse.
func (d *Dota2) RedeemItem(
	ctx context.Context,
	currencyID uint64,
	purchaseDef uint32,
	claimAsPoints bool,
) (*protocol.CMsgDOTARedeemItemResponse, error) {
	req := &protocol.CMsgDOTARedeemItem{
		CurrencyId:    &currencyID,
		PurchaseDef:   &purchaseDef,
		ClaimAsPoints: &claimAsPoints,
	}
	resp := &protocol.CMsgDOTARedeemItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTARedeemItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTARedeemItemResponse),
		resp,
	)
}

// RejoinAllChatChannels is undocumented.
//
// Sends the GC message k_EMsgClientsRejoinChatChannels (CMsgClientsRejoinChatChannels). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RejoinAllChatChannels() {
	req := &protocol.CMsgClientsRejoinChatChannels{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels), req)
}

// ReleaseEditorItemReservation releases a editor item reservation.
//
// Sends the GC message k_EMsgGCItemEditorReleaseReservation (CMsgGCItemEditorReleaseReservation) and awaits the response k_EMsgGCItemEditorReleaseReservationResponse,
// delivered as *CMsgGCItemEditorReleaseReservationResponse.
func (d *Dota2) ReleaseEditorItemReservation(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*protocol.CMsgGCItemEditorReleaseReservationResponse, error) {
	req := &protocol.CMsgGCItemEditorReleaseReservation{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &protocol.CMsgGCItemEditorReleaseReservationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservation),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservationResponse),
		resp,
	)
}

// ReportBattleAcknowledge reports a battle acknowledge.
//
// Sends the GC message k_EMsgClientToGCAcknowledgeBattleReport (CMsgClientToGCAcknowledgeBattleReport) and awaits the response k_EMsgClientToGCAcknowledgeBattleReportResponse,
// delivered as *CMsgClientToGCAcknowledgeBattleReportResponse.
func (d *Dota2) ReportBattleAcknowledge(
	ctx context.Context,
	accountID uint32,
	timestamp uint32,
	duration uint32,
) (*protocol.CMsgClientToGCAcknowledgeBattleReportResponse, error) {
	req := &protocol.CMsgClientToGCAcknowledgeBattleReport{
		AccountId: &accountID,
		Timestamp: &timestamp,
		Duration:  &duration,
	}
	resp := &protocol.CMsgClientToGCAcknowledgeBattleReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcknowledgeBattleReport),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcknowledgeBattleReportResponse),
		resp,
	)
}

// ReportChatPublicSpam reports a chat public spam.
//
// Sends the GC message k_EMsgGCChatReportPublicSpam (CMsgGCChatReportPublicSpam). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ReportChatPublicSpam(
	channelID uint64,
	channelUserID uint32,
) {
	req := &protocol.CMsgGCChatReportPublicSpam{
		ChannelId:     &channelID,
		ChannelUserId: &channelUserID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCChatReportPublicSpam), req)
}

// ReportGuildContent reports a guild content.
//
// Sends the GC message k_EMsgClientToGCReportGuildContent (CMsgClientToGCReportGuildContent) and awaits the response k_EMsgClientToGCReportGuildContentResponse,
// delivered as *CMsgClientToGCReportGuildContentResponse.
func (d *Dota2) ReportGuildContent(
	ctx context.Context,
	guildID uint32,
	guildContentFlags uint32,
) (*protocol.CMsgClientToGCReportGuildContentResponse, error) {
	req := &protocol.CMsgClientToGCReportGuildContent{
		GuildId:           &guildID,
		GuildContentFlags: &guildContentFlags,
	}
	resp := &protocol.CMsgClientToGCReportGuildContentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCReportGuildContent),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCReportGuildContentResponse),
		resp,
	)
}

// RequestAccountGuildEventData requests a account guild event data.
//
// Sends the GC message k_EMsgClientToGCRequestAccountGuildEventData (CMsgClientToGCRequestAccountGuildEventData) and awaits the response k_EMsgClientToGCRequestAccountGuildEventDataResponse,
// delivered as *CMsgClientToGCRequestAccountGuildEventDataResponse.
func (d *Dota2) RequestAccountGuildEventData(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestAccountGuildEventDataResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildEventData{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildEventDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildEventData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildEventDataResponse),
		resp,
	)
}

// RequestAccountGuildPersonaInfo requests a account guild persona info.
//
// Sends the GC message k_EMsgClientToGCRequestAccountGuildPersonaInfo (CMsgClientToGCRequestAccountGuildPersonaInfo) and awaits the response k_EMsgClientToGCRequestAccountGuildPersonaInfoResponse,
// delivered as *CMsgClientToGCRequestAccountGuildPersonaInfoResponse.
func (d *Dota2) RequestAccountGuildPersonaInfo(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCRequestAccountGuildPersonaInfoResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfo{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoResponse),
		resp,
	)
}

// RequestAccountGuildPersonaInfoBatch requests a account guild persona info batch.
//
// Sends the GC message k_EMsgClientToGCRequestAccountGuildPersonaInfoBatch (CMsgClientToGCRequestAccountGuildPersonaInfoBatch) and awaits the response k_EMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse,
// delivered as *CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse.
func (d *Dota2) RequestAccountGuildPersonaInfoBatch(
	ctx context.Context,
	accountIDs []uint32,
) (*protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse, error) {
	req := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatch{
		AccountIds: accountIDs,
	}
	resp := &protocol.CMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoBatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestAccountGuildPersonaInfoBatchResponse),
		resp,
	)
}

// RequestActiveBeaconParties requests active beacon parties.
//
// Sends the GC message k_EMsgClientToGCRequestActiveBeaconParties (CMsgClientToGCRequestActiveBeaconParties) and awaits the response k_EMsgGCToClientRequestActiveBeaconPartiesResponse,
// delivered as *CMsgGCToClientRequestActiveBeaconPartiesResponse.
func (d *Dota2) RequestActiveBeaconParties(
	ctx context.Context,
) (*protocol.CMsgGCToClientRequestActiveBeaconPartiesResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveBeaconParties{}
	resp := &protocol.CMsgGCToClientRequestActiveBeaconPartiesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveBeaconParties),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestActiveBeaconPartiesResponse),
		resp,
	)
}

// RequestActiveGuildChallenge requests a active guild challenge.
//
// Sends the GC message k_EMsgClientToGCRequestActiveGuildChallenge (CMsgClientToGCRequestActiveGuildChallenge) and awaits the response k_EMsgClientToGCRequestActiveGuildChallengeResponse,
// delivered as *CMsgClientToGCRequestActiveGuildChallengeResponse.
func (d *Dota2) RequestActiveGuildChallenge(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestActiveGuildChallengeResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveGuildChallenge{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestActiveGuildChallengeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildChallenge),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildChallengeResponse),
		resp,
	)
}

// RequestActiveGuildContracts requests active guild contracts.
//
// Sends the GC message k_EMsgClientToGCRequestActiveGuildContracts (CMsgClientToGCRequestActiveGuildContracts) and awaits the response k_EMsgClientToGCRequestActiveGuildContractsResponse,
// delivered as *CMsgClientToGCRequestActiveGuildContractsResponse.
func (d *Dota2) RequestActiveGuildContracts(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestActiveGuildContractsResponse, error) {
	req := &protocol.CMsgClientToGCRequestActiveGuildContracts{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestActiveGuildContractsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildContracts),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestActiveGuildContractsResponse),
		resp,
	)
}

// RequestAnchorPhoneNumber requests to check if the target anchor phone number.
//
// Sends the GC message k_EMsgAnchorPhoneNumberRequest (CMsgDOTAAnchorPhoneNumberRequest) and awaits the response k_EMsgAnchorPhoneNumberResponse,
// delivered as *CMsgDOTAAnchorPhoneNumberResponse.
func (d *Dota2) RequestAnchorPhoneNumber(
	ctx context.Context,
) (*protocol.CMsgDOTAAnchorPhoneNumberResponse, error) {
	req := &protocol.CMsgDOTAAnchorPhoneNumberRequest{}
	resp := &protocol.CMsgDOTAAnchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgAnchorPhoneNumberRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgAnchorPhoneNumberResponse),
		resp,
	)
}

// RequestArcanaVotesRemaining requests a arcana votes remaining.
//
// Sends the GC message k_EMsgClientToGCRequestArcanaVotesRemaining (CMsgClientToGCRequestArcanaVotesRemaining) and awaits the response k_EMsgClientToGCRequestArcanaVotesRemainingResponse,
// delivered as *CMsgClientToGCRequestArcanaVotesRemainingResponse.
func (d *Dota2) RequestArcanaVotesRemaining(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestArcanaVotesRemainingResponse, error) {
	req := &protocol.CMsgClientToGCRequestArcanaVotesRemaining{}
	resp := &protocol.CMsgClientToGCRequestArcanaVotesRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemaining),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemainingResponse),
		resp,
	)
}

// RequestBatchGetPlayerCardRoster requests a batch get player card roster.
//
// Sends the GC message k_EMsgClientToGCBatchGetPlayerCardRosterRequest (CMsgClientToGCBatchGetPlayerCardRosterRequest) and awaits the response k_EMsgClientToGCBatchGetPlayerCardRosterResponse,
// delivered as *CMsgClientToGCBatchGetPlayerCardRosterResponse.
func (d *Dota2) RequestBatchGetPlayerCardRoster(
	ctx context.Context,
	leagueTimestamps []*protocol.CMsgClientToGCBatchGetPlayerCardRosterRequest_LeagueTimestamp,
) (*protocol.CMsgClientToGCBatchGetPlayerCardRosterResponse, error) {
	req := &protocol.CMsgClientToGCBatchGetPlayerCardRosterRequest{
		LeagueTimestamps: leagueTimestamps,
	}
	resp := &protocol.CMsgClientToGCBatchGetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBatchGetPlayerCardRosterRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBatchGetPlayerCardRosterResponse),
		resp,
	)
}

// RequestChinaSSAAccepted requests a china ssa accepted.
//
// Sends the GC message k_EMsgClientToGCChinaSSAAcceptedRequest (CMsgClientToGCChinaSSAAcceptedRequest) and awaits the response k_EMsgClientToGCChinaSSAAcceptedResponse,
// delivered as *CMsgClientToGCChinaSSAAcceptedResponse.
func (d *Dota2) RequestChinaSSAAccepted(
	ctx context.Context,
) (*protocol.CMsgClientToGCChinaSSAAcceptedResponse, error) {
	req := &protocol.CMsgClientToGCChinaSSAAcceptedRequest{}
	resp := &protocol.CMsgClientToGCChinaSSAAcceptedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAAcceptedRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAAcceptedResponse),
		resp,
	)
}

// RequestChinaSSAURL requests a china ssaurl.
//
// Sends the GC message k_EMsgClientToGCChinaSSAURLRequest (CMsgClientToGCChinaSSAURLRequest) and awaits the response k_EMsgClientToGCChinaSSAURLResponse,
// delivered as *CMsgClientToGCChinaSSAURLResponse.
func (d *Dota2) RequestChinaSSAURL(
	ctx context.Context,
) (*protocol.CMsgClientToGCChinaSSAURLResponse, error) {
	req := &protocol.CMsgClientToGCChinaSSAURLRequest{}
	resp := &protocol.CMsgClientToGCChinaSSAURLResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAURLRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCChinaSSAURLResponse),
		resp,
	)
}

// RequestCollectorsCacheAvailableData requests a collectors cache available data.
//
// Sends the GC message k_EMsgClientToGCCollectorsCacheAvailableDataRequest (CMsgClientToGCCollectorsCacheAvailableDataRequest) and awaits the response k_EMsgGCToClientCollectorsCacheAvailableDataResponse,
// delivered as *CMsgGCToClientCollectorsCacheAvailableDataResponse.
func (d *Dota2) RequestCollectorsCacheAvailableData(
	ctx context.Context,
	contestID uint32,
) (*protocol.CMsgGCToClientCollectorsCacheAvailableDataResponse, error) {
	req := &protocol.CMsgClientToGCCollectorsCacheAvailableDataRequest{
		ContestId: &contestID,
	}
	resp := &protocol.CMsgGCToClientCollectorsCacheAvailableDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCollectorsCacheAvailableDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCollectorsCacheAvailableDataResponse),
		resp,
	)
}

// RequestCompendiumData requests a compendium data.
//
// Sends the GC message k_EMsgGCCompendiumDataRequest (CMsgDOTACompendiumDataRequest) and awaits the response k_EMsgGCCompendiumDataResponse,
// delivered as *CMsgDOTACompendiumDataResponse.
func (d *Dota2) RequestCompendiumData(
	ctx context.Context,
	accountID uint32,
	leagueid uint32,
) (*protocol.CMsgDOTACompendiumDataResponse, error) {
	req := &protocol.CMsgDOTACompendiumDataRequest{
		AccountId: &accountID,
		Leagueid:  &leagueid,
	}
	resp := &protocol.CMsgDOTACompendiumDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumDataResponse),
		resp,
	)
}

// RequestContestVotes requests contest votes.
//
// Sends the GC message k_EMsgClientToGCRequestContestVotes (CMsgClientToGCRequestContestVotes) and awaits the response k_EMsgClientToGCRequestContestVotesResponse,
// delivered as *CMsgClientToGCRequestContestVotesResponse.
func (d *Dota2) RequestContestVotes(
	ctx context.Context,
	contestID uint32,
) (*protocol.CMsgClientToGCRequestContestVotesResponse, error) {
	req := &protocol.CMsgClientToGCRequestContestVotes{
		ContestId: &contestID,
	}
	resp := &protocol.CMsgClientToGCRequestContestVotesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestContestVotes),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestContestVotesResponse),
		resp,
	)
}

// RequestCrawlCavernMapState requests a crawl cavern map state.
//
// Sends the GC message k_EMsgClientToGCCavernCrawlRequestMapState (CMsgClientToGCCavernCrawlRequestMapState) and awaits the response k_EMsgClientToGCCavernCrawlRequestMapStateResponse,
// delivered as *CMsgClientToGCCavernCrawlRequestMapStateResponse.
func (d *Dota2) RequestCrawlCavernMapState(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCCavernCrawlRequestMapStateResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlRequestMapState{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlRequestMapStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlRequestMapState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlRequestMapStateResponse),
		resp,
	)
}

// RequestCreateStickerbookPage requests to check if the target create stickerbook page.
//
// Sends the GC message k_EMsgClientToGCCreateStickerbookPageRequest (CMsgClientToGCCreateStickerbookPageRequest) and awaits the response k_EMsgClientToGCCreateStickerbookPageResponse,
// delivered as *CMsgClientToGCCreateStickerbookPageResponse.
func (d *Dota2) RequestCreateStickerbookPage(
	ctx context.Context,
	teamID uint32,
	eventID protocol.EEvent,
	pageType protocol.EStickerbookPageType,
) (*protocol.CMsgClientToGCCreateStickerbookPageResponse, error) {
	req := &protocol.CMsgClientToGCCreateStickerbookPageRequest{
		TeamId:   &teamID,
		EventId:  &eventID,
		PageType: &pageType,
	}
	resp := &protocol.CMsgClientToGCCreateStickerbookPageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateStickerbookPageRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCreateStickerbookPageResponse),
		resp,
	)
}

// RequestCustomGamesFriendsPlayed requests a custom games friends played.
//
// Sends the GC message k_EMsgClientToGCCustomGamesFriendsPlayedRequest (CMsgClientToGCCustomGamesFriendsPlayedRequest) and awaits the response k_EMsgGCToClientCustomGamesFriendsPlayedResponse,
// delivered as *CMsgGCToClientCustomGamesFriendsPlayedResponse.
func (d *Dota2) RequestCustomGamesFriendsPlayed(
	ctx context.Context,
) (*protocol.CMsgGCToClientCustomGamesFriendsPlayedResponse, error) {
	req := &protocol.CMsgClientToGCCustomGamesFriendsPlayedRequest{}
	resp := &protocol.CMsgGCToClientCustomGamesFriendsPlayedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCustomGamesFriendsPlayedRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCustomGamesFriendsPlayedResponse),
		resp,
	)
}

// RequestDeleteStickerbookPage requests a delete stickerbook page.
//
// Sends the GC message k_EMsgClientToGCDeleteStickerbookPageRequest (CMsgClientToGCDeleteStickerbookPageRequest) and awaits the response k_EMsgClientToGCDeleteStickerbookPageResponse,
// delivered as *CMsgClientToGCDeleteStickerbookPageResponse.
func (d *Dota2) RequestDeleteStickerbookPage(
	ctx context.Context,
	pageNum uint32,
	stickerCount uint32,
	stickerMax uint32,
) (*protocol.CMsgClientToGCDeleteStickerbookPageResponse, error) {
	req := &protocol.CMsgClientToGCDeleteStickerbookPageRequest{
		PageNum:      &pageNum,
		StickerCount: &stickerCount,
		StickerMax:   &stickerMax,
	}
	resp := &protocol.CMsgClientToGCDeleteStickerbookPageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeleteStickerbookPageRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeleteStickerbookPageResponse),
		resp,
	)
}

// RequestEmoticonData requests a emoticon data.
//
// Sends the GC message k_EMsgClientToGCEmoticonDataRequest (CMsgClientToGCEmoticonDataRequest) and awaits the response k_EMsgGCToClientEmoticonData,
// delivered as *CMsgGCToClientEmoticonData.
func (d *Dota2) RequestEmoticonData(
	ctx context.Context,
) (*protocol.CMsgGCToClientEmoticonData, error) {
	req := &protocol.CMsgClientToGCEmoticonDataRequest{}
	resp := &protocol.CMsgGCToClientEmoticonData{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEmoticonDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientEmoticonData),
		resp,
	)
}

// RequestEventGoals requests event goals.
//
// Sends the GC message k_EMsgClientToGCEventGoalsRequest (CMsgClientToGCGetEventGoals) and awaits the response k_EMsgClientToGCEventGoalsResponse,
// delivered as *CMsgEventGoals.
func (d *Dota2) RequestEventGoals(
	ctx context.Context,
	eventIDs []protocol.EEvent,
) (*protocol.CMsgEventGoals, error) {
	req := &protocol.CMsgClientToGCGetEventGoals{
		EventIds: eventIDs,
	}
	resp := &protocol.CMsgEventGoals{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse),
		resp,
	)
}

// RequestEventPointLogResponseV2 requests a event point log response v 2.
//
// Sends the GC message k_EMsgClientToGCRequestEventPointLogResponseV2 (CMsgClientToGCRequestEventPointLogResponseV2). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RequestEventPointLogResponseV2(
	result bool,
	eventID protocol.EEvent,
	logEntries []*protocol.CMsgClientToGCRequestEventPointLogResponseV2_LogEntry,
) {
	req := &protocol.CMsgClientToGCRequestEventPointLogResponseV2{
		Result:     &result,
		EventId:    &eventID,
		LogEntries: logEntries,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLogResponseV2), req)
}

// RequestEventPointLogV2 requests a event point log v 2.
//
// Sends the GC message k_EMsgClientToGCRequestEventPointLogV2 (CMsgClientToGCRequestEventPointLogV2). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RequestEventPointLogV2(
	eventID uint32,
) {
	req := &protocol.CMsgClientToGCRequestEventPointLogV2{
		EventId: &eventID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLogV2), req)
}

// RequestEventTipsSummary requests a event tips summary.
//
// Sends the GC message k_EMsgClientToGCRequestEventTipsSummary (CMsgEventTipsSummaryRequest) and awaits the response k_EMsgClientToGCRequestEventTipsSummaryResponse,
// delivered as *CMsgEventTipsSummaryResponse.
func (d *Dota2) RequestEventTipsSummary(
	ctx context.Context,
	eventID protocol.EEvent,
	accountID uint32,
) (*protocol.CMsgEventTipsSummaryResponse, error) {
	req := &protocol.CMsgEventTipsSummaryRequest{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgEventTipsSummaryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummary),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummaryResponse),
		resp,
	)
}

// RequestFriendsPlayedCustomGame requests a friends played custom game.
//
// Sends the GC message k_EMsgClientToGCFriendsPlayedCustomGameRequest (CMsgClientToGCFriendsPlayedCustomGameRequest) and awaits the response k_EMsgGCToClientFriendsPlayedCustomGameResponse,
// delivered as *CMsgGCToClientFriendsPlayedCustomGameResponse.
func (d *Dota2) RequestFriendsPlayedCustomGame(
	ctx context.Context,
	customGameID uint64,
) (*protocol.CMsgGCToClientFriendsPlayedCustomGameResponse, error) {
	req := &protocol.CMsgClientToGCFriendsPlayedCustomGameRequest{
		CustomGameId: &customGameID,
	}
	resp := &protocol.CMsgGCToClientFriendsPlayedCustomGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFriendsPlayedCustomGameRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFriendsPlayedCustomGameResponse),
		resp,
	)
}

// RequestGetPlayerCardRoster requests to check if the target get player card roster.
//
// Sends the GC message k_EMsgClientToGCGetPlayerCardRosterRequest (CMsgClientToGCGetPlayerCardRosterRequest) and awaits the response k_EMsgClientToGCGetPlayerCardRosterResponse,
// delivered as *CMsgClientToGCGetPlayerCardRosterResponse.
func (d *Dota2) RequestGetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	fantasyPeriod uint32,
) (*protocol.CMsgClientToGCGetPlayerCardRosterResponse, error) {
	req := &protocol.CMsgClientToGCGetPlayerCardRosterRequest{
		LeagueId:      &leagueID,
		FantasyPeriod: &fantasyPeriod,
	}
	resp := &protocol.CMsgClientToGCGetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterResponse),
		resp,
	)
}

// RequestGetRecentPlayTimeFriends requests to check if the target get recent play time friends.
//
// Sends the GC message k_EMsgGetRecentPlayTimeFriendsRequest (CMsgDOTAGetRecentPlayTimeFriendsRequest) and awaits the response k_EMsgGetRecentPlayTimeFriendsResponse,
// delivered as *CMsgDOTAGetRecentPlayTimeFriendsResponse.
func (d *Dota2) RequestGetRecentPlayTimeFriends(
	ctx context.Context,
) (*protocol.CMsgDOTAGetRecentPlayTimeFriendsResponse, error) {
	req := &protocol.CMsgDOTAGetRecentPlayTimeFriendsRequest{}
	resp := &protocol.CMsgDOTAGetRecentPlayTimeFriendsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGetRecentPlayTimeFriendsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGetRecentPlayTimeFriendsResponse),
		resp,
	)
}

// RequestGetStickerbook requests to check if the target get stickerbook.
//
// Sends the GC message k_EMsgClientToGCGetStickerbookRequest (CMsgClientToGCGetStickerbookRequest) and awaits the response k_EMsgClientToGCGetStickerbookResponse,
// delivered as *CMsgClientToGCGetStickerbookResponse.
func (d *Dota2) RequestGetStickerbook(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetStickerbookResponse, error) {
	req := &protocol.CMsgClientToGCGetStickerbookRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetStickerbookResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetStickerbookRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCGetStickerbookResponse),
		resp,
	)
}

// RequestGuildData requests a guild data.
//
// Sends the GC message k_EMsgClientToGCRequestGuildData (CMsgClientToGCRequestGuildData) and awaits the response k_EMsgClientToGCRequestGuildDataResponse,
// delivered as *CMsgClientToGCRequestGuildDataResponse.
func (d *Dota2) RequestGuildData(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCRequestGuildDataResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildData{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildDataResponse),
		resp,
	)
}

// RequestGuildEventMembers requests guild event members.
//
// Sends the GC message k_EMsgClientToGCRequestGuildEventMembers (CMsgClientToGCRequestGuildEventMembers) and awaits the response k_EMsgClientToGCRequestGuildEventMembersResponse,
// delivered as *CMsgClientToGCRequestGuildEventMembersResponse.
func (d *Dota2) RequestGuildEventMembers(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
) (*protocol.CMsgClientToGCRequestGuildEventMembersResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildEventMembers{
		GuildId: &guildID,
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildEventMembersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildEventMembers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildEventMembersResponse),
		resp,
	)
}

// RequestGuildFeed requests a guild feed.
//
// Sends the GC message k_EMsgClientToGCRequestGuildFeed (CMsgClientToGCGuildFeedRequest) and awaits the response k_EMsgClientToGCRequestGuildFeedResponse,
// delivered as *CMsgClientToGCRequestGuildFeedResponse.
func (d *Dota2) RequestGuildFeed(
	ctx context.Context,
	guildID uint32,
	lastSeenID uint64,
) (*protocol.CMsgClientToGCRequestGuildFeedResponse, error) {
	req := &protocol.CMsgClientToGCGuildFeedRequest{
		GuildId:    &guildID,
		LastSeenId: &lastSeenID,
	}
	resp := &protocol.CMsgClientToGCRequestGuildFeedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeed),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeedResponse),
		resp,
	)
}

// RequestGuildMembership requests a guild membership.
//
// Sends the GC message k_EMsgClientToGCRequestGuildMembership (CMsgClientToGCRequestGuildMembership) and awaits the response k_EMsgClientToGCRequestGuildMembershipResponse,
// delivered as *CMsgClientToGCRequestGuildMembershipResponse.
func (d *Dota2) RequestGuildMembership(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestGuildMembershipResponse, error) {
	req := &protocol.CMsgClientToGCRequestGuildMembership{}
	resp := &protocol.CMsgClientToGCRequestGuildMembershipResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildMembership),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestGuildMembershipResponse),
		resp,
	)
}

// RequestHeroGlobalData requests a hero global data.
//
// Sends the GC message k_EMsgHeroGlobalDataRequest (CMsgHeroGlobalDataRequest) and awaits the response k_EMsgHeroGlobalDataResponse,
// delivered as *CMsgHeroGlobalDataResponse.
func (d *Dota2) RequestHeroGlobalData(
	ctx context.Context,
	heroID int32,
) (*protocol.CMsgHeroGlobalDataResponse, error) {
	req := &protocol.CMsgHeroGlobalDataRequest{
		HeroId: &heroID,
	}
	resp := &protocol.CMsgHeroGlobalDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgHeroGlobalDataRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgHeroGlobalDataResponse),
		resp,
	)
}

// RequestHunterMonsterMaterialsNeededByFriend requests a hunter monster materials needed by friend.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterRequestMaterialsNeededByFriend (CMsgClientToGCMonsterHunterRequestMaterialsNeededByFriend) and awaits the response k_EMsgClientToGCMonsterHunterRequestMaterialsNeededByFriendResponse,
// delivered as *CMsgClientToGCMonsterHunterRequestMaterialsNeededByFriendResponse.
func (d *Dota2) RequestHunterMonsterMaterialsNeededByFriend(
	ctx context.Context,
	friendAccountID uint32,
) (*protocol.CMsgClientToGCMonsterHunterRequestMaterialsNeededByFriendResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterRequestMaterialsNeededByFriend{
		FriendAccountId: &friendAccountID,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterRequestMaterialsNeededByFriendResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterRequestMaterialsNeededByFriend),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterRequestMaterialsNeededByFriendResponse),
		resp,
	)
}

// RequestItemEditorReservations requests item editor reservations.
//
// Sends the GC message k_EMsgGCItemEditorReservationsRequest (CMsgGCItemEditorReservationsRequest) and awaits the response k_EMsgGCItemEditorReservationsResponse,
// delivered as *CMsgGCItemEditorReservationsResponse.
func (d *Dota2) RequestItemEditorReservations(
	ctx context.Context,
) (*protocol.CMsgGCItemEditorReservationsResponse, error) {
	req := &protocol.CMsgGCItemEditorReservationsRequest{}
	resp := &protocol.CMsgGCItemEditorReservationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReservationsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReservationsResponse),
		resp,
	)
}

// RequestJoinableCustomGameModes requests joinable custom game modes.
//
// Sends the GC message k_EMsgGCJoinableCustomGameModesRequest (CMsgJoinableCustomGameModesRequest) and awaits the response k_EMsgGCJoinableCustomGameModesResponse,
// delivered as *CMsgJoinableCustomGameModesResponse.
func (d *Dota2) RequestJoinableCustomGameModes(
	ctx context.Context,
	serverRegion uint32,
) (*protocol.CMsgJoinableCustomGameModesResponse, error) {
	req := &protocol.CMsgJoinableCustomGameModesRequest{
		ServerRegion: &serverRegion,
	}
	resp := &protocol.CMsgJoinableCustomGameModesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesResponse),
		resp,
	)
}

// RequestJoinableCustomLobbies requests joinable custom lobbies.
//
// Sends the GC message k_EMsgGCJoinableCustomLobbiesRequest (CMsgJoinableCustomLobbiesRequest) and awaits the response k_EMsgGCJoinableCustomLobbiesResponse,
// delivered as *CMsgJoinableCustomLobbiesResponse.
func (d *Dota2) RequestJoinableCustomLobbies(
	ctx context.Context,
	serverRegion uint32,
	customGameID uint64,
) (*protocol.CMsgJoinableCustomLobbiesResponse, error) {
	req := &protocol.CMsgJoinableCustomLobbiesRequest{
		ServerRegion: &serverRegion,
		CustomGameId: &customGameID,
	}
	resp := &protocol.CMsgJoinableCustomLobbiesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesResponse),
		resp,
	)
}

// RequestLatestConductScorecard requests a latest conduct scorecard.
//
// Sends the GC message k_EMsgClientToGCLatestConductScorecardRequest (CMsgPlayerConductScorecardRequest) and awaits the response k_EMsgClientToGCLatestConductScorecard,
// delivered as *CMsgPlayerConductScorecard.
func (d *Dota2) RequestLatestConductScorecard(
	ctx context.Context,
) (*protocol.CMsgPlayerConductScorecard, error) {
	req := &protocol.CMsgPlayerConductScorecardRequest{}
	resp := &protocol.CMsgPlayerConductScorecard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard),
		resp,
	)
}

// RequestLeagueAvailableLobbyNodes requests league available lobby nodes.
//
// Sends the GC message k_EMsgDOTALeagueAvailableLobbyNodesRequest (CMsgDOTALeagueAvailableLobbyNodesRequest) and awaits the response k_EMsgDOTALeagueAvailableLobbyNodes,
// delivered as *CMsgDOTALeagueAvailableLobbyNodes.
func (d *Dota2) RequestLeagueAvailableLobbyNodes(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgDOTALeagueAvailableLobbyNodes, error) {
	req := &protocol.CMsgDOTALeagueAvailableLobbyNodesRequest{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgDOTALeagueAvailableLobbyNodes{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodes),
		resp,
	)
}

// RequestMapStats requests map stats.
//
// Sends the GC message k_EMsgClientToGCMapStatsRequest (CMsgClientToGCMapStatsRequest) and awaits the response k_EMsgGCToClientMapStatsResponse,
// delivered as *CMsgGCToClientMapStatsResponse.
func (d *Dota2) RequestMapStats(
	ctx context.Context,
) (*protocol.CMsgGCToClientMapStatsResponse, error) {
	req := &protocol.CMsgClientToGCMapStatsRequest{}
	resp := &protocol.CMsgGCToClientMapStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMapStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMapStatsResponse),
		resp,
	)
}

// RequestMatchDetails requests match details.
//
// Sends the GC message k_EMsgGCMatchDetailsRequest (CMsgGCMatchDetailsRequest) and awaits the response k_EMsgGCMatchDetailsResponse,
// delivered as *CMsgGCMatchDetailsResponse.
func (d *Dota2) RequestMatchDetails(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgGCMatchDetailsResponse, error) {
	req := &protocol.CMsgGCMatchDetailsRequest{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgGCMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchDetailsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchDetailsResponse),
		resp,
	)
}

// RequestMatchesMinimal requests a matches minimal.
//
// Sends the GC message k_EMsgClientToGCMatchesMinimalRequest (CMsgClientToGCMatchesMinimalRequest) and awaits the response k_EMsgClientToGCMatchesMinimalResponse,
// delivered as *CMsgClientToGCMatchesMinimalResponse.
func (d *Dota2) RequestMatchesMinimal(
	ctx context.Context,
	matchIDs []uint64,
) (*protocol.CMsgClientToGCMatchesMinimalResponse, error) {
	req := &protocol.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}
	resp := &protocol.CMsgClientToGCMatchesMinimalResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse),
		resp,
	)
}

// RequestMatchmakingStats requests matchmaking stats.
//
// Sends the GC message k_EMsgGCMatchmakingStatsRequest (CMsgDOTAMatchmakingStatsRequest) and awaits the response k_EMsgGCMatchmakingStatsResponse,
// delivered as *CMsgDOTAMatchmakingStatsResponse.
func (d *Dota2) RequestMatchmakingStats(
	ctx context.Context,
) (*protocol.CMsgDOTAMatchmakingStatsResponse, error) {
	req := &protocol.CMsgDOTAMatchmakingStatsRequest{}
	resp := &protocol.CMsgDOTAMatchmakingStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchmakingStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCMatchmakingStatsResponse),
		resp,
	)
}

// RequestMyTeamInfo requests a my team info.
//
// Sends the GC message k_EMsgClientToGCMyTeamInfoRequest (CMsgDOTAMyTeamInfoRequest) and awaits the response k_EMsgGCToClientTeamInfo,
// delivered as *CMsgDOTATeamInfo.
func (d *Dota2) RequestMyTeamInfo(
	ctx context.Context,
) (*protocol.CMsgDOTATeamInfo, error) {
	req := &protocol.CMsgDOTAMyTeamInfoRequest{}
	resp := &protocol.CMsgDOTATeamInfo{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamInfo),
		resp,
	)
}

// RequestNotifications requests notifications.
//
// Sends the GC message k_EMsgGCNotificationsRequest (CMsgGCNotificationsRequest) and awaits the response k_EMsgGCNotificationsResponse,
// delivered as *CMsgGCNotificationsResponse.
func (d *Dota2) RequestNotifications(
	ctx context.Context,
) (*protocol.CMsgGCNotificationsResponse, error) {
	req := &protocol.CMsgGCNotificationsRequest{}
	resp := &protocol.CMsgGCNotificationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsResponse),
		resp,
	)
}

// RequestNotificationsMarkRead requests a notifications mark read.
//
// Sends the GC message k_EMsgGCNotificationsMarkReadRequest (CMsgGCNotificationsMarkReadRequest). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RequestNotificationsMarkRead() {
	req := &protocol.CMsgGCNotificationsMarkReadRequest{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest), req)
}

// RequestOrderStickerbookTeamPage requests a order stickerbook team page.
//
// Sends the GC message k_EMsgClientToGCOrderStickerbookTeamPageRequest (CMsgClientToGCOrderStickerbookTeamPageRequest) and awaits the response k_EMsgClientToGCOrderStickerbookTeamPageResponse,
// delivered as *CMsgClientToGCOrderStickerbookTeamPageResponse.
func (d *Dota2) RequestOrderStickerbookTeamPage(
	ctx context.Context,
	pageOrderSequence protocol.CMsgStickerbookTeamPageOrderSequence,
) (*protocol.CMsgClientToGCOrderStickerbookTeamPageResponse, error) {
	req := &protocol.CMsgClientToGCOrderStickerbookTeamPageRequest{
		PageOrderSequence: &pageOrderSequence,
	}
	resp := &protocol.CMsgClientToGCOrderStickerbookTeamPageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOrderStickerbookTeamPageRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOrderStickerbookTeamPageResponse),
		resp,
	)
}

// RequestOverworldFortune requests a overworld fortune.
//
// Sends the GC message k_EMsgClientToGCOverworldRequestFortune (CMsgClientToGCOverworldRequestFortune) and awaits the response k_EMsgClientToGCOverworldRequestFortuneResponse,
// delivered as *CMsgClientToGCOverworldRequestFortuneResponse.
func (d *Dota2) RequestOverworldFortune(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldRequestFortuneResponse, error) {
	req := &protocol.CMsgClientToGCOverworldRequestFortune{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldRequestFortuneResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldRequestFortune),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldRequestFortuneResponse),
		resp,
	)
}

// RequestOverworldTokensNeededByFriend requests a overworld tokens needed by friend.
//
// Sends the GC message k_EMsgClientToGCOverworldRequestTokensNeededByFriend (CMsgClientToGCOverworldRequestTokensNeededByFriend) and awaits the response k_EMsgClientToGCOverworldRequestTokensNeededByFriendResponse,
// delivered as *CMsgClientToGCOverworldRequestTokensNeededByFriendResponse.
func (d *Dota2) RequestOverworldTokensNeededByFriend(
	ctx context.Context,
	friendAccountID uint32,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldRequestTokensNeededByFriendResponse, error) {
	req := &protocol.CMsgClientToGCOverworldRequestTokensNeededByFriend{
		FriendAccountId: &friendAccountID,
		OverworldId:     &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldRequestTokensNeededByFriendResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldRequestTokensNeededByFriend),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldRequestTokensNeededByFriendResponse),
		resp,
	)
}

// RequestPlaceCollectionStickers requests place collection stickers.
//
// Sends the GC message k_EMsgClientToGCPlaceCollectionStickersRequest (CMsgClientToGCPlaceCollectionStickersRequest) and awaits the response k_EMsgClientToGCPlaceCollectionStickersResponse,
// delivered as *CMsgClientToGCPlaceCollectionStickersResponse.
func (d *Dota2) RequestPlaceCollectionStickers(
	ctx context.Context,
	slots []*protocol.CMsgClientToGCPlaceCollectionStickersRequest_Slot,
) (*protocol.CMsgClientToGCPlaceCollectionStickersResponse, error) {
	req := &protocol.CMsgClientToGCPlaceCollectionStickersRequest{
		Slots: slots,
	}
	resp := &protocol.CMsgClientToGCPlaceCollectionStickersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlaceCollectionStickersRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlaceCollectionStickersResponse),
		resp,
	)
}

// RequestPlaceStickers requests place stickers.
//
// Sends the GC message k_EMsgClientToGCPlaceStickersRequest (CMsgClientToGCPlaceStickersRequest) and awaits the response k_EMsgClientToGCPlaceStickersResponse,
// delivered as *CMsgClientToGCPlaceStickersResponse.
func (d *Dota2) RequestPlaceStickers(
	ctx context.Context,
	stickerItems []*protocol.CMsgClientToGCPlaceStickersRequest_StickerItem,
) (*protocol.CMsgClientToGCPlaceStickersResponse, error) {
	req := &protocol.CMsgClientToGCPlaceStickersRequest{
		StickerItems: stickerItems,
	}
	resp := &protocol.CMsgClientToGCPlaceStickersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlaceStickersRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlaceStickersResponse),
		resp,
	)
}

// RequestPlayerCoachMatch requests a player coach match.
//
// Sends the GC message k_EMsgClientToGCRequestPlayerCoachMatch (CMsgClientToGCRequestPlayerCoachMatch) and awaits the response k_EMsgClientToGCRequestPlayerCoachMatchResponse,
// delivered as *CMsgClientToGCRequestPlayerCoachMatchResponse.
func (d *Dota2) RequestPlayerCoachMatch(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCRequestPlayerCoachMatchResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerCoachMatch{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerCoachMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatchResponse),
		resp,
	)
}

// RequestPlayerCoachMatches requests player coach matches.
//
// Sends the GC message k_EMsgClientToGCRequestPlayerCoachMatches (CMsgClientToGCRequestPlayerCoachMatches) and awaits the response k_EMsgClientToGCRequestPlayerCoachMatchesResponse,
// delivered as *CMsgClientToGCRequestPlayerCoachMatchesResponse.
func (d *Dota2) RequestPlayerCoachMatches(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestPlayerCoachMatchesResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerCoachMatches{}
	resp := &protocol.CMsgClientToGCRequestPlayerCoachMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatches),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerCoachMatchesResponse),
		resp,
	)
}

// RequestPlayerHeroRecentAccomplishments requests player hero recent accomplishments.
//
// Sends the GC message k_EMsgClientToGCRequestPlayerHeroRecentAccomplishments (CMsgClientToGCRequestPlayerHeroRecentAccomplishments) and awaits the response k_EMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse,
// delivered as *CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse.
func (d *Dota2) RequestPlayerHeroRecentAccomplishments(
	ctx context.Context,
	accountID uint32,
	heroID int32,
) (*protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishments{
		AccountId: &accountID,
		HeroId:    &heroID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerHeroRecentAccomplishments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerHeroRecentAccomplishmentsResponse),
		resp,
	)
}

// RequestPlayerRecentAccomplishments requests player recent accomplishments.
//
// Sends the GC message k_EMsgClientToGCRequestPlayerRecentAccomplishments (CMsgClientToGCRequestPlayerRecentAccomplishments) and awaits the response k_EMsgClientToGCRequestPlayerRecentAccomplishmentsResponse,
// delivered as *CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse.
func (d *Dota2) RequestPlayerRecentAccomplishments(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlayerRecentAccomplishments{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCRequestPlayerRecentAccomplishmentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerRecentAccomplishments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlayerRecentAccomplishmentsResponse),
		resp,
	)
}

// RequestPlayerStats requests player stats.
//
// Sends the GC message k_EMsgClientToGCPlayerStatsRequest (CMsgClientToGCPlayerStatsRequest) and awaits the response k_EMsgGCToClientPlayerStatsResponse,
// delivered as *CMsgGCToClientPlayerStatsResponse.
func (d *Dota2) RequestPlayerStats(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgGCToClientPlayerStatsResponse, error) {
	req := &protocol.CMsgClientToGCPlayerStatsRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgGCToClientPlayerStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPlayerStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlayerStatsResponse),
		resp,
	)
}

// RequestPlusWeeklyChallengeResult requests a plus weekly challenge result.
//
// Sends the GC message k_EMsgClientToGCRequestPlusWeeklyChallengeResult (CMsgClientToGCRequestPlusWeeklyChallengeResult) and awaits the response k_EMsgClientToGCRequestPlusWeeklyChallengeResultResponse,
// delivered as *CMsgClientToGCRequestPlusWeeklyChallengeResultResponse.
func (d *Dota2) RequestPlusWeeklyChallengeResult(
	ctx context.Context,
	eventID protocol.EEvent,
	week uint32,
) (*protocol.CMsgClientToGCRequestPlusWeeklyChallengeResultResponse, error) {
	req := &protocol.CMsgClientToGCRequestPlusWeeklyChallengeResult{
		EventId: &eventID,
		Week:    &week,
	}
	resp := &protocol.CMsgClientToGCRequestPlusWeeklyChallengeResultResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlusWeeklyChallengeResult),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPlusWeeklyChallengeResultResponse),
		resp,
	)
}

// RequestPrivateCoachingSession requests a private coaching session.
//
// Sends the GC message k_EMsgClientToGCRequestPrivateCoachingSession (CMsgClientToGCRequestPrivateCoachingSession) and awaits the response k_EMsgClientToGCRequestPrivateCoachingSessionResponse,
// delivered as *CMsgClientToGCRequestPrivateCoachingSessionResponse.
func (d *Dota2) RequestPrivateCoachingSession(
	ctx context.Context,
	language uint32,
) (*protocol.CMsgClientToGCRequestPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCRequestPrivateCoachingSession{
		Language: &language,
	}
	resp := &protocol.CMsgClientToGCRequestPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestPrivateCoachingSessionResponse),
		resp,
	)
}

// RequestPrivateMetadataKey requests a private metadata key.
//
// Sends the GC message k_EMsgPrivateMetadataKeyRequest (CMsgPrivateMetadataKeyRequest) and awaits the response k_EMsgPrivateMetadataKeyResponse,
// delivered as *CMsgPrivateMetadataKeyResponse.
func (d *Dota2) RequestPrivateMetadataKey(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgPrivateMetadataKeyResponse, error) {
	req := &protocol.CMsgPrivateMetadataKeyRequest{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgPrivateMetadataKeyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPrivateMetadataKeyRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPrivateMetadataKeyResponse),
		resp,
	)
}

// RequestProfile requests a profile.
//
// Sends the GC message k_EMsgProfileRequest (CMsgProfileRequest) and awaits the response k_EMsgProfileResponse,
// delivered as *CMsgProfileResponse.
func (d *Dota2) RequestProfile(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgProfileResponse, error) {
	req := &protocol.CMsgProfileRequest{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgProfileResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileResponse),
		resp,
	)
}

// RequestQuickStats requests quick stats.
//
// Sends the GC message k_EMsgClientToGCQuickStatsRequest (CMsgDOTAClientToGCQuickStatsRequest) and awaits the response k_EMsgClientToGCQuickStatsResponse,
// delivered as *CMsgDOTAClientToGCQuickStatsResponse.
func (d *Dota2) RequestQuickStats(
	ctx context.Context,
	playerAccountID uint32,
	heroID int32,
	itemID int32,
	leagueID uint32,
) (*protocol.CMsgDOTAClientToGCQuickStatsResponse, error) {
	req := &protocol.CMsgDOTAClientToGCQuickStatsRequest{
		PlayerAccountId: &playerAccountID,
		HeroId:          &heroID,
		ItemId:          &itemID,
		LeagueId:        &leagueID,
	}
	resp := &protocol.CMsgDOTAClientToGCQuickStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCQuickStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCQuickStatsResponse),
		resp,
	)
}

// RequestRank requests a rank.
//
// Sends the GC message k_EMsgClientToGCRankRequest (CMsgClientToGCRankRequest) and awaits the response k_EMsgGCToClientRankResponse,
// delivered as *CMsgGCToClientRankResponse.
func (d *Dota2) RequestRank(
	ctx context.Context,
	rankType protocol.ERankType,
) (*protocol.CMsgGCToClientRankResponse, error) {
	req := &protocol.CMsgClientToGCRankRequest{
		RankType: &rankType,
	}
	resp := &protocol.CMsgGCToClientRankResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRankRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRankResponse),
		resp,
	)
}

// RequestReporterUpdates requests reporter updates.
//
// Sends the GC message k_EMsgClientToGCRequestReporterUpdates (CMsgClientToGCRequestReporterUpdates) and awaits the response k_EMsgClientToGCRequestReporterUpdatesResponse,
// delivered as *CMsgClientToGCRequestReporterUpdatesResponse.
func (d *Dota2) RequestReporterUpdates(
	ctx context.Context,
) (*protocol.CMsgClientToGCRequestReporterUpdatesResponse, error) {
	req := &protocol.CMsgClientToGCRequestReporterUpdates{}
	resp := &protocol.CMsgClientToGCRequestReporterUpdatesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestReporterUpdates),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestReporterUpdatesResponse),
		resp,
	)
}

// RequestReportsRemaining requests a reports remaining.
//
// Sends the GC message k_EMsgGCReportsRemainingRequest (CMsgDOTAReportsRemainingRequest) and awaits the response k_EMsgGCReportsRemainingResponse,
// delivered as *CMsgDOTAReportsRemainingResponse.
func (d *Dota2) RequestReportsRemaining(
	ctx context.Context,
) (*protocol.CMsgDOTAReportsRemainingResponse, error) {
	req := &protocol.CMsgDOTAReportsRemainingRequest{}
	resp := &protocol.CMsgDOTAReportsRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCReportsRemainingRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCReportsRemainingResponse),
		resp,
	)
}

// RequestRespondToCoachFriend requests a respond to coach friend.
//
// Sends the GC message k_EMsgClientToGCRespondToCoachFriendRequest (CMsgClientToGCRespondToCoachFriendRequest) and awaits the response k_EMsgClientToGCRespondToCoachFriendRequestResponse,
// delivered as *CMsgClientToGCRespondToCoachFriendRequestResponse.
func (d *Dota2) RequestRespondToCoachFriend(
	ctx context.Context,
	coachAccountID uint32,
	response protocol.ELobbyMemberCoachRequestState,
) (*protocol.CMsgClientToGCRespondToCoachFriendRequestResponse, error) {
	req := &protocol.CMsgClientToGCRespondToCoachFriendRequest{
		CoachAccountId: &coachAccountID,
		Response:       &response,
	}
	resp := &protocol.CMsgClientToGCRespondToCoachFriendRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRespondToCoachFriendRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRespondToCoachFriendRequestResponse),
		resp,
	)
}

// RequestSelectionPriorityChoice requests a selection priority choice.
//
// Sends the GC message k_EMsgSelectionPriorityChoiceRequest (CMsgDOTASelectionPriorityChoiceRequest) and awaits the response k_EMsgSelectionPriorityChoiceResponse,
// delivered as *CMsgDOTASelectionPriorityChoiceResponse.
func (d *Dota2) RequestSelectionPriorityChoice(
	ctx context.Context,
	choice protocol.DOTASelectionPriorityChoice,
) (*protocol.CMsgDOTASelectionPriorityChoiceResponse, error) {
	req := &protocol.CMsgDOTASelectionPriorityChoiceRequest{
		Choice: &choice,
	}
	resp := &protocol.CMsgDOTASelectionPriorityChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceResponse),
		resp,
	)
}

// RequestSetPlayerCardRoster requests to check if the target set player card roster.
//
// Sends the GC message k_EMsgClientToGCSetPlayerCardRosterRequest (CMsgClientToGCSetPlayerCardRosterRequest) and awaits the response k_EMsgClientToGCSetPlayerCardRosterResponse,
// delivered as *CMsgClientToGCSetPlayerCardRosterResponse.
func (d *Dota2) RequestSetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	deprecatedTimestamp uint32,
	slot uint32,
	playerCardItemID uint64,
	eventID uint32,
	fantasyPeriod uint32,
) (*protocol.CMsgClientToGCSetPlayerCardRosterResponse, error) {
	req := &protocol.CMsgClientToGCSetPlayerCardRosterRequest{
		LeagueId:            &leagueID,
		DeprecatedTimestamp: &deprecatedTimestamp,
		Slot:                &slot,
		PlayerCardItemId:    &playerCardItemID,
		EventId:             &eventID,
		FantasyPeriod:       &fantasyPeriod,
	}
	resp := &protocol.CMsgClientToGCSetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterResponse),
		resp,
	)
}

// RequestSocialFeed requests a social feed.
//
// Sends the GC message k_EMsgClientToGCRequestSocialFeed (CMsgSocialFeedRequest) and awaits the response k_EMsgClientToGCRequestSocialFeedResponse,
// delivered as *CMsgSocialFeedResponse.
func (d *Dota2) RequestSocialFeed(
	ctx context.Context,
	accountID uint32,
	selfOnly bool,
) (*protocol.CMsgSocialFeedResponse, error) {
	req := &protocol.CMsgSocialFeedRequest{
		AccountId: &accountID,
		SelfOnly:  &selfOnly,
	}
	resp := &protocol.CMsgSocialFeedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeed),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedResponse),
		resp,
	)
}

// RequestSocialFeedComments requests social feed comments.
//
// Sends the GC message k_EMsgClientToGCRequestSocialFeedComments (CMsgSocialFeedCommentsRequest) and awaits the response k_EMsgClientToGCRequestSocialFeedCommentsResponse,
// delivered as *CMsgSocialFeedCommentsResponse.
func (d *Dota2) RequestSocialFeedComments(
	ctx context.Context,
	feedEventID uint64,
) (*protocol.CMsgSocialFeedCommentsResponse, error) {
	req := &protocol.CMsgSocialFeedCommentsRequest{
		FeedEventId: &feedEventID,
	}
	resp := &protocol.CMsgSocialFeedCommentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedComments),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedCommentsResponse),
		resp,
	)
}

// RequestSocialFeedPostComment requests a social feed post comment.
//
// Sends the GC message k_EMsgClientToGCSocialFeedPostCommentRequest (CMsgClientToGCSocialFeedPostCommentRequest) and awaits the response k_EMsgGCToClientSocialFeedPostCommentResponse,
// delivered as *CMsgGCToClientSocialFeedPostCommentResponse.
func (d *Dota2) RequestSocialFeedPostComment(
	ctx context.Context,
	eventID uint64,
	comment string,
) (*protocol.CMsgGCToClientSocialFeedPostCommentResponse, error) {
	req := &protocol.CMsgClientToGCSocialFeedPostCommentRequest{
		EventId: &eventID,
		Comment: &comment,
	}
	resp := &protocol.CMsgGCToClientSocialFeedPostCommentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostCommentRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostCommentResponse),
		resp,
	)
}

// RequestSocialFeedPostMessage requests a social feed post message.
//
// Sends the GC message k_EMsgClientToGCSocialFeedPostMessageRequest (CMsgClientToGCSocialFeedPostMessageRequest) and awaits the response k_EMsgGCToClientSocialFeedPostMessageResponse,
// delivered as *CMsgGCToClientSocialFeedPostMessageResponse.
func (d *Dota2) RequestSocialFeedPostMessage(
	ctx context.Context,
	message string,
	matchID uint64,
	matchTimestamp uint32,
) (*protocol.CMsgGCToClientSocialFeedPostMessageResponse, error) {
	req := &protocol.CMsgClientToGCSocialFeedPostMessageRequest{
		Message:        &message,
		MatchId:        &matchID,
		MatchTimestamp: &matchTimestamp,
	}
	resp := &protocol.CMsgGCToClientSocialFeedPostMessageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostMessageRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostMessageResponse),
		resp,
	)
}

// RequestSteamDatagramTicket requests a steam datagram ticket.
//
// Sends the GC message k_EMsgClientToGCRequestSteamDatagramTicket (CMsgClientToGCRequestSteamDatagramTicket) and awaits the response k_EMsgClientToGCRequestSteamDatagramTicketResponse,
// delivered as *CMsgClientToGCRequestSteamDatagramTicketResponse.
func (d *Dota2) RequestSteamDatagramTicket(
	ctx context.Context,
	serverSteamID steamid.SteamId,
) (*protocol.CMsgClientToGCRequestSteamDatagramTicketResponse, error) {
	serverSteamIDU64Val := uint64(serverSteamID)
	serverSteamIDU64 := &serverSteamIDU64Val
	req := &protocol.CMsgClientToGCRequestSteamDatagramTicket{
		ServerSteamId: serverSteamIDU64,
	}
	resp := &protocol.CMsgClientToGCRequestSteamDatagramTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicket),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicketResponse),
		resp,
	)
}

// RequestSubmitPlayerAvoid requests to check if the target submit player avoid.
//
// Sends the GC message k_EMsgGCSubmitPlayerAvoidRequest (CMsgDOTASubmitPlayerAvoidRequest) and awaits the response k_EMsgGCSubmitPlayerAvoidRequestResponse,
// delivered as *CMsgDOTASubmitPlayerAvoidRequestResponse.
func (d *Dota2) RequestSubmitPlayerAvoid(
	ctx context.Context,
	targetAccountID uint32,
	lobbyID uint64,
	userNote string,
) (*protocol.CMsgDOTASubmitPlayerAvoidRequestResponse, error) {
	req := &protocol.CMsgDOTASubmitPlayerAvoidRequest{
		TargetAccountId: &targetAccountID,
		LobbyId:         &lobbyID,
		UserNote:        &userNote,
	}
	resp := &protocol.CMsgDOTASubmitPlayerAvoidRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerAvoidRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerAvoidRequestResponse),
		resp,
	)
}

// RequestTeammateStats requests teammate stats.
//
// Sends the GC message k_EMsgClientToGCTeammateStatsRequest (CMsgClientToGCTeammateStatsRequest) and awaits the response k_EMsgClientToGCTeammateStatsResponse,
// delivered as *CMsgClientToGCTeammateStatsResponse.
func (d *Dota2) RequestTeammateStats(
	ctx context.Context,
) (*protocol.CMsgClientToGCTeammateStatsResponse, error) {
	req := &protocol.CMsgClientToGCTeammateStatsRequest{}
	resp := &protocol.CMsgClientToGCTeammateStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsResponse),
		resp,
	)
}

// RequestTopFriendMatches requests top friend matches.
//
// Sends the GC message k_EMsgClientToGCTopFriendMatchesRequest (CMsgClientToGCTopFriendMatchesRequest) and awaits the response k_EMsgGCToClientTopFriendMatchesResponse,
// delivered as *CMsgGCToClientTopFriendMatchesResponse.
func (d *Dota2) RequestTopFriendMatches(
	ctx context.Context,
) (*protocol.CMsgGCToClientTopFriendMatchesResponse, error) {
	req := &protocol.CMsgClientToGCTopFriendMatchesRequest{}
	resp := &protocol.CMsgGCToClientTopFriendMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTopFriendMatchesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTopFriendMatchesResponse),
		resp,
	)
}

// RequestTopLeagueMatches requests top league matches.
//
// Sends the GC message k_EMsgClientToGCTopLeagueMatchesRequest (CMsgClientToGCTopLeagueMatchesRequest) and awaits the response k_EMsgGCToClientTopLeagueMatchesResponse,
// delivered as *CMsgGCToClientTopLeagueMatchesResponse.
func (d *Dota2) RequestTopLeagueMatches(
	ctx context.Context,
) (*protocol.CMsgGCToClientTopLeagueMatchesResponse, error) {
	req := &protocol.CMsgClientToGCTopLeagueMatchesRequest{}
	resp := &protocol.CMsgGCToClientTopLeagueMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTopLeagueMatchesRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTopLeagueMatchesResponse),
		resp,
	)
}

// RequestTransferSeasonalMMR requests to check if the target transfer seasonal mmr.
//
// Sends the GC message k_EMsgClientToGCTransferSeasonalMMRRequest (CMsgClientToGCTransferSeasonalMMRRequest) and awaits the response k_EMsgClientToGCTransferSeasonalMMRResponse,
// delivered as *CMsgClientToGCTransferSeasonalMMRResponse.
func (d *Dota2) RequestTransferSeasonalMMR(
	ctx context.Context,
	isParty bool,
) (*protocol.CMsgClientToGCTransferSeasonalMMRResponse, error) {
	req := &protocol.CMsgClientToGCTransferSeasonalMMRRequest{
		IsParty: &isParty,
	}
	resp := &protocol.CMsgClientToGCTransferSeasonalMMRResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRResponse),
		resp,
	)
}

// RequestUnanchorPhoneNumber requests a unanchor phone number.
//
// Sends the GC message k_EMsgUnanchorPhoneNumberRequest (CMsgDOTAUnanchorPhoneNumberRequest) and awaits the response k_EMsgUnanchorPhoneNumberResponse,
// delivered as *CMsgDOTAUnanchorPhoneNumberResponse.
func (d *Dota2) RequestUnanchorPhoneNumber(
	ctx context.Context,
) (*protocol.CMsgDOTAUnanchorPhoneNumberResponse, error) {
	req := &protocol.CMsgDOTAUnanchorPhoneNumberRequest{}
	resp := &protocol.CMsgDOTAUnanchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberResponse),
		resp,
	)
}

// RequestUnderDraft requests a under draft.
//
// Sends the GC message k_EMsgClientToGCUnderDraftRequest (CMsgClientToGCUnderDraftRequest) and awaits the response k_EMsgClientToGCUnderDraftResponse,
// delivered as *CMsgClientToGCUnderDraftResponse.
func (d *Dota2) RequestUnderDraft(
	ctx context.Context,
	accountID uint32,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRequest{
		AccountId: &accountID,
		EventId:   &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftResponse),
		resp,
	)
}

// RequestWagering requests a wagering.
//
// Sends the GC message k_EMsgClientToGCWageringRequest (CMsgClientToGCWageringRequest) and awaits the response k_EMsgGCToClientWageringResponse,
// delivered as *CMsgGCToClientWageringResponse.
func (d *Dota2) RequestWagering(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgGCToClientWageringResponse, error) {
	req := &protocol.CMsgClientToGCWageringRequest{
		EventId: &eventID,
	}
	resp := &protocol.CMsgGCToClientWageringResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWageringRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientWageringResponse),
		resp,
	)
}

// RerollCraftingFantasyOptions rerolls crafting fantasy options.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingRerollOptions (CMsgClientToGCFantasyCraftingRerollOptions) and awaits the response k_EMsgClientToGCFantasyCraftingRerollOptionsResponse,
// delivered as *CMsgClientToGCFantasyCraftingRerollOptionsResponse.
func (d *Dota2) RerollCraftingFantasyOptions(
	ctx context.Context,
	fantasyLeague uint32,
) (*protocol.CMsgClientToGCFantasyCraftingRerollOptionsResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingRerollOptions{
		FantasyLeague: &fantasyLeague,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingRerollOptionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingRerollOptions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingRerollOptionsResponse),
		resp,
	)
}

// RerollDevBingoCard rerolls a dev bingo card.
//
// Sends the GC message k_EMsgClientToGCBingoDevRerollCard (CMsgClientToGCBingoDevRerollCard) and awaits the response k_EMsgClientToGCBingoDevRerollCardResponse,
// delivered as *CMsgClientToGCBingoDevRerollCardResponse.
func (d *Dota2) RerollDevBingoCard(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
) (*protocol.CMsgClientToGCBingoDevRerollCardResponse, error) {
	req := &protocol.CMsgClientToGCBingoDevRerollCard{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
	}
	resp := &protocol.CMsgClientToGCBingoDevRerollCardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevRerollCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevRerollCardResponse),
		resp,
	)
}

// RerollDraftUnder rerolls a draft under.
//
// Sends the GC message k_EMsgClientToGCUnderDraftReroll (CMsgClientToGCUnderDraftReroll) and awaits the response k_EMsgClientToGCUnderDraftRerollResponse,
// delivered as *CMsgClientToGCUnderDraftRerollResponse.
func (d *Dota2) RerollDraftUnder(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftRerollResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftReroll{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRerollResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftReroll),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRerollResponse),
		resp,
	)
}

// RerollPlayerChallenge rerolls a player challenge.
//
// Sends the GC message k_EMsgClientToGCRerollPlayerChallenge (CMsgClientToGCRerollPlayerChallenge). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RerollPlayerChallenge(
	eventID protocol.EEvent,
	sequenceID uint32,
	heroID int32,
) {
	req := &protocol.CMsgClientToGCRerollPlayerChallenge{
		EventId:    &eventID,
		SequenceId: &sequenceID,
		HeroId:     &heroID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRerollPlayerChallenge), req)
}

// RerollShopCandyRewards rerolls shop candy rewards.
//
// Sends the GC message k_EMsgClientToGCCandyShopRerollRewards (CMsgClientToGCCandyShopRerollRewards) and awaits the response k_EMsgClientToGCCandyShopRerollRewardsResponse,
// delivered as *CMsgClientToGCCandyShopRerollRewardsResponse.
func (d *Dota2) RerollShopCandyRewards(
	ctx context.Context,
	candyShopID uint32,
) (*protocol.CMsgClientToGCCandyShopRerollRewardsResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopRerollRewards{
		CandyShopId: &candyShopID,
	}
	resp := &protocol.CMsgClientToGCCandyShopRerollRewardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopRerollRewards),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopRerollRewardsResponse),
		resp,
	)
}

// ReserveEditorItemItemDef reserves a editor item item def.
//
// Sends the GC message k_EMsgGCItemEditorReserveItemDef (CMsgGCItemEditorReserveItemDef) and awaits the response k_EMsgGCItemEditorReserveItemDefResponse,
// delivered as *CMsgGCItemEditorReserveItemDefResponse.
func (d *Dota2) ReserveEditorItemItemDef(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*protocol.CMsgGCItemEditorReserveItemDefResponse, error) {
	req := &protocol.CMsgGCItemEditorReserveItemDef{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &protocol.CMsgGCItemEditorReserveItemDefResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDef),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDefResponse),
		resp,
	)
}

// RespondToTeamInvite is undocumented.
//
// Sends the GC message k_EMsgGCTeamInvite_InviteeResponseToGC (CMsgDOTATeamInvite_InviteeResponseToGC). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) RespondToTeamInvite(
	result protocol.ETeamInviteResult,
) {
	req := &protocol.CMsgDOTATeamInvite_InviteeResponseToGC{
		Result: &result,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC), req)
}

// SelectCompendiumInGamePrediction selects a compendium in game prediction.
//
// Sends the GC message k_EMsgClientToGCSelectCompendiumInGamePrediction (CMsgClientToGCSelectCompendiumInGamePrediction) and awaits the response k_EMsgClientToGCSelectCompendiumInGamePredictionResponse,
// delivered as *CMsgClientToGCSelectCompendiumInGamePredictionResponse.
func (d *Dota2) SelectCompendiumInGamePrediction(
	ctx context.Context,
	matchID uint64,
	predictions []*protocol.CMsgClientToGCSelectCompendiumInGamePrediction_Prediction,
	leagueID uint32,
) (*protocol.CMsgClientToGCSelectCompendiumInGamePredictionResponse, error) {
	req := &protocol.CMsgClientToGCSelectCompendiumInGamePrediction{
		MatchId:     &matchID,
		Predictions: predictions,
		LeagueId:    &leagueID,
	}
	resp := &protocol.CMsgClientToGCSelectCompendiumInGamePredictionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePrediction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePredictionResponse),
		resp,
	)
}

// SelectCraftingFantasyGlobalPrefix selects a crafting fantasy global prefix.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingSelectGlobalPrefix (CMsgClientToGCFantasyCraftingSelectGlobalPrefix) and awaits the response k_EMsgClientToGCFantasyCraftingSelectGlobalPrefixResponse,
// delivered as *CMsgClientToGCFantasyCraftingSelectGlobalPrefixResponse.
func (d *Dota2) SelectCraftingFantasyGlobalPrefix(
	ctx context.Context,
	fantasyLeague uint32,
	prefix uint32,
) (*protocol.CMsgClientToGCFantasyCraftingSelectGlobalPrefixResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingSelectGlobalPrefix{
		FantasyLeague: &fantasyLeague,
		Prefix:        &prefix,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingSelectGlobalPrefixResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectGlobalPrefix),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectGlobalPrefixResponse),
		resp,
	)
}

// SelectCraftingFantasyGlobalSuffix selects a crafting fantasy global suffix.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingSelectGlobalSuffix (CMsgClientToGCFantasyCraftingSelectGlobalSuffix) and awaits the response k_EMsgClientToGCFantasyCraftingSelectGlobalSuffixResponse,
// delivered as *CMsgClientToGCFantasyCraftingSelectGlobalSuffixResponse.
func (d *Dota2) SelectCraftingFantasyGlobalSuffix(
	ctx context.Context,
	fantasyLeague uint32,
	suffix uint32,
) (*protocol.CMsgClientToGCFantasyCraftingSelectGlobalSuffixResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingSelectGlobalSuffix{
		FantasyLeague: &fantasyLeague,
		Suffix:        &suffix,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingSelectGlobalSuffixResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectGlobalSuffix),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectGlobalSuffixResponse),
		resp,
	)
}

// SelectCraftingFantasyPlayer selects a crafting fantasy player.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingSelectPlayer (CMsgClientToGCFantasyCraftingSelectPlayer) and awaits the response k_EMsgClientToGCFantasyCraftingSelectPlayerResponse,
// delivered as *CMsgClientToGCFantasyCraftingSelectPlayerResponse.
func (d *Dota2) SelectCraftingFantasyPlayer(
	ctx context.Context,
	fantasyLeague uint32,
	accountID uint32,
) (*protocol.CMsgClientToGCFantasyCraftingSelectPlayerResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingSelectPlayer{
		FantasyLeague: &fantasyLeague,
		AccountId:     &accountID,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingSelectPlayerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectPlayer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectPlayerResponse),
		resp,
	)
}

// SelectCraftingFantasyTeam selects a crafting fantasy team.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingSelectTeam (CMsgClientToGCFantasyCraftingSelectTeam) and awaits the response k_EMsgClientToGCFantasyCraftingSelectTeamResponse,
// delivered as *CMsgClientToGCFantasyCraftingSelectTeamResponse.
func (d *Dota2) SelectCraftingFantasyTeam(
	ctx context.Context,
	fantasyLeague uint32,
	role protocol.Fantasy_Roles,
	teamID uint32,
) (*protocol.CMsgClientToGCFantasyCraftingSelectTeamResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingSelectTeam{
		FantasyLeague: &fantasyLeague,
		Role:          &role,
		TeamId:        &teamID,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingSelectTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectTeam),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingSelectTeamResponse),
		resp,
	)
}

// SelectGuildContract selects a guild contract.
//
// Sends the GC message k_EMsgClientToGCSelectGuildContract (CMsgClientToGCSelectGuildContract) and awaits the response k_EMsgClientToGCSelectGuildContractResponse,
// delivered as *CMsgClientToGCSelectGuildContractResponse.
func (d *Dota2) SelectGuildContract(
	ctx context.Context,
	guildID uint32,
	eventID protocol.EEvent,
	contractID uint64,
	contractSlot uint32,
) (*protocol.CMsgClientToGCSelectGuildContractResponse, error) {
	req := &protocol.CMsgClientToGCSelectGuildContract{
		GuildId:      &guildID,
		EventId:      &eventID,
		ContractId:   &contractID,
		ContractSlot: &contractSlot,
	}
	resp := &protocol.CMsgClientToGCSelectGuildContractResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectGuildContract),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSelectGuildContractResponse),
		resp,
	)
}

// SendAcceptInviteToGuild sends a accept invite to guild.
//
// Sends the GC message k_EMsgClientToGCAcceptInviteToGuild (CMsgClientToGCAcceptInviteToGuild) and awaits the response k_EMsgClientToGCAcceptInviteToGuildResponse,
// delivered as *CMsgClientToGCAcceptInviteToGuildResponse.
func (d *Dota2) SendAcceptInviteToGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCAcceptInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCAcceptInviteToGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCAcceptInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptInviteToGuildResponse),
		resp,
	)
}

// SendAcceptPrivateCoachingSession sends a accept private coaching session.
//
// Sends the GC message k_EMsgClientToGCAcceptPrivateCoachingSession (CMsgClientToGCAcceptPrivateCoachingSession) and awaits the response k_EMsgClientToGCAcceptPrivateCoachingSessionResponse,
// delivered as *CMsgClientToGCAcceptPrivateCoachingSessionResponse.
func (d *Dota2) SendAcceptPrivateCoachingSession(
	ctx context.Context,
	coachingSessionID uint64,
) (*protocol.CMsgClientToGCAcceptPrivateCoachingSessionResponse, error) {
	req := &protocol.CMsgClientToGCAcceptPrivateCoachingSession{
		CoachingSessionId: &coachingSessionID,
	}
	resp := &protocol.CMsgClientToGCAcceptPrivateCoachingSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptPrivateCoachingSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcceptPrivateCoachingSessionResponse),
		resp,
	)
}

// SendAcknowledgeReporterUpdates sends acknowledge reporter updates.
//
// Sends the GC message k_EMsgClientToGCAcknowledgeReporterUpdates (CMsgClientToGCAcknowledgeReporterUpdates). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendAcknowledgeReporterUpdates(
	matchIDs []uint64,
) {
	req := &protocol.CMsgClientToGCAcknowledgeReporterUpdates{
		MatchIds: matchIDs,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAcknowledgeReporterUpdates), req)
}

// SendAddGuildRole sends a add guild role.
//
// Sends the GC message k_EMsgClientToGCAddGuildRole (CMsgClientToGCAddGuildRole) and awaits the response k_EMsgClientToGCAddGuildRoleResponse,
// delivered as *CMsgClientToGCAddGuildRoleResponse.
func (d *Dota2) SendAddGuildRole(
	ctx context.Context,
	guildID uint32,
	roleName string,
	roleFlags uint32,
) (*protocol.CMsgClientToGCAddGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCAddGuildRole{
		GuildId:   &guildID,
		RoleName:  &roleName,
		RoleFlags: &roleFlags,
	}
	resp := &protocol.CMsgClientToGCAddGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddGuildRoleResponse),
		resp,
	)
}

// SendAddPlayerToGuildChat sends a add player to guild chat.
//
// Sends the GC message k_EMsgClientToGCAddPlayerToGuildChat (CMsgClientToGCAddPlayerToGuildChat) and awaits the response k_EMsgClientToGCAddPlayerToGuildChatResponse,
// delivered as *CMsgClientToGCAddPlayerToGuildChatResponse.
func (d *Dota2) SendAddPlayerToGuildChat(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCAddPlayerToGuildChatResponse, error) {
	req := &protocol.CMsgClientToGCAddPlayerToGuildChat{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCAddPlayerToGuildChatResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddPlayerToGuildChat),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCAddPlayerToGuildChatResponse),
		resp,
	)
}

// SendBalancedShuffleLobby shuffles the lobby members between teams while
// keeping the teams balanced by MMR.
func (d *Dota2) SendBalancedShuffleLobby() {
	req := &protocol.CMsgBalancedShuffleLobby{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby), req)
}

// SendBingoDevAddTokens sends bingo dev add tokens.
//
// Sends the GC message k_EMsgClientToGCBingoDevAddTokens (CMsgClientToGCBingoDevAddTokens) and awaits the response k_EMsgClientToGCBingoDevAddTokensResponse,
// delivered as *CMsgClientToGCBingoDevAddTokensResponse.
func (d *Dota2) SendBingoDevAddTokens(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
	tokenCount int32,
) (*protocol.CMsgClientToGCBingoDevAddTokensResponse, error) {
	req := &protocol.CMsgClientToGCBingoDevAddTokens{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
		TokenCount:  &tokenCount,
	}
	resp := &protocol.CMsgClientToGCBingoDevAddTokensResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevAddTokens),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevAddTokensResponse),
		resp,
	)
}

// SendBingoDevClearInventory sends a bingo dev clear inventory.
//
// Sends the GC message k_EMsgClientToGCBingoDevClearInventory (CMsgClientToGCBingoDevClearInventory) and awaits the response k_EMsgClientToGCBingoDevClearInventoryResponse,
// delivered as *CMsgClientToGCBingoDevClearInventoryResponse.
func (d *Dota2) SendBingoDevClearInventory(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgClientToGCBingoDevClearInventoryResponse, error) {
	req := &protocol.CMsgClientToGCBingoDevClearInventory{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgClientToGCBingoDevClearInventoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevClearInventory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoDevClearInventoryResponse),
		resp,
	)
}

// SendBingoModifySquare sends a bingo modify square.
//
// Sends the GC message k_EMsgClientToGCBingoModifySquare (CMsgClientToGCBingoModifySquare) and awaits the response k_EMsgClientToGCBingoModifySquareResponse,
// delivered as *CMsgClientToGCBingoModifySquareResponse.
func (d *Dota2) SendBingoModifySquare(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
	squareIndex uint32,
	action protocol.CMsgClientToGCBingoModifySquare_EModifyAction,
) (*protocol.CMsgClientToGCBingoModifySquareResponse, error) {
	req := &protocol.CMsgClientToGCBingoModifySquare{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
		SquareIndex: &squareIndex,
		Action:      &action,
	}
	resp := &protocol.CMsgClientToGCBingoModifySquareResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoModifySquare),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoModifySquareResponse),
		resp,
	)
}

// SendBingoShuffleCard sends a bingo shuffle card.
//
// Sends the GC message k_EMsgClientToGCBingoShuffleCard (CMsgClientToGCBingoShuffleCard) and awaits the response k_EMsgClientToGCBingoShuffleCardResponse,
// delivered as *CMsgClientToGCBingoShuffleCardResponse.
func (d *Dota2) SendBingoShuffleCard(
	ctx context.Context,
	leagueID uint32,
	leaguePhase uint32,
) (*protocol.CMsgClientToGCBingoShuffleCardResponse, error) {
	req := &protocol.CMsgClientToGCBingoShuffleCard{
		LeagueId:    &leagueID,
		LeaguePhase: &leaguePhase,
	}
	resp := &protocol.CMsgClientToGCBingoShuffleCardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoShuffleCard),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCBingoShuffleCardResponse),
		resp,
	)
}

// SendCandyShopDevClearInventory sends a candy shop dev clear inventory.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevClearInventory (CMsgClientToGCCandyShopDevClearInventory) and awaits the response k_EMsgClientToGCCandyShopDevClearInventoryResponse,
// delivered as *CMsgClientToGCCandyShopDevClearInventoryResponse.
func (d *Dota2) SendCandyShopDevClearInventory(
	ctx context.Context,
	candyShopID uint32,
) (*protocol.CMsgClientToGCCandyShopDevClearInventoryResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevClearInventory{
		CandyShopId: &candyShopID,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevClearInventoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevClearInventory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevClearInventoryResponse),
		resp,
	)
}

// SendCandyShopDevResetShop sends a candy shop dev reset shop.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevResetShop (CMsgClientToGCCandyShopDevResetShop) and awaits the response k_EMsgClientToGCCandyShopDevResetShopResponse,
// delivered as *CMsgClientToGCCandyShopDevResetShopResponse.
func (d *Dota2) SendCandyShopDevResetShop(
	ctx context.Context,
	candyShopID uint32,
) (*protocol.CMsgClientToGCCandyShopDevResetShopResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevResetShop{
		CandyShopId: &candyShopID,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevResetShopResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevResetShop),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevResetShopResponse),
		resp,
	)
}

// SendCandyShopDevShuffleExchange sends a candy shop dev shuffle exchange.
//
// Sends the GC message k_EMsgClientToGCCandyShopDevShuffleExchange (CMsgClientToGCCandyShopDevShuffleExchange) and awaits the response k_EMsgClientToGCCandyShopDevShuffleExchangeResponse,
// delivered as *CMsgClientToGCCandyShopDevShuffleExchangeResponse.
func (d *Dota2) SendCandyShopDevShuffleExchange(
	ctx context.Context,
	candyShopID uint32,
) (*protocol.CMsgClientToGCCandyShopDevShuffleExchangeResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDevShuffleExchange{
		CandyShopId: &candyShopID,
	}
	resp := &protocol.CMsgClientToGCCandyShopDevShuffleExchangeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevShuffleExchange),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDevShuffleExchangeResponse),
		resp,
	)
}

// SendCandyShopDoExchange sends a candy shop do exchange.
//
// Sends the GC message k_EMsgClientToGCCandyShopDoExchange (CMsgClientToGCCandyShopDoExchange) and awaits the response k_EMsgClientToGCCandyShopDoExchangeResponse,
// delivered as *CMsgClientToGCCandyShopDoExchangeResponse.
func (d *Dota2) SendCandyShopDoExchange(
	ctx context.Context,
	candyShopID uint32,
	recipeID uint32,
) (*protocol.CMsgClientToGCCandyShopDoExchangeResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDoExchange{
		CandyShopId: &candyShopID,
		RecipeId:    &recipeID,
	}
	resp := &protocol.CMsgClientToGCCandyShopDoExchangeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDoExchange),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDoExchangeResponse),
		resp,
	)
}

// SendCandyShopDoVariableExchange sends a candy shop do variable exchange.
//
// Sends the GC message k_EMsgClientToGCCandyShopDoVariableExchange (CMsgClientToGCCandyShopDoVariableExchange) and awaits the response k_EMsgClientToGCCandyShopDoVariableExchangeResponse,
// delivered as *CMsgClientToGCCandyShopDoVariableExchangeResponse.
func (d *Dota2) SendCandyShopDoVariableExchange(
	ctx context.Context,
	candyShopID uint32,
	input protocol.CMsgCandyShopCandyQuantity,
	output protocol.CMsgCandyShopCandyQuantity,
) (*protocol.CMsgClientToGCCandyShopDoVariableExchangeResponse, error) {
	req := &protocol.CMsgClientToGCCandyShopDoVariableExchange{
		CandyShopId: &candyShopID,
		Input:       &input,
		Output:      &output,
	}
	resp := &protocol.CMsgClientToGCCandyShopDoVariableExchangeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDoVariableExchange),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCandyShopDoVariableExchangeResponse),
		resp,
	)
}

// SendCavernCrawlUseItemOnPath sends a cavern crawl use item on path.
//
// Sends the GC message k_EMsgClientToGCCavernCrawlUseItemOnPath (CMsgClientToGCCavernCrawlUseItemOnPath) and awaits the response k_EMsgClientToGCCavernCrawlUseItemOnPathResponse,
// delivered as *CMsgClientToGCCavernCrawlUseItemOnPathResponse.
func (d *Dota2) SendCavernCrawlUseItemOnPath(
	ctx context.Context,
	eventID uint32,
	pathID uint32,
	itemType uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlUseItemOnPathResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlUseItemOnPath{
		EventId:    &eventID,
		PathId:     &pathID,
		ItemType:   &itemType,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlUseItemOnPathResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnPath),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnPathResponse),
		resp,
	)
}

// SendCavernCrawlUseItemOnRoom sends a cavern crawl use item on room.
//
// Sends the GC message k_EMsgClientToGCCavernCrawlUseItemOnRoom (CMsgClientToGCCavernCrawlUseItemOnRoom) and awaits the response k_EMsgClientToGCCavernCrawlUseItemOnRoomResponse,
// delivered as *CMsgClientToGCCavernCrawlUseItemOnRoomResponse.
func (d *Dota2) SendCavernCrawlUseItemOnRoom(
	ctx context.Context,
	eventID uint32,
	roomID uint32,
	itemType uint32,
	mapVariant uint32,
) (*protocol.CMsgClientToGCCavernCrawlUseItemOnRoomResponse, error) {
	req := &protocol.CMsgClientToGCCavernCrawlUseItemOnRoom{
		EventId:    &eventID,
		RoomId:     &roomID,
		ItemType:   &itemType,
		MapVariant: &mapVariant,
	}
	resp := &protocol.CMsgClientToGCCavernCrawlUseItemOnRoomResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnRoom),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCavernCrawlUseItemOnRoomResponse),
		resp,
	)
}

// SendChatMessage sends a chat message to a joined chat channel or lobby.
// Use SendChannelMessage for plain text messages.
func (d *Dota2) SendChatMessage(
	req *protocol.CMsgDOTAChatMessage,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCChatMessage), req)
}

// SendCoachFriend sends a coach friend.
//
// Sends the GC message k_EMsgClientToGCCoachFriend (CMsgClientToGCCoachFriend) and awaits the response k_EMsgClientToGCCoachFriendResponse,
// delivered as *CMsgClientToGCCoachFriendResponse.
func (d *Dota2) SendCoachFriend(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCCoachFriendResponse, error) {
	req := &protocol.CMsgClientToGCCoachFriend{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCCoachFriendResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCoachFriend),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCoachFriendResponse),
		resp,
	)
}

// SendConvertEventPoints sends convert event points.
//
// Sends the GC message k_EMsgClientToGCConvertEventPoints (CMsgClientToGCConvertEventPoints) and awaits the response k_EMsgClientToGCConvertEventPointsResponse,
// delivered as *CMsgClientToGCConvertEventPointsResponse.
func (d *Dota2) SendConvertEventPoints(
	ctx context.Context,
	eventIDPointsToBuy protocol.EEvent,
	eventIDPointsToSpend protocol.EEvent,
	numPointsToBuy uint32,
	numPointsToSpend uint32,
) (*protocol.CMsgClientToGCConvertEventPointsResponse, error) {
	req := &protocol.CMsgClientToGCConvertEventPoints{
		EventIdPointsToBuy:   &eventIDPointsToBuy,
		EventIdPointsToSpend: &eventIDPointsToSpend,
		NumPointsToBuy:       &numPointsToBuy,
		NumPointsToSpend:     &numPointsToSpend,
	}
	resp := &protocol.CMsgClientToGCConvertEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCConvertEventPoints),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCConvertEventPointsResponse),
		resp,
	)
}

// SendCraftworksCraftRecipe sends a craftworks craft recipe.
//
// Sends the GC message k_EMsgClientToGCCraftworksCraftRecipe (CMsgClientToGCCraftworksCraftRecipe) and awaits the response k_EMsgClientToGCCraftworksCraftRecipeResponse,
// delivered as *CMsgClientToGCCraftworksCraftRecipeResponse.
func (d *Dota2) SendCraftworksCraftRecipe(
	ctx context.Context,
	craftworksID uint32,
	recipeID uint64,
) (*protocol.CMsgClientToGCCraftworksCraftRecipeResponse, error) {
	req := &protocol.CMsgClientToGCCraftworksCraftRecipe{
		CraftworksId: &craftworksID,
		RecipeId:     &recipeID,
	}
	resp := &protocol.CMsgClientToGCCraftworksCraftRecipeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksCraftRecipe),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksCraftRecipeResponse),
		resp,
	)
}

// SendCraftworksDevModifyComponents sends craftworks dev modify components.
//
// Sends the GC message k_EMsgClientToGCCraftworksDevModifyComponents (CMsgClientToGCCraftworksDevModifyComponents) and awaits the response k_EMsgClientToGCCraftworksDevModifyComponentsResponse,
// delivered as *CMsgClientToGCCraftworksDevModifyComponentsResponse.
func (d *Dota2) SendCraftworksDevModifyComponents(
	ctx context.Context,
	craftworksID uint32,
	components protocol.CMsgCraftworksComponents,
	operation protocol.CMsgClientToGCCraftworksDevModifyComponents_EOperation,
) (*protocol.CMsgClientToGCCraftworksDevModifyComponentsResponse, error) {
	req := &protocol.CMsgClientToGCCraftworksDevModifyComponents{
		CraftworksId: &craftworksID,
		Components:   &components,
		Operation:    &operation,
	}
	resp := &protocol.CMsgClientToGCCraftworksDevModifyComponentsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksDevModifyComponents),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCCraftworksDevModifyComponentsResponse),
		resp,
	)
}

// SendCustomGameClientFinishedLoading sends a custom game client finished loading.
//
// Sends the GC message k_EMsgCustomGameClientFinishedLoading (CMsgDOTACustomGameClientFinishedLoading). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendCustomGameClientFinishedLoading(
	lobbyID uint64,
	loadingDuration uint32,
	resultCode int32,
	resultString string,
	signonStates uint32,
	comment string,
) {
	req := &protocol.CMsgDOTACustomGameClientFinishedLoading{
		LobbyId:         &lobbyID,
		LoadingDuration: &loadingDuration,
		ResultCode:      &resultCode,
		ResultString:    &resultString,
		SignonStates:    &signonStates,
		Comment:         &comment,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgCustomGameClientFinishedLoading), req)
}

// SendCustomGameListenServerStartedLoading sends a custom game listen server started loading.
//
// Sends the GC message k_EMsgCustomGameListenServerStartedLoading (CMsgDOTACustomGameListenServerStartedLoading). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendCustomGameListenServerStartedLoading(
	lobbyID uint64,
	customGameID uint64,
	lobbyMembers []uint32,
	startTime uint32,
) {
	req := &protocol.CMsgDOTACustomGameListenServerStartedLoading{
		LobbyId:      &lobbyID,
		CustomGameId: &customGameID,
		LobbyMembers: lobbyMembers,
		StartTime:    &startTime,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgCustomGameListenServerStartedLoading), req)
}

// SendDeclineInviteToGuild sends a decline invite to guild.
//
// Sends the GC message k_EMsgClientToGCDeclineInviteToGuild (CMsgClientToGCDeclineInviteToGuild) and awaits the response k_EMsgClientToGCDeclineInviteToGuildResponse,
// delivered as *CMsgClientToGCDeclineInviteToGuildResponse.
func (d *Dota2) SendDeclineInviteToGuild(
	ctx context.Context,
	guildID uint32,
) (*protocol.CMsgClientToGCDeclineInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCDeclineInviteToGuild{
		GuildId: &guildID,
	}
	resp := &protocol.CMsgClientToGCDeclineInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeclineInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDeclineInviteToGuildResponse),
		resp,
	)
}

// SendDevDeleteEventActions sends dev delete event actions.
//
// Sends the GC message k_EMsgDevDeleteEventActions (CMsgDevDeleteEventActions) and awaits the response k_EMsgDevDeleteEventActionsResponse,
// delivered as *CMsgDevDeleteEventActionsResponse.
func (d *Dota2) SendDevDeleteEventActions(
	ctx context.Context,
	eventID protocol.EEvent,
	startActionID uint32,
	endActionID uint32,
	removeAudit bool,
) (*protocol.CMsgDevDeleteEventActionsResponse, error) {
	req := &protocol.CMsgDevDeleteEventActions{
		EventId:       &eventID,
		StartActionId: &startActionID,
		EndActionId:   &endActionID,
		RemoveAudit:   &removeAudit,
	}
	resp := &protocol.CMsgDevDeleteEventActionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevDeleteEventActions),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevDeleteEventActionsResponse),
		resp,
	)
}

// SendDevReloadAllEvents sends dev reload all events.
//
// Sends the GC message k_EMsgDevReloadAllEvents (CMsgDevReloadAllEvents) and awaits the response k_EMsgDevReloadAllEventsResponse,
// delivered as *CMsgDevReloadAllEventsResponse.
func (d *Dota2) SendDevReloadAllEvents(
	ctx context.Context,
) (*protocol.CMsgDevReloadAllEventsResponse, error) {
	req := &protocol.CMsgDevReloadAllEvents{}
	resp := &protocol.CMsgDevReloadAllEventsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevReloadAllEvents),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevReloadAllEventsResponse),
		resp,
	)
}

// SendDevResetEventState sends a dev reset event state.
//
// Sends the GC message k_EMsgDevResetEventState (CMsgDevResetEventState) and awaits the response k_EMsgDevResetEventStateResponse,
// delivered as *CMsgDevResetEventStateResponse.
func (d *Dota2) SendDevResetEventState(
	ctx context.Context,
	eventID protocol.EEvent,
	removeAudit bool,
) (*protocol.CMsgDevResetEventStateResponse, error) {
	req := &protocol.CMsgDevResetEventState{
		EventId:     &eventID,
		RemoveAudit: &removeAudit,
	}
	resp := &protocol.CMsgDevResetEventStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevResetEventState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgDevResetEventStateResponse),
		resp,
	)
}

// SendDotaLabsFeedback sends a dota labs feedback.
//
// Sends the GC message k_EMsgClientToGCDotaLabsFeedback (CMsgClientToGCDotaLabsFeedback) and awaits the response k_EMsgClientToGCDotaLabsFeedbackResponse,
// delivered as *CMsgClientToGCDotaLabsFeedbackResponse.
func (d *Dota2) SendDotaLabsFeedback(
	ctx context.Context,
	language uint32,
	feedbackItem uint32,
	feedback string,
) (*protocol.CMsgClientToGCDotaLabsFeedbackResponse, error) {
	req := &protocol.CMsgClientToGCDotaLabsFeedback{
		Language:     &language,
		FeedbackItem: &feedbackItem,
		Feedback:     &feedback,
	}
	resp := &protocol.CMsgClientToGCDotaLabsFeedbackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDotaLabsFeedback),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCDotaLabsFeedbackResponse),
		resp,
	)
}

// SendFantasyCraftingDevModifyTablet sends a fantasy crafting dev modify tablet.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingDevModifyTablet (CMsgClientToGCFantasyCraftingDevModifyTablet) and awaits the response k_EMsgClientToGCFantasyCraftingDevModifyTabletResponse,
// delivered as *CMsgClientToGCFantasyCraftingDevModifyTabletResponse.
func (d *Dota2) SendFantasyCraftingDevModifyTablet(
	ctx context.Context,
	fantasyLeague uint32,
	resetTablet bool,
	modifyTokens uint32,
	upgradeTablets bool,
	fantasyPeriod uint32,
) (*protocol.CMsgClientToGCFantasyCraftingDevModifyTabletResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingDevModifyTablet{
		FantasyLeague:  &fantasyLeague,
		ResetTablet:    &resetTablet,
		ModifyTokens:   &modifyTokens,
		UpgradeTablets: &upgradeTablets,
		FantasyPeriod:  &fantasyPeriod,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingDevModifyTabletResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingDevModifyTablet),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingDevModifyTabletResponse),
		resp,
	)
}

// SendFantasyCraftingGenerateTablets sends fantasy crafting generate tablets.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingGenerateTablets (CMsgClientToGCFantasyCraftingGenerateTablets) and awaits the response k_EMsgClientToGCFantasyCraftingGenerateTabletsResponse,
// delivered as *CMsgClientToGCFantasyCraftingGenerateTabletsResponse.
func (d *Dota2) SendFantasyCraftingGenerateTablets(
	ctx context.Context,
	fantasyLeague uint32,
	accountIDs []uint32,
	selectedTeams []*protocol.CMsgClientToGCFantasyCraftingGenerateTablets_TeamChoice,
) (*protocol.CMsgClientToGCFantasyCraftingGenerateTabletsResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingGenerateTablets{
		FantasyLeague: &fantasyLeague,
		AccountIds:    accountIDs,
		SelectedTeams: selectedTeams,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingGenerateTabletsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingGenerateTablets),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingGenerateTabletsResponse),
		resp,
	)
}

// SendFantasyCraftingPerformOperation sends a fantasy crafting perform operation.
//
// Sends the GC message k_EMsgClientToGCFantasyCraftingPerformOperation (CMsgClientToGCFantasyCraftingPerformOperation) and awaits the response k_EMsgClientToGCFantasyCraftingPerformOperationResponse,
// delivered as *CMsgClientToGCFantasyCraftingPerformOperationResponse.
func (d *Dota2) SendFantasyCraftingPerformOperation(
	ctx context.Context,
	fantasyLeague uint32,
	tabletID uint32,
	operationID uint32,
	extraData uint64,
) (*protocol.CMsgClientToGCFantasyCraftingPerformOperationResponse, error) {
	req := &protocol.CMsgClientToGCFantasyCraftingPerformOperation{
		FantasyLeague: &fantasyLeague,
		TabletId:      &tabletID,
		OperationId:   &operationID,
		ExtraData:     &extraData,
	}
	resp := &protocol.CMsgClientToGCFantasyCraftingPerformOperationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingPerformOperation),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFantasyCraftingPerformOperationResponse),
		resp,
	)
}

// SendFightingGameAnswerChallenge sends a fighting game answer challenge.
//
// Sends the GC message k_EMsgClientToGCFightingGameAnswerChallenge (CMsgClientToGCFightingGameAnswerChallenge) and awaits the response k_EMsgClientToGCFightingGameAnswerChallengeResponse,
// delivered as *CMsgClientToGCFightingGameAnswerChallengeResponse.
func (d *Dota2) SendFightingGameAnswerChallenge(
	ctx context.Context,
	challengerAccountID uint32,
	accept bool,
) (*protocol.CMsgClientToGCFightingGameAnswerChallengeResponse, error) {
	req := &protocol.CMsgClientToGCFightingGameAnswerChallenge{
		ChallengerAccountId: &challengerAccountID,
		Accept:              &accept,
	}
	resp := &protocol.CMsgClientToGCFightingGameAnswerChallengeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFightingGameAnswerChallenge),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFightingGameAnswerChallengeResponse),
		resp,
	)
}

// SendFightingGameChallengeFriend sends a fighting game challenge friend.
//
// Sends the GC message k_EMsgClientToGCFightingGameChallengeFriend (CMsgClientToGCFightingGameChallengeFriend) and awaits the response k_EMsgClientToGCFightingGameChallengeFriendResponse,
// delivered as *CMsgClientToGCFightingGameChallengeFriendResponse.
func (d *Dota2) SendFightingGameChallengeFriend(
	ctx context.Context,
	friendAccountID uint32,
) (*protocol.CMsgClientToGCFightingGameChallengeFriendResponse, error) {
	req := &protocol.CMsgClientToGCFightingGameChallengeFriend{
		FriendAccountId: &friendAccountID,
	}
	resp := &protocol.CMsgClientToGCFightingGameChallengeFriendResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFightingGameChallengeFriend),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCFightingGameChallengeFriendResponse),
		resp,
	)
}

// SendH264Unsupported sends a h 264 unsupported.
//
// Sends the GC message k_EMsgClientToGCH264Unsupported (CMsgClientToGCH264Unsupported). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendH264Unsupported() {
	req := &protocol.CMsgClientToGCH264Unsupported{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCH264Unsupported), req)
}

// SendHasPlayerVotedForMVP sends a has player voted for mvp.
//
// Sends the GC message k_EMsgClientToGCHasPlayerVotedForMVP (CMsgClientToGCHasPlayerVotedForMVP) and awaits the response k_EMsgClientToGCHasPlayerVotedForMVPResponse,
// delivered as *CMsgClientToGCHasPlayerVotedForMVPResponse.
func (d *Dota2) SendHasPlayerVotedForMVP(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCHasPlayerVotedForMVPResponse, error) {
	req := &protocol.CMsgClientToGCHasPlayerVotedForMVP{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCHasPlayerVotedForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVP),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVPResponse),
		resp,
	)
}

// SendInitialQuestionnaireResponse sends a initial questionnaire response.
//
// Sends the GC message k_EMsgGCInitialQuestionnaireResponse (CMsgInitialQuestionnaireResponse). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendInitialQuestionnaireResponse(
	initialSkill uint32,
) {
	req := &protocol.CMsgInitialQuestionnaireResponse{
		InitialSkill: &initialSkill,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse), req)
}

// SendInviteToDemoMode sends a invite to demo mode.
//
// Sends the GC message k_EMsgClientToGCInviteToDemoMode (CMsgClientToGCInviteToDemoMode). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendInviteToDemoMode(
	serverID uint64,
	invitedPlayerID uint64,
) {
	req := &protocol.CMsgClientToGCInviteToDemoMode{
		ServerId:        &serverID,
		InvitedPlayerId: &invitedPlayerID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCInviteToDemoMode), req)
}

// SendInviteToGuild sends a invite to guild.
//
// Sends the GC message k_EMsgClientToGCInviteToGuild (CMsgClientToGCInviteToGuild) and awaits the response k_EMsgClientToGCInviteToGuildResponse,
// delivered as *CMsgClientToGCInviteToGuildResponse.
func (d *Dota2) SendInviteToGuild(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*protocol.CMsgClientToGCInviteToGuildResponse, error) {
	req := &protocol.CMsgClientToGCInviteToGuild{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgClientToGCInviteToGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCInviteToGuild),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCInviteToGuildResponse),
		resp,
	)
}

// SendItemBattlerGameAction sends a item battler game action.
//
// Sends the GC message k_EMsgClientToGCItemBattlerGameAction (CMsgClientToGCItemBattlerGameAction) and awaits the response k_EMsgClientToGCItemBattlerGameActionResponse,
// delivered as *CMsgClientToGCItemBattlerGameActionResponse.
func (d *Dota2) SendItemBattlerGameAction(
	ctx context.Context,
	action protocol.CMsgClientToGCItemBattlerGameAction_EAction,
	choiceIndex uint32,
	itemInstanceID uint32,
	itemContainerID uint32,
	itemPositionX uint32,
	itemPositionY uint32,
) (*protocol.CMsgClientToGCItemBattlerGameActionResponse, error) {
	req := &protocol.CMsgClientToGCItemBattlerGameAction{
		Action:          &action,
		ChoiceIndex:     &choiceIndex,
		ItemInstanceId:  &itemInstanceID,
		ItemContainerId: &itemContainerID,
		ItemPositionX:   &itemPositionX,
		ItemPositionY:   &itemPositionY,
	}
	resp := &protocol.CMsgClientToGCItemBattlerGameActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerGameAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCItemBattlerGameActionResponse),
		resp,
	)
}

// SendLatestConductScorecard sends a latest conduct scorecard.
//
// Sends the GC message k_EMsgClientToGCLatestConductScorecard (CMsgPlayerConductScorecard). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLatestConductScorecard(
	req *protocol.CMsgPlayerConductScorecard,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard), req)
}

// SendLeagueAvailableLobbyNodes sends league available lobby nodes.
//
// Sends the GC message k_EMsgDOTALeagueAvailableLobbyNodes (CMsgDOTALeagueAvailableLobbyNodes). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLeagueAvailableLobbyNodes(
	nodeInfos []*protocol.CMsgDOTALeagueAvailableLobbyNodes_NodeInfo,
) {
	req := &protocol.CMsgDOTALeagueAvailableLobbyNodes{
		NodeInfos: nodeInfos,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTALeagueAvailableLobbyNodes), req)
}

// SendLobbyBattleCupVictory sends a lobby battle cup victory.
//
// Sends the GC message k_EMsgLobbyBattleCupVictory (CMsgBattleCupVictory). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyBattleCupVictory(
	req *protocol.CMsgBattleCupVictory,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory), req)
}

// SendLobbyEventGameData sends a lobby event game data.
//
// Sends the GC message k_EMsgLobbyEventGameData (CMsgLobbyEventGameData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyEventGameData(
	gameSeed uint32,
	eventWindowStartTime uint32,
) {
	req := &protocol.CMsgLobbyEventGameData{
		GameSeed:             &gameSeed,
		EventWindowStartTime: &eventWindowStartTime,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventGameData), req)
}

// SendLobbyEventGameDetails sends lobby event game details.
//
// Sends the GC message k_EMsgLobbyEventGameDetails (CMsgLobbyEventGameDetails). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyEventGameDetails(
	kvData []byte,
) {
	req := &protocol.CMsgLobbyEventGameDetails{
		KvData: kvData,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventGameDetails), req)
}

// SendLobbyEventPoints sends lobby event points.
//
// Sends the GC message k_EMsgLobbyEventPoints (CMsgLobbyEventPoints). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyEventPoints(
	eventID uint32,
	accountPoints []*protocol.CMsgLobbyEventPoints_AccountPoints,
) {
	req := &protocol.CMsgLobbyEventPoints{
		EventId:       &eventID,
		AccountPoints: accountPoints,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyEventPoints), req)
}

// SendLobbyFeaturedGamemodeProgress sends lobby featured gamemode progress.
//
// Sends the GC message k_EMsgLobbyFeaturedGamemodeProgress (CMsgLobbyFeaturedGamemodeProgress). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyFeaturedGamemodeProgress(
	accounts []*protocol.CMsgLobbyFeaturedGamemodeProgress_AccountProgress,
) {
	req := &protocol.CMsgLobbyFeaturedGamemodeProgress{
		Accounts: accounts,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyFeaturedGamemodeProgress), req)
}

// SendLobbyPlaytestDetails sends lobby playtest details.
//
// Sends the GC message k_EMsgLobbyPlaytestDetails (CMsgLobbyPlaytestDetails). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyPlaytestDetails(
	jSON string,
) {
	req := &protocol.CMsgLobbyPlaytestDetails{
		Json: &jSON,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyPlaytestDetails), req)
}

// SendLobbyRoadToTIMatchQuestData sends a lobby road to ti match quest data.
//
// Sends the GC message k_EMsgLobbyRoadToTIMatchQuestData (CMsgLobbyRoadToTIMatchQuestData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendLobbyRoadToTIMatchQuestData(
	questData protocol.CMsgRoadToTIAssignedQuest,
	questPeriod uint32,
	questNumber uint32,
) {
	req := &protocol.CMsgLobbyRoadToTIMatchQuestData{
		QuestData:   &questData,
		QuestPeriod: &questPeriod,
		QuestNumber: &questNumber,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgLobbyRoadToTIMatchQuestData), req)
}

// SendMMInfo sends a mm info.
//
// Sends the GC message k_EMsgClientToGCMMInfo (CMsgClientToGCMMInfo). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendMMInfo(
	laneSelectionFlags uint32,
	highPriorityDisabled bool,
) {
	req := &protocol.CMsgClientToGCMMInfo{
		LaneSelectionFlags:   &laneSelectionFlags,
		HighPriorityDisabled: &highPriorityDisabled,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMMInfo), req)
}

// SendManageFavorites sends manage favorites.
//
// Sends the GC message k_EMsgClientToGCManageFavorites (CMsgClientToGCManageFavorites) and awaits the response k_EMsgGCToClientManageFavoritesResponse,
// delivered as *CMsgGCToClientManageFavoritesResponse.
func (d *Dota2) SendManageFavorites(
	ctx context.Context,
	action protocol.CMsgClientToGCManageFavorites_Action,
	accountID uint32,
	favoriteName string,
	inviteResponse bool,
	fromFriendlist bool,
	lobbyID uint64,
) (*protocol.CMsgGCToClientManageFavoritesResponse, error) {
	req := &protocol.CMsgClientToGCManageFavorites{
		Action:         &action,
		AccountId:      &accountID,
		FavoriteName:   &favoriteName,
		InviteResponse: &inviteResponse,
		FromFriendlist: &fromFriendlist,
		LobbyId:        &lobbyID,
	}
	resp := &protocol.CMsgGCToClientManageFavoritesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCManageFavorites),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientManageFavoritesResponse),
		resp,
	)
}

// SendMatchMatchmakingStats sends match matchmaking stats.
//
// Sends the GC message k_EMsgMatchMatchmakingStats (CMsgMatchMatchmakingStats). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendMatchMatchmakingStats(
	averageQueueTime uint32,
	maximumQueueTime uint32,
	behaviorScoreVariance protocol.EMatchBehaviorScoreVariance,
) {
	req := &protocol.CMsgMatchMatchmakingStats{
		AverageQueueTime:      &averageQueueTime,
		MaximumQueueTime:      &maximumQueueTime,
		BehaviorScoreVariance: &behaviorScoreVariance,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgMatchMatchmakingStats), req)
}

// SendMergePartyInvite sends a merge party invite.
//
// Sends the GC message k_EMsgClientToGCMergePartyInvite (CMsgDOTAGroupMergeInvite). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendMergePartyInvite(
	otherGroupID uint64,
) {
	req := &protocol.CMsgDOTAGroupMergeInvite{
		OtherGroupId: &otherGroupID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite), req)
}

// SendModifyGuildRole sends a modify guild role.
//
// Sends the GC message k_EMsgClientToGCModifyGuildRole (CMsgClientToGCModifyGuildRole) and awaits the response k_EMsgClientToGCModifyGuildRoleResponse,
// delivered as *CMsgClientToGCModifyGuildRoleResponse.
func (d *Dota2) SendModifyGuildRole(
	ctx context.Context,
	guildID uint32,
	roleID uint32,
	roleName string,
	roleFlags uint32,
) (*protocol.CMsgClientToGCModifyGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCModifyGuildRole{
		GuildId:   &guildID,
		RoleId:    &roleID,
		RoleName:  &roleName,
		RoleFlags: &roleFlags,
	}
	resp := &protocol.CMsgClientToGCModifyGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCModifyGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCModifyGuildRoleResponse),
		resp,
	)
}

// SendMonsterHunterDevClearInventory sends a monster hunter dev clear inventory.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterDevClearInventory (CMsgClientToGCMonsterHunterDevClearInventory) and awaits the response k_EMsgClientToGCMonsterHunterDevClearInventoryResponse,
// delivered as *CMsgClientToGCMonsterHunterDevClearInventoryResponse.
func (d *Dota2) SendMonsterHunterDevClearInventory(
	ctx context.Context,
) (*protocol.CMsgClientToGCMonsterHunterDevClearInventoryResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterDevClearInventory{}
	resp := &protocol.CMsgClientToGCMonsterHunterDevClearInventoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevClearInventory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevClearInventoryResponse),
		resp,
	)
}

// SendMonsterHunterDevModifyHeroCodex sends a monster hunter dev modify hero codex.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterDevModifyHeroCodex (CMsgClientToGCMonsterHunterDevModifyHeroCodex) and awaits the response k_EMsgClientToGCMonsterHunterDevModifyHeroCodexResponse,
// delivered as *CMsgClientToGCMonsterHunterDevModifyHeroCodexResponse.
func (d *Dota2) SendMonsterHunterDevModifyHeroCodex(
	ctx context.Context,
	actions []*protocol.CMsgDevModifyCodexAction,
) (*protocol.CMsgClientToGCMonsterHunterDevModifyHeroCodexResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterDevModifyHeroCodex{
		Actions: actions,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterDevModifyHeroCodexResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevModifyHeroCodex),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevModifyHeroCodexResponse),
		resp,
	)
}

// SendMonsterHunterDevResetAll sends a monster hunter dev reset all.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterDevResetAll (CMsgClientToGCMonsterHunterDevResetAll) and awaits the response k_EMsgClientToGCMonsterHunterDevResetAllResponse,
// delivered as *CMsgClientToGCMonsterHunterDevResetAllResponse.
func (d *Dota2) SendMonsterHunterDevResetAll(
	ctx context.Context,
	resetCodexOnly bool,
) (*protocol.CMsgClientToGCMonsterHunterDevResetAllResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterDevResetAll{
		ResetCodexOnly: &resetCodexOnly,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterDevResetAllResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevResetAll),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterDevResetAllResponse),
		resp,
	)
}

// SendMonsterHunterGiftMaterials sends monster hunter gift materials.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterGiftMaterials (CMsgClientToGCMonsterHunterGiftMaterials) and awaits the response k_EMsgClientToGCMonsterHunterGiftMaterialsResponse,
// delivered as *CMsgClientToGCMonsterHunterGiftMaterialsResponse.
func (d *Dota2) SendMonsterHunterGiftMaterials(
	ctx context.Context,
	tokenGift protocol.CMsgMonsterHunterMaterialCount,
	recipientAccountID uint32,
	periodicResourceID uint32,
) (*protocol.CMsgClientToGCMonsterHunterGiftMaterialsResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterGiftMaterials{
		TokenGift:          &tokenGift,
		RecipientAccountId: &recipientAccountID,
		PeriodicResourceId: &periodicResourceID,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterGiftMaterialsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterGiftMaterials),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterGiftMaterialsResponse),
		resp,
	)
}

// SendMonsterHunterTradeMaterials sends monster hunter trade materials.
//
// Sends the GC message k_EMsgClientToGCMonsterHunterTradeMaterials (CMsgClientToGCMonsterHunterTradeMaterials) and awaits the response k_EMsgClientToGCMonsterHunterTradeMaterialsResponse,
// delivered as *CMsgClientToGCMonsterHunterTradeMaterialsResponse.
func (d *Dota2) SendMonsterHunterTradeMaterials(
	ctx context.Context,
	materialOffer protocol.CMsgMonsterHunterMaterialQuantity,
	materialRequest protocol.CMsgMonsterHunterMaterialQuantity,
	recipeID uint32,
) (*protocol.CMsgClientToGCMonsterHunterTradeMaterialsResponse, error) {
	req := &protocol.CMsgClientToGCMonsterHunterTradeMaterials{
		MaterialOffer:   &materialOffer,
		MaterialRequest: &materialRequest,
		RecipeId:        &recipeID,
	}
	resp := &protocol.CMsgClientToGCMonsterHunterTradeMaterialsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterTradeMaterials),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterTradeMaterialsResponse),
		resp,
	)
}

// SendNeutralItemStats sends neutral item stats.
//
// Sends the GC message k_EMsgNeutralItemStats (CMsgNeutralItemStats). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendNeutralItemStats(
	neutralItems []*protocol.CMsgNeutralItemStats_NeutralItem,
) {
	req := &protocol.CMsgNeutralItemStats{
		NeutralItems: neutralItems,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgNeutralItemStats), req)
}

// SendNewBloomGift sends a new bloom gift.
//
// Sends the GC message k_EMsgClientToGCNewBloomGift (CMsgClientToGCNewBloomGift) and awaits the response k_EMsgClientToGCNewBloomGiftResponse,
// delivered as *CMsgClientToGCNewBloomGiftResponse.
func (d *Dota2) SendNewBloomGift(
	ctx context.Context,
	defindex uint32,
	lobbyID uint64,
	targetAccountIDs []uint32,
) (*protocol.CMsgClientToGCNewBloomGiftResponse, error) {
	req := &protocol.CMsgClientToGCNewBloomGift{
		Defindex:         &defindex,
		LobbyId:          &lobbyID,
		TargetAccountIds: targetAccountIDs,
	}
	resp := &protocol.CMsgClientToGCNewBloomGiftResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCNewBloomGift),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCNewBloomGiftResponse),
		resp,
	)
}

// SendOverwatchReplayError sends a overwatch replay error.
//
// Sends the GC message k_EMsgClientToGCOverwatchReplayError (CMsgClientToGCOverwatchReplayError). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverwatchReplayError(
	overwatchReplayID uint64,
) {
	req := &protocol.CMsgClientToGCOverwatchReplayError{
		OverwatchReplayId: &overwatchReplayID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverwatchReplayError), req)
}

// SendOverworldCompletePath sends a overworld complete path.
//
// Sends the GC message k_EMsgClientToGCOverworldCompletePath (CMsgClientToGCOverworldCompletePath) and awaits the response k_EMsgClientToGCOverworldCompletePathResponse,
// delivered as *CMsgClientToGCOverworldCompletePathResponse.
func (d *Dota2) SendOverworldCompletePath(
	ctx context.Context,
	overworldID uint32,
	pathID uint32,
	usePathUnlocker bool,
	devIgnoreReleaseSchedule bool,
) (*protocol.CMsgClientToGCOverworldCompletePathResponse, error) {
	req := &protocol.CMsgClientToGCOverworldCompletePath{
		OverworldId:              &overworldID,
		PathId:                   &pathID,
		UsePathUnlocker:          &usePathUnlocker,
		DevIgnoreReleaseSchedule: &devIgnoreReleaseSchedule,
	}
	resp := &protocol.CMsgClientToGCOverworldCompletePathResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldCompletePath),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldCompletePathResponse),
		resp,
	)
}

// SendOverworldDevClearFortune sends a overworld dev clear fortune.
//
// Sends the GC message k_EMsgClientToGCOverworldDevClearFortune (CMsgClientToGCOverworldDevClearFortune) and awaits the response k_EMsgClientToGCOverworldDevClearFortuneResponse,
// delivered as *CMsgClientToGCOverworldDevClearFortuneResponse.
func (d *Dota2) SendOverworldDevClearFortune(
	ctx context.Context,
	overworldID uint32,
	fortuneID uint32,
) (*protocol.CMsgClientToGCOverworldDevClearFortuneResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevClearFortune{
		OverworldId: &overworldID,
		FortuneId:   &fortuneID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevClearFortuneResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevClearFortune),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevClearFortuneResponse),
		resp,
	)
}

// SendOverworldDevClearInventory sends a overworld dev clear inventory.
//
// Sends the GC message k_EMsgClientToGCOverworldDevClearInventory (CMsgClientToGCOverworldDevClearInventory) and awaits the response k_EMsgClientToGCOverworldDevClearInventoryResponse,
// delivered as *CMsgClientToGCOverworldDevClearInventoryResponse.
func (d *Dota2) SendOverworldDevClearInventory(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldDevClearInventoryResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevClearInventory{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevClearInventoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevClearInventory),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevClearInventoryResponse),
		resp,
	)
}

// SendOverworldDevResetAll sends a overworld dev reset all.
//
// Sends the GC message k_EMsgClientToGCOverworldDevResetAll (CMsgClientToGCOverworldDevResetAll) and awaits the response k_EMsgClientToGCOverworldDevResetAllResponse,
// delivered as *CMsgClientToGCOverworldDevResetAllResponse.
func (d *Dota2) SendOverworldDevResetAll(
	ctx context.Context,
	overworldID uint32,
) (*protocol.CMsgClientToGCOverworldDevResetAllResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevResetAll{
		OverworldId: &overworldID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevResetAllResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevResetAll),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevResetAllResponse),
		resp,
	)
}

// SendOverworldDevResetNode sends a overworld dev reset node.
//
// Sends the GC message k_EMsgClientToGCOverworldDevResetNode (CMsgClientToGCOverworldDevResetNode) and awaits the response k_EMsgClientToGCOverworldDevResetNodeResponse,
// delivered as *CMsgClientToGCOverworldDevResetNodeResponse.
func (d *Dota2) SendOverworldDevResetNode(
	ctx context.Context,
	overworldID uint32,
	nodeID uint32,
) (*protocol.CMsgClientToGCOverworldDevResetNodeResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevResetNode{
		OverworldId: &overworldID,
		NodeId:      &nodeID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevResetNodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevResetNode),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevResetNodeResponse),
		resp,
	)
}

// SendOverworldEncounterChooseHeroData sends a overworld encounter choose hero data.
//
// Sends the GC message k_EMsgOverworldEncounterChooseHeroData (CMsgOverworldEncounterChooseHeroData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverworldEncounterChooseHeroData(
	heroList protocol.CMsgOverworldHeroList,
	additive bool,
) {
	req := &protocol.CMsgOverworldEncounterChooseHeroData{
		HeroList: &heroList,
		Additive: &additive,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgOverworldEncounterChooseHeroData), req)
}

// SendOverworldEncounterPitFighterRewardData sends a overworld encounter pit fighter reward data.
//
// Sends the GC message k_EMsgOverworldEncounterPitFighterRewardData (CMsgOverworldEncounterPitFighterRewardData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverworldEncounterPitFighterRewardData(
	tokenID uint32,
	choice uint32,
) {
	req := &protocol.CMsgOverworldEncounterPitFighterRewardData{
		TokenId: &tokenID,
		Choice:  &choice,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgOverworldEncounterPitFighterRewardData), req)
}

// SendOverworldEncounterProgressData sends a overworld encounter progress data.
//
// Sends the GC message k_EMsgOverworldEncounterProgressData (CMsgOverworldEncounterProgressData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverworldEncounterProgressData(
	choice int32,
	progress int32,
	maxProgress int32,
	visited bool,
) {
	req := &protocol.CMsgOverworldEncounterProgressData{
		Choice:      &choice,
		Progress:    &progress,
		MaxProgress: &maxProgress,
		Visited:     &visited,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgOverworldEncounterProgressData), req)
}

// SendOverworldEncounterTokenQuestData sends a overworld encounter token quest data.
//
// Sends the GC message k_EMsgOverworldEncounterTokenQuestData (CMsgOverworldEncounterTokenQuestData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverworldEncounterTokenQuestData(
	quests []*protocol.CMsgOverworldEncounterTokenQuestData_Quest,
) {
	req := &protocol.CMsgOverworldEncounterTokenQuestData{
		Quests: quests,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgOverworldEncounterTokenQuestData), req)
}

// SendOverworldEncounterTokenTreasureData sends a overworld encounter token treasure data.
//
// Sends the GC message k_EMsgOverworldEncounterTokenTreasureData (CMsgOverworldEncounterTokenTreasureData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendOverworldEncounterTokenTreasureData(
	rewardOptions []*protocol.CMsgOverworldEncounterTokenTreasureData_RewardOption,
) {
	req := &protocol.CMsgOverworldEncounterTokenTreasureData{
		RewardOptions: rewardOptions,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgOverworldEncounterTokenTreasureData), req)
}

// SendOverworldFeedback sends a overworld feedback.
//
// Sends the GC message k_EMsgClientToGCOverworldFeedback (CMsgClientToGCOverworldFeedback) and awaits the response k_EMsgClientToGCOverworldFeedbackResponse,
// delivered as *CMsgClientToGCOverworldFeedbackResponse.
func (d *Dota2) SendOverworldFeedback(
	ctx context.Context,
	language uint32,
	overworldID uint32,
	feedback string,
) (*protocol.CMsgClientToGCOverworldFeedbackResponse, error) {
	req := &protocol.CMsgClientToGCOverworldFeedback{
		Language:    &language,
		OverworldId: &overworldID,
		Feedback:    &feedback,
	}
	resp := &protocol.CMsgClientToGCOverworldFeedbackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldFeedback),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldFeedbackResponse),
		resp,
	)
}

// SendOverworldGiftTokens sends overworld gift tokens.
//
// Sends the GC message k_EMsgClientToGCOverworldGiftTokens (CMsgClientToGCOverworldGiftTokens) and awaits the response k_EMsgClientToGCOverworldGiftTokensResponse,
// delivered as *CMsgClientToGCOverworldGiftTokensResponse.
func (d *Dota2) SendOverworldGiftTokens(
	ctx context.Context,
	overworldID uint32,
	tokenGift protocol.CMsgOverworldTokenCount,
	recipientAccountID uint32,
	periodicResourceID uint32,
) (*protocol.CMsgClientToGCOverworldGiftTokensResponse, error) {
	req := &protocol.CMsgClientToGCOverworldGiftTokens{
		OverworldId:        &overworldID,
		TokenGift:          &tokenGift,
		RecipientAccountId: &recipientAccountID,
		PeriodicResourceId: &periodicResourceID,
	}
	resp := &protocol.CMsgClientToGCOverworldGiftTokensResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGiftTokens),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldGiftTokensResponse),
		resp,
	)
}

// SendOverworldMinigameAction sends a overworld minigame action.
//
// Sends the GC message k_EMsgClientToGCOverworldMinigameAction (CMsgClientToGCOverworldMinigameAction) and awaits the response k_EMsgClientToGCOverworldMinigameActionResponse,
// delivered as *CMsgClientToGCOverworldMinigameActionResponse.
func (d *Dota2) SendOverworldMinigameAction(
	ctx context.Context,
	overworldID uint32,
	nodeID uint32,
	action protocol.EOverworldMinigameAction,
	selection uint32,
	optionValue uint32,
	currencyAmount uint32,
) (*protocol.CMsgClientToGCOverworldMinigameActionResponse, error) {
	req := &protocol.CMsgClientToGCOverworldMinigameAction{
		OverworldId:    &overworldID,
		NodeId:         &nodeID,
		Action:         &action,
		Selection:      &selection,
		OptionValue:    &optionValue,
		CurrencyAmount: &currencyAmount,
	}
	resp := &protocol.CMsgClientToGCOverworldMinigameActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldMinigameAction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldMinigameActionResponse),
		resp,
	)
}

// SendOverworldMoveToNode sends a overworld move to node.
//
// Sends the GC message k_EMsgClientToGCOverworldMoveToNode (CMsgClientToGCOverworldMoveToNode) and awaits the response k_EMsgClientToGCOverworldMoveToNodeResponse,
// delivered as *CMsgClientToGCOverworldMoveToNodeResponse.
func (d *Dota2) SendOverworldMoveToNode(
	ctx context.Context,
	overworldID uint32,
	nodeID uint32,
) (*protocol.CMsgClientToGCOverworldMoveToNodeResponse, error) {
	req := &protocol.CMsgClientToGCOverworldMoveToNode{
		OverworldId: &overworldID,
		NodeId:      &nodeID,
	}
	resp := &protocol.CMsgClientToGCOverworldMoveToNodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldMoveToNode),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldMoveToNodeResponse),
		resp,
	)
}

// SendOverworldTradeTokens sends overworld trade tokens.
//
// Sends the GC message k_EMsgClientToGCOverworldTradeTokens (CMsgClientToGCOverworldTradeTokens) and awaits the response k_EMsgClientToGCOverworldTradeTokensResponse,
// delivered as *CMsgClientToGCOverworldTradeTokensResponse.
func (d *Dota2) SendOverworldTradeTokens(
	ctx context.Context,
	overworldID uint32,
	tokenOffer protocol.CMsgOverworldTokenQuantity,
	tokenRequest protocol.CMsgOverworldTokenQuantity,
	recipe uint32,
	encounterID uint32,
) (*protocol.CMsgClientToGCOverworldTradeTokensResponse, error) {
	req := &protocol.CMsgClientToGCOverworldTradeTokens{
		OverworldId:  &overworldID,
		TokenOffer:   &tokenOffer,
		TokenRequest: &tokenRequest,
		Recipe:       &recipe,
		EncounterId:  &encounterID,
	}
	resp := &protocol.CMsgClientToGCOverworldTradeTokensResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldTradeTokens),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldTradeTokensResponse),
		resp,
	)
}

// SendOverworldVisitEncounter sends a overworld visit encounter.
//
// Sends the GC message k_EMsgClientToGCOverworldVisitEncounter (CMsgClientToGCOverworldVisitEncounter) and awaits the response k_EMsgClientToGCOverworldVisitEncounterResponse,
// delivered as *CMsgClientToGCOverworldVisitEncounterResponse.
func (d *Dota2) SendOverworldVisitEncounter(
	ctx context.Context,
	overworldID uint32,
	nodeID uint32,
) (*protocol.CMsgClientToGCOverworldVisitEncounterResponse, error) {
	req := &protocol.CMsgClientToGCOverworldVisitEncounter{
		OverworldId: &overworldID,
		NodeId:      &nodeID,
	}
	resp := &protocol.CMsgClientToGCOverworldVisitEncounterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldVisitEncounter),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldVisitEncounterResponse),
		resp,
	)
}

// SendPartyReadyCheck starts a ready check across the whole party.
func (d *Dota2) SendPartyReadyCheck(
	ctx context.Context,
) (*protocol.CMsgPartyReadyCheckResponse, error) {
	req := &protocol.CMsgPartyReadyCheckRequest{}
	resp := &protocol.CMsgPartyReadyCheckResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckResponse),
		resp,
	)
}

// SendPeriodicResourceUpdated sends a periodic resource updated.
//
// Sends the GC message k_EMsgDOTAPeriodicResourceUpdated (CMsgDOTAPeriodicResourceUpdated). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendPeriodicResourceUpdated(
	periodicResourceKey protocol.CMsgDOTAGetPeriodicResource,
	periodicResourceValue protocol.CMsgDOTAGetPeriodicResourceResponse,
) {
	req := &protocol.CMsgDOTAPeriodicResourceUpdated{
		PeriodicResourceKey:   &periodicResourceKey,
		PeriodicResourceValue: &periodicResourceValue,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTAPeriodicResourceUpdated), req)
}

// SendPingData sends a ping data.
//
// Sends the GC message k_EMsgClientToGCPingData (CMsgClientPingData). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendPingData(
	relayCodes []uint32,
	relayPings []uint32,
	regionCodes []uint32,
	regionPings []uint32,
	regionPingFailedBitmask uint32,
) {
	req := &protocol.CMsgClientPingData{
		RelayCodes:              relayCodes,
		RelayPings:              relayPings,
		RegionCodes:             regionCodes,
		RegionPings:             regionPings,
		RegionPingFailedBitmask: &regionPingFailedBitmask,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCPingData), req)
}

// SendProfileUpdate sends a profile update.
//
// Sends the GC message k_EMsgProfileUpdate (CMsgProfileUpdate) and awaits the response k_EMsgProfileUpdateResponse,
// delivered as *CMsgProfileUpdateResponse.
func (d *Dota2) SendProfileUpdate(
	ctx context.Context,
	backgroundItemID uint64,
	featuredHeroIDs []int32,
) (*protocol.CMsgProfileUpdateResponse, error) {
	req := &protocol.CMsgProfileUpdate{
		BackgroundItemId: &backgroundItemID,
		FeaturedHeroIds:  featuredHeroIDs,
	}
	resp := &protocol.CMsgProfileUpdateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileUpdate),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgProfileUpdateResponse),
		resp,
	)
}

// SendReadyUp accepts (or declines) an incoming match. The GC assigns teams
// and broadcasts the game setup state once all players accept.
//
// A player who fails to ready up in time causes the match search to restart.
func (d *Dota2) SendReadyUp(
	state protocol.DOTALobbyReadyState,
	readyUpKey uint64,
	hardwareSpecs protocol.CDOTAClientHardwareSpecs,
) {
	req := &protocol.CMsgReadyUp{
		State:         &state,
		ReadyUpKey:    &readyUpKey,
		HardwareSpecs: &hardwareSpecs,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCReadyUp), req)
}

// SendRecalibrateMMR sends a recalibrate mmr.
//
// Sends the GC message k_EMsgClientToGCRecalibrateMMR (CMsgClientToGCRecalibrateMMR) and awaits the response k_EMsgClientToGCRecalibrateMMRResponse,
// delivered as *CMsgClientToGCRecalibrateMMRResponse.
func (d *Dota2) SendRecalibrateMMR(
	ctx context.Context,
) (*protocol.CMsgClientToGCRecalibrateMMRResponse, error) {
	req := &protocol.CMsgClientToGCRecalibrateMMR{}
	resp := &protocol.CMsgClientToGCRecalibrateMMRResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecalibrateMMR),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRecalibrateMMRResponse),
		resp,
	)
}

// SendRemoveFilteredPlayer sends a remove filtered player.
//
// Sends the GC message k_EMsgClientToGCRemoveFilteredPlayer (CMsgClientToGCRemoveFilteredPlayer) and awaits the response k_EMsgGCToClientRemoveFilteredPlayerResponse,
// delivered as *CMsgGCToClientRemoveFilteredPlayerResponse.
func (d *Dota2) SendRemoveFilteredPlayer(
	ctx context.Context,
	accountIDToRemove uint32,
) (*protocol.CMsgGCToClientRemoveFilteredPlayerResponse, error) {
	req := &protocol.CMsgClientToGCRemoveFilteredPlayer{
		AccountIdToRemove: &accountIDToRemove,
	}
	resp := &protocol.CMsgGCToClientRemoveFilteredPlayerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveFilteredPlayer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRemoveFilteredPlayerResponse),
		resp,
	)
}

// SendRemoveGuildRole sends a remove guild role.
//
// Sends the GC message k_EMsgClientToGCRemoveGuildRole (CMsgClientToGCRemoveGuildRole) and awaits the response k_EMsgClientToGCRemoveGuildRoleResponse,
// delivered as *CMsgClientToGCRemoveGuildRoleResponse.
func (d *Dota2) SendRemoveGuildRole(
	ctx context.Context,
	guildID uint32,
	roleID uint32,
) (*protocol.CMsgClientToGCRemoveGuildRoleResponse, error) {
	req := &protocol.CMsgClientToGCRemoveGuildRole{
		GuildId: &guildID,
		RoleId:  &roleID,
	}
	resp := &protocol.CMsgClientToGCRemoveGuildRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveGuildRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRemoveGuildRoleResponse),
		resp,
	)
}

// SendRoadToTIDevForceQuest sends a road to ti dev force quest.
//
// Sends the GC message k_EMsgClientToGCRoadToTIDevForceQuest (CMsgClientToGCRoadToTIDevForceQuest). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendRoadToTIDevForceQuest(
	eventID uint32,
	forceMatchType bool,
	forceID uint32,
) {
	req := &protocol.CMsgClientToGCRoadToTIDevForceQuest{
		EventId:        &eventID,
		ForceMatchType: &forceMatchType,
		ForceId:        &forceID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIDevForceQuest), req)
}

// SendRoadToTIUseItem sends a road to ti use item.
//
// Sends the GC message k_EMsgClientToGCRoadToTIUseItem (CMsgClientToGCRoadToTIUseItem) and awaits the response k_EMsgClientToGCRoadToTIUseItemResponse,
// delivered as *CMsgClientToGCRoadToTIUseItemResponse.
func (d *Dota2) SendRoadToTIUseItem(
	ctx context.Context,
	eventID uint32,
	itemType uint32,
	heroIndex uint32,
) (*protocol.CMsgClientToGCRoadToTIUseItemResponse, error) {
	req := &protocol.CMsgClientToGCRoadToTIUseItem{
		EventId:   &eventID,
		ItemType:  &itemType,
		HeroIndex: &heroIndex,
	}
	resp := &protocol.CMsgClientToGCRoadToTIUseItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIUseItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCRoadToTIUseItemResponse),
		resp,
	)
}

// SendShowcaseAdminConvict sends a showcase admin convict.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminConvict (CMsgClientToGCShowcaseAdminConvict) and awaits the response k_EMsgClientToGCShowcaseAdminConvictResponse,
// delivered as *CMsgClientToGCShowcaseAdminConvictResponse.
func (d *Dota2) SendShowcaseAdminConvict(
	ctx context.Context,
	targetAccountID uint32,
	showcaseType protocol.EShowcaseType,
) (*protocol.CMsgClientToGCShowcaseAdminConvictResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminConvict{
		TargetAccountId: &targetAccountID,
		ShowcaseType:    &showcaseType,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminConvictResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminConvict),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminConvictResponse),
		resp,
	)
}

// SendShowcaseAdminExonerate sends a showcase admin exonerate.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminExonerate (CMsgClientToGCShowcaseAdminExonerate) and awaits the response k_EMsgClientToGCShowcaseAdminExonerateResponse,
// delivered as *CMsgClientToGCShowcaseAdminExonerateResponse.
func (d *Dota2) SendShowcaseAdminExonerate(
	ctx context.Context,
	targetAccountID uint32,
	showcaseType protocol.EShowcaseType,
) (*protocol.CMsgClientToGCShowcaseAdminExonerateResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminExonerate{
		TargetAccountId: &targetAccountID,
		ShowcaseType:    &showcaseType,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminExonerateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminExonerate),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminExonerateResponse),
		resp,
	)
}

// SendShowcaseAdminLockAccount sends a showcase admin lock account.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminLockAccount (CMsgClientToGCShowcaseAdminLockAccount) and awaits the response k_EMsgClientToGCShowcaseAdminLockAccountResponse,
// delivered as *CMsgClientToGCShowcaseAdminLockAccountResponse.
func (d *Dota2) SendShowcaseAdminLockAccount(
	ctx context.Context,
	targetAccountID uint32,
	lockedUntilTimestamp uint32,
) (*protocol.CMsgClientToGCShowcaseAdminLockAccountResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminLockAccount{
		TargetAccountId:      &targetAccountID,
		LockedUntilTimestamp: &lockedUntilTimestamp,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminLockAccountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminLockAccount),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminLockAccountResponse),
		resp,
	)
}

// SendShowcaseAdminReset sends a showcase admin reset.
//
// Sends the GC message k_EMsgClientToGCShowcaseAdminReset (CMsgClientToGCShowcaseAdminReset) and awaits the response k_EMsgClientToGCShowcaseAdminResetResponse,
// delivered as *CMsgClientToGCShowcaseAdminResetResponse.
func (d *Dota2) SendShowcaseAdminReset(
	ctx context.Context,
	targetAccountID uint32,
	showcaseType protocol.EShowcaseType,
) (*protocol.CMsgClientToGCShowcaseAdminResetResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseAdminReset{
		TargetAccountId: &targetAccountID,
		ShowcaseType:    &showcaseType,
	}
	resp := &protocol.CMsgClientToGCShowcaseAdminResetResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminReset),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseAdminResetResponse),
		resp,
	)
}

// SendSpectatorLobbyGameDetails sends spectator lobby game details.
//
// Sends the GC message k_EMsgSpectatorLobbyGameDetails (CMsgSpectatorLobbyGameDetails). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendSpectatorLobbyGameDetails(
	req *protocol.CMsgSpectatorLobbyGameDetails,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgSpectatorLobbyGameDetails), req)
}

// SendTeamInvite_GCResponseToInvitee sends a team invite _ gc response to invitee.
//
// Sends the GC message k_EMsgGCTeamInvite_GCResponseToInvitee (CMsgDOTATeamInvite_GCResponseToInvitee). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendTeamInvite_GCResponseToInvitee(
	result protocol.ETeamInviteResult,
	teamName string,
) {
	req := &protocol.CMsgDOTATeamInvite_GCResponseToInvitee{
		Result:   &result,
		TeamName: &teamName,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee), req)
}

// SendUnderDraftBuy sends a under draft buy.
//
// Sends the GC message k_EMsgClientToGCUnderDraftBuy (CMsgClientToGCUnderDraftBuy) and awaits the response k_EMsgClientToGCUnderDraftBuyResponse,
// delivered as *CMsgClientToGCUnderDraftBuyResponse.
func (d *Dota2) SendUnderDraftBuy(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
) (*protocol.CMsgClientToGCUnderDraftBuyResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftBuy{
		EventId: &eventID,
		SlotId:  &slotID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftBuyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftBuy),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftBuyResponse),
		resp,
	)
}

// SendUnderDraftRollBackBench sends a under draft roll back bench.
//
// Sends the GC message k_EMsgClientToGCUnderDraftRollBackBench (CMsgClientToGCUnderDraftRollBackBench) and awaits the response k_EMsgClientToGCUnderDraftRollBackBenchResponse,
// delivered as *CMsgClientToGCUnderDraftRollBackBenchResponse.
func (d *Dota2) SendUnderDraftRollBackBench(
	ctx context.Context,
	eventID uint32,
) (*protocol.CMsgClientToGCUnderDraftRollBackBenchResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftRollBackBench{
		EventId: &eventID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftRollBackBenchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRollBackBench),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftRollBackBenchResponse),
		resp,
	)
}

// SendUnderDraftSell sends a under draft sell.
//
// Sends the GC message k_EMsgClientToGCUnderDraftSell (CMsgClientToGCUnderDraftSell) and awaits the response k_EMsgClientToGCUnderDraftSellResponse,
// delivered as *CMsgClientToGCUnderDraftSellResponse.
func (d *Dota2) SendUnderDraftSell(
	ctx context.Context,
	eventID uint32,
	slotID uint32,
) (*protocol.CMsgClientToGCUnderDraftSellResponse, error) {
	req := &protocol.CMsgClientToGCUnderDraftSell{
		EventId: &eventID,
		SlotId:  &slotID,
	}
	resp := &protocol.CMsgClientToGCUnderDraftSellResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftSell),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUnderDraftSellResponse),
		resp,
	)
}

// SendUpdateComicBookStats sends update comic book stats.
//
// Sends the GC message k_EMsgClientToGCUpdateComicBookStats (CMsgClientToGCUpdateComicBookStats). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendUpdateComicBookStats(
	comicID uint32,
	stats []*protocol.CMsgClientToGCUpdateComicBookStats_SingleStat,
	languageStats protocol.CMsgClientToGCUpdateComicBookStats_LanguageStats,
) {
	req := &protocol.CMsgClientToGCUpdateComicBookStats{
		ComicId:       &comicID,
		Stats:         stats,
		LanguageStats: &languageStats,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUpdateComicBookStats), req)
}

// SendUpdateFilteredPlayerNote sends a update filtered player note.
//
// Sends the GC message k_EMsgClientToGCUpdateFilteredPlayerNote (CMsgClientToGCUpdateFilteredPlayerNote) and awaits the response k_EMsgGCToClientUpdateFilteredPlayerNoteResponse,
// delivered as *CMsgGCToClientUpdateFilteredPlayerNoteResponse.
func (d *Dota2) SendUpdateFilteredPlayerNote(
	ctx context.Context,
	targetAccountID uint32,
	newNote string,
) (*protocol.CMsgGCToClientUpdateFilteredPlayerNoteResponse, error) {
	req := &protocol.CMsgClientToGCUpdateFilteredPlayerNote{
		TargetAccountId: &targetAccountID,
		NewNote:         &newNote,
	}
	resp := &protocol.CMsgGCToClientUpdateFilteredPlayerNoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUpdateFilteredPlayerNote),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientUpdateFilteredPlayerNoteResponse),
		resp,
	)
}

// SendUpdatePartyBeacon sends a update party beacon.
//
// Sends the GC message k_EMsgClientToGCUpdatePartyBeacon (CMsgClientToGCUpdatePartyBeacon). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendUpdatePartyBeacon(
	action protocol.CMsgClientToGCUpdatePartyBeacon_Action,
) {
	req := &protocol.CMsgClientToGCUpdatePartyBeacon{
		Action: &action,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUpdatePartyBeacon), req)
}

// SendUploadMatchClip sends a upload match clip.
//
// Sends the GC message k_EMsgClientToGCUploadMatchClip (CMsgClientToGCUploadMatchClip) and awaits the response k_EMsgGCToClientUploadMatchClipResponse,
// delivered as *CMsgGCToClientUploadMatchClipResponse.
func (d *Dota2) SendUploadMatchClip(
	ctx context.Context,
	matchClip protocol.CMatchClip,
) (*protocol.CMsgGCToClientUploadMatchClipResponse, error) {
	req := &protocol.CMsgClientToGCUploadMatchClip{
		MatchClip: &matchClip,
	}
	resp := &protocol.CMsgGCToClientUploadMatchClipResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCUploadMatchClip),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientUploadMatchClipResponse),
		resp,
	)
}

// SendVerifyFavoritePlayers sends verify favorite players.
//
// Sends the GC message k_EMsgClientToGCVerifyFavoritePlayers (CMsgClientToGCVerifyFavoritePlayers) and awaits the response k_EMsgGCToClientVerifyFavoritePlayersResponse,
// delivered as *CMsgGCToClientVerifyFavoritePlayersResponse.
func (d *Dota2) SendVerifyFavoritePlayers(
	ctx context.Context,
	accountIDs []uint32,
) (*protocol.CMsgGCToClientVerifyFavoritePlayersResponse, error) {
	req := &protocol.CMsgClientToGCVerifyFavoritePlayers{
		AccountIds: accountIDs,
	}
	resp := &protocol.CMsgGCToClientVerifyFavoritePlayersResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVerifyFavoritePlayers),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientVerifyFavoritePlayersResponse),
		resp,
	)
}

// SendWatchGame sends a watch game.
//
// Sends the GC message k_EMsgGCWatchGame (CMsgWatchGame) and awaits the response k_EMsgGCWatchGameResponse,
// delivered as *CMsgWatchGameResponse.
func (d *Dota2) SendWatchGame(
	ctx context.Context,
	serverSteamid steamid.SteamId,
	watchServerSteamid steamid.SteamId,
	lobbyID uint64,
	regions []uint32,
) (*protocol.CMsgWatchGameResponse, error) {
	serverSteamidU64Val := uint64(serverSteamid)
	serverSteamidU64 := &serverSteamidU64Val
	watchServerSteamidU64Val := uint64(watchServerSteamid)
	watchServerSteamidU64 := &watchServerSteamidU64Val
	req := &protocol.CMsgWatchGame{
		ServerSteamid:      serverSteamidU64,
		WatchServerSteamid: watchServerSteamidU64,
		LobbyId:            &lobbyID,
		Regions:            regions,
	}
	resp := &protocol.CMsgWatchGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchGame),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchGameResponse),
		resp,
	)
}

// SendWatchingBroadcast sends a watching broadcast.
//
// Sends the GC message k_EMsgClientToGCWatchingBroadcast (CMsgClientToGCWatchingBroadcast). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SendWatchingBroadcast(
	matchID uint64,
) {
	req := &protocol.CMsgClientToGCWatchingBroadcast{
		MatchId: &matchID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCWatchingBroadcast), req)
}

// SetBannedHeroes sets banned heroes.
//
// Sends the GC message k_EMsgClientToGCSetBannedHeroes (CMsgClientToGCSetBannedHeroes). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SetBannedHeroes(
	bannedHeroIDs []int32,
) {
	req := &protocol.CMsgClientToGCSetBannedHeroes{
		BannedHeroIds: bannedHeroIDs,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetBannedHeroes), req)
}

// SetCompendiumSelection sets a compendium selection.
//
// Sends the GC message k_EMsgGCCompendiumSetSelection (CMsgDOTACompendiumSelection) and awaits the response k_EMsgGCCompendiumSetSelectionResponse,
// delivered as *CMsgDOTACompendiumSelectionResponse.
func (d *Dota2) SetCompendiumSelection(
	ctx context.Context,
	selectionIndex uint32,
	selection uint32,
	leagueid uint32,
) (*protocol.CMsgDOTACompendiumSelectionResponse, error) {
	req := &protocol.CMsgDOTACompendiumSelection{
		SelectionIndex: &selectionIndex,
		Selection:      &selection,
		Leagueid:       &leagueid,
	}
	resp := &protocol.CMsgDOTACompendiumSelectionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse),
		resp,
	)
}

// SetDPCFavoriteState sets a dpc favorite state.
//
// Sends the GC message k_EMsgClientToGCSetDPCFavoriteState (CMsgClientToGCSetDPCFavoriteState) and awaits the response k_EMsgClientToGCSetDPCFavoriteStateResponse,
// delivered as *CMsgClientToGCSetDPCFavoriteStateResponse.
func (d *Dota2) SetDPCFavoriteState(
	ctx context.Context,
	favoriteType protocol.EDPCFavoriteType,
	favoriteID uint32,
	enabled bool,
) (*protocol.CMsgClientToGCSetDPCFavoriteStateResponse, error) {
	req := &protocol.CMsgClientToGCSetDPCFavoriteState{
		FavoriteType: &favoriteType,
		FavoriteId:   &favoriteID,
		Enabled:      &enabled,
	}
	resp := &protocol.CMsgClientToGCSetDPCFavoriteStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetDPCFavoriteState),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetDPCFavoriteStateResponse),
		resp,
	)
}

// SetDevOverworldFortune sets a dev overworld fortune.
//
// Sends the GC message k_EMsgClientToGCOverworldDevSetFortune (CMsgClientToGCOverworldDevSetFortune) and awaits the response k_EMsgClientToGCOverworldDevSetFortuneResponse,
// delivered as *CMsgClientToGCOverworldDevSetFortuneResponse.
func (d *Dota2) SetDevOverworldFortune(
	ctx context.Context,
	overworldID uint32,
	fortuneID uint32,
) (*protocol.CMsgClientToGCOverworldDevSetFortuneResponse, error) {
	req := &protocol.CMsgClientToGCOverworldDevSetFortune{
		OverworldId: &overworldID,
		FortuneId:   &fortuneID,
	}
	resp := &protocol.CMsgClientToGCOverworldDevSetFortuneResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevSetFortune),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCOverworldDevSetFortuneResponse),
		resp,
	)
}

// SetEventActiveSeasonID sets a event active season id.
//
// Sends the GC message k_EMsgClientToGCSetEventActiveSeasonID (CMsgClientToGCSetEventActiveSeasonID) and awaits the response k_EMsgClientToGCSetEventActiveSeasonIDResponse,
// delivered as *CMsgClientToGCSetEventActiveSeasonIDResponse.
func (d *Dota2) SetEventActiveSeasonID(
	ctx context.Context,
	eventID uint32,
	activeSeasonID uint32,
) (*protocol.CMsgClientToGCSetEventActiveSeasonIDResponse, error) {
	req := &protocol.CMsgClientToGCSetEventActiveSeasonID{
		EventId:        &eventID,
		ActiveSeasonId: &activeSeasonID,
	}
	resp := &protocol.CMsgClientToGCSetEventActiveSeasonIDResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetEventActiveSeasonID),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetEventActiveSeasonIDResponse),
		resp,
	)
}

// SetFavoritePage sets a favorite page.
//
// Sends the GC message k_EMsgClientToGCSetFavoritePage (CMsgClientToGCSetFavoritePage) and awaits the response k_EMsgClientToGCSetFavoritePageResponse,
// delivered as *CMsgClientToGCSetFavoritePageResponse.
func (d *Dota2) SetFavoritePage(
	ctx context.Context,
	pageNum uint32,
	clear bool,
) (*protocol.CMsgClientToGCSetFavoritePageResponse, error) {
	req := &protocol.CMsgClientToGCSetFavoritePage{
		PageNum: &pageNum,
		Clear:   &clear,
	}
	resp := &protocol.CMsgClientToGCSetFavoritePageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetFavoritePage),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetFavoritePageResponse),
		resp,
	)
}

// SetFavoriteTeam sets a favorite team.
//
// Sends the GC message k_EMsgDOTASetFavoriteTeam (CMsgDOTASetFavoriteTeam). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SetFavoriteTeam(
	teamID uint32,
	eventID uint32,
) {
	req := &protocol.CMsgDOTASetFavoriteTeam{
		TeamId:  &teamID,
		EventId: &eventID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgDOTASetFavoriteTeam), req)
}

// SetGuildInfo sets a guild info.
//
// Sends the GC message k_EMsgClientToGCSetGuildInfo (CMsgClientToGCSetGuildInfo) and awaits the response k_EMsgClientToGCSetGuildInfoResponse,
// delivered as *CMsgClientToGCSetGuildInfoResponse.
func (d *Dota2) SetGuildInfo(
	ctx context.Context,
	guildID uint32,
	guildInfo protocol.CMsgGuildInfo,
	guildChatType protocol.EGuildChatType,
) (*protocol.CMsgClientToGCSetGuildInfoResponse, error) {
	req := &protocol.CMsgClientToGCSetGuildInfo{
		GuildId:       &guildID,
		GuildInfo:     &guildInfo,
		GuildChatType: &guildChatType,
	}
	resp := &protocol.CMsgClientToGCSetGuildInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildInfo),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildInfoResponse),
		resp,
	)
}

// SetGuildMemberRole sets a guild member role.
//
// Sends the GC message k_EMsgClientToGCSetGuildMemberRole (CMsgClientToGCSetGuildMemberRole) and awaits the response k_EMsgClientToGCSetGuildMemberRoleResponse,
// delivered as *CMsgClientToGCSetGuildMemberRoleResponse.
func (d *Dota2) SetGuildMemberRole(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
	targetRoleID uint32,
) (*protocol.CMsgClientToGCSetGuildMemberRoleResponse, error) {
	req := &protocol.CMsgClientToGCSetGuildMemberRole{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
		TargetRoleId:    &targetRoleID,
	}
	resp := &protocol.CMsgClientToGCSetGuildMemberRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildMemberRole),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildMemberRoleResponse),
		resp,
	)
}

// SetGuildRoleOrder sets a guild role order.
//
// Sends the GC message k_EMsgClientToGCSetGuildRoleOrder (CMsgClientToGCSetGuildRoleOrder) and awaits the response k_EMsgClientToGCSetGuildRoleOrderResponse,
// delivered as *CMsgClientToGCSetGuildRoleOrderResponse.
func (d *Dota2) SetGuildRoleOrder(
	ctx context.Context,
	guildID uint32,
	requestedRoleIDs []uint32,
	previousRoleIDs []uint32,
) (*protocol.CMsgClientToGCSetGuildRoleOrderResponse, error) {
	req := &protocol.CMsgClientToGCSetGuildRoleOrder{
		GuildId:          &guildID,
		RequestedRoleIds: requestedRoleIDs,
		PreviousRoleIds:  previousRoleIDs,
	}
	resp := &protocol.CMsgClientToGCSetGuildRoleOrderResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildRoleOrder),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetGuildRoleOrderResponse),
		resp,
	)
}

// SetHeroSticker sets a hero sticker.
//
// Sends the GC message k_EMsgClientToGCSetHeroSticker (CMsgClientToGCSetHeroSticker) and awaits the response k_EMsgClientToGCSetHeroStickerResponse,
// delivered as *CMsgClientToGCSetHeroStickerResponse.
func (d *Dota2) SetHeroSticker(
	ctx context.Context,
	heroID int32,
	newItemID uint64,
	oldItemID uint64,
) (*protocol.CMsgClientToGCSetHeroStickerResponse, error) {
	req := &protocol.CMsgClientToGCSetHeroSticker{
		HeroId:    &heroID,
		NewItemId: &newItemID,
		OldItemId: &oldItemID,
	}
	resp := &protocol.CMsgClientToGCSetHeroStickerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetHeroSticker),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetHeroStickerResponse),
		resp,
	)
}

// SetLobbyCoach requests the coach slot for your team in the lobby.
func (d *Dota2) SetLobbyCoach(
	team protocol.DOTA_GC_TEAM,
) {
	req := &protocol.CMsgPracticeLobbySetCoach{
		Team: &team,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbySetCoach), req)
}

// SetLobbyDetails updates the details of the current practice lobby: game
// mode, map, server region, pass key, spectator policy, and cheat/bot settings.
func (d *Dota2) SetLobbyDetails(
	req *protocol.CMsgPracticeLobbySetDetails,
) {
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails), req)
}

// SetMatchHistoryAccess sets match history access.
//
// Sends the GC message k_EMsgGCSetMatchHistoryAccess (CMsgDOTASetMatchHistoryAccess) and awaits the response k_EMsgGCSetMatchHistoryAccessResponse,
// delivered as *CMsgDOTASetMatchHistoryAccessResponse.
func (d *Dota2) SetMatchHistoryAccess(
	ctx context.Context,
	allow3RdPartyMatchHistory bool,
) (*protocol.CMsgDOTASetMatchHistoryAccessResponse, error) {
	req := &protocol.CMsgDOTASetMatchHistoryAccess{
		Allow_3RdPartyMatchHistory: &allow3RdPartyMatchHistory,
	}
	resp := &protocol.CMsgDOTASetMatchHistoryAccessResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccess),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccessResponse),
		resp,
	)
}

// SetMemberPartyCoach sets a member party coach.
//
// Sends the GC message k_EMsgGCPartyMemberSetCoach (CMsgDOTAPartyMemberSetCoach). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SetMemberPartyCoach(
	wantsCoach bool,
) {
	req := &protocol.CMsgDOTAPartyMemberSetCoach{
		WantsCoach: &wantsCoach,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPartyMemberSetCoach), req)
}

// SetPartyBuilderOptions configures how the party enters matchmaking, such
// as the selected lanes or roles in role-queue modes.
func (d *Dota2) SetPartyBuilderOptions(
	additionalSlots uint32,
	matchType protocol.MatchType,
	matchgroups uint32,
	language protocol.MatchLanguages,
) {
	req := &protocol.CMsgPartyBuilderOptions{
		AdditionalSlots: &additionalSlots,
		MatchType:       &matchType,
		Matchgroups:     &matchgroups,
		Language:        &language,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions), req)
}

// SetPartyLeader makes the given party member the party leader.
func (d *Dota2) SetPartyLeader(
	newLeaderSteamid steamid.SteamId,
) {
	newLeaderSteamidU64Val := uint64(newLeaderSteamid)
	newLeaderSteamidU64 := &newLeaderSteamidU64Val
	req := &protocol.CMsgDOTASetGroupLeader{
		NewLeaderSteamid: newLeaderSteamidU64,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader), req)
}

// SetPartyOpen opens or closes the party so that friends can join freely.
func (d *Dota2) SetPartyOpen(
	open bool,
) {
	req := &protocol.CMsgDOTASetGroupOpenStatus{
		Open: &open,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen), req)
}

// SetProfileCardSlots sets profile card slots.
//
// Sends the GC message k_EMsgClientToGCSetProfileCardSlots (CMsgClientToGCSetProfileCardSlots). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SetProfileCardSlots(
	slots []*protocol.CMsgClientToGCSetProfileCardSlots_CardSlot,
) {
	req := &protocol.CMsgClientToGCSetProfileCardSlots{
		Slots: slots,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetProfileCardSlots), req)
}

// SetProfilePrivacy sets a profile privacy.
//
// Sends the GC message k_EMsgGCSetProfilePrivacy (CMsgDOTASetProfilePrivacy) and awaits the response k_EMsgGCSetProfilePrivacyResponse,
// delivered as *CMsgDOTASetProfilePrivacyResponse.
func (d *Dota2) SetProfilePrivacy(
	ctx context.Context,
	profilePrivate bool,
) (*protocol.CMsgDOTASetProfilePrivacyResponse, error) {
	req := &protocol.CMsgDOTASetProfilePrivacy{
		ProfilePrivate: &profilePrivate,
	}
	resp := &protocol.CMsgDOTASetProfilePrivacyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetProfilePrivacy),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSetProfilePrivacyResponse),
		resp,
	)
}

// SetShowcaseUserData sets a showcase user data.
//
// Sends the GC message k_EMsgClientToGCShowcaseSetUserData (CMsgClientToGCShowcaseSetUserData) and awaits the response k_EMsgClientToGCShowcaseSetUserDataResponse,
// delivered as *CMsgClientToGCShowcaseSetUserDataResponse.
func (d *Dota2) SetShowcaseUserData(
	ctx context.Context,
	showcaseType protocol.EShowcaseType,
	showcase protocol.CMsgShowcase,
	formatVersion uint32,
) (*protocol.CMsgClientToGCShowcaseSetUserDataResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseSetUserData{
		ShowcaseType:  &showcaseType,
		Showcase:      &showcase,
		FormatVersion: &formatVersion,
	}
	resp := &protocol.CMsgClientToGCShowcaseSetUserDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseSetUserData),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseSetUserDataResponse),
		resp,
	)
}

// SetSpectatorLobbyDetails sets spectator lobby details.
//
// Sends the GC message k_EMsgClientToGCSetSpectatorLobbyDetails (CMsgSetSpectatorLobbyDetails). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SetSpectatorLobbyDetails(
	lobbyID uint64,
	lobbyName string,
	passKey string,
	gameDetails protocol.CMsgSpectatorLobbyGameDetails,
) {
	req := &protocol.CMsgSetSpectatorLobbyDetails{
		LobbyId:     &lobbyID,
		LobbyName:   &lobbyName,
		PassKey:     &passKey,
		GameDetails: &gameDetails,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetails), req)
}

// SpectateFriendGame spectates a friend game.
//
// Sends the GC message k_EMsgGCSpectateFriendGame (CMsgSpectateFriendGame) and awaits the response k_EMsgGCSpectateFriendGameResponse,
// delivered as *CMsgSpectateFriendGameResponse.
func (d *Dota2) SpectateFriendGame(
	ctx context.Context,
	steamID steamid.SteamId,
	live bool,
) (*protocol.CMsgSpectateFriendGameResponse, error) {
	steamIDU64Val := uint64(steamID)
	steamIDU64 := &steamIDU64Val
	req := &protocol.CMsgSpectateFriendGame{
		SteamId: steamIDU64,
		Live:    &live,
	}
	resp := &protocol.CMsgSpectateFriendGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSpectateFriendGame),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSpectateFriendGameResponse),
		resp,
	)
}

// StartFindingMatch enters the matchmaking queue with the given options:
// match mode, map, team desirability, and lobby type.
//
// The GC confirms entry through the response and later reports the found match
// through the ready-up flow; see SendReadyUp.
func (d *Dota2) StartFindingMatch(
	ctx context.Context,
	req *protocol.CMsgStartFindingMatch,
) (*protocol.CMsgStartFindingMatchResult, error) {
	resp := &protocol.CMsgStartFindingMatchResult{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCStartFindingMatch),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCStartFindingMatchResponse),
		resp,
	)
}

// StartTriviaSession starts a trivia session.
//
// Sends the GC message k_EMsgStartTriviaSession (CMsgDOTAStartTriviaSession) and awaits the response k_EMsgStartTriviaSessionResponse,
// delivered as *CMsgDOTAStartTriviaSessionResponse.
func (d *Dota2) StartTriviaSession(
	ctx context.Context,
) (*protocol.CMsgDOTAStartTriviaSessionResponse, error) {
	req := &protocol.CMsgDOTAStartTriviaSession{}
	resp := &protocol.CMsgDOTAStartTriviaSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgStartTriviaSession),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgStartTriviaSessionResponse),
		resp,
	)
}

// StartWatchingOverwatch starts a watching overwatch.
//
// Sends the GC message k_EMsgClientToGCStartWatchingOverwatch (CMsgClientToGCStartWatchingOverwatch). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) StartWatchingOverwatch(
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
) {
	req := &protocol.CMsgClientToGCStartWatchingOverwatch{
		OverwatchReplayId: &overwatchReplayID,
		TargetPlayerSlot:  &targetPlayerSlot,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCStartWatchingOverwatch), req)
}

// StopFindingMatch leaves the matchmaking queue.
func (d *Dota2) StopFindingMatch(
	acceptCooldown bool,
) {
	req := &protocol.CMsgStopFindingMatch{
		AcceptCooldown: &acceptCooldown,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCStopFindingMatch), req)
}

// StopWatchingOverwatch stops a watching overwatch.
//
// Sends the GC message k_EMsgClientToGCStopWatchingOverwatch (CMsgClientToGCStopWatchingOverwatch). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) StopWatchingOverwatch(
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
) {
	req := &protocol.CMsgClientToGCStopWatchingOverwatch{
		OverwatchReplayId: &overwatchReplayID,
		TargetPlayerSlot:  &targetPlayerSlot,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCStopWatchingOverwatch), req)
}

// SubmitCoachTeammateRating submits a coach teammate rating.
//
// Sends the GC message k_EMsgClientToGCSubmitCoachTeammateRating (CMsgClientToGCSubmitCoachTeammateRating) and awaits the response k_EMsgClientToGCSubmitCoachTeammateRatingResponse,
// delivered as *CMsgClientToGCSubmitCoachTeammateRatingResponse.
func (d *Dota2) SubmitCoachTeammateRating(
	ctx context.Context,
	matchID uint64,
	coachAccountID uint32,
	rating protocol.ECoachTeammateRating,
	reason string,
) (*protocol.CMsgClientToGCSubmitCoachTeammateRatingResponse, error) {
	req := &protocol.CMsgClientToGCSubmitCoachTeammateRating{
		MatchId:        &matchID,
		CoachAccountId: &coachAccountID,
		Rating:         &rating,
		Reason:         &reason,
	}
	resp := &protocol.CMsgClientToGCSubmitCoachTeammateRatingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitCoachTeammateRating),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitCoachTeammateRatingResponse),
		resp,
	)
}

// SubmitDraftTriviaMatchAnswer submits a draft trivia match answer.
//
// Sends the GC message k_EMsgClientToGCSubmitDraftTriviaMatchAnswer (CMsgClientToGCSubmitDraftTriviaMatchAnswer) and awaits the response k_EMsgClientToGCSubmitDraftTriviaMatchAnswerResponse,
// delivered as *CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse.
func (d *Dota2) SubmitDraftTriviaMatchAnswer(
	ctx context.Context,
	choseRadiantAsWinner bool,
	eventID uint32,
	endTime uint32,
) (*protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse, error) {
	req := &protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswer{
		ChoseRadiantAsWinner: &choseRadiantAsWinner,
		EventId:              &eventID,
		EndTime:              &endTime,
	}
	resp := &protocol.CMsgClientToGCSubmitDraftTriviaMatchAnswerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswerResponse),
		resp,
	)
}

// SubmitInfoPlayer submits a info player.
//
// Sends the GC message k_EMsgGCPlayerInfoSubmit (CMsgGCPlayerInfoSubmit) and awaits the response k_EMsgGCPlayerInfoSubmitResponse,
// delivered as *CMsgGCPlayerInfoSubmitResponse.
func (d *Dota2) SubmitInfoPlayer(
	ctx context.Context,
	req *protocol.CMsgGCPlayerInfoSubmit,
) (*protocol.CMsgGCPlayerInfoSubmitResponse, error) {
	resp := &protocol.CMsgGCPlayerInfoSubmitResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmit),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmitResponse),
		resp,
	)
}

// SubmitLobbyMVPVote submits a lobby mvp vote.
//
// Sends the GC message k_EMsgGCSubmitLobbyMVPVote (CMsgDOTASubmitLobbyMVPVote) and awaits the response k_EMsgGCSubmitLobbyMVPVoteResponse,
// delivered as *CMsgDOTASubmitLobbyMVPVoteResponse.
func (d *Dota2) SubmitLobbyMVPVote(
	ctx context.Context,
	targetAccountID uint32,
) (*protocol.CMsgDOTASubmitLobbyMVPVoteResponse, error) {
	req := &protocol.CMsgDOTASubmitLobbyMVPVote{
		TargetAccountId: &targetAccountID,
	}
	resp := &protocol.CMsgDOTASubmitLobbyMVPVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVote),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVoteResponse),
		resp,
	)
}

// SubmitOWConviction submits a ow conviction.
//
// Sends the GC message k_EMsgClientToGCSubmitOWConviction (CMsgClientToGCSubmitOWConviction) and awaits the response k_EMsgClientToGCSubmitOWConvictionResponse,
// delivered as *CMsgClientToGCSubmitOWConvictionResponse.
func (d *Dota2) SubmitOWConviction(
	ctx context.Context,
	overwatchReplayID uint64,
	targetPlayerSlot uint32,
	cheatingConviction protocol.EOverwatchConviction,
	griefingConviction protocol.EOverwatchConviction,
) (*protocol.CMsgClientToGCSubmitOWConvictionResponse, error) {
	req := &protocol.CMsgClientToGCSubmitOWConviction{
		OverwatchReplayId:  &overwatchReplayID,
		TargetPlayerSlot:   &targetPlayerSlot,
		CheatingConviction: &cheatingConviction,
		GriefingConviction: &griefingConviction,
	}
	resp := &protocol.CMsgClientToGCSubmitOWConvictionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitOWConviction),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitOWConvictionResponse),
		resp,
	)
}

// SubmitPlayerInfoRanked submits a player info ranked.
//
// Sends the GC message k_EMsgGCRankedPlayerInfoSubmit (CMsgGCRankedPlayerInfoSubmit) and awaits the response k_EMsgGCRankedPlayerInfoSubmitResponse,
// delivered as *CMsgGCRankedPlayerInfoSubmitResponse.
func (d *Dota2) SubmitPlayerInfoRanked(
	ctx context.Context,
	name string,
) (*protocol.CMsgGCRankedPlayerInfoSubmitResponse, error) {
	req := &protocol.CMsgGCRankedPlayerInfoSubmit{
		Name: &name,
	}
	resp := &protocol.CMsgGCRankedPlayerInfoSubmitResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRankedPlayerInfoSubmit),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCRankedPlayerInfoSubmitResponse),
		resp,
	)
}

// SubmitPlayerMatchSurvey submits a player match survey.
//
// Sends the GC message k_EMsgClientToGCSubmitPlayerMatchSurvey (CMsgClientToGCSubmitPlayerMatchSurvey) and awaits the response k_EMsgClientToGCSubmitPlayerMatchSurveyResponse,
// delivered as *CMsgClientToGCSubmitPlayerMatchSurveyResponse.
func (d *Dota2) SubmitPlayerMatchSurvey(
	ctx context.Context,
	matchID uint64,
	rating int32,
	flags uint32,
) (*protocol.CMsgClientToGCSubmitPlayerMatchSurveyResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPlayerMatchSurvey{
		MatchId: &matchID,
		Rating:  &rating,
		Flags:   &flags,
	}
	resp := &protocol.CMsgClientToGCSubmitPlayerMatchSurveyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPlayerMatchSurvey),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPlayerMatchSurveyResponse),
		resp,
	)
}

// SubmitPlayerReport submits a player report.
//
// Sends the GC message k_EMsgGCSubmitPlayerReport (CMsgDOTASubmitPlayerReport) and awaits the response k_EMsgGCSubmitPlayerReportResponse,
// delivered as *CMsgDOTASubmitPlayerReportResponse.
func (d *Dota2) SubmitPlayerReport(
	ctx context.Context,
	targetAccountID uint32,
	reportFlags uint32,
	lobbyID uint64,
	comment string,
) (*protocol.CMsgDOTASubmitPlayerReportResponse, error) {
	req := &protocol.CMsgDOTASubmitPlayerReport{
		TargetAccountId: &targetAccountID,
		ReportFlags:     &reportFlags,
		LobbyId:         &lobbyID,
		Comment:         &comment,
	}
	resp := &protocol.CMsgDOTASubmitPlayerReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReport),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReportResponse),
		resp,
	)
}

// SubmitPlayerReportResponseV2 submits a player report response v 2.
//
// Sends the GC message k_EMsgGCSubmitPlayerReportResponseV2 (CMsgDOTASubmitPlayerReportResponseV2). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SubmitPlayerReportResponseV2(
	targetAccountID uint32,
	reportReason []uint32,
	debugMessage string,
	enumResult protocol.CMsgDOTASubmitPlayerReportResponseV2_EResult,
) {
	req := &protocol.CMsgDOTASubmitPlayerReportResponseV2{
		TargetAccountId: &targetAccountID,
		ReportReason:    reportReason,
		DebugMessage:    &debugMessage,
		EnumResult:      &enumResult,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReportResponseV2), req)
}

// SubmitPlayerReportV2 submits a player report v 2.
//
// Sends the GC message k_EMsgGCSubmitPlayerReportV2 (CMsgDOTASubmitPlayerReportV2). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) SubmitPlayerReportV2(
	targetAccountID uint32,
	reportReason []uint32,
	lobbyID uint64,
	gameTime float32,
	debugSlot uint32,
	debugMatchID uint64,
) {
	req := &protocol.CMsgDOTASubmitPlayerReportV2{
		TargetAccountId: &targetAccountID,
		ReportReason:    reportReason,
		LobbyId:         &lobbyID,
		GameTime:        &gameTime,
		DebugSlot:       &debugSlot,
		DebugMatchId:    &debugMatchID,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCSubmitPlayerReportV2), req)
}

// SubmitPrivateCoachingSessionRating submits a private coaching session rating.
//
// Sends the GC message k_EMsgClientToGCSubmitPrivateCoachingSessionRating (CMsgClientToGCSubmitPrivateCoachingSessionRating) and awaits the response k_EMsgClientToGCSubmitPrivateCoachingSessionRatingResponse,
// delivered as *CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse.
func (d *Dota2) SubmitPrivateCoachingSessionRating(
	ctx context.Context,
	coachingSessionID uint64,
	sessionRating protocol.ECoachTeammateRating,
) (*protocol.CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse, error) {
	req := &protocol.CMsgClientToGCSubmitPrivateCoachingSessionRating{
		CoachingSessionId: &coachingSessionID,
		SessionRating:     &sessionRating,
	}
	resp := &protocol.CMsgClientToGCSubmitPrivateCoachingSessionRatingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPrivateCoachingSessionRating),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCSubmitPrivateCoachingSessionRatingResponse),
		resp,
	)
}

// SubmitShowcaseReport submits a showcase report.
//
// Sends the GC message k_EMsgClientToGCShowcaseSubmitReport (CMsgClientToGCShowcaseSubmitReport) and awaits the response k_EMsgClientToGCShowcaseSubmitReportResponse,
// delivered as *CMsgClientToGCShowcaseSubmitReportResponse.
func (d *Dota2) SubmitShowcaseReport(
	ctx context.Context,
	targetAccountID uint32,
	showcaseType protocol.EShowcaseType,
	reportComment string,
) (*protocol.CMsgClientToGCShowcaseSubmitReportResponse, error) {
	req := &protocol.CMsgClientToGCShowcaseSubmitReport{
		TargetAccountId: &targetAccountID,
		ShowcaseType:    &showcaseType,
		ReportComment:   &reportComment,
	}
	resp := &protocol.CMsgClientToGCShowcaseSubmitReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseSubmitReport),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCShowcaseSubmitReportResponse),
		resp,
	)
}

// SubmitTriviaQuestionAnswer submits a trivia question answer.
//
// Sends the GC message k_EMsgSubmitTriviaQuestionAnswer (CMsgDOTASubmitTriviaQuestionAnswer) and awaits the response k_EMsgSubmitTriviaQuestionAnswerResponse,
// delivered as *CMsgDOTASubmitTriviaQuestionAnswerResponse.
func (d *Dota2) SubmitTriviaQuestionAnswer(
	ctx context.Context,
	questionID uint32,
	answerIndex uint32,
) (*protocol.CMsgDOTASubmitTriviaQuestionAnswerResponse, error) {
	req := &protocol.CMsgDOTASubmitTriviaQuestionAnswer{
		QuestionId:  &questionID,
		AnswerIndex: &answerIndex,
	}
	resp := &protocol.CMsgDOTASubmitTriviaQuestionAnswerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswer),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswerResponse),
		resp,
	)
}

// ToggleLobbyBroadcastChannelCameramanStatus toggles lobby broadcast channel cameraman status.
//
// Sends the GC message k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus (CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus). No response is tracked; any result
// arrives through the client event stream.
func (d *Dota2) ToggleLobbyBroadcastChannelCameramanStatus() {
	req := &protocol.CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus{}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus), req)
}

// TransferTeamAdmin transfers a team admin.
//
// Sends the GC message k_EMsgGCTransferTeamAdmin (CMsgDOTATransferTeamAdmin) and awaits the response k_EMsgGCTransferTeamAdminResponse,
// delivered as *CMsgDOTATransferTeamAdminResponse.
func (d *Dota2) TransferTeamAdmin(
	ctx context.Context,
	newAdminAccountID uint32,
	teamID uint32,
) (*protocol.CMsgDOTATransferTeamAdminResponse, error) {
	req := &protocol.CMsgDOTATransferTeamAdmin{
		NewAdminAccountId: &newAdminAccountID,
		TeamId:            &teamID,
	}
	resp := &protocol.CMsgDOTATransferTeamAdminResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTransferTeamAdmin),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgGCTransferTeamAdminResponse),
		resp,
	)
}

// UpgradeLeagueItem upgrades a league item.
//
// Sends the GC message k_EMsgUpgradeLeagueItem (CMsgUpgradeLeagueItem) and awaits the response k_EMsgUpgradeLeagueItemResponse,
// delivered as *CMsgUpgradeLeagueItemResponse.
func (d *Dota2) UpgradeLeagueItem(
	ctx context.Context,
	matchID uint64,
	leagueID uint32,
) (*protocol.CMsgUpgradeLeagueItemResponse, error) {
	req := &protocol.CMsgUpgradeLeagueItem{
		MatchId:  &matchID,
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgUpgradeLeagueItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgUpgradeLeagueItem),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgUpgradeLeagueItemResponse),
		resp,
	)
}

// UpgradeToGcFantasyCraftingClientTablets upgrades to gc fantasy crafting client tablets.
//
// Sends the GC message k_EMsgClientToGcFantasyCraftingUpgradeTablets (CMsgClientToGcFantasyCraftingUpgradeTablets) and awaits the response k_EMsgClientToGcFantasyCraftingUpgradeTabletsResponse,
// delivered as *CMsgClientToGcFantasyCraftingUpgradeTabletsResponse.
func (d *Dota2) UpgradeToGcFantasyCraftingClientTablets(
	ctx context.Context,
	fantasyLeague uint32,
) (*protocol.CMsgClientToGcFantasyCraftingUpgradeTabletsResponse, error) {
	req := &protocol.CMsgClientToGcFantasyCraftingUpgradeTablets{
		FantasyLeague: &fantasyLeague,
	}
	resp := &protocol.CMsgClientToGcFantasyCraftingUpgradeTabletsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGcFantasyCraftingUpgradeTablets),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGcFantasyCraftingUpgradeTabletsResponse),
		resp,
	)
}

// VoteForArcana votes a for arcana.
//
// Sends the GC message k_EMsgClientToGCVoteForArcana (CMsgClientToGCVoteForArcana) and awaits the response k_EMsgClientToGCVoteForArcanaResponse,
// delivered as *CMsgClientToGCVoteForArcanaResponse.
func (d *Dota2) VoteForArcana(
	ctx context.Context,
	matches []*protocol.CMsgArcanaVoteMatchVotes,
) (*protocol.CMsgClientToGCVoteForArcanaResponse, error) {
	req := &protocol.CMsgClientToGCVoteForArcana{
		Matches: matches,
	}
	resp := &protocol.CMsgClientToGCVoteForArcanaResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForArcana),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForArcanaResponse),
		resp,
	)
}

// VoteForMVP votes a for mvp.
//
// Sends the GC message k_EMsgClientToGCVoteForMVP (CMsgClientToGCVoteForMVP) and awaits the response k_EMsgClientToGCVoteForMVPResponse,
// delivered as *CMsgClientToGCVoteForMVPResponse.
func (d *Dota2) VoteForMVP(
	ctx context.Context,
	matchID uint64,
	accountID uint32,
) (*protocol.CMsgClientToGCVoteForMVPResponse, error) {
	req := &protocol.CMsgClientToGCVoteForMVP{
		MatchId:   &matchID,
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCVoteForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForMVP),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCVoteForMVPResponse),
		resp,
	)
}

// VoteMVPTimeout votes a mvp timeout.
//
// Sends the GC message k_EMsgClientToGCMVPVoteTimeout (CMsgClientToGCMVPVoteTimeout) and awaits the response k_EMsgClientToGCMVPVoteTimeoutResponse,
// delivered as *CMsgClientToGCMVPVoteTimeoutResponse.
func (d *Dota2) VoteMVPTimeout(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCMVPVoteTimeoutResponse, error) {
	req := &protocol.CMsgClientToGCMVPVoteTimeout{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCMVPVoteTimeoutResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMVPVoteTimeout),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgClientToGCMVPVoteTimeoutResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Dota2) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientAccountGuildEventDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.AccountGuildEventDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildChallengeUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ActiveGuildChallengeUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildContractsUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ActiveGuildContractsUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientArcanaVotesUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.ArcanaVotesUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupListRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupListRequest{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupRequest{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientBingoUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.BingoUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCBroadcastNotification)] = d.getEventEmitter(func() events.Event {
		return &events.BroadcastNotification{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCandyShopUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.CandyShopUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapPathCompleted)] = d.getEventEmitter(func() events.Event {
		return &events.CavernCrawlMapPathCompleted{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.CavernCrawlMapUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCChatModeratorBan)] = d.getEventEmitter(func() events.Event {
		return &events.ChatModeratorBan{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled)] = d.getEventEmitter(func() events.Event {
		return &events.ChatRegionsEnabled{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientClaimEventActionUsingItemCompleted)] = d.getEventEmitter(func() events.Event {
		return &events.ClaimEventActionUsingItemCompleted{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCClientSuspended)] = d.getEventEmitter(func() events.Event {
		return &events.ClientSuspended{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCoachTeammateRatingsChanged)] = d.getEventEmitter(func() events.Event {
		return &events.CoachTeammateRatingsChanged{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCommendNotification)] = d.getEventEmitter(func() events.Event {
		return &events.CommendNotification{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCCompendiumRemoveAllSelections)] = d.getEventEmitter(func() events.Event {
		return &events.CompendiumRemoveAllSelections{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientCraftworksUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.CraftworksUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule)] = d.getEventEmitter(func() events.Event {
		return &events.DOTAWeekendTourneySchedule{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientEmoticonData)] = d.getEventEmitter(func() events.Event {
		return &events.EmoticonData{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyFinalPlayerStats{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameChallenge)] = d.getEventEmitter(func() events.Event {
		return &events.FightingGameChallenge{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameChallengeCanceled)] = d.getEventEmitter(func() events.Event {
		return &events.FightingGameChallengeCanceled{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameStartMatch)] = d.getEventEmitter(func() events.Event {
		return &events.FightingGameStartMatch{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildFeedUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildFeedUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembersDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildMembersDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembershipUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.GuildMembershipUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientHeroStatueCreateResult)] = d.getEventEmitter(func() events.Event {
		return &events.HeroStatueCreateResult{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientInviteToDemoMode)] = d.getEventEmitter(func() events.Event {
		return &events.InviteToDemoMode{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientItemBattlerUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ItemBattlerUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue)] = d.getEventEmitter(func() events.Event {
		return &events.KickedFromMatchmakingQueue{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCLeagueAdminList)] = d.getEventEmitter(func() events.Event {
		return &events.LeagueAdminList{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyMVPAwarded{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCLobbyUpdateBroadcastChannelInfo)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyUpdateBroadcastChannelInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMatchGroupsVersion)] = d.getEventEmitter(func() events.Event {
		return &events.MatchGroupsVersion{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMatchSignedOut)] = d.getEventEmitter(func() events.Event {
		return &events.MatchSignedOut{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergeGroupInviteReply{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergePartyResponseReply{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientMonsterHunterUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.MonsterHunterUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientNotificationsUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.NotificationsUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCNotifyAccountFlagsChange)] = d.getEventEmitter(func() events.Event {
		return &events.NotifyAccountFlagsChange{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientOverwatchCasesAvailable)] = d.getEventEmitter(func() events.Event {
		return &events.OverwatchCasesAvailable{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientOverworldUserDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.OverworldUserDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartyBeaconUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.PartyBeaconUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCPartyLeaderWatchGamePrompt)] = d.getEventEmitter(func() events.Event {
		return &events.PartyLeaderWatchGamePrompt{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvite)] = d.getEventEmitter(func() events.Event {
		return &events.PartySearchInvite{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvites)] = d.getEventEmitter(func() events.Event {
		return &events.PartySearchInvites{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlayerBeaconState)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerBeaconState{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerJoinedChannel{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCOtherLeftChannel)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerLeftChannel{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPlaytestStatus)] = d.getEventEmitter(func() events.Event {
		return &events.PlaytestStatus{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCPopup)] = d.getEventEmitter(func() events.Event {
		return &events.Popup{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientPrivateCoachingSessionUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.PrivateCoachingSessionUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ProfileCardUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientQuestProgressUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.QuestProgressUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRankUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.RankUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCReadyUpStatus)] = d.getEventEmitter(func() events.Event {
		return &events.ReadyUpStatus{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestLaneSelection)] = d.getEventEmitter(func() events.Event {
		return &events.RequestLaneSelection{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRequestMMInfo)] = d.getEventEmitter(func() events.Event {
		return &events.RequestMMInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientRoadToTIQuestDataUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.RoadToTIQuestDataUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientSteamDatagramTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SteamDatagramTicket{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteGCImmediateResponseToInviter{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteReceived{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteResponseReceived{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamsInfo{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTournamentItemDrop)] = d.getEventEmitter(func() events.Event {
		return &events.TournamentItemDrop{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientTrophyAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.TrophyAwarded{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientUnderDraftGoldUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.UnderDraftGoldUpdated{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientVACReminder)] = d.getEventEmitter(func() events.Event {
		return &events.VACReminder{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCToClientWageringUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.WageringUpdate{}
	})
	d.handlers[uint32(protocol.EDOTAGCMsg_k_EMsgGCWatchDownloadedReplay)] = d.getEventEmitter(func() events.Event {
		return &events.WatchDownloadedReplay{}
	})
}
