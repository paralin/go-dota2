package dota2

import (
	"context"
	"github.com/faceit/go-steam/steamid"
	"github.com/paralin/go-dota2/events"
	"github.com/paralin/go-dota2/protocol/base_gcmessages"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_fantasy"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_guild"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_team"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_tournament"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_watch"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	"github.com/paralin/go-dota2/protocol/dota_shared_enums"
	"github.com/paralin/go-dota2/protocol/econ_gcmessages"
)

// AbandonLobby abandons a lobby.
// Request ID: k_EMsgGCAbandonCurrentGame
// Request type: CMsgAbandonCurrentGame
func (d *Dota2) AbandonLobby() {
	req := &dota_gcmessages_client_match_management.CMsgAbandonCurrentGame{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame), req)
}

// ApplyGemCombiner applys a gem combiner.
// Request ID: k_EMsgClientToGCApplyGemCombiner
// Request type: CMsgClientToGCApplyGemCombiner
func (d *Dota2) ApplyGemCombiner(
	itemID1 uint64,
	itemID2 uint64,
) {
	req := &dota_gcmessages_client.CMsgClientToGCApplyGemCombiner{
		ItemId_1: &itemID1,
		ItemId_2: &itemID2,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCApplyGemCombiner), req)
}

// ApplyTeamToLobby applys a team to lobby.
// Request ID: k_EMsgGCApplyTeamToPracticeLobby
// Request type: CMsgApplyTeamToPracticeLobby
func (d *Dota2) ApplyTeamToLobby(
	teamID uint32,
) {
	req := &dota_gcmessages_client_match_management.CMsgApplyTeamToPracticeLobby{
		TeamId: &teamID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCApplyTeamToPracticeLobby), req)
}

// AutographReward autographs a reward.
// Request ID: k_EMsgGameAutographReward
// Response ID: k_EMsgGameAutographRewardResponse
// Request type: CMsgDOTAGameAutographReward
// Response type: CMsgDOTAGameAutographRewardResponse
func (d *Dota2) AutographReward(
	ctx context.Context,
	badgeID string,
) (*dota_gcmessages_client.CMsgDOTAGameAutographRewardResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAGameAutographReward{
		BadgeId: &badgeID,
	}
	resp := &dota_gcmessages_client.CMsgDOTAGameAutographRewardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGameAutographReward),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGameAutographRewardResponse),
		resp,
	)
}

// CancelFantasyTeamTrade cancels a fantasy team trade.
// Request ID: k_EMsgGCFantasyTeamTradeCancelRequest
// Response ID: k_EMsgGCFantasyTeamTradeCancelResponse
// Request type: CMsgDOTAFantasyTeamTradeCancelRequest
// Response type: CMsgDOTAFantasyTeamTradeCancelResponse
func (d *Dota2) CancelFantasyTeamTrade(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex1 uint32,
	ownerAccountID2 uint32,
	teamIndex2 uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradeCancelResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradeCancelRequest{
		FantasyLeagueId:  &fantasyLeagueID,
		TeamIndex_1:      &teamIndex1,
		OwnerAccountId_2: &ownerAccountID2,
		TeamIndex_2:      &teamIndex2,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradeCancelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamTradeCancelRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamTradeCancelResponse),
		resp,
	)
}

// CancelPartyInvites cancels party invites.
// Request ID: k_EMsgClientToGCCancelPartyInvites
// Request type: CMsgDOTACancelGroupInvites
func (d *Dota2) CancelPartyInvites(
	invitedSteamids []uint64,
	invitedGroupids []uint64,
) {
	req := &dota_gcmessages_client_match_management.CMsgDOTACancelGroupInvites{
		InvitedSteamids: invitedSteamids,
		InvitedGroupids: invitedGroupids,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites), req)
}

// CancelWatchGame cancels a watch game.
// Request ID: k_EMsgGCCancelWatchGame
// Request type: CMsgCancelWatchGame
func (d *Dota2) CancelWatchGame() {
	req := &dota_gcmessages_client_watch.CMsgCancelWatchGame{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCancelWatchGame), req)
}

// CastMatchVote votes for/from a cast match.
// Request ID: k_EMsgCastMatchVote
// Request type: CMsgCastMatchVote
func (d *Dota2) CastMatchVote(
	matchID uint64,
	vote dota_shared_enums.DOTAMatchVote,
) {
	req := &dota_gcmessages_client.CMsgCastMatchVote{
		MatchId: &matchID,
		Vote:    &vote,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgCastMatchVote), req)
}

// ClearSuccessfulReportNotification is undocumented.
// Request ID: k_EMsgGCDOTAClearNotifySuccessfulReport
// Request type: CMsgDOTAClearNotifySuccessfulReport
func (d *Dota2) ClearSuccessfulReportNotification() {
	req := &dota_gcmessages_client.CMsgDOTAClearNotifySuccessfulReport{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCDOTAClearNotifySuccessfulReport), req)
}

// CloseLobbyBroadcastChannel closes a lobby broadcast channel.
// Request ID: k_EMsgGCPracticeLobbyCloseBroadcastChannel
// Request type: CMsgPracticeLobbyCloseBroadcastChannel
func (d *Dota2) CloseLobbyBroadcastChannel(
	channel uint32,
) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyCloseBroadcastChannel{
		Channel: &channel,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyCloseBroadcastChannel), req)
}

// CreateBotGame creates a bot game.
// Request ID: k_EMsgGCBotGameCreate
// Request type: CMsgBotGameCreate
func (d *Dota2) CreateBotGame(
	searchKey string,
	difficultyRadiant dota_shared_enums.DOTABotDifficulty,
	team dota_shared_enums.DOTA_GC_TEAM,
	gameMode uint32,
	difficultyDire dota_shared_enums.DOTABotDifficulty,
) {
	req := &dota_gcmessages_client_match_management.CMsgBotGameCreate{
		SearchKey:         &searchKey,
		DifficultyRadiant: &difficultyRadiant,
		Team:              &team,
		GameMode:          &gameMode,
		DifficultyDire:    &difficultyDire,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCBotGameCreate), req)
}

// CreateDOTAStaticRecipe creates a dota static recipe.
// Request ID: k_EMsgClientToGCDOTACreateStaticRecipe
// Response ID: k_EMsgClientToGCDOTACreateStaticRecipeResponse
// Request type: CMsgClientToGCCreateStaticRecipe
// Response type: CMsgClientToGCCreateStaticRecipeResponse
func (d *Dota2) CreateDOTAStaticRecipe(
	ctx context.Context,
	items []*econ_gcmessages.CMsgClientToGCCreateStaticRecipe_Item,
	recipeDefIndex uint32,
) (*econ_gcmessages.CMsgClientToGCCreateStaticRecipeResponse, error) {
	req := &econ_gcmessages.CMsgClientToGCCreateStaticRecipe{
		Items:          items,
		RecipeDefIndex: &recipeDefIndex,
	}
	resp := &econ_gcmessages.CMsgClientToGCCreateStaticRecipeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCDOTACreateStaticRecipe),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCDOTACreateStaticRecipeResponse),
		resp,
	)
}

// CreateFantasyLeague creates a fantasy league.
// Request ID: k_EMsgGCFantasyLeagueCreateRequest
// Response ID: k_EMsgGCFantasyLeagueCreateResponse
// Request type: CMsgDOTAFantasyLeagueCreateRequest
// Response type: CMsgDOTAFantasyLeagueCreateResponse
func (d *Dota2) CreateFantasyLeague(
	ctx context.Context,
	seasonID uint32,
	fantasyLeagueName string,
	password string,
	teamName string,
	logo uint64,
	ticketItemID uint64,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueCreateResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueCreateRequest{
		SeasonId:          &seasonID,
		FantasyLeagueName: &fantasyLeagueName,
		Password:          &password,
		TeamName:          &teamName,
		Logo:              &logo,
		TicketItemId:      &ticketItemID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateResponse),
		resp,
	)
}

// CreateFantasyTeam creates a fantasy team.
// Request ID: k_EMsgGCFantasyTeamCreateRequest
// Response ID: k_EMsgGCFantasyTeamCreateResponse
// Request type: CMsgDOTAFantasyTeamCreateRequest
// Response type: CMsgDOTAFantasyTeamCreateResponse
func (d *Dota2) CreateFantasyTeam(
	ctx context.Context,
	fantasyLeagueID uint32,
	password string,
	teamName string,
	logo uint64,
	ticketItemID uint64,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamCreateResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamCreateRequest{
		FantasyLeagueId: &fantasyLeagueID,
		Password:        &password,
		TeamName:        &teamName,
		Logo:            &logo,
		TicketItemId:    &ticketItemID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamCreateRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamCreateResponse),
		resp,
	)
}

// CreateGameCustom creates a game custom.
// Request ID: k_EMsgGCCustomGameCreate
// Request type: CMsgCustomGameCreate
func (d *Dota2) CreateGameCustom(
	searchKey string,
	difficulty uint32,
	gameMode string,
	gameMap string,
	customGameID uint64,
) {
	req := &dota_gcmessages_client_match_management.CMsgCustomGameCreate{
		SearchKey:    &searchKey,
		Difficulty:   &difficulty,
		GameMode:     &gameMode,
		Map:          &gameMap,
		CustomGameId: &customGameID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCustomGameCreate), req)
}

// CreateGameEvent creates a game event.
// Request ID: k_EMsgGCEventGameCreate
// Request type: CMsgEventGameCreate
func (d *Dota2) CreateGameEvent(
	searchKey string,
	difficulty uint32,
	gameMode string,
	gameMap string,
	customGameID uint64,
) {
	req := &dota_gcmessages_client_match_management.CMsgEventGameCreate{
		SearchKey:    &searchKey,
		Difficulty:   &difficulty,
		GameMode:     &gameMode,
		Map:          &gameMap,
		CustomGameId: &customGameID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCEventGameCreate), req)
}

// CreateGuild creates a guild.
// Request ID: k_EMsgGCGuildCreateRequest
// Response ID: k_EMsgGCGuildCreateResponse
// Request type: CMsgDOTAGuildCreateRequest
// Response type: CMsgDOTAGuildCreateResponse
func (d *Dota2) CreateGuild(
	ctx context.Context,
	name string,
	tag string,
	logo uint64,
	baseLogo uint64,
	bannerLogo uint64,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildCreateResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildCreateRequest{
		Name:       &name,
		Tag:        &tag,
		Logo:       &logo,
		BaseLogo:   &baseLogo,
		BannerLogo: &bannerLogo,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildCreateRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildCreateResponse),
		resp,
	)
}

// CreateHeroStatue creates a hero statue.
// Request ID: k_EMsgClientToGCCreateHeroStatue
// Request type: CMsgClientToGCCreateHeroStatue
func (d *Dota2) CreateHeroStatue(
	req *dota_gcmessages_client.CMsgClientToGCCreateHeroStatue,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCreateHeroStatue), req)
}

// CreatePlayerCardPack creates a player card pack.
// Request ID: k_EMsgClientToGCCreatePlayerCardPack
// Response ID: k_EMsgClientToGCCreatePlayerCardPackResponse
// Request type: CMsgClientToGCCreatePlayerCardPack
// Response type: CMsgClientToGCCreatePlayerCardPackResponse
func (d *Dota2) CreatePlayerCardPack(
	ctx context.Context,
	cardDustItemID uint64,
	eventID uint32,
	premiumPack bool,
) (*dota_gcmessages_client.CMsgClientToGCCreatePlayerCardPackResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCCreatePlayerCardPack{
		CardDustItemId: &cardDustItemID,
		EventId:        &eventID,
		PremiumPack:    &premiumPack,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCCreatePlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse),
		resp,
	)
}

// CreateSpectatorLobby creates a spectator lobby.
// Request ID: k_EMsgClientToGCCreateSpectatorLobby
// Request type: CMsgCreateSpectatorLobby
func (d *Dota2) CreateSpectatorLobby(
	details dota_gcmessages_client_match_management.CMsgSetSpectatorLobbyDetails,
) {
	req := &dota_gcmessages_client_match_management.CMsgCreateSpectatorLobby{
		Details: &details,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobby), req)
}

// CreateTeam creates a team.
// Request ID: k_EMsgGCCreateTeam
// Response ID: k_EMsgGCCreateTeamResponse
// Request type: CMsgDOTACreateTeam
// Response type: CMsgDOTACreateTeamResponse
func (d *Dota2) CreateTeam(
	ctx context.Context,
	req *dota_gcmessages_client_team.CMsgDOTACreateTeam,
) (*dota_gcmessages_client_team.CMsgDOTACreateTeamResponse, error) {
	resp := &dota_gcmessages_client_team.CMsgDOTACreateTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCreateTeam),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCreateTeamResponse),
		resp,
	)
}

// DemotePrivateChatMember demotes a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatDemote
// Request type: CMsgClientToGCPrivateChatDemote
func (d *Dota2) DemotePrivateChatMember(
	privateChatChannelName string,
	demoteAccountID uint32,
) {
	req := &dota_gcmessages_client_chat.CMsgClientToGCPrivateChatDemote{
		PrivateChatChannelName: &privateChatChannelName,
		DemoteAccountId:        &demoteAccountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote), req)
}

// DestroyLobby destroys a lobby.
// Request ID: k_EMsgDestroyLobbyRequest
// Response ID: k_EMsgDestroyLobbyResponse
// Request type: CMsgDOTADestroyLobbyRequest
// Response type: CMsgDOTADestroyLobbyResponse
func (d *Dota2) DestroyLobby(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTADestroyLobbyResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTADestroyLobbyRequest{}
	resp := &dota_gcmessages_client.CMsgDOTADestroyLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDestroyLobbyRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDestroyLobbyResponse),
		resp,
	)
}

// EditFantasyTeam edits a fantasy team.
// Request ID: k_EMsgGCEditFantasyTeamRequest
// Response ID: k_EMsgGCEditFantasyTeamResponse
// Request type: CMsgDOTAEditFantasyTeamRequest
// Response type: CMsgDOTAEditFantasyTeamResponse
func (d *Dota2) EditFantasyTeam(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex uint32,
	teamName string,
	teamLogo uint64,
) (*dota_gcmessages_client_fantasy.CMsgDOTAEditFantasyTeamResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAEditFantasyTeamRequest{
		FantasyLeagueId: &fantasyLeagueID,
		TeamIndex:       &teamIndex,
		TeamName:        &teamName,
		TeamLogo:        &teamLogo,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAEditFantasyTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCEditFantasyTeamRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCEditFantasyTeamResponse),
		resp,
	)
}

// EditTeamDetails edits team details.
// Request ID: k_EMsgGCEditTeamDetails
// Response ID: k_EMsgGCEditTeamDetailsResponse
// Request type: CMsgDOTAEditTeamDetails
// Response type: CMsgDOTAEditTeamDetailsResponse
func (d *Dota2) EditTeamDetails(
	ctx context.Context,
	req *dota_gcmessages_client_team.CMsgDOTAEditTeamDetails,
) (*dota_gcmessages_client_team.CMsgDOTAEditTeamDetailsResponse, error) {
	resp := &dota_gcmessages_client_team.CMsgDOTAEditTeamDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCEditTeamDetails),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCEditTeamDetailsResponse),
		resp,
	)
}

// FindFantasyLeague finds a fantasy league.
// Request ID: k_EMsgDOTAFantasyLeagueFindRequest
// Response ID: k_EMsgDOTAFantasyLeagueFindResponse
// Request type: CMsgDOTAFantasyLeagueFindRequest
// Response type: CMsgDOTAFantasyLeagueFindResponse
func (d *Dota2) FindFantasyLeague(
	ctx context.Context,
	fantasyLeagueID uint32,
	password string,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueFindResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueFindRequest{
		FantasyLeagueId: &fantasyLeagueID,
		Password:        &password,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueFindResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAFantasyLeagueFindRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAFantasyLeagueFindResponse),
		resp,
	)
}

// FindTopSourceTVGames finds top source tv games.
// Request ID: k_EMsgClientToGCFindTopSourceTVGames
// Response ID: k_EMsgGCToClientFindTopSourceTVGamesResponse
// Request type: CMsgClientToGCFindTopSourceTVGames
// Response type: CMsgGCToClientFindTopSourceTVGamesResponse
func (d *Dota2) FindTopSourceTVGames(
	ctx context.Context,
	searchKey string,
	leagueID uint32,
	heroID uint32,
	startGame uint32,
	gameListIndex uint32,
	lobbyIDs []uint64,
) (*dota_gcmessages_client_watch.CMsgGCToClientFindTopSourceTVGamesResponse, error) {
	req := &dota_gcmessages_client_watch.CMsgClientToGCFindTopSourceTVGames{
		SearchKey:     &searchKey,
		LeagueId:      &leagueID,
		HeroId:        &heroID,
		StartGame:     &startGame,
		GameListIndex: &gameListIndex,
		LobbyIds:      lobbyIDs,
	}
	resp := &dota_gcmessages_client_watch.CMsgGCToClientFindTopSourceTVGamesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCFindTopSourceTVGames),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientFindTopSourceTVGamesResponse),
		resp,
	)
}

// FlipLobbyTeams flips lobby teams.
// Request ID: k_EMsgGCFlipLobbyTeams
// Request type: CMsgFlipLobbyTeams
func (d *Dota2) FlipLobbyTeams() {
	req := &dota_gcmessages_client.CMsgFlipLobbyTeams{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFlipLobbyTeams), req)
}

// GetAdditionalEquips gets additional equips.
// Request ID: k_EMsgClientToGCGetAdditionalEquips
// Response ID: k_EMsgClientToGCGetAdditionalEquipsResponse
// Request type: CMsgClientToGCGetAdditionalEquips
// Response type: CMsgClientToGCGetAdditionalEquipsResponse
func (d *Dota2) GetAdditionalEquips(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgClientToGCGetAdditionalEquipsResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetAdditionalEquips{}
	resp := &dota_gcmessages_client.CMsgClientToGCGetAdditionalEquipsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAdditionalEquips),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAdditionalEquipsResponse),
		resp,
	)
}

// GetAllHeroOrder gets all hero order.
// Request ID: k_EMsgClientToGCGetAllHeroOrder
// Response ID: k_EMsgClientToGCGetAllHeroOrderResponse
// Request type: CMsgClientToGCGetAllHeroOrder
// Response type: CMsgClientToGCGetAllHeroOrderResponse
func (d *Dota2) GetAllHeroOrder(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgClientToGCGetAllHeroOrderResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetAllHeroOrder{}
	resp := &dota_gcmessages_client.CMsgClientToGCGetAllHeroOrderResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrder),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroOrderResponse),
		resp,
	)
}

// GetAllHeroProgress gets all hero progress.
// Request ID: k_EMsgClientToGCGetAllHeroProgress
// Response ID: k_EMsgClientToGCGetAllHeroProgressResponse
// Request type: CMsgClientToGCGetAllHeroProgress
// Response type: CMsgClientToGCGetAllHeroProgressResponse
func (d *Dota2) GetAllHeroProgress(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_client.CMsgClientToGCGetAllHeroProgressResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetAllHeroProgress{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCGetAllHeroProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgress),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetAllHeroProgressResponse),
		resp,
	)
}

// GetChatMemberCount gets a chat member count.
// Request ID: k_EMsgDOTAChatGetMemberCount
// Response ID: k_EMsgDOTAChatGetMemberCountResponse
// Request type: CMsgDOTAChatGetMemberCount
// Response type: CMsgDOTAChatGetMemberCountResponse
func (d *Dota2) GetChatMemberCount(
	ctx context.Context,
	channelName string,
	channelType dota_shared_enums.DOTAChatChannelTypeT,
) (*dota_gcmessages_client_chat.CMsgDOTAChatGetMemberCountResponse, error) {
	req := &dota_gcmessages_client_chat.CMsgDOTAChatGetMemberCount{
		ChannelName: &channelName,
		ChannelType: &channelType,
	}
	resp := &dota_gcmessages_client_chat.CMsgDOTAChatGetMemberCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCount),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAChatGetMemberCountResponse),
		resp,
	)
}

