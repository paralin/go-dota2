package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol/base_gcmessages"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_fantasy"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_team"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_tournament"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	dm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// msgSenderOverrides overrides the heuristic-generated sender parties for each message
// Most of the MsgSenderNone messages are not GC<->Client messages.
var msgSenderOverrides = map[dm.EDOTAGCMsg]MsgSender{
	dm.EDOTAGCMsg_k_EMsgGCLiveScoreboardUpdate: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCGeneralResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCTeamData: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGC_TournamentItemEvent:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGC_TournamentItemEventResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCFeaturedItems: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCSetFeaturedItems: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCMatchHistoryList: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGetRecentMatches: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyList: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame: MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCStopFindingMatch:   MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCReadyUp:            MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCLeaverDetected:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCLeaverDetectedResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRequestSaveGamesServer: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTALiveLeagueGameUpdate: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCTeamMemberProfileRequest: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRequestPlayerResources:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCRequestPlayerResourcesResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCBanStatusRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCBanStatusResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCGenerateDiretidePrizeList:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGenerateDiretidePrizeListResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCPassportDataRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPassportDataResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateInfoRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueCreateInfoResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueInviteInfoRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueInviteInfoResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueFriendJoinListRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCFantasyLeagueFriendJoinListResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRequestBatchPlayerResources:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCRequestBatchPlayerResourcesResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientLeaguePredictionsResponse: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCLeaguePredictions:         MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeaveResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetailsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobbyResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptionsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientEconNotification_Job: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgTeamFanfare:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgResponseTeamFanfare: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTAAwardEventPoints: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTAFrostivusTimeElapsed: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCDev_GrantWarKill: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCLobbyList:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCLobbyListResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyLeagueRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyLeagueResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCConnectedPlayers: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCLeagueAdminList:  MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCLeagueAdminState: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCChatMessage: MsgSenderClient,

	// Hand-written lobby code
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate:      MsgSenderNone,
}

// msgMethodNameOverrides overrides the generated client method names.
var msgMethodNameOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGameAutographReward:               "AutographReward",
	dm.EDOTAGCMsg_k_EMsgDestroyLobbyRequest:               "DestroyLobby",
	dm.EDOTAGCMsg_k_EMsgGCReadyUp:                         "SendReadyUp",
	dm.EDOTAGCMsg_k_EMsgGCRequestInternatinalTicketEmail:  "RequestInternationalTicketEmail",
	dm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame:              "AbandonLobby",
	dm.EDOTAGCMsg_k_EMsgGCDOTAClearNotifySuccessfulReport: "ClearSuccessfulReportNotification",
	dm.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList:           "ListTrophies",
	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:        "CreateFantasyTeam",
	dm.EDOTAGCMsg_k_EMsgGCEditFantasyTeamRequest:          "EditFantasyTeam",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick:         "KickPrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote:      "PromotePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote:       "DemotePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInfoRequest:  "RequestPrivateChatInfo",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite:       "InvitePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgCastMatchVote:                     "CastMatchVote",
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick:               "KickLobbyMember",
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam:       "KickLobbyMemberFromTeam",
	dm.EDOTAGCMsg_k_EMsgGCBotGameCreate:                   "CreateBotGame",
}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[dm.EDOTAGCMsg]dm.EDOTAGCMsg{
	// Example:
	// dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack: dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse,
	dm.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest: 0,
	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest:    dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo,
}

// msgProtoTypeOverrides overrides the GC message to proto mapping.
var msgProtoTypeOverrides = map[dm.EDOTAGCMsg]proto.Message{
	dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo: &dota_gcmessages_client_team.CMsgDOTATeamInfo{},

	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamRequest:  &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamCreateRequest{},
	dm.EDOTAGCMsg_k_EMsgGCCreateFantasyTeamResponse: &dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamCreateResponse{},

	dm.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection:         &dota_gcmessages_client.CMsgDOTACompendiumSelection{},
	dm.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse: &dota_gcmessages_client.CMsgDOTACompendiumSelectionResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard:        &dota_gcmessages_client.CMsgPlayerConductScorecard{},
	dm.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest: &dota_gcmessages_client.CMsgPlayerConductScorecardRequest{},

	dm.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse: &dota_gcmessages_client.CMsgEventGoals{},

	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse:           &dota_gcmessages_client_tournament.CMsgWeekendTourneyOpts{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave:                  &dota_gcmessages_client_tournament.CMsgWeekendTourneyLeave{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse: &dota_gcmessages_client_tournament.CMsgDOTAWeekendTourneyPlayerStats{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats:         &dota_gcmessages_client_tournament.CMsgDOTAWeekendTourneyPlayerStatsRequest{},
	dm.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule:                  &dota_gcmessages_client_tournament.CMsgRequestWeekendTourneySchedule{},

	dm.EDOTAGCMsg_k_EMsgGCRequestInternatinalTicketEmail: &dota_gcmessages_client.CMsgRequestInternationalTicket{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader:     &dota_gcmessages_client_match_management.CMsgDOTASetGroupLeader{},
	dm.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites: &dota_gcmessages_client_match_management.CMsgDOTACancelGroupInvites{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen: &dota_gcmessages_client_match_management.CMsgDOTASetGroupOpenStatus{},

	dm.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite:        &dota_gcmessages_client_match_management.CMsgDOTAGroupMergeInvite{},
	dm.EDOTAGCMsg_k_EMsgClientToGCMergePartyResponse:      &dota_gcmessages_client_match_management.CMsgDOTAGroupMergeResponse{},
	dm.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply: &dota_gcmessages_client_match_management.CMsgDOTAGroupMergeReply{},

	dm.EDOTAGCMsg_k_EMsgClientToGCPingData: &base_gcmessages.CMsgClientPingData{},

	dm.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest: &dota_gcmessages_client.CMsgClientToGCGetEventGoals{},

	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: &dota_gcmessages_client_team.CMsgDOTAMyTeamInfoRequest{},

	dm.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory: &dota_gcmessages_common.CMsgBattleCupVictory{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions: &dota_gcmessages_client_match_management.CMsgPartyBuilderOptions{},

	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: &dota_gcmessages_client_chat.CMsgDOTAOtherJoinedChatChannel{},
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   &dota_gcmessages_client_chat.CMsgDOTAOtherLeftChatChannel{},
}

var msgArgAsParameterOverrides = map[dm.EDOTAGCMsg]bool{
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: true,
}
