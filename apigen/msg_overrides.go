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
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails:          true,
	dm.EDOTAGCMsg_k_EMsgClientToGCMonsterHunterClaimReward: true,
}

var msgEventNameOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
	dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter: "TeamInviteResponseReceived",
	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel:             "PlayerJoinedChannel",
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel:               "PlayerLeftChannel",
}

// msgDocOverrides contains handwritten godoc comments for request methods.
// An entry replaces the entire generated comment for the method.
var msgDocOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGCStartFindingMatch: `StartFindingMatch enters the matchmaking queue with the given options:
match mode, map, team desirability, and lobby type.

The GC confirms entry through the response and later reports the found match
through the ready-up flow; see SendReadyUp.`,
	dm.EDOTAGCMsg_k_EMsgGCStopFindingMatch: `StopFindingMatch leaves the matchmaking queue.`,
	dm.EDOTAGCMsg_k_EMsgGCReadyUp: `SendReadyUp accepts (or declines) an incoming match. The GC assigns teams
and broadcasts the game setup state once all players accept.

A player who fails to ready up in time causes the match search to restart.`,
	dm.EDOTAGCMsg_k_EMsgGCAbandonCurrentGame: `AbandonLobby abandons the current practice lobby or live game.`,
	dm.EDOTAGCMsg_k_EMsgDestroyLobbyRequest: `DestroyLobby destroys the practice lobby you lead. Only the lobby leader
can destroy it.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyLeave:  `LeaveLobby leaves the current practice lobby without destroying it.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyLaunch: `LaunchLobby starts the match for the practice lobby you lead.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyJoin:   `JoinLobby joins a practice lobby by ID, optionally with a pass key.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbyKick: `KickLobbyMember kicks a member from your practice lobby. Lobby leaders
only.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: `SetLobbyDetails updates the details of the current practice lobby: game
mode, map, server region, pass key, spectator policy, and cheat/bot settings.`,
	dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetCoach: `SetLobbyCoach requests the coach slot for your team in the lobby.`,
	dm.EDOTAGCMsg_k_EMsgGCFlipLobbyTeams: `FlipLobbyTeams swaps every member of the lobby between the Radiant and
Dire teams.`,
	dm.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby: `SendBalancedShuffleLobby shuffles the lobby members between teams while
keeping the teams balanced by MMR.`,
	dm.EDOTAGCMsg_k_EMsgGCChatMessage: `SendChatMessage sends a chat message to a joined chat channel or lobby.
Use SendChannelMessage for plain text messages.`,
	dm.EDOTAGCMsg_k_EMsgGCJoinChatChannel: `JoinChatChannel joins a chat channel by name and type. The response lists
the channel members; other joins and leaves arrive as events.`,
	dm.EDOTAGCMsg_k_EMsgGCLeaveChatChannel:       `LeaveChatChannel leaves a joined chat channel.`,
	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyLeader: `SetPartyLeader makes the given party member the party leader.`,
	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyOpen:   `SetPartyOpen opens or closes the party so that friends can join freely.`,
	dm.EDOTAGCMsg_k_EMsgClientToGCCancelPartyInvites: `CancelPartyInvites cancels all outstanding party invites sent by your
account.`,
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge: `AckPartyReadyCheck acknowledges a party-wide ready check on behalf of your
account.`,
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest: `SendPartyReadyCheck starts a ready check across the whole party.`,
	dm.EDOTAGCMsg_k_EMsgClientToGCSetPartyBuilderOptions: `SetPartyBuilderOptions configures how the party enters matchmaking, such
as the selected lanes or roles in role-queue modes.`,
	dm.EDOTAGCMsg_k_EMsgClientToGCGetProfileCard: `GetProfileCard requests the profile card of an account.

The card is not returned directly; watch for a ProfileCardUpdated event after
sending this request.`,
	dm.EDOTAGCMsg_k_EMsgDOTAGetPlayerMatchHistory: `GetPlayerMatchHistory returns a page of an account's match history with
per-match results. Pass start_at_match_id from the previous response to page
through older matches.`,
}

// msgEventDocOverrides contains handwritten godoc comments for events.
// An entry replaces the entire generated comment for the event.
var msgEventDocOverrides = map[dm.EDOTAGCMsg]string{
	dm.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel: `OtherJoinedChannel is emitted when another member joins a chat channel you
are in.`,
	dm.EDOTAGCMsg_k_EMsgGCOtherLeftChannel: `OtherLeftChannel is emitted when another member leaves a chat channel you
are in.`,
	dm.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded: `LobbyMVPAwarded is emitted when the MVP votes for a finished lobby match
are awarded.`,
	dm.EDOTAGCMsg_k_EMsgGCToClientCommendNotification: `CommendNotification is emitted when another player commends your account.`,
	dm.EDOTAGCMsg_k_EMsgGCClientSuspended: `ClientSuspended is emitted when the GC suspends new sessions, usually for
a scheduled update or maintenance.`,
	dm.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue: `KickedFromMatchmakingQueue is emitted when the GC removes the account from
a matchmaking queue it was waiting in.`,
}