// GetChatUserList gets a chat user list.
// Request ID: k_EMsgDOTAChatGetUserList
// Response ID: k_EMsgDOTAChatGetUserListResponse
// Request type: CMsgDOTAChatGetUserList
// Response type: CMsgDOTAChatGetUserListResponse
func (d *Dota2) GetChatUserList(
	ctx context.Context,
	channelID uint64,
) (*dota_gcmessages_client_chat.CMsgDOTAChatGetUserListResponse, error) {
	req := &dota_gcmessages_client_chat.CMsgDOTAChatGetUserList{
		ChannelId: &channelID,
	}
	resp := &dota_gcmessages_client_chat.CMsgDOTAChatGetUserListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAChatGetUserList),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAChatGetUserListResponse),
		resp,
	)
}

// GetEventPoints gets event points.
// Request ID: k_EMsgDOTAGetEventPoints
// Response ID: k_EMsgDOTAGetEventPointsResponse
// Request type: CMsgDOTAGetEventPoints
// Response type: CMsgDOTAGetEventPointsResponse
func (d *Dota2) GetEventPoints(
	ctx context.Context,
	eventID uint32,
	accountID uint32,
) (*dota_gcmessages_client.CMsgDOTAGetEventPointsResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAGetEventPoints{
		EventId:   &eventID,
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgDOTAGetEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetEventPoints),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetEventPointsResponse),
		resp,
	)
}

// GetGiftPermissions gets gift permissions.
// Request ID: k_EMsgClientToGCGetGiftPermissions
// Response ID: k_EMsgClientToGCGetGiftPermissionsResponse
// Request type: CMsgClientToGCGetGiftPermissions
// Response type: CMsgClientToGCGetGiftPermissionsResponse
func (d *Dota2) GetGiftPermissions(
	ctx context.Context,
) (*econ_gcmessages.CMsgClientToGCGetGiftPermissionsResponse, error) {
	req := &econ_gcmessages.CMsgClientToGCGetGiftPermissions{}
	resp := &econ_gcmessages.CMsgClientToGCGetGiftPermissionsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissions),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetGiftPermissionsResponse),
		resp,
	)
}

// GetHeroStandings gets hero standings.
// Request ID: k_EMsgGCGetHeroStandings
// Response ID: k_EMsgGCGetHeroStandingsResponse
// Request type: CMsgGCGetHeroStandings
// Response type: CMsgGCGetHeroStandingsResponse
func (d *Dota2) GetHeroStandings(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCGetHeroStandingsResponse, error) {
	req := &dota_gcmessages_client.CMsgGCGetHeroStandings{}
	resp := &dota_gcmessages_client.CMsgGCGetHeroStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroStandings),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroStandingsResponse),
		resp,
	)
}

// GetHeroStatsHistory gets a hero stats history.
// Request ID: k_EMsgGCGetHeroStatsHistory
// Response ID: k_EMsgGCGetHeroStatsHistoryResponse
// Request type: CMsgGCGetHeroStatsHistory
// Response type: CMsgGCGetHeroStatsHistoryResponse
func (d *Dota2) GetHeroStatsHistory(
	ctx context.Context,
	heroID uint32,
) (*dota_gcmessages_client.CMsgGCGetHeroStatsHistoryResponse, error) {
	req := &dota_gcmessages_client.CMsgGCGetHeroStatsHistory{
		HeroId: &heroID,
	}
	resp := &dota_gcmessages_client.CMsgGCGetHeroStatsHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistory),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroStatsHistoryResponse),
		resp,
	)
}

// GetHeroTimedStats gets hero timed stats.
// Request ID: k_EMsgGCGetHeroTimedStats
// Response ID: k_EMsgGCGetHeroTimedStatsResponse
// Request type: CMsgGCGetHeroTimedStats
// Response type: CMsgGCGetHeroTimedStatsResponse
func (d *Dota2) GetHeroTimedStats(
	ctx context.Context,
	accountID uint32,
	heroID uint32,
) (*dota_gcmessages_client.CMsgGCGetHeroTimedStatsResponse, error) {
	req := &dota_gcmessages_client.CMsgGCGetHeroTimedStats{
		AccountId: &accountID,
		HeroId:    &heroID,
	}
	resp := &dota_gcmessages_client.CMsgGCGetHeroTimedStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroTimedStats),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetHeroTimedStatsResponse),
		resp,
	)
}

// GetLeagueSeries gets league series.
// Request ID: k_EMsgClientToGCGetLeagueSeries
// Response ID: k_EMsgClientToGCGetLeagueSeriesResponse
// Request type: CMsgClientToGCGetLeagueSeries
// Response type: CMsgClientToGCGetLeagueSeriesResponse
func (d *Dota2) GetLeagueSeries(
	ctx context.Context,
	leagueID uint32,
) (*dota_gcmessages_client.CMsgClientToGCGetLeagueSeriesResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetLeagueSeries{
		LeagueId: &leagueID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCGetLeagueSeriesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetLeagueSeries),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetLeagueSeriesResponse),
		resp,
	)
}

// GetPeriodicResource gets a periodic resource.
// Request ID: k_EMsgDOTAGetPeriodicResource
// Response ID: k_EMsgDOTAGetPeriodicResourceResponse
// Request type: CMsgDOTAGetPeriodicResource
// Response type: CMsgDOTAGetPeriodicResourceResponse
func (d *Dota2) GetPeriodicResource(
	ctx context.Context,
	accountID uint32,
	periodicResourceID uint32,
	timestamp uint32,
) (*dota_gcmessages_client.CMsgDOTAGetPeriodicResourceResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAGetPeriodicResource{
		AccountId:          &accountID,
		PeriodicResourceId: &periodicResourceID,
		Timestamp:          &timestamp,
	}
	resp := &dota_gcmessages_client.CMsgDOTAGetPeriodicResourceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResource),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetPeriodicResourceResponse),
		resp,
	)
}

// GetPlayerCardItemInfo gets a player card item info.
// Request ID: k_EMsgGCGetPlayerCardItemInfo
// Response ID: k_EMsgGCGetPlayerCardItemInfoResponse
// Request type: CMsgGCGetPlayerCardItemInfo
// Response type: CMsgGCGetPlayerCardItemInfoResponse
func (d *Dota2) GetPlayerCardItemInfo(
	ctx context.Context,
	accountID uint32,
	playerCardItemIDs []uint64,
) (*dota_gcmessages_common.CMsgGCGetPlayerCardItemInfoResponse, error) {
	req := &dota_gcmessages_common.CMsgGCGetPlayerCardItemInfo{
		AccountId:         &accountID,
		PlayerCardItemIds: playerCardItemIDs,
	}
	resp := &dota_gcmessages_common.CMsgGCGetPlayerCardItemInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfo),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGetPlayerCardItemInfoResponse),
		resp,
	)
}

// GetPlayerMatchHistory gets a player match history.
// Request ID: k_EMsgDOTAGetPlayerMatchHistory
// Response ID: k_EMsgDOTAGetPlayerMatchHistoryResponse
// Request type: CMsgDOTAGetPlayerMatchHistory
// Response type: CMsgDOTAGetPlayerMatchHistoryResponse
func (d *Dota2) GetPlayerMatchHistory(
	ctx context.Context,
	req *dota_gcmessages_client.CMsgDOTAGetPlayerMatchHistory,
) (*dota_gcmessages_client.CMsgDOTAGetPlayerMatchHistoryResponse, error) {
	resp := &dota_gcmessages_client.CMsgDOTAGetPlayerMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistory),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistoryResponse),
		resp,
	)
}

// GetProfileCard gets a profile card.
// Request ID: k_EMsgClientToGCGetProfileCard
// Response ID: k_EMsgClientToGCGetProfileCardResponse
// Request type: CMsgClientToGCGetProfileCard
// Response type: CMsgDOTAProfileCard
func (d *Dota2) GetProfileCard(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_common.CMsgDOTAProfileCard, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetProfileCard{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_common.CMsgDOTAProfileCard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse),
		resp,
	)
}

// GetProfileTickets gets profile tickets.
// Request ID: k_EMsgClientToGCGetProfileTickets
// Response ID: k_EMsgClientToGCGetProfileTicketsResponse
// Request type: CMsgClientToGCGetProfileTickets
// Response type: CMsgDOTAProfileTickets
func (d *Dota2) GetProfileTickets(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_client.CMsgDOTAProfileTickets, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetProfileTickets{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgDOTAProfileTickets{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetProfileTickets),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse),
		resp,
	)
}

// GetQuestProgress gets quest progress.
// Request ID: k_EMsgClientToGCGetQuestProgress
// Response ID: k_EMsgClientToGCGetQuestProgressResponse
// Request type: CMsgClientToGCGetQuestProgress
// Response type: CMsgClientToGCGetQuestProgressResponse
func (d *Dota2) GetQuestProgress(
	ctx context.Context,
	questIDs []uint32,
) (*dota_gcmessages_client.CMsgClientToGCGetQuestProgressResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetQuestProgress{
		QuestIds: questIDs,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCGetQuestProgressResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgress),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetQuestProgressResponse),
		resp,
	)
}

// GetTourneyWeekendPlayerStats gets tourney weekend player stats.
// Request ID: k_EMsgClientToGCWeekendTourneyGetPlayerStats
// Response ID: k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse
// Request type: CMsgDOTAWeekendTourneyPlayerStatsRequest
// Response type: CMsgDOTAWeekendTourneyPlayerStats
func (d *Dota2) GetTourneyWeekendPlayerStats(
	ctx context.Context,
	accountID uint32,
	seasonTrophyID uint32,
) (*dota_gcmessages_client_tournament.CMsgDOTAWeekendTourneyPlayerStats, error) {
	req := &dota_gcmessages_client_tournament.CMsgDOTAWeekendTourneyPlayerStatsRequest{
		AccountId:      &accountID,
		SeasonTrophyId: &seasonTrophyID,
	}
	resp := &dota_gcmessages_client_tournament.CMsgDOTAWeekendTourneyPlayerStats{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse),
		resp,
	)
}

// GetWeekendTourneySchedule gets a weekend tourney schedule.
// Request ID: k_EMsgDOTAGetWeekendTourneySchedule
// Request type: CMsgRequestWeekendTourneySchedule
func (d *Dota2) GetWeekendTourneySchedule() {
	req := &dota_gcmessages_client_tournament.CMsgRequestWeekendTourneySchedule{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule), req)
}

// GiveTip gives a tip.
// Request ID: k_EMsgClientToGCGiveTip
// Response ID: k_EMsgClientToGCGiveTipResponse
// Request type: CMsgClientToGCGiveTip
// Response type: CMsgClientToGCGiveTipResponse
func (d *Dota2) GiveTip(
	ctx context.Context,
	recipientAccountID uint32,
	matchID uint64,
	eventID uint32,
) (*dota_gcmessages_client.CMsgClientToGCGiveTipResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGiveTip{
		RecipientAccountId: &recipientAccountID,
		MatchId:            &matchID,
		EventId:            &eventID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCGiveTipResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGiveTip),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGiveTipResponse),
		resp,
	)
}

// InvitePlayerToTeam is undocumented.
// Request ID: k_EMsgGCTeamInvite_InviterToGC
// Response ID: k_EMsgGCTeamInvite_GCImmediateResponseToInviter
// Request type: CMsgDOTATeamInvite_InviterToGC
// Response type: CMsgDOTATeamInvite_GCImmediateResponseToInviter
func (d *Dota2) InvitePlayerToTeam(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*dota_gcmessages_client_team.CMsgDOTATeamInvite_GCImmediateResponseToInviter, error) {
	req := &dota_gcmessages_client_team.CMsgDOTATeamInvite_InviterToGC{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &dota_gcmessages_client_team.CMsgDOTATeamInvite_GCImmediateResponseToInviter{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter),
		resp,
	)
}

// InvitePrivateChatMember is undocumented.
// Request ID: k_EMsgClientToGCPrivateChatInvite
// Request type: CMsgClientToGCPrivateChatInvite
func (d *Dota2) InvitePrivateChatMember(
	privateChatChannelName string,
	invitedAccountID uint32,
) {
	req := &dota_gcmessages_client_chat.CMsgClientToGCPrivateChatInvite{
		PrivateChatChannelName: &privateChatChannelName,
		InvitedAccountId:       &invitedAccountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite), req)
}

// JoinChatChannel joins a chat channel.
// Request ID: k_EMsgGCJoinChatChannel
// Response ID: k_EMsgGCJoinChatChannelResponse
// Request type: CMsgDOTAJoinChatChannel
// Response type: CMsgDOTAJoinChatChannelResponse
func (d *Dota2) JoinChatChannel(
	ctx context.Context,
	channelName string,
	channelType dota_shared_enums.DOTAChatChannelTypeT,
) (*dota_gcmessages_client_chat.CMsgDOTAJoinChatChannelResponse, error) {
	req := &dota_gcmessages_client_chat.CMsgDOTAJoinChatChannel{
		ChannelName: &channelName,
		ChannelType: &channelType,
	}
	resp := &dota_gcmessages_client_chat.CMsgDOTAJoinChatChannelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinChatChannel),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinChatChannelResponse),
		resp,
	)
}

// JoinLobby joins a lobby.
// Request ID: k_EMsgGCPracticeLobbyJoin
// Response ID: k_EMsgGCPracticeLobbyJoinResponse
// Request type: CMsgPracticeLobbyJoin
// Response type: CMsgPracticeLobbyJoinResponse
func (d *Dota2) JoinLobby(
	ctx context.Context,
	lobbyID uint64,
	passKey string,
	customGameCrc uint64,
	customGameTimestamp uint32,
) (*dota_gcmessages_client_match_management.CMsgPracticeLobbyJoinResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyJoin{
		LobbyId:             &lobbyID,
		PassKey:             &passKey,
		CustomGameCrc:       &customGameCrc,
		CustomGameTimestamp: &customGameTimestamp,
	}
	resp := &dota_gcmessages_client_match_management.CMsgPracticeLobbyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoin),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinResponse),
		resp,
	)
}

// JoinLobbyBroadcastChannel joins a lobby broadcast channel.
// Request ID: k_EMsgGCPracticeLobbyJoinBroadcastChannel
// Request type: CMsgPracticeLobbyJoinBroadcastChannel
func (d *Dota2) JoinLobbyBroadcastChannel(
	channel uint32,
	preferredDescription string,
	preferredCountryCode string,
	preferredLanguageCode string,
) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyJoinBroadcastChannel{
		Channel:               &channel,
		PreferredDescription:  &preferredDescription,
		PreferredCountryCode:  &preferredCountryCode,
		PreferredLanguageCode: &preferredLanguageCode,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoinBroadcastChannel), req)
}

// JoinPlaytest joins a playtest.
// Request ID: k_EMsgClientToGCJoinPlaytest
// Response ID: k_EMsgClientToGCJoinPlaytestResponse
// Request type: CMsgClientToGCJoinPlaytest
// Response type: CMsgClientToGCJoinPlaytestResponse
func (d *Dota2) JoinPlaytest(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgClientToGCJoinPlaytestResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCJoinPlaytest{}
	resp := &dota_gcmessages_client.CMsgClientToGCJoinPlaytestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCJoinPlaytestResponse),
		resp,
	)
}

// JoinQuickCustomLobby joins a quick custom lobby.
// Request ID: k_EMsgGCQuickJoinCustomLobby
// Response ID: k_EMsgGCQuickJoinCustomLobbyResponse
// Request type: CMsgQuickJoinCustomLobby
// Response type: CMsgQuickJoinCustomLobbyResponse
func (d *Dota2) JoinQuickCustomLobby(
	ctx context.Context,
	legacyServerRegion uint32,
	customGameID uint64,
	createLobbyDetails dota_gcmessages_client_match_management.CMsgPracticeLobbySetDetails,
	allowAnyMap bool,
	legacyRegionPings []*dota_gcmessages_client_match_management.CMsgQuickJoinCustomLobby_LegacyRegionPing,
	pingData base_gcmessages.CMsgClientPingData,
) (*dota_gcmessages_client_match_management.CMsgQuickJoinCustomLobbyResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgQuickJoinCustomLobby{
		LegacyServerRegion: &legacyServerRegion,
		CustomGameId:       &customGameID,
		CreateLobbyDetails: &createLobbyDetails,
		AllowAnyMap:        &allowAnyMap,
		LegacyRegionPings:  legacyRegionPings,
		PingData:           &pingData,
	}
	resp := &dota_gcmessages_client_match_management.CMsgQuickJoinCustomLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobby),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCQuickJoinCustomLobbyResponse),
		resp,
	)
}

// KickLobbyMember kicks a lobby member.
// Request ID: k_EMsgGCPracticeLobbyKick
// Request type: CMsgPracticeLobbyKick
func (d *Dota2) KickLobbyMember(
	accountID uint32,
) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyKick{
		AccountId: &accountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick), req)
}

// KickLobbyMemberFromTeam kicks a lobby member from team.
// Request ID: k_EMsgGCPracticeLobbyKickFromTeam
// Request type: CMsgPracticeLobbyKickFromTeam
func (d *Dota2) KickLobbyMemberFromTeam(
	accountID uint32,
) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyKickFromTeam{
		AccountId: &accountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam), req)
}

// KickPrivateChatMember kicks a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatKick
// Request type: CMsgClientToGCPrivateChatKick
func (d *Dota2) KickPrivateChatMember(
	privateChatChannelName string,
	kickAccountID uint32,
) {
	req := &dota_gcmessages_client_chat.CMsgClientToGCPrivateChatKick{
		PrivateChatChannelName: &privateChatChannelName,
		KickAccountId:          &kickAccountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick), req)
}

// KickTeamMember kicks a team member.
// Request ID: k_EMsgGCKickTeamMember
// Response ID: k_EMsgGCKickTeamMemberResponse
// Request type: CMsgDOTAKickTeamMember
// Response type: CMsgDOTAKickTeamMemberResponse
func (d *Dota2) KickTeamMember(
	ctx context.Context,
	accountID uint32,
	teamID uint32,
) (*dota_gcmessages_client_team.CMsgDOTAKickTeamMemberResponse, error) {
	req := &dota_gcmessages_client_team.CMsgDOTAKickTeamMember{
		AccountId: &accountID,
		TeamId:    &teamID,
	}
	resp := &dota_gcmessages_client_team.CMsgDOTAKickTeamMemberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCKickTeamMember),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCKickTeamMemberResponse),
		resp,
	)
}

// LaunchLobby launchs a lobby.
// Request ID: k_EMsgGCPracticeLobbyLaunch
// Request type: CMsgPracticeLobbyLaunch
func (d *Dota2) LaunchLobby() {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyLaunch{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyLaunch), req)
}

// LeaveChatChannel leaves a chat channel.
// Request ID: k_EMsgGCLeaveChatChannel
// Request type: CMsgDOTALeaveChatChannel
func (d *Dota2) LeaveChatChannel(
	channelID uint64,
) {
	req := &dota_gcmessages_client_chat.CMsgDOTALeaveChatChannel{
		ChannelId: &channelID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeaveChatChannel), req)
}

// LeaveLobby leaves a lobby.
// Request ID: k_EMsgGCPracticeLobbyLeave
// Request type: CMsgPracticeLobbyLeave
func (d *Dota2) LeaveLobby() {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyLeave{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyLeave), req)
}

// LeaveTeam leaves a team.
// Request ID: k_EMsgGCLeaveTeam
// Response ID: k_EMsgGCLeaveTeamResponse
// Request type: CMsgDOTALeaveTeam
// Response type: CMsgDOTALeaveTeamResponse
func (d *Dota2) LeaveTeam(
	ctx context.Context,
	teamID uint32,
) (*dota_gcmessages_client_team.CMsgDOTALeaveTeamResponse, error) {
	req := &dota_gcmessages_client_team.CMsgDOTALeaveTeam{
		TeamId: &teamID,
	}
	resp := &dota_gcmessages_client_team.CMsgDOTALeaveTeamResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeaveTeam),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeaveTeamResponse),
		resp,
	)
}

