package main

import (
	"github.com/golang/protobuf/proto"
	dm "github.com/paralin/go-dota2/protocol"
)

// msgSenderOverrides overrides the heuristic-generated sender parties for each message
// Most of the MsgSenderNone messages are not GC<->Client messages.
var msgSenderOverrides = map[dm.EDOTAGCMsg]MsgSender{
	dm.EDOTAGCMsg_k_EMsgGCLiveScoreboardUpdate:          MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPlayerFailedToConnect:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToLANServerRelayConnect:     MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToRelayConnect:              MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCGCToRelayConnectresponse:      MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCOverworldDevResetPath: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCSurvivorsTelemetry:    MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGC_TournamentItemEvent:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGC_TournamentItemEventResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyList: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCInitialQuestionnaireResponse: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame: MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCStopFindingMatch:   MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCReadyUp:            MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCLeaverDetected:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCLeaverDetectedResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCRequestSaveGamesServer:   MsgSenderNone, // server
	dm.EDOTAGCMsg_k_EMsgGCRequestSaveGames:         MsgSenderNone, // client
	dm.EDOTAGCMsg_k_EMsgGCRequestSaveGamesResponse: MsgSenderNone, // gc

	dm.EDOTAGCMsg_k_EMsgGCBanStatusRequest:  MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCBanStatusResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeaveResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCSetSpectatorLobbyDetailsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCCreateSpectatorLobbyResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptionsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgTeamFanfare:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgResponseTeamFanfare: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgDOTAAwardEventPoints: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCLobbyList:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCLobbyListResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCConnectedPlayers: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCLeagueAdminList: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCChatMessage: MsgSenderClient,

	// Hand-written lobby code
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetTeamSlot: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyCreate:      MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCClearPracticeLobbyTeam:   MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCFantasyLivePlayerStats: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStats:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardStatsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientProfileCardStatsUpdated: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientAutomatedTournamentStateChange: MsgSenderNone,

	// todo: determine who sends the CMsgWeekendTourneyOpts and what the response is
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOpts:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:                  MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter: MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:           MsgSenderGC,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC:          MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee:          MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter:          MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgDOTALeagueInfoListAdminsRequest: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgDOTALeagueInfoListAdminsReponse: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby: MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgGCWatchGame:            MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCtoGCRequestRecalibrationCheck:      MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgGCtoGCAssociatedExploiterAccountInfo: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeed: MsgSenderClient,

	dm.EDOTAGCMsg_k_EMsgGCToClientPostGameItemAwardNotification: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswer:         MsgSenderClient,
	dm.EDOTAGCMsg_k_EMsgClientToGCSubmitDraftTriviaMatchAnswerResponse: MsgSenderGC,

	dm.EDOTAGCMsg_k_EMsgWebapiDPCSeasonResults: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgClientToGCGiveTip: MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgLobbyGauntletProgress:         MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgClientToGCApplyGauntletTicket: MsgSenderNone,
	dm.EDOTAGCMsg_k_EMsgLobbyAdditionalAccountData:    MsgSenderNone,

	dm.EDOTAGCMsg_k_EMsgGCToClientFantasyCraftingGetDataUpdated: MsgSenderNone,
}

// msgMethodNameOverrides overrides the generated client method names.
var msgMethodNameOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGameAutographReward:              "AutographReward",
	dm.EDOTAGCMsg_k_EMsgDestroyLobbyRequest:              "DestroyLobby",
	dm.EDOTAGCMsg_k_EMsgGCReadyUp:                        "SendReadyUp",
	dm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame:             "AbandonLobby",
	dm.EDOTAGCMsg_k_EMsgClientToGCGetTrophyList:          "ListTrophies",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatKick:        "KickPrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatPromote:     "PromotePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatDemote:      "DemotePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgClientToGCPrivateChatInvite:      "InvitePrivateChatMember",
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick:              "KickLobbyMember",
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKickFromTeam:      "KickLobbyMemberFromTeam",
	dm.EDOTAGCMsg_k_EMsgGCBotGameCreate:                  "CreateBotGame",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:         "InvitePlayerToTeam",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC: "RespondToTeamInvite",
	dm.EDOTAGCMsg_k_EMsgClientsRejoinChatChannels:        "RejoinAllChatChannels",
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest:           "SendPartyReadyCheck",
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge:       "AckPartyReadyCheck",
}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[dm.EDOTAGCMsg]dm.EDOTAGCMsg{
	// Example:
	// dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPack: dm.EDOTAGCMsg_k_EMsgClientToGCCreatePlayerCardPackResponse,
	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo,

	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviterToGC:         dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter,
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_InviteeResponseToGC: dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInvitee,
	dm.EDOTAGCMsg_k_EMsgGCWatchGame:                      dm.EDOTAGCMsg_k_EMsgGCWatchGameResponse,

	dm.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby:         0,
	dm.EDOTAGCMsg_k_EMsgGCNotificationsMarkReadRequest: 0,
}

