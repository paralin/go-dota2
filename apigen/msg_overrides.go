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
	dm.EDOTAGCMsg_k_EMsgGCLiveScoreboardUpdate:      MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPlayerReports:             MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPlayerFailedToConnect:     MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToLANServerRelayConnect: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToRelayConnect:          MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToRelayConnectresponse:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCSuggestTeamMatchmaking:    MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCUpdateClippy:              MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCConsumeFantasyTicket:        MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCConsumeFantasyTicketFailure: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCPlayerHeroesFavoritesAdd:    MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPlayerHeroesFavoritesRemove: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRewardDiretidePrizes: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCRewardTutorialPrizes: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRequestGuildData: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGuildData:        MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCNotInGuildData:   MsgSenderNone,

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

	dm.EDOTAGCMsg_k_EMsgGCFantasyLivePlayerStats: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCFantasyMatch:             MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPCBangTimedRewardMessage: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgCastMatchVoteResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStats:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStatsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientProfileCardStatsUpdated: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientAutomatedTournamentStateChange: MsgSenderNone,

	// todo: determine who sends the CMsgWeekendTourneyOpts and what the response is
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOpts:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient: MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded:         MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:                  MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter: MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:           MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC:          MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee:          MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter:          MsgSenderGC,
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
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:          "InvitePlayerToTeam",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC:  "RespondToTeamInvite",
	dm.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels:         "RejoinAllChatChannels",
}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[dm.EDOTAGCMsg]dm.EDOTAGCMsg{
	// Example:
	// dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack: dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse,
	dm.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest: 0,
	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest:    dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo,

	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:         dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC: dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee,
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
	dm.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply:   &dota_gcmessages_client_match_management.CMsgDOTAGroupMergeReply{},

	dm.EDOTAGCMsg_k_EMsgClientToGCPingData: &base_gcmessages.CMsgClientPingData{},

	dm.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest: &dota_gcmessages_client.CMsgClientToGCGetEventGoals{},

	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: &dota_gcmessages_client_team.CMsgDOTAMyTeamInfoRequest{},

	dm.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory: &dota_gcmessages_common.CMsgBattleCupVictory{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions: &dota_gcmessages_client_match_management.CMsgPartyBuilderOptions{},

	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: &dota_gcmessages_client_chat.CMsgDOTAOtherJoinedChatChannel{},
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   &dota_gcmessages_client_chat.CMsgDOTAOtherLeftChatChannel{},

	dm.EDOTAGCMsg_k_EMsgGCCompendiumDataChanged: &dota_gcmessages_client.CMsgDOTACompendiumData{},

	dm.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated:   &dota_gcmessages_common.CMsgDOTAProfileCard{},
	dm.EDOTAGCMsg_k_EMsgGCToClientNewNotificationAdded: &dota_gcmessages_client.CMsgGCNotificationsResponse_Notification{},

	dm.EDOTAGCMsg_k_EMsgGCUpdateClientClippy: &dota_gcmessages_client.CMsgGCToClientUpdateClientClippy{},

	dm.EDOTAGCMsg_k_EMsgRetrieveMatchVoteResponse: &dota_gcmessages_client.CMsgMatchVoteResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse: &dota_gcmessages_common.CMsgDOTAProfileCard{},

	dm.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled: &dota_gcmessages_client_chat.CMsgDOTAChatRegionsEnabled{},

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse: &dota_gcmessages_client.CMsgDOTAProfileTickets{},

	// Experimental
	dm.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats: &dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats{},

	dm.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo: &dota_gcmessages_client_team.CMsgDOTATeamsInfo{},

	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient: &dota_gcmessages_client.CMsgDOTALobbyMVPNotifyRecipient{},
	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded:         &dota_gcmessages_client.CMsgDOTALobbyMVPAwarded{},
}

var msgArgAsParameterOverrides = map[dm.EDOTAGCMsg]bool{
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: true,
}

var msgEventNameOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter: "TeamInviteResponseReceived",
}