// LeaveTourneyWeekend leaves a tourney weekend.
// Request ID: k_EMsgClientToGCWeekendTourneyLeave
// Request type: CMsgWeekendTourneyLeave
func (d *Dota2) LeaveTourneyWeekend() {
	req := &dota_gcmessages_client_tournament.CMsgWeekendTourneyLeave{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave), req)
}

// ListChatChannel lists a chat channel.
// Request ID: k_EMsgGCRequestChatChannelList
// Response ID: k_EMsgGCRequestChatChannelListResponse
// Request type: CMsgDOTARequestChatChannelList
// Response type: CMsgDOTARequestChatChannelListResponse
func (d *Dota2) ListChatChannel(
	ctx context.Context,
) (*dota_gcmessages_client_chat.CMsgDOTARequestChatChannelListResponse, error) {
	req := &dota_gcmessages_client_chat.CMsgDOTARequestChatChannelList{}
	resp := &dota_gcmessages_client_chat.CMsgDOTARequestChatChannelListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestChatChannelList),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestChatChannelListResponse),
		resp,
	)
}

// ListCustomGamesTop lists a custom games top.
// Request ID: k_EMsgGCTopCustomGamesList
// Request type: CMsgGCTopCustomGamesList
func (d *Dota2) ListCustomGamesTop(
	topCustomGames []uint64,
	gameOfTheDay uint64,
) {
	req := &dota_gcmessages_common.CMsgGCTopCustomGamesList{
		TopCustomGames: topCustomGames,
		GameOfTheDay:   &gameOfTheDay,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTopCustomGamesList), req)
}

// ListFriendLobby lists a friend lobby.
// Request ID: k_EMsgGCFriendPracticeLobbyListRequest
// Response ID: k_EMsgGCFriendPracticeLobbyListResponse
// Request type: CMsgFriendPracticeLobbyListRequest
// Response type: CMsgFriendPracticeLobbyListResponse
func (d *Dota2) ListFriendLobby(
	ctx context.Context,
	friends []uint32,
) (*dota_gcmessages_client_match_management.CMsgFriendPracticeLobbyListResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgFriendPracticeLobbyListRequest{
		Friends: friends,
	}
	resp := &dota_gcmessages_client_match_management.CMsgFriendPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFriendPracticeLobbyListResponse),
		resp,
	)
}

// ListGuildmateLobby lists a guildmate lobby.
// Request ID: k_EMsgGCGuildmatePracticeLobbyListRequest
// Response ID: k_EMsgGCGuildmatePracticeLobbyListResponse
// Request type: CMsgGuildmatePracticeLobbyListRequest
// Response type: CMsgGuildmatePracticeLobbyListResponse
func (d *Dota2) ListGuildmateLobby(
	ctx context.Context,
	guilds []uint32,
) (*dota_gcmessages_client_match_management.CMsgGuildmatePracticeLobbyListResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgGuildmatePracticeLobbyListRequest{
		Guilds: guilds,
	}
	resp := &dota_gcmessages_client_match_management.CMsgGuildmatePracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildmatePracticeLobbyListRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildmatePracticeLobbyListResponse),
		resp,
	)
}

// ListLobbies lists lobbies.
// Request ID: k_EMsgGCPracticeLobbyList
// Response ID: k_EMsgGCPracticeLobbyListResponse
// Request type: CMsgPracticeLobbyList
// Response type: CMsgPracticeLobbyListResponse
func (d *Dota2) ListLobbies(
	ctx context.Context,
	tournamentGames bool,
	passKey string,
	region uint32,
	gameMode dota_shared_enums.DOTA_GameMode,
) (*dota_gcmessages_client_match_management.CMsgPracticeLobbyListResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyList{
		TournamentGames: &tournamentGames,
		PassKey:         &passKey,
		Region:          &region,
		GameMode:        &gameMode,
	}
	resp := &dota_gcmessages_client_match_management.CMsgPracticeLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyList),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyListResponse),
		resp,
	)
}

// ListLobbySpectator lists a lobby spectator.
// Request ID: k_EMsgClientToGCSpectatorLobbyList
// Response ID: k_EMsgClientToGCSpectatorLobbyListResponse
// Request type: CMsgSpectatorLobbyList
// Response type: CMsgSpectatorLobbyListResponse
func (d *Dota2) ListLobbySpectator(
	ctx context.Context,
) (*dota_gcmessages_client_match_management.CMsgSpectatorLobbyListResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgSpectatorLobbyList{}
	resp := &dota_gcmessages_client_match_management.CMsgSpectatorLobbyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyList),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSpectatorLobbyListResponse),
		resp,
	)
}

// ListNotificationMarkRead lists a notification mark read.
// Request ID: k_EMsgClientToGCMarkNotificationListRead
// Request type: CMsgClientToGCMarkNotificationListRead
func (d *Dota2) ListNotificationMarkRead(
	notificationIDs []uint64,
) {
	req := &dota_gcmessages_client.CMsgClientToGCMarkNotificationListRead{
		NotificationIds: notificationIDs,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCMarkNotificationListRead), req)
}

// ListProTeam lists a pro team.
// Request ID: k_EMsgGCProTeamListRequest
// Response ID: k_EMsgGCProTeamListResponse
// Request type: CMsgDOTAProTeamListRequest
// Response type: CMsgDOTAProTeamListResponse
func (d *Dota2) ListProTeam(
	ctx context.Context,
) (*dota_gcmessages_client_team.CMsgDOTAProTeamListResponse, error) {
	req := &dota_gcmessages_client_team.CMsgDOTAProTeamListRequest{}
	resp := &dota_gcmessages_client_team.CMsgDOTAProTeamListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProTeamListRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProTeamListResponse),
		resp,
	)
}

// ListTrophies lists trophies.
// Request ID: k_EMsgClientToGCGetTrophyList
// Response ID: k_EMsgClientToGCGetTrophyListResponse
// Request type: CMsgClientToGCGetTrophyList
// Response type: CMsgClientToGCGetTrophyListResponse
func (d *Dota2) ListTrophies(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_client.CMsgClientToGCGetTrophyListResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetTrophyList{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCGetTrophyListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetTrophyListResponse),
		resp,
	)
}

// OpenGuildPartyRefresh opens a guild party refresh.
// Request ID: k_EMsgGCGuildOpenPartyRefresh
// Request type: CMsgDOTAGuildOpenPartyRefresh
func (d *Dota2) OpenGuildPartyRefresh(
	guildID uint32,
	openParties []*dota_gcmessages_client_guild.CMsgDOTAGuildOpenPartyRefresh_OpenParty,
) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildOpenPartyRefresh{
		GuildId:     &guildID,
		OpenParties: openParties,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildOpenPartyRefresh), req)
}

// OpenPlayerCardPack opens a player card pack.
// Request ID: k_EMsgClientToGCOpenPlayerCardPack
// Response ID: k_EMsgClientToGCOpenPlayerCardPackResponse
// Request type: CMsgClientToGCOpenPlayerCardPack
// Response type: CMsgClientToGCOpenPlayerCardPackResponse
func (d *Dota2) OpenPlayerCardPack(
	ctx context.Context,
	playerCardPackItemID uint64,
) (*dota_gcmessages_client.CMsgClientToGCOpenPlayerCardPackResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCOpenPlayerCardPack{
		PlayerCardPackItemId: &playerCardPackItemID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCOpenPlayerCardPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPack),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCOpenPlayerCardPackResponse),
		resp,
	)
}

// PromotePrivateChatMember promotes a private chat member.
// Request ID: k_EMsgClientToGCPrivateChatPromote
// Request type: CMsgClientToGCPrivateChatPromote
func (d *Dota2) PromotePrivateChatMember(
	privateChatChannelName string,
	promoteAccountID uint32,
) {
	req := &dota_gcmessages_client_chat.CMsgClientToGCPrivateChatPromote{
		PrivateChatChannelName: &privateChatChannelName,
		PromoteAccountId:       &promoteAccountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote), req)
}

// PublishUserStat publishs a user stat.
// Request ID: k_EMsgClientToGCPublishUserStat
// Request type: CMsgClientToGCPublishUserStat
func (d *Dota2) PublishUserStat(
	userStatsEvent uint32,
	referenceData uint64,
) {
	req := &dota_gcmessages_client.CMsgClientToGCPublishUserStat{
		UserStatsEvent: &userStatsEvent,
		ReferenceData:  &referenceData,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPublishUserStat), req)
}

// PurchaseItemWithEventPoints purchases item with event points.
// Request ID: k_EMsgPurchaseItemWithEventPoints
// Response ID: k_EMsgPurchaseItemWithEventPointsResponse
// Request type: CMsgPurchaseItemWithEventPoints
// Response type: CMsgPurchaseItemWithEventPointsResponse
func (d *Dota2) PurchaseItemWithEventPoints(
	ctx context.Context,
	itemDef uint32,
	quantity uint32,
	eventID dota_shared_enums.EEvent,
	usePremiumPoints bool,
) (*dota_gcmessages_client.CMsgPurchaseItemWithEventPointsResponse, error) {
	req := &dota_gcmessages_client.CMsgPurchaseItemWithEventPoints{
		ItemDef:          &itemDef,
		Quantity:         &quantity,
		EventId:          &eventID,
		UsePremiumPoints: &usePremiumPoints,
	}
	resp := &dota_gcmessages_client.CMsgPurchaseItemWithEventPointsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPoints),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgPurchaseItemWithEventPointsResponse),
		resp,
	)
}

// QueryHasItem queries to check if the target has item.
// Request ID: k_EMsgGCHasItemQuery
// Request type: CMsgDOTAHasItemQuery
func (d *Dota2) QueryHasItem(
	accountID uint32,
	itemID uint64,
) {
	req := &dota_gcmessages_common.CMsgDOTAHasItemQuery{
		AccountId: &accountID,
		ItemId:    &itemID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHasItemQuery), req)
}

// QueryHasItemDefs queries to check if the target has item defs.
// Request ID: k_EMsgGCHasItemDefsQuery
// Request type: CMsgDOTAHasItemDefsQuery
func (d *Dota2) QueryHasItemDefs(
	accountID uint32,
	itemdefIDs []uint32,
) {
	req := &dota_gcmessages_common.CMsgDOTAHasItemDefsQuery{
		AccountId:  &accountID,
		ItemdefIds: itemdefIDs,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHasItemDefsQuery), req)
}

// QueryIsPro queries to check if the target is pro.
// Request ID: k_EMsgGCIsProQuery
// Request type: CMsgGCIsProQuery
func (d *Dota2) QueryIsPro(
	accountID uint32,
) {
	req := &dota_gcmessages_common.CMsgGCIsProQuery{
		AccountId: &accountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCIsProQuery), req)
}

// RecordCompendiumStats records compendium stats.
// Request ID: k_EMsgClientToGCRecordCompendiumStats
// Request type: CMsgClientToGCRecordCompendiumStats
func (d *Dota2) RecordCompendiumStats(
	leagueID uint32,
	viewDurationS uint32,
	videosViewed uint32,
	pageTurns uint32,
	linksFollowed uint32,
) {
	req := &dota_gcmessages_client.CMsgClientToGCRecordCompendiumStats{
		LeagueId:      &leagueID,
		ViewDurationS: &viewDurationS,
		VideosViewed:  &videosViewed,
		PageTurns:     &pageTurns,
		LinksFollowed: &linksFollowed,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRecordCompendiumStats), req)
}

// RecyclePlayerCard recycles a player card.
// Request ID: k_EMsgClientToGCRecyclePlayerCard
// Response ID: k_EMsgClientToGCRecyclePlayerCardResponse
// Request type: CMsgClientToGCRecyclePlayerCard
// Response type: CMsgClientToGCRecyclePlayerCardResponse
func (d *Dota2) RecyclePlayerCard(
	ctx context.Context,
	playerCardItemIDs []uint64,
	eventID uint32,
) (*dota_gcmessages_client.CMsgClientToGCRecyclePlayerCardResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRecyclePlayerCard{
		PlayerCardItemIds: playerCardItemIDs,
		EventId:           &eventID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCRecyclePlayerCardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCard),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRecyclePlayerCardResponse),
		resp,
	)
}

// RedeemItem redeems a item.
// Request ID: k_EMsgDOTARedeemItem
// Response ID: k_EMsgDOTARedeemItemResponse
// Request type: CMsgDOTARedeemItem
// Response type: CMsgDOTARedeemItemResponse
func (d *Dota2) RedeemItem(
	ctx context.Context,
	currencyID uint64,
	purchaseDef uint32,
) (*dota_gcmessages_client.CMsgDOTARedeemItemResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTARedeemItem{
		CurrencyId:  &currencyID,
		PurchaseDef: &purchaseDef,
	}
	resp := &dota_gcmessages_client.CMsgDOTARedeemItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTARedeemItem),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTARedeemItemResponse),
		resp,
	)
}

// RefreshPartnerAccountLink refreshs a partner account link.
// Request ID: k_EMsgRefreshPartnerAccountLink
// Request type: CMsgRefreshPartnerAccountLink
func (d *Dota2) RefreshPartnerAccountLink(
	partnerType int32,
) {
	req := &dota_gcmessages_client.CMsgRefreshPartnerAccountLink{
		PartnerType: &partnerType,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgRefreshPartnerAccountLink), req)
}

// RejoinAllChatChannels is undocumented.
// Request ID: k_EMsgClientsRejoinChatChannels
// Request type: CMsgClientsRejoinChatChannels
func (d *Dota2) RejoinAllChatChannels() {
	req := &dota_gcmessages_client.CMsgClientsRejoinChatChannels{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels), req)
}

// ReleaseEditorItemReservation releases a editor item reservation.
// Request ID: k_EMsgGCItemEditorReleaseReservation
// Response ID: k_EMsgGCItemEditorReleaseReservationResponse
// Request type: CMsgGCItemEditorReleaseReservation
// Response type: CMsgGCItemEditorReleaseReservationResponse
func (d *Dota2) ReleaseEditorItemReservation(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*dota_gcmessages_client.CMsgGCItemEditorReleaseReservationResponse, error) {
	req := &dota_gcmessages_client.CMsgGCItemEditorReleaseReservation{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &dota_gcmessages_client.CMsgGCItemEditorReleaseReservationResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservation),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReleaseReservationResponse),
		resp,
	)
}

// ReportChatPublicSpam reports a chat public spam.
// Request ID: k_EMsgGCChatReportPublicSpam
// Request type: CMsgGCChatReportPublicSpam
func (d *Dota2) ReportChatPublicSpam(
	channelID uint64,
	channelUserID uint32,
) {
	req := &dota_gcmessages_client_chat.CMsgGCChatReportPublicSpam{
		ChannelId:     &channelID,
		ChannelUserId: &channelUserID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCChatReportPublicSpam), req)
}

// RequestAnchorPhoneNumber requests a anchor phone number.
// Request ID: k_EMsgAnchorPhoneNumberRequest
// Response ID: k_EMsgAnchorPhoneNumberResponse
// Request type: CMsgDOTAAnchorPhoneNumberRequest
// Response type: CMsgDOTAAnchorPhoneNumberResponse
func (d *Dota2) RequestAnchorPhoneNumber(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTAAnchorPhoneNumberResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAAnchorPhoneNumberRequest{}
	resp := &dota_gcmessages_client.CMsgDOTAAnchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgAnchorPhoneNumberRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgAnchorPhoneNumberResponse),
		resp,
	)
}

// RequestArcanaVotesRemaining requests a arcana votes remaining.
// Request ID: k_EMsgClientToGCRequestArcanaVotesRemaining
// Response ID: k_EMsgClientToGCRequestArcanaVotesRemainingResponse
// Request type: CMsgClientToGCRequestArcanaVotesRemaining
// Response type: CMsgClientToGCRequestArcanaVotesRemainingResponse
func (d *Dota2) RequestArcanaVotesRemaining(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgClientToGCRequestArcanaVotesRemainingResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRequestArcanaVotesRemaining{}
	resp := &dota_gcmessages_client.CMsgClientToGCRequestArcanaVotesRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemaining),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestArcanaVotesRemainingResponse),
		resp,
	)
}

// RequestCompendiumData requests a compendium data.
// Request ID: k_EMsgGCCompendiumDataRequest
// Response ID: k_EMsgGCCompendiumDataResponse
// Request type: CMsgDOTACompendiumDataRequest
// Response type: CMsgDOTACompendiumDataResponse
func (d *Dota2) RequestCompendiumData(
	ctx context.Context,
	accountID uint32,
	leagueid uint32,
) (*dota_gcmessages_client.CMsgDOTACompendiumDataResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTACompendiumDataRequest{
		AccountId: &accountID,
		Leagueid:  &leagueid,
	}
	resp := &dota_gcmessages_client.CMsgDOTACompendiumDataResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumDataRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumDataResponse),
		resp,
	)
}

// RequestCustomGamePlayerCount requests a custom game player count.
// Request ID: k_EMsgClientToGCCustomGamePlayerCountRequest
// Response ID: k_EMsgGCToClientCustomGamePlayerCountResponse
// Request type: CMsgClientToGCCustomGamePlayerCountRequest
// Response type: CMsgGCToClientCustomGamePlayerCountResponse
func (d *Dota2) RequestCustomGamePlayerCount(
	ctx context.Context,
	customGameID uint64,
) (*dota_gcmessages_client.CMsgGCToClientCustomGamePlayerCountResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCCustomGamePlayerCountRequest{
		CustomGameId: &customGameID,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientCustomGamePlayerCountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCustomGamePlayerCountRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientCustomGamePlayerCountResponse),
		resp,
	)
}

// RequestCustomGamesFriendsPlayed requests a custom games friends played.
// Request ID: k_EMsgClientToGCCustomGamesFriendsPlayedRequest
// Response ID: k_EMsgGCToClientCustomGamesFriendsPlayedResponse
// Request type: CMsgClientToGCCustomGamesFriendsPlayedRequest
// Response type: CMsgGCToClientCustomGamesFriendsPlayedResponse
func (d *Dota2) RequestCustomGamesFriendsPlayed(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCToClientCustomGamesFriendsPlayedResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCCustomGamesFriendsPlayedRequest{}
	resp := &dota_gcmessages_client.CMsgGCToClientCustomGamesFriendsPlayedResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCCustomGamesFriendsPlayedRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientCustomGamesFriendsPlayedResponse),
		resp,
	)
}

// RequestEditorItemLeagueInfo requests a editor item league info.
// Request ID: k_EMsgGCItemEditorRequestLeagueInfo
// Request type: CMsgGCItemEditorRequestLeagueInfo
func (d *Dota2) RequestEditorItemLeagueInfo(
	leagueID uint32,
) {
	req := &dota_gcmessages_client.CMsgGCItemEditorRequestLeagueInfo{
		LeagueId: &leagueID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorRequestLeagueInfo), req)
}

// RequestEmoticonData requests a emoticon data.
// Request ID: k_EMsgClientToGCEmoticonDataRequest
// Response ID: k_EMsgGCToClientEmoticonData
// Request type: CMsgClientToGCEmoticonDataRequest
// Response type: CMsgGCToClientEmoticonData
func (d *Dota2) RequestEmoticonData(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCToClientEmoticonData, error) {
	req := &dota_gcmessages_client.CMsgClientToGCEmoticonDataRequest{}
	resp := &dota_gcmessages_client.CMsgGCToClientEmoticonData{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCEmoticonDataRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientEmoticonData),
		resp,
	)
}

// RequestEventGoals requests event goals.
// Request ID: k_EMsgClientToGCEventGoalsRequest
// Response ID: k_EMsgClientToGCEventGoalsResponse
// Request type: CMsgClientToGCGetEventGoals
// Response type: CMsgEventGoals
func (d *Dota2) RequestEventGoals(
	ctx context.Context,
	eventIDs []dota_shared_enums.EEvent,
) (*dota_gcmessages_client.CMsgEventGoals, error) {
	req := &dota_gcmessages_client.CMsgClientToGCGetEventGoals{
		EventIds: eventIDs,
	}
	resp := &dota_gcmessages_client.CMsgEventGoals{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse),
		resp,
	)
}