// msgProtoTypeOverrides overrides the GC message to proto mapping.
var msgProtoTypeOverrides = map[dm.EDOTAGCMsg]proto.Message{
	dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo: &dm.CMsgDOTATeamInfo{},

	dm.EDOTAGCMsg_k_EMsgGCCompendiumSetSelection:         &dm.CMsgDOTACompendiumSelection{},
	dm.EDOTAGCMsg_k_EMsgGCCompendiumSetSelectionResponse: &dm.CMsgDOTACompendiumSelectionResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecard:        &dm.CMsgPlayerConductScorecard{},
	dm.EDOTAGCMsg_k_EMsgClientToGCLatestConductScorecardRequest: &dm.CMsgPlayerConductScorecardRequest{},

	dm.EDOTAGCMsg_k_EMsgClientToGCEventGoalsResponse: &dm.CMsgEventGoals{},

	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyOptsResponse:           &dm.CMsgWeekendTourneyOpts{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyLeave:                  &dm.CMsgWeekendTourneyLeave{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStatsResponse: &dm.CMsgDOTAWeekendTourneyPlayerStats{},
	dm.EDOTAGCMsg_k_EMsgClientToGCWeekendTourneyGetPlayerStats:         &dm.CMsgDOTAWeekendTourneyPlayerStatsRequest{},
	dm.EDOTAGCMsg_k_EMsgDOTAGetWeekendTourneySchedule:                  &dm.CMsgRequestWeekendTourneySchedule{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader:     &dm.CMsgDOTASetGroupLeader{},
	dm.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites: &dm.CMsgDOTACancelGroupInvites{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen: &dm.CMsgDOTASetGroupOpenStatus{},

	dm.EDOTAGCMsg_k_EMsgClientToGCMergePartyInvite:        &dm.CMsgDOTAGroupMergeInvite{},
	dm.EDOTAGCMsg_k_EMsgClientToGCMergePartyResponse:      &dm.CMsgDOTAGroupMergeResponse{},
	dm.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply: &dm.CMsgDOTAGroupMergeReply{},
	dm.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply:   &dm.CMsgDOTAGroupMergeReply{},

	dm.EDOTAGCMsg_k_EMsgClientToGCPingData: &dm.CMsgClientPingData{},

	dm.EDOTAGCMsg_k_EMsgClientToGCEventGoalsRequest: &dm.CMsgClientToGCGetEventGoals{},

	dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest: &dm.CMsgDOTAMyTeamInfoRequest{},

	dm.EDOTAGCMsg_k_EMsgLobbyBattleCupVictory: &dm.CMsgBattleCupVictory{},

	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions: &dm.CMsgPartyBuilderOptions{},

	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: &dm.CMsgDOTAOtherJoinedChatChannel{},
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:   &dm.CMsgDOTAOtherLeftChatChannel{},

	dm.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated:   &dm.CMsgDOTAProfileCard{},
	dm.EDOTAGCMsg_k_EMsgGCToClientNotificationsUpdated: &dm.CMsgGCNotificationsResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCardResponse: &dm.CMsgDOTAProfileCard{},

	dm.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled: &dm.CMsgDOTAChatRegionsEnabled{},

	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileTicketsResponse: &dm.CMsgDOTAProfileTickets{},

	// Experimental

	dm.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo: &dm.CMsgDOTATeamsInfo{},

	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded: &dm.CMsgDOTALobbyMVPAwarded{},

	dm.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummary:         &dm.CMsgEventTipsSummaryRequest{},
	dm.EDOTAGCMsg_k_EMsgClientToGCRequestEventTipsSummaryResponse: &dm.CMsgEventTipsSummaryResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeed:         &dm.CMsgSocialFeedRequest{},
	dm.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedResponse: &dm.CMsgSocialFeedResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedComments:         &dm.CMsgSocialFeedCommentsRequest{},
	dm.EDOTAGCMsg_k_EMsgClientToGCRequestSocialFeedCommentsResponse: &dm.CMsgSocialFeedCommentsResponse{},

	dm.EDOTAGCMsg_k_EMsgClientToGCRequestGuildFeed:      &dm.CMsgClientToGCGuildFeedRequest{},
	dm.EDOTAGCMsg_k_EMsgGCToClientUnderDraftGoldUpdated: &dm.CMsgGCToClientGuildUnderDraftGoldUpdated{},

	dm.EDOTAGCMsg_k_EMsgGCToClientClaimSwagResponse: &dm.CMsgClientToGCClaimSwagResponse{},
}

var msgArgAsParameterOverrides = map[dm.EDOTAGCMsg]bool{
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: true,
}

var msgEventNameOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter: "TeamInviteResponseReceived",
	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel:             "PlayerJoinedChannel",
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:               "PlayerLeftChannel",
}