// RequestEventPointLog requests a event point log.
// Request ID: k_EMsgClientToGCRequestEventPointLog
// Response ID: k_EMsgClientToGCRequestEventPointLogResponse
// Request type: CMsgClientToGCRequestEventPointLog
// Response type: CMsgClientToGCRequestEventPointLogResponse
func (d *Dota2) RequestEventPointLog(
	ctx context.Context,
	eventID uint32,
) (*dota_gcmessages_client.CMsgClientToGCRequestEventPointLogResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRequestEventPointLog{
		EventId: &eventID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCRequestEventPointLogResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLog),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestEventPointLogResponse),
		resp,
	)
}

// RequestFantasyLeagueDraftPlayer requests a fantasy league draft player.
// Request ID: k_EMsgGCFantasyLeagueDraftPlayerRequest
// Response ID: k_EMsgGCFantasyLeagueDraftPlayerResponse
// Request type: CMsgDOTAFantasyLeagueDraftPlayerRequest
// Response type: CMsgDOTAFantasyLeagueDraftPlayerResponse
func (d *Dota2) RequestFantasyLeagueDraftPlayer(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex uint32,
	playerAccountID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftPlayerResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftPlayerRequest{
		FantasyLeagueId: &fantasyLeagueID,
		TeamIndex:       &teamIndex,
		PlayerAccountId: &playerAccountID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftPlayerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftPlayerRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftPlayerResponse),
		resp,
	)
}

// RequestFantasyLeagueDraftStatus requests fantasy league draft status.
// Request ID: k_EMsgGCFantasyLeagueDraftStatusRequest
// Response ID: k_EMsgGCFantasyLeagueDraftStatus
// Request type: CMsgDOTAFantasyLeagueDraftStatusRequest
// Response type: CMsgDOTAFantasyLeagueDraftStatus
func (d *Dota2) RequestFantasyLeagueDraftStatus(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftStatus, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftStatusRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftStatus{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftStatusRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftStatus),
		resp,
	)
}

// RequestFantasyLeagueEditInfo requests a fantasy league edit info.
// Request ID: k_EMsgGCFantasyLeagueEditInfoRequest
// Response ID: k_EMsgGCFantasyLeagueEditInfoResponse
// Request type: CMsgDOTAFantasyLeagueEditInfoRequest
// Response type: CMsgDOTAFantasyLeagueEditInfoResponse
func (d *Dota2) RequestFantasyLeagueEditInfo(
	ctx context.Context,
	fantasyLeagueID uint32,
	editInfo dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueInfo,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInfoResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInfoRequest{
		FantasyLeagueId: &fantasyLeagueID,
		EditInfo:        &editInfo,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueEditInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueEditInfoResponse),
		resp,
	)
}

// RequestFantasyLeagueEditInvites requests fantasy league edit invites.
// Request ID: k_EMsgGCFantasyLeagueEditInvitesRequest
// Response ID: k_EMsgGCFantasyLeagueEditInvitesResponse
// Request type: CMsgDOTAFantasyLeagueEditInvitesRequest
// Response type: CMsgDOTAFantasyLeagueEditInvitesResponse
func (d *Dota2) RequestFantasyLeagueEditInvites(
	ctx context.Context,
	fantasyLeagueID uint32,
	password string,
	inviteChange []*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInvitesRequest_InviteChange,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInvitesResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInvitesRequest{
		FantasyLeagueId: &fantasyLeagueID,
		Password:        &password,
		InviteChange:    inviteChange,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueEditInvitesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueEditInvitesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueEditInvitesResponse),
		resp,
	)
}

// RequestFantasyLeagueInfo requests a fantasy league info.
// Request ID: k_EMsgGCFantasyLeagueInfoRequest
// Response ID: k_EMsgGCFantasyLeagueInfoResponse
// Request type: CMsgDOTAFantasyLeagueInfoRequest
// Response type: CMsgDOTAFantasyLeagueInfoResponse
func (d *Dota2) RequestFantasyLeagueInfo(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueInfoResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueInfoRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueInfoResponse),
		resp,
	)
}

// RequestFantasyLeagueMatchups requests fantasy league matchups.
// Request ID: k_EMsgGCFantasyLeagueMatchupsRequest
// Response ID: k_EMsgGCFantasyLeagueMatchupsResponse
// Request type: CMsgDOTAFantasyLeagueMatchupsRequest
// Response type: CMsgDOTAFantasyLeagueMatchupsResponse
func (d *Dota2) RequestFantasyLeagueMatchups(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueMatchupsResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueMatchupsRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueMatchupsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueMatchupsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueMatchupsResponse),
		resp,
	)
}

// RequestFantasyLeaveLeague requests a fantasy leave league.
// Request ID: k_EMsgGCFantasyLeaveLeagueRequest
// Response ID: k_EMsgGCFantasyLeaveLeagueResponse
// Request type: CMsgDOTAFantasyLeaveLeagueRequest
// Response type: CMsgDOTAFantasyLeaveLeagueResponse
func (d *Dota2) RequestFantasyLeaveLeague(
	ctx context.Context,
	fantasyLeagueID uint32,
	fantasyTeamIndex uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeaveLeagueResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeaveLeagueRequest{
		FantasyLeagueId:  &fantasyLeagueID,
		FantasyTeamIndex: &fantasyTeamIndex,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeaveLeagueResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeaveLeagueRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeaveLeagueResponse),
		resp,
	)
}

// RequestFantasyMessages requests fantasy messages.
// Request ID: k_EMsgGCFantasyMessagesRequest
// Response ID: k_EMsgGCFantasyMessagesResponse
// Request type: CMsgDOTAFantasyMessagesRequest
// Response type: CMsgDOTAFantasyMessagesResponse
func (d *Dota2) RequestFantasyMessages(
	ctx context.Context,
	fantasyLeagueID uint32,
	startMessage uint32,
	endMessage uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyMessagesResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyMessagesRequest{
		FantasyLeagueId: &fantasyLeagueID,
		StartMessage:    &startMessage,
		EndMessage:      &endMessage,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyMessagesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyMessagesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyMessagesResponse),
		resp,
	)
}

// RequestFantasyPlayerHisoricalStats requests fantasy player hisorical stats.
// Request ID: k_EMsgGCFantasyPlayerHisoricalStatsRequest
// Response ID: k_EMsgGCFantasyPlayerHisoricalStatsResponse
// Request type: CMsgDOTAFantasyPlayerHisoricalStatsRequest
// Response type: CMsgDOTAFantasyPlayerHisoricalStatsResponse
func (d *Dota2) RequestFantasyPlayerHisoricalStats(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerHisoricalStatsResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerHisoricalStatsRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerHisoricalStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerHisoricalStatsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerHisoricalStatsResponse),
		resp,
	)
}

// RequestFantasyPlayerInfo requests a fantasy player info.
// Request ID: k_EMsgGCFantasyPlayerInfoRequest
// Response ID: k_EMsgGCFantasyPlayerInfoResponse
// Request type: CMsgDOTAFantasyPlayerInfoRequest
// Response type: CMsgDOTAFantasyPlayerInfoResponse
func (d *Dota2) RequestFantasyPlayerInfo(
	ctx context.Context,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerInfoResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerInfoRequest{}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerInfoResponse),
		resp,
	)
}

// RequestFantasyPlayerScore requests a fantasy player score.
// Request ID: k_EMsgGCFantasyPlayerScoreRequest
// Response ID: k_EMsgGCFantasyPlayerScoreResponse
// Request type: CMsgDOTAFantasyPlayerScoreRequest
// Response type: CMsgDOTAFantasyPlayerScoreResponse
func (d *Dota2) RequestFantasyPlayerScore(
	ctx context.Context,
	fantasyLeagueID uint32,
	playerAccountID uint32,
	filterStartTime uint32,
	filterEndTime uint32,
	filterMatchID uint64,
	filterLastMatch bool,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreRequest{
		FantasyLeagueId: &fantasyLeagueID,
		PlayerAccountId: &playerAccountID,
		FilterStartTime: &filterStartTime,
		FilterEndTime:   &filterEndTime,
		FilterMatchId:   &filterMatchID,
		FilterLastMatch: &filterLastMatch,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerScoreRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerScoreResponse),
		resp,
	)
}

// RequestFantasyPlayerScoreDetails requests fantasy player score details.
// Request ID: k_EMsgGCFantasyPlayerScoreDetailsRequest
// Response ID: k_EMsgGCFantasyPlayerScoreDetailsResponse
// Request type: CMsgDOTAFantasyPlayerScoreDetailsRequest
// Response type: CMsgDOTAFantasyPlayerScoreDetailsResponse
func (d *Dota2) RequestFantasyPlayerScoreDetails(
	ctx context.Context,
	fantasyLeagueID uint32,
	playerAccountID uint32,
	startTime uint32,
	endTime uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreDetailsResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreDetailsRequest{
		FantasyLeagueId: &fantasyLeagueID,
		PlayerAccountId: &playerAccountID,
		StartTime:       &startTime,
		EndTime:         &endTime,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerScoreDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerScoreDetailsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerScoreDetailsResponse),
		resp,
	)
}

// RequestFantasyPlayerStandings requests fantasy player standings.
// Request ID: k_EMsgGCFantasyPlayerStandingsRequest
// Response ID: k_EMsgGCFantasyPlayerStandingsResponse
// Request type: CMsgDOTAFantasyPlayerStandingsRequest
// Response type: CMsgDOTAFantasyPlayerStandingsResponse
func (d *Dota2) RequestFantasyPlayerStandings(
	ctx context.Context,
	fantasyLeagueID uint32,
	count uint32,
	role uint32,
	filterStartTime uint32,
	filterEndTime uint32,
	filterMatchID uint64,
	filterLastMatch bool,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerStandingsResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerStandingsRequest{
		FantasyLeagueId: &fantasyLeagueID,
		Count:           &count,
		Role:            &role,
		FilterStartTime: &filterStartTime,
		FilterEndTime:   &filterEndTime,
		FilterMatchId:   &filterMatchID,
		FilterLastMatch: &filterLastMatch,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerStandingsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyPlayerStandingsResponse),
		resp,
	)
}

// RequestFantasyScheduledMatches requests fantasy scheduled matches.
// Request ID: k_EMsgGCFantasyScheduledMatchesRequest
// Response ID: k_EMsgGCFantasyScheduledMatchesResponse
// Request type: CMsgDOTAFantasyScheduledMatchesRequest
// Response type: CMsgDOTAFantasyScheduledMatchesResponse
func (d *Dota2) RequestFantasyScheduledMatches(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyScheduledMatchesResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyScheduledMatchesRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyScheduledMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyScheduledMatchesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyScheduledMatchesResponse),
		resp,
	)
}

// RequestFantasyTeamRoster requests a fantasy team roster.
// Request ID: k_EMsgGCFantasyTeamRosterRequest
// Response ID: k_EMsgGCFantasyTeamRosterResponse
// Request type: CMsgDOTAFantasyTeamRosterRequest
// Response type: CMsgDOTAFantasyTeamRosterResponse
func (d *Dota2) RequestFantasyTeamRoster(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex uint32,
	ownerAccountID uint32,
	timestamp uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterRequest{
		FantasyLeagueId: &fantasyLeagueID,
		TeamIndex:       &teamIndex,
		OwnerAccountId:  &ownerAccountID,
		Timestamp:       &timestamp,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterResponse),
		resp,
	)
}

// RequestFantasyTeamRosterAddDrop requests a fantasy team roster add drop.
// Request ID: k_EMsgGCFantasyTeamRosterAddDropRequest
// Response ID: k_EMsgGCFantasyTeamRosterAddDropResponse
// Request type: CMsgDOTAFantasyTeamRosterAddDropRequest
// Response type: CMsgDOTAFantasyTeamRosterAddDropResponse
func (d *Dota2) RequestFantasyTeamRosterAddDrop(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex uint32,
	addAccountID uint32,
	dropAccountID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterAddDropResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterAddDropRequest{
		FantasyLeagueId: &fantasyLeagueID,
		TeamIndex:       &teamIndex,
		AddAccountId:    &addAccountID,
		DropAccountId:   &dropAccountID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterAddDropResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterAddDropRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterAddDropResponse),
		resp,
	)
}

// RequestFantasyTeamScore requests a fantasy team score.
// Request ID: k_EMsgGCFantasyTeamScoreRequest
// Response ID: k_EMsgGCFantasyTeamScoreResponse
// Request type: CMsgDOTAFantasyTeamScoreRequest
// Response type: CMsgDOTAFantasyTeamScoreResponse
func (d *Dota2) RequestFantasyTeamScore(
	ctx context.Context,
	fantasyLeagueID uint32,
	ownerAccountID uint32,
	fantasyTeamIndex uint32,
	filterMatchID uint64,
	filterStartTime uint32,
	filterEndTime uint32,
	includeBench bool,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamScoreResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamScoreRequest{
		FantasyLeagueId:  &fantasyLeagueID,
		OwnerAccountId:   &ownerAccountID,
		FantasyTeamIndex: &fantasyTeamIndex,
		FilterMatchId:    &filterMatchID,
		FilterStartTime:  &filterStartTime,
		FilterEndTime:    &filterEndTime,
		IncludeBench:     &includeBench,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamScoreResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamScoreRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamScoreResponse),
		resp,
	)
}

// RequestFantasyTeamStandings requests fantasy team standings.
// Request ID: k_EMsgGCFantasyTeamStandingsRequest
// Response ID: k_EMsgGCFantasyTeamStandingsResponse
// Request type: CMsgDOTAFantasyTeamStandingsRequest
// Response type: CMsgDOTAFantasyTeamStandingsResponse
func (d *Dota2) RequestFantasyTeamStandings(
	ctx context.Context,
	fantasyLeagueID uint32,
	count uint32,
	filterStartTime uint32,
	filterEndTime uint32,
	filterMatchID uint64,
	filterLastMatch bool,
	filterInHall bool,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamStandingsResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamStandingsRequest{
		FantasyLeagueId: &fantasyLeagueID,
		Count:           &count,
		FilterStartTime: &filterStartTime,
		FilterEndTime:   &filterEndTime,
		FilterMatchId:   &filterMatchID,
		FilterLastMatch: &filterLastMatch,
		FilterInHall:    &filterInHall,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamStandingsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamStandingsResponse),
		resp,
	)
}

// RequestFantasyTeamTrades requests fantasy team trades.
// Request ID: k_EMsgGCFantasyTeamTradesRequest
// Response ID: k_EMsgGCFantasyTeamTradesResponse
// Request type: CMsgDOTAFantasyTeamTradesRequest
// Response type: CMsgDOTAFantasyTeamTradesResponse
func (d *Dota2) RequestFantasyTeamTrades(
	ctx context.Context,
	fantasyLeagueID uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradesResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradesRequest{
		FantasyLeagueId: &fantasyLeagueID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamTradesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamTradesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamTradesResponse),
		resp,
	)
}

// RequestFeaturedHeroes requests featured heroes.
// Request ID: k_EMsgClientToGCFeaturedHeroesRequest
// Response ID: k_EMsgGCToClientFeaturedHeroesResponse
// Request type: CMsgClientToGCFeaturedHeroesRequest
// Response type: CMsgGCToClientFeaturedHeroesResponse
func (d *Dota2) RequestFeaturedHeroes(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCToClientFeaturedHeroesResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCFeaturedHeroesRequest{}
	resp := &dota_gcmessages_client.CMsgGCToClientFeaturedHeroesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCFeaturedHeroesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientFeaturedHeroesResponse),
		resp,
	)
}

// RequestFriendRecruits requests friend recruits.
// Request ID: k_EMsgDOTAFriendRecruitsRequest
// Response ID: k_EMsgDOTAFriendRecruitsResponse
// Request type: CMsgDOTAFriendRecruitsRequest
// Response type: CMsgDOTAFriendRecruitsResponse
func (d *Dota2) RequestFriendRecruits(
	ctx context.Context,
	accountIDs []uint32,
) (*dota_gcmessages_client.CMsgDOTAFriendRecruitsResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAFriendRecruitsRequest{
		AccountIds: accountIDs,
	}
	resp := &dota_gcmessages_client.CMsgDOTAFriendRecruitsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAFriendRecruitsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAFriendRecruitsResponse),
		resp,
	)
}

// RequestFriendsPlayedCustomGame requests a friends played custom game.
// Request ID: k_EMsgClientToGCFriendsPlayedCustomGameRequest
// Response ID: k_EMsgGCToClientFriendsPlayedCustomGameResponse
// Request type: CMsgClientToGCFriendsPlayedCustomGameRequest
// Response type: CMsgGCToClientFriendsPlayedCustomGameResponse
func (d *Dota2) RequestFriendsPlayedCustomGame(
	ctx context.Context,
	customGameID uint64,
) (*dota_gcmessages_client.CMsgGCToClientFriendsPlayedCustomGameResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCFriendsPlayedCustomGameRequest{
		CustomGameId: &customGameID,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientFriendsPlayedCustomGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCFriendsPlayedCustomGameRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientFriendsPlayedCustomGameResponse),
		resp,
	)
}

// RequestGetPlayerCardRoster requests to check if the target get player card roster.
// Request ID: k_EMsgClientToGCGetPlayerCardRosterRequest
// Response ID: k_EMsgClientToGCGetPlayerCardRosterResponse
// Request type: CMsgClientToGCGetPlayerCardRosterRequest
// Response type: CMsgClientToGCGetPlayerCardRosterResponse
func (d *Dota2) RequestGetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	timestamp uint32,
) (*dota_gcmessages_client_fantasy.CMsgClientToGCGetPlayerCardRosterResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgClientToGCGetPlayerCardRosterRequest{
		LeagueId:  &leagueID,
		Timestamp: &timestamp,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgClientToGCGetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCGetPlayerCardRosterResponse),
		resp,
	)
}

// RequestGuildCancelInvite requests a guild cancel invite.
// Request ID: k_EMsgGCGuildCancelInviteRequest
// Response ID: k_EMsgGCGuildCancelInviteResponse
// Request type: CMsgDOTAGuildCancelInviteRequest
// Response type: CMsgDOTAGuildCancelInviteResponse
func (d *Dota2) RequestGuildCancelInvite(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildCancelInviteResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildCancelInviteRequest{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildCancelInviteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildCancelInviteRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildCancelInviteResponse),
		resp,
	)
}

// RequestGuildEditLogo requests a guild edit logo.
// Request ID: k_EMsgGCGuildEditLogoRequest
// Response ID: k_EMsgGCGuildEditLogoResponse
// Request type: CMsgDOTAGuildEditLogoRequest
// Response type: CMsgDOTAGuildEditLogoResponse
func (d *Dota2) RequestGuildEditLogo(
	ctx context.Context,
	guildID uint32,
	logo uint64,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildEditLogoResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildEditLogoRequest{
		GuildId: &guildID,
		Logo:    &logo,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildEditLogoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildEditLogoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildEditLogoResponse),
		resp,
	)
}

// RequestGuildInviteAccount requests a guild invite account.
// Request ID: k_EMsgGCGuildInviteAccountRequest
// Response ID: k_EMsgGCGuildInviteAccountResponse
// Request type: CMsgDOTAGuildInviteAccountRequest
// Response type: CMsgDOTAGuildInviteAccountResponse
func (d *Dota2) RequestGuildInviteAccount(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildInviteAccountResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildInviteAccountRequest{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildInviteAccountResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildInviteAccountRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildInviteAccountResponse),
		resp,
	)
}

// RequestGuildSetAccountRole requests a guild set account role.
// Request ID: k_EMsgGCGuildSetAccountRoleRequest
// Response ID: k_EMsgGCGuildSetAccountRoleResponse
// Request type: CMsgDOTAGuildSetAccountRoleRequest
// Response type: CMsgDOTAGuildSetAccountRoleResponse
func (d *Dota2) RequestGuildSetAccountRole(
	ctx context.Context,
	guildID uint32,
	targetAccountID uint32,
	targetRole uint32,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildSetAccountRoleResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildSetAccountRoleRequest{
		GuildId:         &guildID,
		TargetAccountId: &targetAccountID,
		TargetRole:      &targetRole,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildSetAccountRoleResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildSetAccountRoleRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildSetAccountRoleResponse),
		resp,
	)
}

// RequestGuildUpdateDetails requests guild update details.
// Request ID: k_EMsgGCGuildUpdateDetailsRequest
// Response ID: k_EMsgGCGuildUpdateDetailsResponse
// Request type: CMsgDOTAGuildUpdateDetailsRequest
// Response type: CMsgDOTAGuildUpdateDetailsResponse
func (d *Dota2) RequestGuildUpdateDetails(
	ctx context.Context,
	guildID uint32,
	logo uint64,
	baseLogo uint64,
	bannerLogo uint64,
) (*dota_gcmessages_client_guild.CMsgDOTAGuildUpdateDetailsResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAGuildUpdateDetailsRequest{
		GuildId:    &guildID,
		Logo:       &logo,
		BaseLogo:   &baseLogo,
		BannerLogo: &bannerLogo,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAGuildUpdateDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildUpdateDetailsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildUpdateDetailsResponse),
		resp,
	)
}

// RequestH264Support requests a h 264 support.
// Request ID: k_EMsgClientToGCRequestH264Support
// Request type: CMsgClientToGCRequestH264Support
func (d *Dota2) RequestH264Support() {
	req := &dota_gcmessages_client.CMsgClientToGCRequestH264Support{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestH264Support), req)
}

// RequestHallOfFame requests a hall of fame.
// Request ID: k_EMsgGCHallOfFameRequest
// Response ID: k_EMsgGCHallOfFameResponse
// Request type: CMsgDOTAHallOfFameRequest
// Response type: CMsgDOTAHallOfFameResponse
func (d *Dota2) RequestHallOfFame(
	ctx context.Context,
	week uint32,
) (*dota_gcmessages_client.CMsgDOTAHallOfFameResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAHallOfFameRequest{
		Week: &week,
	}
	resp := &dota_gcmessages_client.CMsgDOTAHallOfFameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHallOfFameRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHallOfFameResponse),
		resp,
	)
}

// RequestHalloweenHighScore requests a halloween high score.
// Request ID: k_EMsgGCHalloweenHighScoreRequest
// Response ID: k_EMsgGCHalloweenHighScoreResponse
// Request type: CMsgDOTAHalloweenHighScoreRequest
// Response type: CMsgDOTAHalloweenHighScoreResponse
func (d *Dota2) RequestHalloweenHighScore(
	ctx context.Context,
	round int32,
) (*dota_gcmessages_client.CMsgDOTAHalloweenHighScoreResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAHalloweenHighScoreRequest{
		Round: &round,
	}
	resp := &dota_gcmessages_client.CMsgDOTAHalloweenHighScoreResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHalloweenHighScoreRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHalloweenHighScoreResponse),
		resp,
	)
}

// RequestInternationalTicketEmail requests a international ticket email.
// Request ID: k_EMsgGCRequestInternatinalTicketEmail
// Request type: CMsgRequestInternationalTicket
func (d *Dota2) RequestInternationalTicketEmail() {
	req := &dota_gcmessages_client.CMsgRequestInternationalTicket{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestInternatinalTicketEmail), req)
}

// RequestItemEditorReservations requests item editor reservations.
// Request ID: k_EMsgGCItemEditorReservationsRequest
// Response ID: k_EMsgGCItemEditorReservationsResponse
// Request type: CMsgGCItemEditorReservationsRequest
// Response type: CMsgGCItemEditorReservationsResponse
func (d *Dota2) RequestItemEditorReservations(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCItemEditorReservationsResponse, error) {
	req := &dota_gcmessages_client.CMsgGCItemEditorReservationsRequest{}
	resp := &dota_gcmessages_client.CMsgGCItemEditorReservationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReservationsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReservationsResponse),
		resp,
	)
}

// RequestJoinOpenGuildParty requests to check if the target join open guild party.
// Request ID: k_EMsgGCJoinOpenGuildPartyRequest
// Response ID: k_EMsgGCJoinOpenGuildPartyResponse
// Request type: CMsgDOTAJoinOpenGuildPartyRequest
// Response type: CMsgDOTAJoinOpenGuildPartyResponse
func (d *Dota2) RequestJoinOpenGuildParty(
	ctx context.Context,
	partyID uint64,
) (*dota_gcmessages_client_guild.CMsgDOTAJoinOpenGuildPartyResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAJoinOpenGuildPartyRequest{
		PartyId: &partyID,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAJoinOpenGuildPartyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinOpenGuildPartyRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinOpenGuildPartyResponse),
		resp,
	)
}

// RequestJoinableCustomGameModes requests joinable custom game modes.
// Request ID: k_EMsgGCJoinableCustomGameModesRequest
// Response ID: k_EMsgGCJoinableCustomGameModesResponse
// Request type: CMsgJoinableCustomGameModesRequest
// Response type: CMsgJoinableCustomGameModesResponse
func (d *Dota2) RequestJoinableCustomGameModes(
	ctx context.Context,
	serverRegion uint32,
) (*dota_gcmessages_client_match_management.CMsgJoinableCustomGameModesResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgJoinableCustomGameModesRequest{
		ServerRegion: &serverRegion,
	}
	resp := &dota_gcmessages_client_match_management.CMsgJoinableCustomGameModesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinableCustomGameModesResponse),
		resp,
	)
}

// RequestJoinableCustomLobbies requests joinable custom lobbies.
// Request ID: k_EMsgGCJoinableCustomLobbiesRequest
// Response ID: k_EMsgGCJoinableCustomLobbiesResponse
// Request type: CMsgJoinableCustomLobbiesRequest
// Response type: CMsgJoinableCustomLobbiesResponse
func (d *Dota2) RequestJoinableCustomLobbies(
	ctx context.Context,
	serverRegion uint32,
	customGameID uint64,
) (*dota_gcmessages_client_match_management.CMsgJoinableCustomLobbiesResponse, error) {
	req := &dota_gcmessages_client_match_management.CMsgJoinableCustomLobbiesRequest{
		ServerRegion: &serverRegion,
		CustomGameId: &customGameID,
	}
	resp := &dota_gcmessages_client_match_management.CMsgJoinableCustomLobbiesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCJoinableCustomLobbiesResponse),
		resp,
	)
}

// RequestLastHitChallengeHighScore requests a last hit challenge high score.
// Request ID: k_EMsgGCLastHitChallengeHighScoreRequest
// Response ID: k_EMsgGCLastHitChallengeHighScoreResponse
// Request type: CMsgDOTALastHitChallengeHighScoreRequest
// Response type: CMsgDOTALastHitChallengeHighScoreResponse
func (d *Dota2) RequestLastHitChallengeHighScore(
	ctx context.Context,
	heroID uint32,
) (*dota_gcmessages_client.CMsgDOTALastHitChallengeHighScoreResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTALastHitChallengeHighScoreRequest{
		HeroId: &heroID,
	}
	resp := &dota_gcmessages_client.CMsgDOTALastHitChallengeHighScoreResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLastHitChallengeHighScoreRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLastHitChallengeHighScoreResponse),
		resp,
	)
}

// RequestLatestConductScorecard requests a latest conduct scorecard.
// Request ID: k_EMsgClientToGCLatestConductScorecardRequest
// Response ID: k_EMsgClientToGCLatestConductScorecard
// Request type: CMsgPlayerConductScorecardRequest
// Response type: CMsgPlayerConductScorecard
func (d *Dota2) RequestLatestConductScorecard(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgPlayerConductScorecard, error) {
	req := &dota_gcmessages_client.CMsgPlayerConductScorecardRequest{}
	resp := &dota_gcmessages_client.CMsgPlayerConductScorecard{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard),
		resp,
	)
}

// RequestLeagueInfo requests a league info.
// Request ID: k_EMsgRequestLeagueInfo
// Request type: CMsgRequestLeagueInfo
func (d *Dota2) RequestLeagueInfo() {
	req := &dota_gcmessages_client.CMsgRequestLeagueInfo{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgRequestLeagueInfo), req)
}

// RequestLeaguePrizePool requests a league prize pool.
// Request ID: k_EMsgGCRequestLeaguePrizePool
// Response ID: k_EMsgGCRequestLeaguePrizePoolResponse
// Request type: CMsgRequestLeaguePrizePool
// Response type: CMsgRequestLeaguePrizePoolResponse
func (d *Dota2) RequestLeaguePrizePool(
	ctx context.Context,
	leagueID uint32,
) (*dota_gcmessages_client.CMsgRequestLeaguePrizePoolResponse, error) {
	req := &dota_gcmessages_client.CMsgRequestLeaguePrizePool{
		LeagueId: &leagueID,
	}
	resp := &dota_gcmessages_client.CMsgRequestLeaguePrizePoolResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestLeaguePrizePool),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestLeaguePrizePoolResponse),
		resp,
	)
}

// RequestLinaGameResult requests a lina game result.
// Request ID: k_EMsgClientToGCRequestLinaGameResult
// Response ID: k_EMsgClientToGCRequestLinaGameResultResponse
// Request type: CMsgClientToGCRequestLinaGameResult
// Response type: CMsgClientToGCRequestLinaGameResultResponse
func (d *Dota2) RequestLinaGameResult(
	ctx context.Context,
	eventID dota_shared_enums.EEvent,
	slotChosen uint32,
) (*dota_gcmessages_client.CMsgClientToGCRequestLinaGameResultResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRequestLinaGameResult{
		EventId:    &eventID,
		SlotChosen: &slotChosen,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCRequestLinaGameResultResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestLinaGameResult),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestLinaGameResultResponse),
		resp,
	)
}

// RequestLinaPlaysRemaining requests a lina plays remaining.
// Request ID: k_EMsgClientToGCRequestLinaPlaysRemaining
// Response ID: k_EMsgClientToGCRequestLinaPlaysRemainingResponse
// Request type: CMsgClientToGCRequestLinaPlaysRemaining
// Response type: CMsgClientToGCRequestLinaPlaysRemainingResponse
func (d *Dota2) RequestLinaPlaysRemaining(
	ctx context.Context,
	eventID dota_shared_enums.EEvent,
) (*dota_gcmessages_client.CMsgClientToGCRequestLinaPlaysRemainingResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRequestLinaPlaysRemaining{
		EventId: &eventID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCRequestLinaPlaysRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestLinaPlaysRemaining),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestLinaPlaysRemainingResponse),
		resp,
	)
}

// RequestMatchDetails requests match details.
// Request ID: k_EMsgGCMatchDetailsRequest
// Response ID: k_EMsgGCMatchDetailsResponse
// Request type: CMsgGCMatchDetailsRequest
// Response type: CMsgGCMatchDetailsResponse
func (d *Dota2) RequestMatchDetails(
	ctx context.Context,
	matchID uint64,
) (*dota_gcmessages_client.CMsgGCMatchDetailsResponse, error) {
	req := &dota_gcmessages_client.CMsgGCMatchDetailsRequest{
		MatchId: &matchID,
	}
	resp := &dota_gcmessages_client.CMsgGCMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMatchDetailsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMatchDetailsResponse),
		resp,
	)
}

// RequestMatches requests matches.
// Request ID: k_EMsgGCRequestMatches
// Response ID: k_EMsgGCRequestMatchesResponse
// Request type: CMsgDOTARequestMatches
// Response type: CMsgDOTARequestMatchesResponse
func (d *Dota2) RequestMatches(
	ctx context.Context,
	req *dota_gcmessages_client.CMsgDOTARequestMatches,
) (*dota_gcmessages_client.CMsgDOTARequestMatchesResponse, error) {
	resp := &dota_gcmessages_client.CMsgDOTARequestMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestMatches),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestMatchesResponse),
		resp,
	)
}

// RequestMatchesMinimal requests a matches minimal.
// Request ID: k_EMsgClientToGCMatchesMinimalRequest
// Response ID: k_EMsgClientToGCMatchesMinimalResponse
// Request type: CMsgClientToGCMatchesMinimalRequest
// Response type: CMsgClientToGCMatchesMinimalResponse
func (d *Dota2) RequestMatchesMinimal(
	ctx context.Context,
	matchIDs []uint64,
) (*dota_gcmessages_client_watch.CMsgClientToGCMatchesMinimalResponse, error) {
	req := &dota_gcmessages_client_watch.CMsgClientToGCMatchesMinimalRequest{
		MatchIds: matchIDs,
	}
	resp := &dota_gcmessages_client_watch.CMsgClientToGCMatchesMinimalResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCMatchesMinimalResponse),
		resp,
	)
}

// RequestMatchmakingStats requests matchmaking stats.
// Request ID: k_EMsgGCMatchmakingStatsRequest
// Response ID: k_EMsgGCMatchmakingStatsResponse
// Request type: CMsgDOTAMatchmakingStatsRequest
// Response type: CMsgDOTAMatchmakingStatsResponse
func (d *Dota2) RequestMatchmakingStats(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTAMatchmakingStatsResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAMatchmakingStatsRequest{}
	resp := &dota_gcmessages_client.CMsgDOTAMatchmakingStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMatchmakingStatsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMatchmakingStatsResponse),
		resp,
	)
}

// RequestMyTeamInfo requests a my team info.
// Request ID: k_EMsgClientToGCMyTeamInfoRequest
// Response ID: k_EMsgGCToClientTeamInfo
// Request type: CMsgDOTAMyTeamInfoRequest
// Response type: CMsgDOTATeamInfo
func (d *Dota2) RequestMyTeamInfo(
	ctx context.Context,
) (*dota_gcmessages_client_team.CMsgDOTATeamInfo, error) {
	req := &dota_gcmessages_client_team.CMsgDOTAMyTeamInfoRequest{}
	resp := &dota_gcmessages_client_team.CMsgDOTATeamInfo{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTeamInfo),
		resp,
	)
}

// RequestNotifications requests notifications.
// Request ID: k_EMsgGCNotificationsRequest
// Response ID: k_EMsgGCNotificationsResponse
// Request type: CMsgGCNotificationsRequest
// Response type: CMsgGCNotificationsResponse
func (d *Dota2) RequestNotifications(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgGCNotificationsResponse, error) {
	req := &dota_gcmessages_client.CMsgGCNotificationsRequest{}
	resp := &dota_gcmessages_client.CMsgGCNotificationsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNotificationsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNotificationsResponse),
		resp,
	)
}

// RequestNotificationsMarkRead requests a notifications mark read.
// Request ID: k_EMsgGCNotificationsMarkReadRequest
// Request type: CMsgGCNotificationsMarkReadRequest
func (d *Dota2) RequestNotificationsMarkRead() {
	req := &dota_gcmessages_client.CMsgGCNotificationsMarkReadRequest{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest), req)
}

// RequestOfferings requests offerings.
// Request ID: k_EMsgGCRequestOfferings
// Response ID: k_EMsgGCRequestOfferingsResponse
// Request type: CMsgRequestOfferings
// Response type: CMsgRequestOfferingsResponse
func (d *Dota2) RequestOfferings(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgRequestOfferingsResponse, error) {
	req := &dota_gcmessages_client.CMsgRequestOfferings{}
	resp := &dota_gcmessages_client.CMsgRequestOfferingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestOfferings),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestOfferingsResponse),
		resp,
	)
}

// RequestPartySetOpenGuild requests a party set open guild.
// Request ID: k_EMsgGCPartySetOpenGuildRequest
// Response ID: k_EMsgGCPartySetOpenGuildResponse
// Request type: CMsgDOTAPartySetOpenGuildRequest
// Response type: CMsgDOTAPartySetOpenGuildResponse
func (d *Dota2) RequestPartySetOpenGuild(
	ctx context.Context,
	guildID uint32,
	description string,
) (*dota_gcmessages_client_guild.CMsgDOTAPartySetOpenGuildResponse, error) {
	req := &dota_gcmessages_client_guild.CMsgDOTAPartySetOpenGuildRequest{
		GuildId:     &guildID,
		Description: &description,
	}
	resp := &dota_gcmessages_client_guild.CMsgDOTAPartySetOpenGuildResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPartySetOpenGuildRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPartySetOpenGuildResponse),
		resp,
	)
}

// RequestPerfectWorldUserLookup requests a perfect world user lookup.
// Request ID: k_EMsgGCPerfectWorldUserLookupRequest
// Response ID: k_EMsgGCPerfectWorldUserLookupResponse
// Request type: CMsgPerfectWorldUserLookupRequest
// Response type: CMsgPerfectWorldUserLookupResponse
func (d *Dota2) RequestPerfectWorldUserLookup(
	ctx context.Context,
	userName string,
) (*dota_gcmessages_client.CMsgPerfectWorldUserLookupResponse, error) {
	req := &dota_gcmessages_client.CMsgPerfectWorldUserLookupRequest{
		UserName: &userName,
	}
	resp := &dota_gcmessages_client.CMsgPerfectWorldUserLookupResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPerfectWorldUserLookupRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPerfectWorldUserLookupResponse),
		resp,
	)
}

// RequestPlayerInfo requests a player info.
// Request ID: k_EMsgGCPlayerInfoRequest
// Response ID: k_EMsgGCPlayerInfo
// Request type: CMsgGCPlayerInfoRequest
// Response type: CMsgGCPlayerInfo
func (d *Dota2) RequestPlayerInfo(
	ctx context.Context,
	playerInfos []*dota_gcmessages_client.CMsgGCPlayerInfoRequest_PlayerInfo,
) (*dota_gcmessages_client_fantasy.CMsgGCPlayerInfo, error) {
	req := &dota_gcmessages_client.CMsgGCPlayerInfoRequest{
		PlayerInfos: playerInfos,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgGCPlayerInfo{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfo),
		resp,
	)
}

// RequestPlayerStats requests player stats.
// Request ID: k_EMsgClientToGCPlayerStatsRequest
// Response ID: k_EMsgGCToClientPlayerStatsResponse
// Request type: CMsgClientToGCPlayerStatsRequest
// Response type: CMsgGCToClientPlayerStatsResponse
func (d *Dota2) RequestPlayerStats(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_client.CMsgGCToClientPlayerStatsResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCPlayerStatsRequest{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientPlayerStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPlayerStatsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientPlayerStatsResponse),
		resp,
	)
}

// RequestPrivateChatInfo requests a private chat info.
// Request ID: k_EMsgClientToGCPrivateChatInfoRequest
// Response ID: k_EMsgGCToClientPrivateChatInfoResponse
// Request type: CMsgClientToGCPrivateChatInfoRequest
// Response type: CMsgGCToClientPrivateChatInfoResponse
func (d *Dota2) RequestPrivateChatInfo(
	ctx context.Context,
	privateChatChannelName string,
) (*dota_gcmessages_client_chat.CMsgGCToClientPrivateChatInfoResponse, error) {
	req := &dota_gcmessages_client_chat.CMsgClientToGCPrivateChatInfoRequest{
		PrivateChatChannelName: &privateChatChannelName,
	}
	resp := &dota_gcmessages_client_chat.CMsgGCToClientPrivateChatInfoResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInfoRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientPrivateChatInfoResponse),
		resp,
	)
}

// RequestProfile requests a profile.
// Request ID: k_EMsgGCProfileRequest
// Response ID: k_EMsgGCProfileResponse
// Request type: CMsgDOTAProfileRequest
// Response type: CMsgDOTAProfileResponse
func (d *Dota2) RequestProfile(
	ctx context.Context,
	accountID uint32,
	requestName bool,
) (*dota_gcmessages_client.CMsgDOTAProfileResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAProfileRequest{
		AccountId:   &accountID,
		RequestName: &requestName,
	}
	resp := &dota_gcmessages_client.CMsgDOTAProfileResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProfileRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProfileResponse),
		resp,
	)
}

// RequestQuickStats requests quick stats.
// Request ID: k_EMsgClientToGCQuickStatsRequest
// Response ID: k_EMsgClientToGCQuickStatsResponse
// Request type: CMsgDOTAClientToGCQuickStatsRequest
// Response type: CMsgDOTAClientToGCQuickStatsResponse
func (d *Dota2) RequestQuickStats(
	ctx context.Context,
	playerAccountID uint32,
	heroID uint32,
	itemID uint32,
	leagueID uint32,
) (*dota_gcmessages_client.CMsgDOTAClientToGCQuickStatsResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAClientToGCQuickStatsRequest{
		PlayerAccountId: &playerAccountID,
		HeroId:          &heroID,
		ItemId:          &itemID,
		LeagueId:        &leagueID,
	}
	resp := &dota_gcmessages_client.CMsgDOTAClientToGCQuickStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCQuickStatsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCQuickStatsResponse),
		resp,
	)
}

// RequestReportCounts requests to check if the target report counts.
// Request ID: k_EMsgGCReportCountsRequest
// Response ID: k_EMsgGCReportCountsResponse
// Request type: CMsgDOTAReportCountsRequest
// Response type: CMsgDOTAReportCountsResponse
func (d *Dota2) RequestReportCounts(
	ctx context.Context,
	targetAccountID uint32,
) (*dota_gcmessages_client.CMsgDOTAReportCountsResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAReportCountsRequest{
		TargetAccountId: &targetAccountID,
	}
	resp := &dota_gcmessages_client.CMsgDOTAReportCountsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReportCountsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReportCountsResponse),
		resp,
	)
}

// RequestReportsRemaining requests a reports remaining.
// Request ID: k_EMsgGCReportsRemainingRequest
// Response ID: k_EMsgGCReportsRemainingResponse
// Request type: CMsgDOTAReportsRemainingRequest
// Response type: CMsgDOTAReportsRemainingResponse
func (d *Dota2) RequestReportsRemaining(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTAReportsRemainingResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAReportsRemainingRequest{}
	resp := &dota_gcmessages_client.CMsgDOTAReportsRemainingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReportsRemainingRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReportsRemainingResponse),
		resp,
	)
}

// RequestSaveGames requests save games.
// Request ID: k_EMsgGCRequestSaveGames
// Response ID: k_EMsgGCRequestSaveGamesResponse
// Request type: CMsgDOTARequestSaveGames
// Response type: CMsgDOTARequestSaveGamesResponse
func (d *Dota2) RequestSaveGames(
	ctx context.Context,
	serverRegion uint32,
) (*dota_gcmessages_client.CMsgDOTARequestSaveGamesResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTARequestSaveGames{
		ServerRegion: &serverRegion,
	}
	resp := &dota_gcmessages_client.CMsgDOTARequestSaveGamesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestSaveGames),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCRequestSaveGamesResponse),
		resp,
	)
}

// RequestSelectionPriorityChoice requests a selection priority choice.
// Request ID: k_EMsgSelectionPriorityChoiceRequest
// Response ID: k_EMsgSelectionPriorityChoiceResponse
// Request type: CMsgDOTASelectionPriorityChoiceRequest
// Response type: CMsgDOTASelectionPriorityChoiceResponse
func (d *Dota2) RequestSelectionPriorityChoice(
	ctx context.Context,
	choice dota_shared_enums.DOTASelectionPriorityChoice,
) (*dota_gcmessages_client.CMsgDOTASelectionPriorityChoiceResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASelectionPriorityChoiceRequest{
		Choice: &choice,
	}
	resp := &dota_gcmessages_client.CMsgDOTASelectionPriorityChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgSelectionPriorityChoiceResponse),
		resp,
	)
}

// RequestSetPlayerCardRoster requests to check if the target set player card roster.
// Request ID: k_EMsgClientToGCSetPlayerCardRosterRequest
// Response ID: k_EMsgClientToGCSetPlayerCardRosterResponse
// Request type: CMsgClientToGCSetPlayerCardRosterRequest
// Response type: CMsgClientToGCSetPlayerCardRosterResponse
func (d *Dota2) RequestSetPlayerCardRoster(
	ctx context.Context,
	leagueID uint32,
	timestamp uint32,
	slot uint32,
	playerCardItemID uint64,
	eventID uint32,
) (*dota_gcmessages_client_fantasy.CMsgClientToGCSetPlayerCardRosterResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgClientToGCSetPlayerCardRosterRequest{
		LeagueId:         &leagueID,
		Timestamp:        &timestamp,
		Slot:             &slot,
		PlayerCardItemId: &playerCardItemID,
		EventId:          &eventID,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgClientToGCSetPlayerCardRosterResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetPlayerCardRosterResponse),
		resp,
	)
}

// RequestSlarkGameResult requests a slark game result.
// Request ID: k_EMsgClientToGCRequestSlarkGameResult
// Response ID: k_EMsgClientToGCRequestSlarkGameResultResponse
// Request type: CMsgClientToGCRequestSlarkGameResult
// Response type: CMsgClientToGCRequestSlarkGameResultResponse
func (d *Dota2) RequestSlarkGameResult(
	ctx context.Context,
	eventID dota_shared_enums.EEvent,
	slotChosen uint32,
	week uint32,
) (*dota_gcmessages_client.CMsgClientToGCRequestSlarkGameResultResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCRequestSlarkGameResult{
		EventId:    &eventID,
		SlotChosen: &slotChosen,
		Week:       &week,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCRequestSlarkGameResultResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestSlarkGameResult),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestSlarkGameResultResponse),
		resp,
	)
}

// RequestSocialFeedPostComment requests a social feed post comment.
// Request ID: k_EMsgClientToGCSocialFeedPostCommentRequest
// Response ID: k_EMsgGCToClientSocialFeedPostCommentResponse
// Request type: CMsgClientToGCSocialFeedPostCommentRequest
// Response type: CMsgGCToClientSocialFeedPostCommentResponse
func (d *Dota2) RequestSocialFeedPostComment(
	ctx context.Context,
	eventID uint64,
	comment string,
) (*dota_gcmessages_client.CMsgGCToClientSocialFeedPostCommentResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSocialFeedPostCommentRequest{
		EventId: &eventID,
		Comment: &comment,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientSocialFeedPostCommentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostCommentRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostCommentResponse),
		resp,
	)
}

// RequestSocialFeedPostMessage requests a social feed post message.
// Request ID: k_EMsgClientToGCSocialFeedPostMessageRequest
// Response ID: k_EMsgGCToClientSocialFeedPostMessageResponse
// Request type: CMsgClientToGCSocialFeedPostMessageRequest
// Response type: CMsgGCToClientSocialFeedPostMessageResponse
func (d *Dota2) RequestSocialFeedPostMessage(
	ctx context.Context,
	message string,
	matchID uint64,
	matchTimestamp uint32,
) (*dota_gcmessages_client.CMsgGCToClientSocialFeedPostMessageResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSocialFeedPostMessageRequest{
		Message:        &message,
		MatchId:        &matchID,
		MatchTimestamp: &matchTimestamp,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientSocialFeedPostMessageResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSocialFeedPostMessageRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSocialFeedPostMessageResponse),
		resp,
	)
}

// RequestSocialMatchDetails requests social match details.
// Request ID: k_EMsgClientToGCSocialMatchDetailsRequest
// Response ID: k_EMsgGCToClientSocialMatchDetailsResponse
// Request type: CMsgClientToGCSocialMatchDetailsRequest
// Response type: CMsgGCToClientSocialMatchDetailsResponse
func (d *Dota2) RequestSocialMatchDetails(
	ctx context.Context,
	matchID uint64,
	paginationTimestamp uint32,
) (*dota_gcmessages_client.CMsgGCToClientSocialMatchDetailsResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSocialMatchDetailsRequest{
		MatchId:             &matchID,
		PaginationTimestamp: &paginationTimestamp,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientSocialMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSocialMatchDetailsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSocialMatchDetailsResponse),
		resp,
	)
}

// RequestSocialMatchPostComment requests a social match post comment.
// Request ID: k_EMsgClientToGCSocialMatchPostCommentRequest
// Response ID: k_EMsgGCToClientSocialMatchPostCommentResponse
// Request type: CMsgClientToGCSocialMatchPostCommentRequest
// Response type: CMsgGCToClientSocialMatchPostCommentResponse
func (d *Dota2) RequestSocialMatchPostComment(
	ctx context.Context,
	matchID uint64,
	comment string,
) (*dota_gcmessages_client.CMsgGCToClientSocialMatchPostCommentResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSocialMatchPostCommentRequest{
		MatchId: &matchID,
		Comment: &comment,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientSocialMatchPostCommentResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSocialMatchPostCommentRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSocialMatchPostCommentResponse),
		resp,
	)
}

// RequestSteamDatagramTicket requests a steam datagram ticket.
// Request ID: k_EMsgClientToGCRequestSteamDatagramTicket
// Response ID: k_EMsgClientToGCRequestSteamDatagramTicketResponse
// Request type: CMsgClientToGCRequestSteamDatagramTicket
// Response type: CMsgClientToGCRequestSteamDatagramTicketResponse
func (d *Dota2) RequestSteamDatagramTicket(
	ctx context.Context,
	serverSteamID steamid.SteamId,
) (*dota_gcmessages_client_match_management.CMsgClientToGCRequestSteamDatagramTicketResponse, error) {
	serverSteamIDU64Val := uint64(serverSteamID)
	serverSteamIDU64 := &serverSteamIDU64Val
	req := &dota_gcmessages_client_match_management.CMsgClientToGCRequestSteamDatagramTicket{
		ServerSteamId: serverSteamIDU64,
	}
	resp := &dota_gcmessages_client_match_management.CMsgClientToGCRequestSteamDatagramTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicket),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRequestSteamDatagramTicketResponse),
		resp,
	)
}

// RequestSteamProfile requests a steam profile.
// Request ID: k_EMsgGCSteamProfileRequest
// Response ID: k_EMsgGCSteamProfileRequestResponse
// Request type: CMsgGCSteamProfileRequest
// Response type: CMsgGCSteamProfileRequestResponse
func (d *Dota2) RequestSteamProfile(
	ctx context.Context,
	accountID uint32,
) (*dota_gcmessages_client.CMsgGCSteamProfileRequestResponse, error) {
	req := &dota_gcmessages_client.CMsgGCSteamProfileRequest{
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgGCSteamProfileRequestResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSteamProfileRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSteamProfileRequestResponse),
		resp,
	)
}

// RequestStorePromoPages requests store promo pages.
// Request ID: k_EMsgGCStorePromoPagesRequest
// Response ID: k_EMsgGCStorePromoPagesResponse
// Request type: CMsgDOTAStorePromoPagesRequest
// Response type: CMsgDOTAStorePromoPagesResponse
func (d *Dota2) RequestStorePromoPages(
	ctx context.Context,
	versionSeen uint32,
) (*dota_gcmessages_client.CMsgDOTAStorePromoPagesResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAStorePromoPagesRequest{
		VersionSeen: &versionSeen,
	}
	resp := &dota_gcmessages_client.CMsgDOTAStorePromoPagesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCStorePromoPagesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCStorePromoPagesResponse),
		resp,
	)
}

// RequestTeamInfoFantasyByFantasyLeagueID requests a team info fantasy by fantasy league id.
// Request ID: k_EMsgGCFantasyTeamInfoRequestByFantasyLeagueID
// Request type: CMsgDOTAFantasyTeamInfoRequestByFantasyLeagueID
func (d *Dota2) RequestTeamInfoFantasyByFantasyLeagueID(
	fantasyLeagueID uint32,
) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamInfoRequestByFantasyLeagueID{
		FantasyLeagueId: &fantasyLeagueID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamInfoRequestByFantasyLeagueID), req)
}

// RequestTeamInfoFantasyByOwnerAccountID requests a team info fantasy by owner account id.
// Request ID: k_EMsgGCFantasyTeamInfoRequestByOwnerAccountID
// Request type: CMsgDOTAFantasyTeamInfoRequestByOwnerAccountID
func (d *Dota2) RequestTeamInfoFantasyByOwnerAccountID(
	ownerAccountID uint32,
) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamInfoRequestByOwnerAccountID{
		OwnerAccountId: &ownerAccountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamInfoRequestByOwnerAccountID), req)
}

// RequestTeammateStats requests teammate stats.
// Request ID: k_EMsgClientToGCTeammateStatsRequest
// Response ID: k_EMsgClientToGCTeammateStatsResponse
// Request type: CMsgClientToGCTeammateStatsRequest
// Response type: CMsgClientToGCTeammateStatsResponse
func (d *Dota2) RequestTeammateStats(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgClientToGCTeammateStatsResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCTeammateStatsRequest{}
	resp := &dota_gcmessages_client.CMsgClientToGCTeammateStatsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTeammateStatsResponse),
		resp,
	)
}

// RequestTopFriendMatches requests top friend matches.
// Request ID: k_EMsgClientToGCTopFriendMatchesRequest
// Response ID: k_EMsgGCToClientTopFriendMatchesResponse
// Request type: CMsgClientToGCTopFriendMatchesRequest
// Response type: CMsgGCToClientTopFriendMatchesResponse
func (d *Dota2) RequestTopFriendMatches(
	ctx context.Context,
) (*dota_gcmessages_client_watch.CMsgGCToClientTopFriendMatchesResponse, error) {
	req := &dota_gcmessages_client_watch.CMsgClientToGCTopFriendMatchesRequest{}
	resp := &dota_gcmessages_client_watch.CMsgGCToClientTopFriendMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTopFriendMatchesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTopFriendMatchesResponse),
		resp,
	)
}

// RequestTopLeagueMatches requests top league matches.
// Request ID: k_EMsgClientToGCTopLeagueMatchesRequest
// Response ID: k_EMsgGCToClientTopLeagueMatchesResponse
// Request type: CMsgClientToGCTopLeagueMatchesRequest
// Response type: CMsgGCToClientTopLeagueMatchesResponse
func (d *Dota2) RequestTopLeagueMatches(
	ctx context.Context,
) (*dota_gcmessages_client_watch.CMsgGCToClientTopLeagueMatchesResponse, error) {
	req := &dota_gcmessages_client_watch.CMsgClientToGCTopLeagueMatchesRequest{}
	resp := &dota_gcmessages_client_watch.CMsgGCToClientTopLeagueMatchesResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTopLeagueMatchesRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTopLeagueMatchesResponse),
		resp,
	)
}

// RequestTransferSeasonalMMR requests to check if the target transfer seasonal mmr.
// Request ID: k_EMsgClientToGCTransferSeasonalMMRRequest
// Response ID: k_EMsgClientToGCTransferSeasonalMMRResponse
// Request type: CMsgClientToGCTransferSeasonalMMRRequest
// Response type: CMsgClientToGCTransferSeasonalMMRResponse
func (d *Dota2) RequestTransferSeasonalMMR(
	ctx context.Context,
	isParty bool,
) (*dota_gcmessages_client.CMsgClientToGCTransferSeasonalMMRResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCTransferSeasonalMMRRequest{
		IsParty: &isParty,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCTransferSeasonalMMRResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTransferSeasonalMMRResponse),
		resp,
	)
}

// RequestUnanchorPhoneNumber requests a unanchor phone number.
// Request ID: k_EMsgUnanchorPhoneNumberRequest
// Response ID: k_EMsgUnanchorPhoneNumberResponse
// Request type: CMsgDOTAUnanchorPhoneNumberRequest
// Response type: CMsgDOTAUnanchorPhoneNumberResponse
func (d *Dota2) RequestUnanchorPhoneNumber(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTAUnanchorPhoneNumberResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAUnanchorPhoneNumberRequest{}
	resp := &dota_gcmessages_client.CMsgDOTAUnanchorPhoneNumberResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgUnanchorPhoneNumberResponse),
		resp,
	)
}

// RequestWagering requests a wagering.
// Request ID: k_EMsgClientToGCWageringRequest
// Response ID: k_EMsgGCToClientWageringResponse
// Request type: CMsgClientToGCWageringRequest
// Response type: CMsgGCToClientWageringResponse
func (d *Dota2) RequestWagering(
	ctx context.Context,
	eventID uint32,
) (*dota_gcmessages_client.CMsgGCToClientWageringResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCWageringRequest{
		EventId: &eventID,
	}
	resp := &dota_gcmessages_client.CMsgGCToClientWageringResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCWageringRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientWageringResponse),
		resp,
	)
}

// RerollPlayerChallenge rerolls a player challenge.
// Request ID: k_EMsgClientToGCRerollPlayerChallenge
// Request type: CMsgClientToGCRerollPlayerChallenge
func (d *Dota2) RerollPlayerChallenge(
	eventID uint32,
	sequenceID uint32,
) {
	req := &dota_gcmessages_common.CMsgClientToGCRerollPlayerChallenge{
		EventId:    &eventID,
		SequenceId: &sequenceID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCRerollPlayerChallenge), req)
}

// ReserveEditorItemItemDef reserves a editor item item def.
// Request ID: k_EMsgGCItemEditorReserveItemDef
// Response ID: k_EMsgGCItemEditorReserveItemDefResponse
// Request type: CMsgGCItemEditorReserveItemDef
// Response type: CMsgGCItemEditorReserveItemDefResponse
func (d *Dota2) ReserveEditorItemItemDef(
	ctx context.Context,
	defIndex uint32,
	username string,
) (*dota_gcmessages_client.CMsgGCItemEditorReserveItemDefResponse, error) {
	req := &dota_gcmessages_client.CMsgGCItemEditorReserveItemDef{
		DefIndex: &defIndex,
		Username: &username,
	}
	resp := &dota_gcmessages_client.CMsgGCItemEditorReserveItemDefResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDef),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCItemEditorReserveItemDefResponse),
		resp,
	)
}

// RespondToTeamInvite is undocumented.
// Request ID: k_EMsgGCTeamInvite_InviteeResponseToGC
// Request type: CMsgDOTATeamInvite_InviteeResponseToGC
func (d *Dota2) RespondToTeamInvite(
	result dota_gcmessages_client_team.ETeamInviteResult,
) {
	req := &dota_gcmessages_client_team.CMsgDOTATeamInvite_InviteeResponseToGC{
		Result: &result,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC), req)
}

// RetrieveMatchVote retrieves a match vote.
// Request ID: k_EMsgRetrieveMatchVote
// Response ID: k_EMsgRetrieveMatchVoteResponse
// Request type: CMsgRetrieveMatchVote
// Response type: CMsgMatchVoteResponse
func (d *Dota2) RetrieveMatchVote(
	ctx context.Context,
	matchID uint64,
	incremental uint32,
) (*dota_gcmessages_client.CMsgMatchVoteResponse, error) {
	req := &dota_gcmessages_client.CMsgRetrieveMatchVote{
		MatchId:     &matchID,
		Incremental: &incremental,
	}
	resp := &dota_gcmessages_client.CMsgMatchVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgRetrieveMatchVote),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgRetrieveMatchVoteResponse),
		resp,
	)
}

// SelectCompendiumInGamePrediction selects a compendium in game prediction.
// Request ID: k_EMsgClientToGCSelectCompendiumInGamePrediction
// Response ID: k_EMsgClientToGCSelectCompendiumInGamePredictionResponse
// Request type: CMsgClientToGCSelectCompendiumInGamePrediction
// Response type: CMsgClientToGCSelectCompendiumInGamePredictionResponse
func (d *Dota2) SelectCompendiumInGamePrediction(
	ctx context.Context,
	matchID uint64,
	predictions []*dota_gcmessages_client.CMsgClientToGCSelectCompendiumInGamePrediction_Prediction,
	leagueID uint32,
) (*dota_gcmessages_client.CMsgClientToGCSelectCompendiumInGamePredictionResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSelectCompendiumInGamePrediction{
		MatchId:     &matchID,
		Predictions: predictions,
		LeagueId:    &leagueID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCSelectCompendiumInGamePredictionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePrediction),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSelectCompendiumInGamePredictionResponse),
		resp,
	)
}

// SendAddTI6TreeProgress sends add ti 6 tree progress.
// Request ID: k_EMsgClientToGCAddTI6TreeProgress
// Request type: CMsgClientToGCAddTI6TreeProgress
func (d *Dota2) SendAddTI6TreeProgress(
	trees uint32,
) {
	req := &dota_gcmessages_client.CMsgClientToGCAddTI6TreeProgress{
		Trees: &trees,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCAddTI6TreeProgress), req)
}

// SendChatChannelMemberUpdate sends a chat channel member update.
// Request ID: k_EMsgDOTAChatChannelMemberUpdate
// Request type: CMsgDOTAChatChannelMemberUpdate
func (d *Dota2) SendChatChannelMemberUpdate(
	channelID uint64,
	leftSteamIDs []uint64,
	joinedMembers []*dota_gcmessages_client_chat.CMsgDOTAChatChannelMemberUpdate_JoinedMember,
) {
	req := &dota_gcmessages_client_chat.CMsgDOTAChatChannelMemberUpdate{
		ChannelId:     &channelID,
		LeftSteamIds:  leftSteamIDs,
		JoinedMembers: joinedMembers,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAChatChannelMemberUpdate), req)
}

// SendChatMessage sends a chat message.
// Request ID: k_EMsgGCChatMessage
// Request type: CMsgDOTAChatMessage
func (d *Dota2) SendChatMessage(
	req *dota_gcmessages_client_chat.CMsgDOTAChatMessage,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCChatMessage), req)
}

// SendClaimEventAction sends a claim event action.
// Request ID: k_EMsgDOTAClaimEventAction
// Response ID: k_EMsgDOTAClaimEventActionResponse
// Request type: CMsgDOTAClaimEventAction
// Response type: CMsgDOTAClaimEventActionResponse
func (d *Dota2) SendClaimEventAction(
	ctx context.Context,
	eventID uint32,
	actionID uint32,
	quantity uint32,
) (*dota_gcmessages_client.CMsgDOTAClaimEventActionResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAClaimEventAction{
		EventId:  &eventID,
		ActionId: &actionID,
		Quantity: &quantity,
	}
	resp := &dota_gcmessages_client.CMsgDOTAClaimEventActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAClaimEventAction),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAClaimEventActionResponse),
		resp,
	)
}

// SendClientProvideSurveyResult sends a client provide survey result.
// Request ID: k_EMsgClientProvideSurveyResult
// Request type: CMsgClientProvideSurveyResult
func (d *Dota2) SendClientProvideSurveyResult(
	responses []*dota_gcmessages_client.CMsgClientProvideSurveyResult_Response,
	surveyKey uint64,
) {
	req := &dota_gcmessages_client.CMsgClientProvideSurveyResult{
		Responses: responses,
		SurveyKey: &surveyKey,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientProvideSurveyResult), req)
}

// SendCustomGameClientFinishedLoading sends a custom game client finished loading.
// Request ID: k_EMsgCustomGameClientFinishedLoading
// Request type: CMsgDOTACustomGameClientFinishedLoading
func (d *Dota2) SendCustomGameClientFinishedLoading(
	lobbyID uint64,
	loadingDuration uint32,
	resultCode int32,
	resultString string,
	signonStates uint32,
	comment string,
) {
	req := &dota_gcmessages_client.CMsgDOTACustomGameClientFinishedLoading{
		LobbyId:         &lobbyID,
		LoadingDuration: &loadingDuration,
		ResultCode:      &resultCode,
		ResultString:    &resultString,
		SignonStates:    &signonStates,
		Comment:         &comment,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgCustomGameClientFinishedLoading), req)
}

// SendCustomGameListenServerStartedLoading sends a custom game listen server started loading.
// Request ID: k_EMsgCustomGameListenServerStartedLoading
// Request type: CMsgDOTACustomGameListenServerStartedLoading
func (d *Dota2) SendCustomGameListenServerStartedLoading(
	lobbyID uint64,
	customGameID uint64,
	lobbyMembers []uint32,
	startTime uint32,
) {
	req := &dota_gcmessages_client.CMsgDOTACustomGameListenServerStartedLoading{
		LobbyId:      &lobbyID,
		CustomGameId: &customGameID,
		LobbyMembers: lobbyMembers,
		StartTime:    &startTime,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgCustomGameListenServerStartedLoading), req)
}

// SendFriendRecruitInviteAcceptDecline sends a friend recruit invite accept decline.
// Request ID: k_EMsgDOTAFriendRecruitInviteAcceptDecline
// Request type: CMsgDOTAFriendRecruitInviteAcceptDecline
func (d *Dota2) SendFriendRecruitInviteAcceptDecline(
	accepted bool,
	accountID uint32,
) {
	req := &dota_gcmessages_client.CMsgDOTAFriendRecruitInviteAcceptDecline{
		Accepted:  &accepted,
		AccountId: &accountID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAFriendRecruitInviteAcceptDecline), req)
}

// SendFriendRecruits sends friend recruits.
// Request ID: k_EMsgDOTASendFriendRecruits
// Request type: CMsgDOTASendFriendRecruits
func (d *Dota2) SendFriendRecruits(
	recruits []uint32,
) {
	req := &dota_gcmessages_client.CMsgDOTASendFriendRecruits{
		Recruits: recruits,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTASendFriendRecruits), req)
}

// SendH264Unsupported sends a h 264 unsupported.
// Request ID: k_EMsgClientToGCH264Unsupported
// Request type: CMsgClientToGCH264Unsupported
func (d *Dota2) SendH264Unsupported() {
	req := &dota_gcmessages_client.CMsgClientToGCH264Unsupported{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCH264Unsupported), req)
}

// SendHasPlayerVotedForMVP sends a has player voted for mvp.
// Request ID: k_EMsgClientToGCHasPlayerVotedForMVP
// Response ID: k_EMsgClientToGCHasPlayerVotedForMVPResponse
// Request type: CMsgClientToGCHasPlayerVotedForMVP
// Response type: CMsgClientToGCHasPlayerVotedForMVPResponse
func (d *Dota2) SendHasPlayerVotedForMVP(
	ctx context.Context,
	matchID uint64,
) (*dota_gcmessages_client.CMsgClientToGCHasPlayerVotedForMVPResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCHasPlayerVotedForMVP{
		MatchId: &matchID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCHasPlayerVotedForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVP),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCHasPlayerVotedForMVPResponse),
		resp,
	)
}

// SendInitialQuestionnaireResponse sends a initial questionnaire response.
// Request ID: k_EMsgGCInitialQuestionnaireResponse
// Request type: CMsgInitialQuestionnaireResponse
func (d *Dota2) SendInitialQuestionnaireResponse(
	initialSkill uint32,
) {
	req := &dota_gcmessages_client.CMsgInitialQuestionnaireResponse{
		InitialSkill: &initialSkill,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse), req)
}

// SendLatestConductScorecard sends a latest conduct scorecard.
// Request ID: k_EMsgClientToGCLatestConductScorecard
// Request type: CMsgPlayerConductScorecard
func (d *Dota2) SendLatestConductScorecard(
	req *dota_gcmessages_client.CMsgPlayerConductScorecard,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard), req)
}

// SendLobbyBattleCupVictory sends a lobby battle cup victory.
// Request ID: k_EMsgLobbyBattleCupVictory
// Request type: CMsgBattleCupVictory
func (d *Dota2) SendLobbyBattleCupVictory(
	req *dota_gcmessages_common.CMsgBattleCupVictory,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory), req)
}

// SendLobbyEventPoints sends lobby event points.
// Request ID: k_EMsgLobbyEventPoints
// Request type: CMsgLobbyEventPoints
func (d *Dota2) SendLobbyEventPoints(
	eventID uint32,
	accountPoints []*dota_gcmessages_common.CMsgLobbyEventPoints_AccountPoints,
) {
	req := &dota_gcmessages_common.CMsgLobbyEventPoints{
		EventId:       &eventID,
		AccountPoints: accountPoints,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgLobbyEventPoints), req)
}

// SendLobbyPlaytestDetails sends lobby playtest details.
// Request ID: k_EMsgLobbyPlaytestDetails
// Request type: CMsgLobbyPlaytestDetails
func (d *Dota2) SendLobbyPlaytestDetails(
	jSON string,
) {
	req := &dota_gcmessages_common_match_management.CMsgLobbyPlaytestDetails{
		Json: &jSON,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgLobbyPlaytestDetails), req)
}

// SendMergePartyInvite sends a merge party invite.
// Request ID: k_EMsgClientToGCMergePartyInvite
// Request type: CMsgDOTAGroupMergeInvite
func (d *Dota2) SendMergePartyInvite(
	otherGroupID uint64,
) {
	req := &dota_gcmessages_client_match_management.CMsgDOTAGroupMergeInvite{
		OtherGroupId: &otherGroupID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite), req)
}

// SendPeriodicResourceUpdated sends a periodic resource updated.
// Request ID: k_EMsgDOTAPeriodicResourceUpdated
// Request type: CMsgDOTAPeriodicResourceUpdated
func (d *Dota2) SendPeriodicResourceUpdated(
	periodicResourceKey dota_gcmessages_client.CMsgDOTAGetPeriodicResource,
	periodicResourceValue dota_gcmessages_client.CMsgDOTAGetPeriodicResourceResponse,
) {
	req := &dota_gcmessages_client.CMsgDOTAPeriodicResourceUpdated{
		PeriodicResourceKey:   &periodicResourceKey,
		PeriodicResourceValue: &periodicResourceValue,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAPeriodicResourceUpdated), req)
}

// SendPingData sends a ping data.
// Request ID: k_EMsgClientToGCPingData
// Request type: CMsgClientPingData
func (d *Dota2) SendPingData(
	relayCodes []uint32,
	relayPings []uint32,
	regionCodes []uint32,
	regionPings []uint32,
	regionPingFailedBitmask uint32,
) {
	req := &base_gcmessages.CMsgClientPingData{
		RelayCodes:              relayCodes,
		RelayPings:              relayPings,
		RegionCodes:             regionCodes,
		RegionPings:             regionPings,
		RegionPingFailedBitmask: &regionPingFailedBitmask,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCPingData), req)
}

// SendPresentedClientTerminateDlg sends a presented client terminate dlg.
// Request ID: k_EMsgPresentedClientTerminateDlg
// Request type: CMsgPresentedClientTerminateDlg
func (d *Dota2) SendPresentedClientTerminateDlg() {
	req := &dota_gcmessages_client.CMsgPresentedClientTerminateDlg{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgPresentedClientTerminateDlg), req)
}

// SendReadyUp sends a ready up.
// Request ID: k_EMsgGCReadyUp
// Request type: CMsgReadyUp
func (d *Dota2) SendReadyUp(
	state dota_shared_enums.DOTALobbyReadyState,
	readyUpKey uint64,
	hardwareSpecs dota_shared_enums.CDOTAClientHardwareSpecs,
) {
	req := &dota_gcmessages_client_match_management.CMsgReadyUp{
		State:         &state,
		ReadyUpKey:    &readyUpKey,
		HardwareSpecs: &hardwareSpecs,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReadyUp), req)
}

// SendResponseLeagueInfo sends a response league info.
// Request ID: k_EMsgResponseLeagueInfo
// Request type: CMsgResponseLeagueInfo
func (d *Dota2) SendResponseLeagueInfo(
	leagues []*dota_gcmessages_client.CDynamicLeagueData,
) {
	req := &dota_gcmessages_client.CMsgResponseLeagueInfo{
		Leagues: leagues,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgResponseLeagueInfo), req)
}

// SendSpectatorLobbyGameDetails sends spectator lobby game details.
// Request ID: k_EMsgSpectatorLobbyGameDetails
// Request type: CMsgSpectatorLobbyGameDetails
func (d *Dota2) SendSpectatorLobbyGameDetails(
	req *dota_gcmessages_client_match_management.CMsgSpectatorLobbyGameDetails,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgSpectatorLobbyGameDetails), req)
}

// SendStopFindingMatch sends a stop finding match.
// Request ID: k_EMsgGCStopFindingMatch
// Request type: CMsgStopFindingMatch
func (d *Dota2) SendStopFindingMatch() {
	req := &dota_gcmessages_client_match_management.CMsgStopFindingMatch{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCStopFindingMatch), req)
}

// SendSuspiciousActivity sends a suspicious activity.
// Request ID: k_EMsgClientToGCSuspiciousActivity
// Request type: CMsgClientToGCSuspiciousActivity
func (d *Dota2) SendSuspiciousActivity(
	appData uint64,
) {
	req := &dota_gcmessages_client.CMsgClientToGCSuspiciousActivity{
		AppData: &appData,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSuspiciousActivity), req)
}

// SendTeamInvite_GCResponseToInvitee sends a team invite _ gc response to invitee.
// Request ID: k_EMsgGCTeamInvite_GCResponseToInvitee
// Request type: CMsgDOTATeamInvite_GCResponseToInvitee
func (d *Dota2) SendTeamInvite_GCResponseToInvitee(
	result dota_gcmessages_client_team.ETeamInviteResult,
	teamName string,
) {
	req := &dota_gcmessages_client_team.CMsgDOTATeamInvite_GCResponseToInvitee{
		Result:   &result,
		TeamName: &teamName,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee), req)
}

// SetAdditionalEquips sets additional equips.
// Request ID: k_EMsgClientToGCSetAdditionalEquips
// Response ID: k_EMsgClientToGCSetAdditionalEquipsResponse
// Request type: CMsgClientToGCSetAdditionalEquips
// Response type: CMsgClientToGCSetAdditionalEquipsResponse
func (d *Dota2) SetAdditionalEquips(
	ctx context.Context,
	equips []*dota_gcmessages_common.CAdditionalEquipSlot,
) (*dota_gcmessages_client.CMsgClientToGCSetAdditionalEquipsResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCSetAdditionalEquips{
		Equips: equips,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCSetAdditionalEquipsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetAdditionalEquips),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetAdditionalEquipsResponse),
		resp,
	)
}

// SetCompendiumSelection sets a compendium selection.
// Request ID: k_EMsgGCCompendiumSetSelection
// Response ID: k_EMsgGCCompendiumSetSelectionResponse
// Request type: CMsgDOTACompendiumSelection
// Response type: CMsgDOTACompendiumSelectionResponse
func (d *Dota2) SetCompendiumSelection(
	ctx context.Context,
	selectionIndex uint32,
	selection uint32,
	leagueid uint32,
) (*dota_gcmessages_client.CMsgDOTACompendiumSelectionResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTACompendiumSelection{
		SelectionIndex: &selectionIndex,
		Selection:      &selection,
		Leagueid:       &leagueid,
	}
	resp := &dota_gcmessages_client.CMsgDOTACompendiumSelectionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse),
		resp,
	)
}

// SetFavoriteTeam sets a favorite team.
// Request ID: k_EMsgDOTASetFavoriteTeam
// Request type: CMsgDOTASetFavoriteTeam
func (d *Dota2) SetFavoriteTeam(
	teamID uint32,
	eventID uint32,
) {
	req := &dota_gcmessages_client.CMsgDOTASetFavoriteTeam{
		TeamId:  &teamID,
		EventId: &eventID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTASetFavoriteTeam), req)
}

// SetFeaturedItems sets featured items.
// Request ID: k_EMsgGCSetFeaturedItems
// Request type: CMsgSetFeaturedItems
func (d *Dota2) SetFeaturedItems(
	featuredItemID []uint64,
) {
	req := &dota_gcmessages_client.CMsgSetFeaturedItems{
		FeaturedItemId: featuredItemID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetFeaturedItems), req)
}

// SetLobbyCoach sets a lobby coach.
// Request ID: k_EMsgGCPracticeLobbySetCoach
// Request type: CMsgPracticeLobbySetCoach
func (d *Dota2) SetLobbyCoach(
	team dota_shared_enums.DOTA_GC_TEAM,
) {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbySetCoach{
		Team: &team,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbySetCoach), req)
}

// SetLobbyDetails sets lobby details.
// Request ID: k_EMsgGCPracticeLobbySetDetails
// Request type: CMsgPracticeLobbySetDetails
func (d *Dota2) SetLobbyDetails(
	req *dota_gcmessages_client_match_management.CMsgPracticeLobbySetDetails,
) {
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails), req)
}

// SetMapLocationState sets a map location state.
// Request ID: k_EMsgGCSetMapLocationState
// Response ID: k_EMsgGCSetMapLocationStateResponse
// Request type: CMsgSetMapLocationState
// Response type: CMsgSetMapLocationStateResponse
func (d *Dota2) SetMapLocationState(
	ctx context.Context,
	locationID int32,
	completed bool,
) (*dota_gcmessages_client.CMsgSetMapLocationStateResponse, error) {
	req := &dota_gcmessages_client.CMsgSetMapLocationState{
		LocationId: &locationID,
		Completed:  &completed,
	}
	resp := &dota_gcmessages_client.CMsgSetMapLocationStateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetMapLocationState),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetMapLocationStateResponse),
		resp,
	)
}

// SetMatchHistoryAccess sets match history access.
// Request ID: k_EMsgGCSetMatchHistoryAccess
// Response ID: k_EMsgGCSetMatchHistoryAccessResponse
// Request type: CMsgDOTASetMatchHistoryAccess
// Response type: CMsgDOTASetMatchHistoryAccessResponse
func (d *Dota2) SetMatchHistoryAccess(
	ctx context.Context,
	allow3RdPartyMatchHistory bool,
) (*dota_gcmessages_client.CMsgDOTASetMatchHistoryAccessResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASetMatchHistoryAccess{
		Allow_3RdPartyMatchHistory: &allow3RdPartyMatchHistory,
	}
	resp := &dota_gcmessages_client.CMsgDOTASetMatchHistoryAccessResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccess),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetMatchHistoryAccessResponse),
		resp,
	)
}

// SetMemberPartyCoach sets a member party coach.
// Request ID: k_EMsgGCPartyMemberSetCoach
// Request type: CMsgDOTAPartyMemberSetCoach
func (d *Dota2) SetMemberPartyCoach(
	wantsCoach bool,
) {
	req := &dota_gcmessages_client_match_management.CMsgDOTAPartyMemberSetCoach{
		WantsCoach: &wantsCoach,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPartyMemberSetCoach), req)
}

// SetPartyBuilderOptions sets party builder options.
// Request ID: k_EMsgClientToGCSetPartyBuilderOptions
// Request type: CMsgPartyBuilderOptions
func (d *Dota2) SetPartyBuilderOptions(
	additionalSlots uint32,
	matchType dota_shared_enums.MatchType,
	matchgroups uint32,
	language dota_shared_enums.MatchLanguages,
) {
	req := &dota_gcmessages_client_match_management.CMsgPartyBuilderOptions{
		AdditionalSlots: &additionalSlots,
		MatchType:       &matchType,
		Matchgroups:     &matchgroups,
		Language:        &language,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions), req)
}

// SetPartyLeader sets a party leader.
// Request ID: k_EMsgClientToGCSetPartyLeader
// Request type: CMsgDOTASetGroupLeader
func (d *Dota2) SetPartyLeader(
	newLeaderSteamid steamid.SteamId,
) {
	newLeaderSteamidU64Val := uint64(newLeaderSteamid)
	newLeaderSteamidU64 := &newLeaderSteamidU64Val
	req := &dota_gcmessages_client_match_management.CMsgDOTASetGroupLeader{
		NewLeaderSteamid: newLeaderSteamidU64,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader), req)
}

// SetPartyOpen sets a party open.
// Request ID: k_EMsgClientToGCSetPartyOpen
// Request type: CMsgDOTASetGroupOpenStatus
func (d *Dota2) SetPartyOpen(
	open bool,
) {
	req := &dota_gcmessages_client_match_management.CMsgDOTASetGroupOpenStatus{
		Open: &open,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen), req)
}

// SetProfileCardSlots sets profile card slots.
// Request ID: k_EMsgClientToGCSetProfileCardSlots
// Request type: CMsgClientToGCSetProfileCardSlots
func (d *Dota2) SetProfileCardSlots(
	slots []*dota_gcmessages_client.CMsgClientToGCSetProfileCardSlots_CardSlot,
) {
	req := &dota_gcmessages_client.CMsgClientToGCSetProfileCardSlots{
		Slots: slots,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetProfileCardSlots), req)
}

// SetProfilePrivacy sets a profile privacy.
// Request ID: k_EMsgGCSetProfilePrivacy
// Response ID: k_EMsgGCSetProfilePrivacyResponse
// Request type: CMsgDOTASetProfilePrivacy
// Response type: CMsgDOTASetProfilePrivacyResponse
func (d *Dota2) SetProfilePrivacy(
	ctx context.Context,
	profilePrivate bool,
) (*dota_gcmessages_client.CMsgDOTASetProfilePrivacyResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASetProfilePrivacy{
		ProfilePrivate: &profilePrivate,
	}
	resp := &dota_gcmessages_client.CMsgDOTASetProfilePrivacyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetProfilePrivacy),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetProfilePrivacyResponse),
		resp,
	)
}

// SetShowcaseHero sets a showcase hero.
// Request ID: k_EMsgGCSetShowcaseHero
// Request type: CMsgSetShowcaseHero
func (d *Dota2) SetShowcaseHero(
	showcaseHeroID uint32,
) {
	req := &dota_gcmessages_client.CMsgSetShowcaseHero{
		ShowcaseHeroId: &showcaseHeroID,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSetShowcaseHero), req)
}

// SetSpectatorLobbyDetails sets spectator lobby details.
// Request ID: k_EMsgClientToGCSetSpectatorLobbyDetails
// Request type: CMsgSetSpectatorLobbyDetails
func (d *Dota2) SetSpectatorLobbyDetails(
	lobbyID uint64,
	lobbyName string,
	passKey string,
	gameDetails dota_gcmessages_client_match_management.CMsgSpectatorLobbyGameDetails,
) {
	req := &dota_gcmessages_client_match_management.CMsgSetSpectatorLobbyDetails{
		LobbyId:     &lobbyID,
		LobbyName:   &lobbyName,
		PassKey:     &passKey,
		GameDetails: &gameDetails,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetails), req)
}

// SpectateFriendGame spectates a friend game.
// Request ID: k_EMsgGCSpectateFriendGame
// Response ID: k_EMsgGCSpectateFriendGameResponse
// Request type: CMsgSpectateFriendGame
// Response type: CMsgSpectateFriendGameResponse
func (d *Dota2) SpectateFriendGame(
	ctx context.Context,
	steamID steamid.SteamId,
) (*dota_gcmessages_client_watch.CMsgSpectateFriendGameResponse, error) {
	steamIDU64Val := uint64(steamID)
	steamIDU64 := &steamIDU64Val
	req := &dota_gcmessages_client_watch.CMsgSpectateFriendGame{
		SteamId: steamIDU64,
	}
	resp := &dota_gcmessages_client_watch.CMsgSpectateFriendGameResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSpectateFriendGame),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSpectateFriendGameResponse),
		resp,
	)
}

// StartFindingMatch starts a finding match.
// Request ID: k_EMsgGCStartFindingMatch
// Response ID: k_EMsgGCStartFindingMatchResponse
// Request type: CMsgStartFindingMatch
// Response type: CMsgStartFindingMatchResult
func (d *Dota2) StartFindingMatch(
	ctx context.Context,
	req *dota_gcmessages_client_match_management.CMsgStartFindingMatch,
) (*dota_gcmessages_client_match_management.CMsgStartFindingMatchResult, error) {
	resp := &dota_gcmessages_client_match_management.CMsgStartFindingMatchResult{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCStartFindingMatch),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCStartFindingMatchResponse),
		resp,
	)
}

// StartTriviaSession starts a trivia session.
// Request ID: k_EMsgStartTriviaSession
// Response ID: k_EMsgStartTriviaSessionResponse
// Request type: CMsgDOTAStartTriviaSession
// Response type: CMsgDOTAStartTriviaSessionResponse
func (d *Dota2) StartTriviaSession(
	ctx context.Context,
) (*dota_gcmessages_client.CMsgDOTAStartTriviaSessionResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTAStartTriviaSession{}
	resp := &dota_gcmessages_client.CMsgDOTAStartTriviaSessionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgStartTriviaSession),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgStartTriviaSessionResponse),
		resp,
	)
}

// SubmitInfoPlayer submits a info player.
// Request ID: k_EMsgGCPlayerInfoSubmit
// Response ID: k_EMsgGCPlayerInfoSubmitResponse
// Request type: CMsgGCPlayerInfoSubmit
// Response type: CMsgGCPlayerInfoSubmitResponse
func (d *Dota2) SubmitInfoPlayer(
	ctx context.Context,
	name string,
	countryCode string,
	fantasyRole uint32,
	teamID uint32,
	sponsor string,
) (*dota_gcmessages_client.CMsgGCPlayerInfoSubmitResponse, error) {
	req := &dota_gcmessages_client.CMsgGCPlayerInfoSubmit{
		Name:        &name,
		CountryCode: &countryCode,
		FantasyRole: &fantasyRole,
		TeamId:      &teamID,
		Sponsor:     &sponsor,
	}
	resp := &dota_gcmessages_client.CMsgGCPlayerInfoSubmitResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmit),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfoSubmitResponse),
		resp,
	)
}

// SubmitLobbyMVPVote submits a lobby mvp vote.
// Request ID: k_EMsgGCSubmitLobbyMVPVote
// Response ID: k_EMsgGCSubmitLobbyMVPVoteResponse
// Request type: CMsgDOTASubmitLobbyMVPVote
// Response type: CMsgDOTASubmitLobbyMVPVoteResponse
func (d *Dota2) SubmitLobbyMVPVote(
	ctx context.Context,
	targetAccountID uint32,
) (*dota_gcmessages_client.CMsgDOTASubmitLobbyMVPVoteResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASubmitLobbyMVPVote{
		TargetAccountId: &targetAccountID,
	}
	resp := &dota_gcmessages_client.CMsgDOTASubmitLobbyMVPVoteResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVote),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSubmitLobbyMVPVoteResponse),
		resp,
	)
}

// SubmitPlayerReport submits a player report.
// Request ID: k_EMsgGCSubmitPlayerReport
// Response ID: k_EMsgGCSubmitPlayerReportResponse
// Request type: CMsgDOTASubmitPlayerReport
// Response type: CMsgDOTASubmitPlayerReportResponse
func (d *Dota2) SubmitPlayerReport(
	ctx context.Context,
	targetAccountID uint32,
	reportFlags uint32,
	lobbyID uint64,
	comment string,
) (*dota_gcmessages_client.CMsgDOTASubmitPlayerReportResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASubmitPlayerReport{
		TargetAccountId: &targetAccountID,
		ReportFlags:     &reportFlags,
		LobbyId:         &lobbyID,
		Comment:         &comment,
	}
	resp := &dota_gcmessages_client.CMsgDOTASubmitPlayerReportResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSubmitPlayerReport),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCSubmitPlayerReportResponse),
		resp,
	)
}

// SubmitTriviaQuestionAnswer submits a trivia question answer.
// Request ID: k_EMsgSubmitTriviaQuestionAnswer
// Response ID: k_EMsgSubmitTriviaQuestionAnswerResponse
// Request type: CMsgDOTASubmitTriviaQuestionAnswer
// Response type: CMsgDOTASubmitTriviaQuestionAnswerResponse
func (d *Dota2) SubmitTriviaQuestionAnswer(
	ctx context.Context,
	questionID uint32,
	answerIndex uint32,
) (*dota_gcmessages_client.CMsgDOTASubmitTriviaQuestionAnswerResponse, error) {
	req := &dota_gcmessages_client.CMsgDOTASubmitTriviaQuestionAnswer{
		QuestionId:  &questionID,
		AnswerIndex: &answerIndex,
	}
	resp := &dota_gcmessages_client.CMsgDOTASubmitTriviaQuestionAnswerResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswer),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgSubmitTriviaQuestionAnswerResponse),
		resp,
	)
}

// SwapFantasyTeamRoster swaps a fantasy team roster.
// Request ID: k_EMsgGCFantasyTeamRosterSwapRequest
// Response ID: k_EMsgGCFantasyTeamRosterSwapResponse
// Request type: CMsgDOTAFantasyTeamRosterSwapRequest
// Response type: CMsgDOTAFantasyTeamRosterSwapResponse
func (d *Dota2) SwapFantasyTeamRoster(
	ctx context.Context,
	fantasyLeagueID uint32,
	teamIndex uint32,
	timestamp uint32,
	slot1 uint32,
	slot2 uint32,
) (*dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterSwapResponse, error) {
	req := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterSwapRequest{
		FantasyLeagueId: &fantasyLeagueID,
		TeamIndex:       &teamIndex,
		Timestamp:       &timestamp,
		Slot_1:          &slot1,
		Slot_2:          &slot2,
	}
	resp := &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamRosterSwapResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterSwapRequest),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamRosterSwapResponse),
		resp,
	)
}

// ToggleLobbyBroadcastChannelCameramanStatus toggles lobby broadcast channel cameraman status.
// Request ID: k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus
// Request type: CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus
func (d *Dota2) ToggleLobbyBroadcastChannelCameramanStatus() {
	req := &dota_gcmessages_client_match_management.CMsgPracticeLobbyToggleBroadcastChannelCameramanStatus{}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPracticeLobbyToggleBroadcastChannelCameramanStatus), req)
}

// TrackDialogResult tracks a dialog result.
// Request ID: k_EMsgClientToGCTrackDialogResult
// Request type: CMsgClientToGCTrackDialogResult
func (d *Dota2) TrackDialogResult(
	dialogID uint32,
	value uint32,
) {
	req := &dota_gcmessages_client.CMsgClientToGCTrackDialogResult{
		DialogId: &dialogID,
		Value:    &value,
	}
	d.write(uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCTrackDialogResult), req)
}

// TransferTeamAdmin transfers a team admin.
// Request ID: k_EMsgGCTransferTeamAdmin
// Response ID: k_EMsgGCTransferTeamAdminResponse
// Request type: CMsgDOTATransferTeamAdmin
// Response type: CMsgDOTATransferTeamAdminResponse
func (d *Dota2) TransferTeamAdmin(
	ctx context.Context,
	newAdminAccountID uint32,
	teamID uint32,
) (*dota_gcmessages_client_team.CMsgDOTATransferTeamAdminResponse, error) {
	req := &dota_gcmessages_client_team.CMsgDOTATransferTeamAdmin{
		NewAdminAccountId: &newAdminAccountID,
		TeamId:            &teamID,
	}
	resp := &dota_gcmessages_client_team.CMsgDOTATransferTeamAdminResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTransferTeamAdmin),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTransferTeamAdminResponse),
		resp,
	)
}

// UpgradeLeagueItem upgrades a league item.
// Request ID: k_EMsgUpgradeLeagueItem
// Response ID: k_EMsgUpgradeLeagueItemResponse
// Request type: CMsgUpgradeLeagueItem
// Response type: CMsgUpgradeLeagueItemResponse
func (d *Dota2) UpgradeLeagueItem(
	ctx context.Context,
	matchID uint64,
	leagueID uint32,
) (*dota_gcmessages_client.CMsgUpgradeLeagueItemResponse, error) {
	req := &dota_gcmessages_client.CMsgUpgradeLeagueItem{
		MatchId:  &matchID,
		LeagueId: &leagueID,
	}
	resp := &dota_gcmessages_client.CMsgUpgradeLeagueItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgUpgradeLeagueItem),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgUpgradeLeagueItemResponse),
		resp,
	)
}

// VoteForArcana votes a for arcana.
// Request ID: k_EMsgClientToGCVoteForArcana
// Response ID: k_EMsgClientToGCVoteForArcanaResponse
// Request type: CMsgClientToGCVoteForArcana
// Response type: CMsgClientToGCVoteForArcanaResponse
func (d *Dota2) VoteForArcana(
	ctx context.Context,
	matches []*dota_gcmessages_client.CMsgClientToGCVoteForArcana_MatchVote,
) (*dota_gcmessages_client.CMsgClientToGCVoteForArcanaResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCVoteForArcana{
		Matches: matches,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCVoteForArcanaResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCVoteForArcana),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCVoteForArcanaResponse),
		resp,
	)
}

// VoteForMVP votes a for mvp.
// Request ID: k_EMsgClientToGCVoteForMVP
// Response ID: k_EMsgClientToGCVoteForMVPResponse
// Request type: CMsgClientToGCVoteForMVP
// Response type: CMsgClientToGCVoteForMVPResponse
func (d *Dota2) VoteForMVP(
	ctx context.Context,
	matchID uint64,
	accountID uint32,
) (*dota_gcmessages_client.CMsgClientToGCVoteForMVPResponse, error) {
	req := &dota_gcmessages_client.CMsgClientToGCVoteForMVP{
		MatchId:   &matchID,
		AccountId: &accountID,
	}
	resp := &dota_gcmessages_client.CMsgClientToGCVoteForMVPResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCVoteForMVP),
		req,
		uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgClientToGCVoteForMVPResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Dota2) registerGeneratedHandlers() {
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesReply)] = d.getEventEmitter(func() events.Event {
		return &events.AllStarVotesReply{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesRequest)] = d.getEventEmitter(func() events.Event {
		return &events.AllStarVotesRequest{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesSubmit)] = d.getEventEmitter(func() events.Event {
		return &events.AllStarVotesSubmit{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientAllStarVotesSubmitReply)] = d.getEventEmitter(func() events.Event {
		return &events.AllStarVotesSubmitReply{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientArcanaVotesUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.ArcanaVotesUpdate{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby)] = d.getEventEmitter(func() events.Event {
		return &events.BalancedShuffleLobby{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupListRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupListRequest{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupRequest)] = d.getEventEmitter(func() events.Event {
		return &events.BattlePassRollupRequest{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCBroadcastNotification)] = d.getEventEmitter(func() events.Event {
		return &events.BroadcastNotification{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCChangeTeamSub)] = d.getEventEmitter(func() events.Event {
		return &events.ChangeTeamSub{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled)] = d.getEventEmitter(func() events.Event {
		return &events.ChatRegionsEnabled{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClearPracticeLobbyTeam)] = d.getEventEmitter(func() events.Event {
		return &events.ClearPracticeLobbyTeam{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClientIgnoredUser)] = d.getEventEmitter(func() events.Event {
		return &events.ClientIgnoredUser{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClientSuspended)] = d.getEventEmitter(func() events.Event {
		return &events.ClientSuspended{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumDataChanged)] = d.getEventEmitter(func() events.Event {
		return &events.CompendiumDataChanged{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTALiveLeagueGameUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.DOTALiveLeagueGameUpdate{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule)] = d.getEventEmitter(func() events.Event {
		return &events.DOTAWeekendTourneySchedule{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientEmoticonData)] = d.getEventEmitter(func() events.Event {
		return &events.EmoticonData{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientEventStatusChanged)] = d.getEventEmitter(func() events.Event {
		return &events.EventStatusChanged{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyFinalPlayerStats{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftStatus)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyLeagueDraftStatus{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueInfo)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyLeagueInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyMessageAdd)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyMessageAdd{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyRemoveOwner)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyRemoveOwner{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamInfo)] = d.getEventEmitter(func() events.Event {
		return &events.FantasyTeamInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFeaturedItems)] = d.getEventEmitter(func() events.Event {
		return &events.FeaturedItems{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildInviteData)] = d.getEventEmitter(func() events.Event {
		return &events.GuildInviteData{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildUpdateMessage)] = d.getEventEmitter(func() events.Event {
		return &events.GuildUpdateMessage{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHallOfFame)] = d.getEventEmitter(func() events.Event {
		return &events.HallOfFame{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientHeroStatueCreateResult)] = d.getEventEmitter(func() events.Event {
		return &events.HeroStatueCreateResult{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue)] = d.getEventEmitter(func() events.Event {
		return &events.KickedFromMatchmakingQueue{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLastHitChallengeHighScorePost)] = d.getEventEmitter(func() events.Event {
		return &events.LastHitChallengeHighScorePost{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeagueAdminList)] = d.getEventEmitter(func() events.Event {
		return &events.LeagueAdminList{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeagueAdminState)] = d.getEventEmitter(func() events.Event {
		return &events.LeagueAdminState{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyMVPAwarded{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyMVPNotifyRecipient{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLobbyUpdateBroadcastChannelInfo)] = d.getEventEmitter(func() events.Event {
		return &events.LobbyUpdateBroadcastChannelInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMakeOffering)] = d.getEventEmitter(func() events.Event {
		return &events.MakeOffering{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMatchGroupsVersion)] = d.getEventEmitter(func() events.Event {
		return &events.MatchGroupsVersion{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMatchSignedOut)] = d.getEventEmitter(func() events.Event {
		return &events.MatchSignedOut{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergeGroupInviteReply{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply)] = d.getEventEmitter(func() events.Event {
		return &events.MergePartyResponseReply{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientNewNotificationAdded)] = d.getEventEmitter(func() events.Event {
		return &events.NewNotificationAdded{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNexonPartnerUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.NexonPartnerUpdate{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNotifyAccountFlagsChange)] = d.getEventEmitter(func() events.Event {
		return &events.NotifyAccountFlagsChange{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel)] = d.getEventEmitter(func() events.Event {
		return &events.OtherJoinedChannel{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCOtherLeftChannel)] = d.getEventEmitter(func() events.Event {
		return &events.OtherLeftChannel{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPartyLeaderWatchGamePrompt)] = d.getEventEmitter(func() events.Event {
		return &events.PartyLeaderWatchGamePrompt{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfo)] = d.getEventEmitter(func() events.Event {
		return &events.PlayerInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientPlaytestStatus)] = d.getEventEmitter(func() events.Event {
		return &events.PlaytestStatus{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPopup)] = d.getEventEmitter(func() events.Event {
		return &events.Popup{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProcessFantasyScheduledEvent)] = d.getEventEmitter(func() events.Event {
		return &events.ProcessFantasyScheduledEvent{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.ProfileCardUpdated{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientQuestProgressUpdated)] = d.getEventEmitter(func() events.Event {
		return &events.QuestProgressUpdated{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReadyUpStatus)] = d.getEventEmitter(func() events.Event {
		return &events.ReadyUpStatus{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCResetMapLocations)] = d.getEventEmitter(func() events.Event {
		return &events.ResetMapLocations{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSteamDatagramTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SteamDatagramTicket{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTeamInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteGCImmediateResponseToInviter{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteReceived{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter)] = d.getEventEmitter(func() events.Event {
		return &events.TeamInviteResponseReceived{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo)] = d.getEventEmitter(func() events.Event {
		return &events.TeamsInfo{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTipNotification)] = d.getEventEmitter(func() events.Event {
		return &events.TipNotification{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTournamentItemDrop)] = d.getEventEmitter(func() events.Event {
		return &events.TournamentItemDrop{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTrophyAwarded)] = d.getEventEmitter(func() events.Event {
		return &events.TrophyAwarded{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCUpdateClientClippy)] = d.getEventEmitter(func() events.Event {
		return &events.UpdateClientClippy{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientWageringUpdate)] = d.getEventEmitter(func() events.Event {
		return &events.WageringUpdate{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCWatchDownloadedReplay)] = d.getEventEmitter(func() events.Event {
		return &events.WatchDownloadedReplay{}
	})
	d.handlers[uint32(dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCWatchGame)] = d.getEventEmitter(func() events.Event {
		return &events.WatchGame{}
	})
}
