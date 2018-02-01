package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_chat"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_fantasy"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_guild"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_match_management"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_team"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_tournament"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_client_watch"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_common"
	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// ArcanaVotesUpdate event.
// MessageID: k_EMsgGCToClientArcanaVotesUpdate
type ArcanaVotesUpdate struct {
	dota_gcmessages_client.CMsgGCToClientArcanaVotesUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ArcanaVotesUpdate) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientArcanaVotesUpdate
}

// GetEventBody returns the event body.
func (e *ArcanaVotesUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientArcanaVotesUpdate
}

// GetEventName returns the event name.
func (e *ArcanaVotesUpdate) GetEventName() string {
	return "ArcanaVotesUpdate"
}

// BalancedShuffleLobby event.
// MessageID: k_EMsgGCBalancedShuffleLobby
type BalancedShuffleLobby struct {
	dota_gcmessages_client.CMsgBalancedShuffleLobby
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BalancedShuffleLobby) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCBalancedShuffleLobby
}

// GetEventBody returns the event body.
func (e *BalancedShuffleLobby) GetEventBody() proto.Message {
	return &e.CMsgBalancedShuffleLobby
}

// GetEventName returns the event name.
func (e *BalancedShuffleLobby) GetEventName() string {
	return "BalancedShuffleLobby"
}

// BattlePassRollupListRequest event.
// MessageID: k_EMsgGCToClientBattlePassRollupListRequest
type BattlePassRollupListRequest struct {
	dota_gcmessages_client.CMsgGCToClientBattlePassRollupListRequest
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BattlePassRollupListRequest) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupListRequest
}

// GetEventBody returns the event body.
func (e *BattlePassRollupListRequest) GetEventBody() proto.Message {
	return &e.CMsgGCToClientBattlePassRollupListRequest
}

// GetEventName returns the event name.
func (e *BattlePassRollupListRequest) GetEventName() string {
	return "BattlePassRollupListRequest"
}

// BattlePassRollupRequest event.
// MessageID: k_EMsgGCToClientBattlePassRollupRequest
type BattlePassRollupRequest struct {
	dota_gcmessages_client.CMsgGCToClientBattlePassRollupRequest
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BattlePassRollupRequest) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupRequest
}

// GetEventBody returns the event body.
func (e *BattlePassRollupRequest) GetEventBody() proto.Message {
	return &e.CMsgGCToClientBattlePassRollupRequest
}

// GetEventName returns the event name.
func (e *BattlePassRollupRequest) GetEventName() string {
	return "BattlePassRollupRequest"
}

// BroadcastNotification event.
// MessageID: k_EMsgGCBroadcastNotification
type BroadcastNotification struct {
	dota_gcmessages_common.CMsgDOTABroadcastNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BroadcastNotification) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCBroadcastNotification
}

// GetEventBody returns the event body.
func (e *BroadcastNotification) GetEventBody() proto.Message {
	return &e.CMsgDOTABroadcastNotification
}

// GetEventName returns the event name.
func (e *BroadcastNotification) GetEventName() string {
	return "BroadcastNotification"
}

// ChangeTeamSub event.
// MessageID: k_EMsgGCChangeTeamSub
type ChangeTeamSub struct {
	dota_gcmessages_client_team.CMsgDOTAChangeTeamSub
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChangeTeamSub) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCChangeTeamSub
}

// GetEventBody returns the event body.
func (e *ChangeTeamSub) GetEventBody() proto.Message {
	return &e.CMsgDOTAChangeTeamSub
}

// GetEventName returns the event name.
func (e *ChangeTeamSub) GetEventName() string {
	return "ChangeTeamSub"
}

// ChatRegionsEnabled event.
// MessageID: k_EMsgGCToClientChatRegionsEnabled
type ChatRegionsEnabled struct {
	dota_gcmessages_client_chat.CMsgDOTAChatRegionsEnabled
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatRegionsEnabled) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled
}

// GetEventBody returns the event body.
func (e *ChatRegionsEnabled) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatRegionsEnabled
}

// GetEventName returns the event name.
func (e *ChatRegionsEnabled) GetEventName() string {
	return "ChatRegionsEnabled"
}

// ClearPracticeLobbyTeam event.
// MessageID: k_EMsgGCClearPracticeLobbyTeam
type ClearPracticeLobbyTeam struct {
	dota_gcmessages_client_match_management.CMsgClearPracticeLobbyTeam
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClearPracticeLobbyTeam) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClearPracticeLobbyTeam
}

// GetEventBody returns the event body.
func (e *ClearPracticeLobbyTeam) GetEventBody() proto.Message {
	return &e.CMsgClearPracticeLobbyTeam
}

// GetEventName returns the event name.
func (e *ClearPracticeLobbyTeam) GetEventName() string {
	return "ClearPracticeLobbyTeam"
}

// ClientIgnoredUser event.
// MessageID: k_EMsgGCClientIgnoredUser
type ClientIgnoredUser struct {
	dota_gcmessages_client_chat.CMsgDOTAClientIgnoredUser
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClientIgnoredUser) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClientIgnoredUser
}

// GetEventBody returns the event body.
func (e *ClientIgnoredUser) GetEventBody() proto.Message {
	return &e.CMsgDOTAClientIgnoredUser
}

// GetEventName returns the event name.
func (e *ClientIgnoredUser) GetEventName() string {
	return "ClientIgnoredUser"
}

// ClientSuspended event.
// MessageID: k_EMsgGCClientSuspended
type ClientSuspended struct {
	dota_gcmessages_client.CMsgClientSuspended
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClientSuspended) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCClientSuspended
}

// GetEventBody returns the event body.
func (e *ClientSuspended) GetEventBody() proto.Message {
	return &e.CMsgClientSuspended
}

// GetEventName returns the event name.
func (e *ClientSuspended) GetEventName() string {
	return "ClientSuspended"
}

// CommendNotification event.
// MessageID: k_EMsgGCToClientCommendNotification
type CommendNotification struct {
	dota_gcmessages_client.CMsgGCToClientCommendNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CommendNotification) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientCommendNotification
}

// GetEventBody returns the event body.
func (e *CommendNotification) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCommendNotification
}

// GetEventName returns the event name.
func (e *CommendNotification) GetEventName() string {
	return "CommendNotification"
}

// CompendiumDataChanged event.
// MessageID: k_EMsgGCCompendiumDataChanged
type CompendiumDataChanged struct {
	dota_gcmessages_client.CMsgDOTACompendiumData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CompendiumDataChanged) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCCompendiumDataChanged
}

// GetEventBody returns the event body.
func (e *CompendiumDataChanged) GetEventBody() proto.Message {
	return &e.CMsgDOTACompendiumData
}

// GetEventName returns the event name.
func (e *CompendiumDataChanged) GetEventName() string {
	return "CompendiumDataChanged"
}

// DOTALiveLeagueGameUpdate event.
// MessageID: k_EMsgDOTALiveLeagueGameUpdate
type DOTALiveLeagueGameUpdate struct {
	dota_gcmessages_client.CMsgDOTALiveLeagueGameUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DOTALiveLeagueGameUpdate) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTALiveLeagueGameUpdate
}

// GetEventBody returns the event body.
func (e *DOTALiveLeagueGameUpdate) GetEventBody() proto.Message {
	return &e.CMsgDOTALiveLeagueGameUpdate
}

// GetEventName returns the event name.
func (e *DOTALiveLeagueGameUpdate) GetEventName() string {
	return "DOTALiveLeagueGameUpdate"
}

// DOTAWeekendTourneySchedule event.
// MessageID: k_EMsgDOTAWeekendTourneySchedule
type DOTAWeekendTourneySchedule struct {
	dota_gcmessages_client_tournament.CMsgWeekendTourneySchedule
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DOTAWeekendTourneySchedule) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule
}

// GetEventBody returns the event body.
func (e *DOTAWeekendTourneySchedule) GetEventBody() proto.Message {
	return &e.CMsgWeekendTourneySchedule
}

// GetEventName returns the event name.
func (e *DOTAWeekendTourneySchedule) GetEventName() string {
	return "DOTAWeekendTourneySchedule"
}

// EmoticonData event.
// MessageID: k_EMsgGCToClientEmoticonData
type EmoticonData struct {
	dota_gcmessages_client.CMsgGCToClientEmoticonData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *EmoticonData) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientEmoticonData
}

// GetEventBody returns the event body.
func (e *EmoticonData) GetEventBody() proto.Message {
	return &e.CMsgGCToClientEmoticonData
}

// GetEventName returns the event name.
func (e *EmoticonData) GetEventName() string {
	return "EmoticonData"
}

// EventStatusChanged event.
// MessageID: k_EMsgGCToClientEventStatusChanged
type EventStatusChanged struct {
	dota_gcmessages_client.CMsgGCToClientEventStatusChanged
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *EventStatusChanged) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientEventStatusChanged
}

// GetEventBody returns the event body.
func (e *EventStatusChanged) GetEventBody() proto.Message {
	return &e.CMsgGCToClientEventStatusChanged
}

// GetEventName returns the event name.
func (e *EventStatusChanged) GetEventName() string {
	return "EventStatusChanged"
}

// FantasyFinalPlayerStats event.
// MessageID: k_EMsgGCFantasyFinalPlayerStats
type FantasyFinalPlayerStats struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyFinalPlayerStats) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats
}

// GetEventBody returns the event body.
func (e *FantasyFinalPlayerStats) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats
}

// GetEventName returns the event name.
func (e *FantasyFinalPlayerStats) GetEventName() string {
	return "FantasyFinalPlayerStats"
}

// FantasyLeagueDraftStatus event.
// MessageID: k_EMsgGCFantasyLeagueDraftStatus
type FantasyLeagueDraftStatus struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueDraftStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyLeagueDraftStatus) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftStatus
}

// GetEventBody returns the event body.
func (e *FantasyLeagueDraftStatus) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyLeagueDraftStatus
}

// GetEventName returns the event name.
func (e *FantasyLeagueDraftStatus) GetEventName() string {
	return "FantasyLeagueDraftStatus"
}

// FantasyLeagueInfo event.
// MessageID: k_EMsgGCFantasyLeagueInfo
type FantasyLeagueInfo struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyLeagueInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyLeagueInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyLeagueInfo
}

// GetEventBody returns the event body.
func (e *FantasyLeagueInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyLeagueInfo
}

// GetEventName returns the event name.
func (e *FantasyLeagueInfo) GetEventName() string {
	return "FantasyLeagueInfo"
}

// FantasyMessageAdd event.
// MessageID: k_EMsgGCFantasyMessageAdd
type FantasyMessageAdd struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyMessageAdd
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyMessageAdd) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyMessageAdd
}

// GetEventBody returns the event body.
func (e *FantasyMessageAdd) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyMessageAdd
}

// GetEventName returns the event name.
func (e *FantasyMessageAdd) GetEventName() string {
	return "FantasyMessageAdd"
}

// FantasyRemoveOwner event.
// MessageID: k_EMsgGCFantasyRemoveOwner
type FantasyRemoveOwner struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyRemoveOwner
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyRemoveOwner) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyRemoveOwner
}

// GetEventBody returns the event body.
func (e *FantasyRemoveOwner) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyRemoveOwner
}

// GetEventName returns the event name.
func (e *FantasyRemoveOwner) GetEventName() string {
	return "FantasyRemoveOwner"
}

// FantasyTeamInfo event.
// MessageID: k_EMsgGCFantasyTeamInfo
type FantasyTeamInfo struct {
	dota_gcmessages_client_fantasy.CMsgDOTAFantasyTeamInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyTeamInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFantasyTeamInfo
}

// GetEventBody returns the event body.
func (e *FantasyTeamInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyTeamInfo
}

// GetEventName returns the event name.
func (e *FantasyTeamInfo) GetEventName() string {
	return "FantasyTeamInfo"
}

// FeaturedItems event.
// MessageID: k_EMsgGCFeaturedItems
type FeaturedItems struct {
	dota_gcmessages_client.CMsgDOTAFeaturedItems
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FeaturedItems) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCFeaturedItems
}

// GetEventBody returns the event body.
func (e *FeaturedItems) GetEventBody() proto.Message {
	return &e.CMsgDOTAFeaturedItems
}

// GetEventName returns the event name.
func (e *FeaturedItems) GetEventName() string {
	return "FeaturedItems"
}

// GuildInviteData event.
// MessageID: k_EMsgGCGuildInviteData
type GuildInviteData struct {
	dota_gcmessages_client_guild.CMsgDOTAGuildInviteData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildInviteData) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildInviteData
}

// GetEventBody returns the event body.
func (e *GuildInviteData) GetEventBody() proto.Message {
	return &e.CMsgDOTAGuildInviteData
}

// GetEventName returns the event name.
func (e *GuildInviteData) GetEventName() string {
	return "GuildInviteData"
}

// GuildUpdateMessage event.
// MessageID: k_EMsgGCGuildUpdateMessage
type GuildUpdateMessage struct {
	dota_gcmessages_client_guild.CMsgDOTAGuildUpdateMessage
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildUpdateMessage) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCGuildUpdateMessage
}

// GetEventBody returns the event body.
func (e *GuildUpdateMessage) GetEventBody() proto.Message {
	return &e.CMsgDOTAGuildUpdateMessage
}

// GetEventName returns the event name.
func (e *GuildUpdateMessage) GetEventName() string {
	return "GuildUpdateMessage"
}

// HallOfFame event.
// MessageID: k_EMsgGCHallOfFame
type HallOfFame struct {
	dota_gcmessages_client.CMsgDOTAHallOfFame
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *HallOfFame) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCHallOfFame
}

// GetEventBody returns the event body.
func (e *HallOfFame) GetEventBody() proto.Message {
	return &e.CMsgDOTAHallOfFame
}

// GetEventName returns the event name.
func (e *HallOfFame) GetEventName() string {
	return "HallOfFame"
}

// HeroStatueCreateResult event.
// MessageID: k_EMsgGCToClientHeroStatueCreateResult
type HeroStatueCreateResult struct {
	dota_gcmessages_client.CMsgGCToClientHeroStatueCreateResult
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *HeroStatueCreateResult) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientHeroStatueCreateResult
}

// GetEventBody returns the event body.
func (e *HeroStatueCreateResult) GetEventBody() proto.Message {
	return &e.CMsgGCToClientHeroStatueCreateResult
}

// GetEventName returns the event name.
func (e *HeroStatueCreateResult) GetEventName() string {
	return "HeroStatueCreateResult"
}

// KickedFromMatchmakingQueue event.
// MessageID: k_EMsgGCKickedFromMatchmakingQueue
type KickedFromMatchmakingQueue struct {
	dota_gcmessages_client.CMsgDOTAKickedFromMatchmakingQueue
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *KickedFromMatchmakingQueue) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue
}

// GetEventBody returns the event body.
func (e *KickedFromMatchmakingQueue) GetEventBody() proto.Message {
	return &e.CMsgDOTAKickedFromMatchmakingQueue
}

// GetEventName returns the event name.
func (e *KickedFromMatchmakingQueue) GetEventName() string {
	return "KickedFromMatchmakingQueue"
}

// LastHitChallengeHighScorePost event.
// MessageID: k_EMsgGCLastHitChallengeHighScorePost
type LastHitChallengeHighScorePost struct {
	dota_gcmessages_client.CMsgDOTALastHitChallengeHighScorePost
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LastHitChallengeHighScorePost) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLastHitChallengeHighScorePost
}

// GetEventBody returns the event body.
func (e *LastHitChallengeHighScorePost) GetEventBody() proto.Message {
	return &e.CMsgDOTALastHitChallengeHighScorePost
}

// GetEventName returns the event name.
func (e *LastHitChallengeHighScorePost) GetEventName() string {
	return "LastHitChallengeHighScorePost"
}

// LeagueAdminList event.
// MessageID: k_EMsgGCLeagueAdminList
type LeagueAdminList struct {
	dota_gcmessages_common.CMsgLeagueAdminList
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LeagueAdminList) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeagueAdminList
}

// GetEventBody returns the event body.
func (e *LeagueAdminList) GetEventBody() proto.Message {
	return &e.CMsgLeagueAdminList
}

// GetEventName returns the event name.
func (e *LeagueAdminList) GetEventName() string {
	return "LeagueAdminList"
}

// LeagueAdminState event.
// MessageID: k_EMsgGCLeagueAdminState
type LeagueAdminState struct {
	dota_gcmessages_client.CMsgGCLeagueAdminState
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LeagueAdminState) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLeagueAdminState
}

// GetEventBody returns the event body.
func (e *LeagueAdminState) GetEventBody() proto.Message {
	return &e.CMsgGCLeagueAdminState
}

// GetEventName returns the event name.
func (e *LeagueAdminState) GetEventName() string {
	return "LeagueAdminState"
}

// LobbyMVPAwarded event.
// MessageID: k_EMsgGCToClientLobbyMVPAwarded
type LobbyMVPAwarded struct {
	dota_gcmessages_client.CMsgDOTALobbyMVPAwarded
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyMVPAwarded) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded
}

// GetEventBody returns the event body.
func (e *LobbyMVPAwarded) GetEventBody() proto.Message {
	return &e.CMsgDOTALobbyMVPAwarded
}

// GetEventName returns the event name.
func (e *LobbyMVPAwarded) GetEventName() string {
	return "LobbyMVPAwarded"
}

// LobbyMVPNotifyRecipient event.
// MessageID: k_EMsgGCToClientLobbyMVPNotifyRecipient
type LobbyMVPNotifyRecipient struct {
	dota_gcmessages_client.CMsgDOTALobbyMVPNotifyRecipient
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyMVPNotifyRecipient) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient
}

// GetEventBody returns the event body.
func (e *LobbyMVPNotifyRecipient) GetEventBody() proto.Message {
	return &e.CMsgDOTALobbyMVPNotifyRecipient
}

// GetEventName returns the event name.
func (e *LobbyMVPNotifyRecipient) GetEventName() string {
	return "LobbyMVPNotifyRecipient"
}

// LobbyUpdateBroadcastChannelInfo event.
// MessageID: k_EMsgGCLobbyUpdateBroadcastChannelInfo
type LobbyUpdateBroadcastChannelInfo struct {
	dota_gcmessages_client.CMsgGCLobbyUpdateBroadcastChannelInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyUpdateBroadcastChannelInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCLobbyUpdateBroadcastChannelInfo
}

// GetEventBody returns the event body.
func (e *LobbyUpdateBroadcastChannelInfo) GetEventBody() proto.Message {
	return &e.CMsgGCLobbyUpdateBroadcastChannelInfo
}

// GetEventName returns the event name.
func (e *LobbyUpdateBroadcastChannelInfo) GetEventName() string {
	return "LobbyUpdateBroadcastChannelInfo"
}

// MakeOffering event.
// MessageID: k_EMsgGCMakeOffering
type MakeOffering struct {
	dota_gcmessages_client.CMsgMakeOffering
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MakeOffering) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCMakeOffering
}

// GetEventBody returns the event body.
func (e *MakeOffering) GetEventBody() proto.Message {
	return &e.CMsgMakeOffering
}

// GetEventName returns the event name.
func (e *MakeOffering) GetEventName() string {
	return "MakeOffering"
}

// MatchGroupsVersion event.
// MessageID: k_EMsgGCToClientMatchGroupsVersion
type MatchGroupsVersion struct {
	dota_gcmessages_common.CMsgGCToClientMatchGroupsVersion
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MatchGroupsVersion) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMatchGroupsVersion
}

// GetEventBody returns the event body.
func (e *MatchGroupsVersion) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchGroupsVersion
}

// GetEventName returns the event name.
func (e *MatchGroupsVersion) GetEventName() string {
	return "MatchGroupsVersion"
}

// MatchSignedOut event.
// MessageID: k_EMsgGCToClientMatchSignedOut
type MatchSignedOut struct {
	dota_gcmessages_client.CMsgGCToClientMatchSignedOut
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MatchSignedOut) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMatchSignedOut
}

// GetEventBody returns the event body.
func (e *MatchSignedOut) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchSignedOut
}

// GetEventName returns the event name.
func (e *MatchSignedOut) GetEventName() string {
	return "MatchSignedOut"
}

// MergeGroupInviteReply event.
// MessageID: k_EMsgGCToClientMergeGroupInviteReply
type MergeGroupInviteReply struct {
	dota_gcmessages_client_match_management.CMsgDOTAGroupMergeReply
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MergeGroupInviteReply) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply
}

// GetEventBody returns the event body.
func (e *MergeGroupInviteReply) GetEventBody() proto.Message {
	return &e.CMsgDOTAGroupMergeReply
}

// GetEventName returns the event name.
func (e *MergeGroupInviteReply) GetEventName() string {
	return "MergeGroupInviteReply"
}

// MergePartyResponseReply event.
// MessageID: k_EMsgGCToClientMergePartyResponseReply
type MergePartyResponseReply struct {
	dota_gcmessages_client_match_management.CMsgDOTAGroupMergeReply
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MergePartyResponseReply) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply
}

// GetEventBody returns the event body.
func (e *MergePartyResponseReply) GetEventBody() proto.Message {
	return &e.CMsgDOTAGroupMergeReply
}

// GetEventName returns the event name.
func (e *MergePartyResponseReply) GetEventName() string {
	return "MergePartyResponseReply"
}

// NewNotificationAdded event.
// MessageID: k_EMsgGCToClientNewNotificationAdded
type NewNotificationAdded struct {
	dota_gcmessages_client.CMsgGCNotificationsResponse_Notification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NewNotificationAdded) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientNewNotificationAdded
}

// GetEventBody returns the event body.
func (e *NewNotificationAdded) GetEventBody() proto.Message {
	return &e.CMsgGCNotificationsResponse_Notification
}

// GetEventName returns the event name.
func (e *NewNotificationAdded) GetEventName() string {
	return "NewNotificationAdded"
}

// NexonPartnerUpdate event.
// MessageID: k_EMsgGCNexonPartnerUpdate
type NexonPartnerUpdate struct {
	dota_gcmessages_client.CMsgNexonPartnerUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NexonPartnerUpdate) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNexonPartnerUpdate
}

// GetEventBody returns the event body.
func (e *NexonPartnerUpdate) GetEventBody() proto.Message {
	return &e.CMsgNexonPartnerUpdate
}

// GetEventName returns the event name.
func (e *NexonPartnerUpdate) GetEventName() string {
	return "NexonPartnerUpdate"
}

// NotifyAccountFlagsChange event.
// MessageID: k_EMsgGCNotifyAccountFlagsChange
type NotifyAccountFlagsChange struct {
	dota_gcmessages_client.CMsgDOTANotifyAccountFlagsChange
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NotifyAccountFlagsChange) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCNotifyAccountFlagsChange
}

// GetEventBody returns the event body.
func (e *NotifyAccountFlagsChange) GetEventBody() proto.Message {
	return &e.CMsgDOTANotifyAccountFlagsChange
}

// GetEventName returns the event name.
func (e *NotifyAccountFlagsChange) GetEventName() string {
	return "NotifyAccountFlagsChange"
}

// PartyLeaderWatchGamePrompt event.
// MessageID: k_EMsgGCPartyLeaderWatchGamePrompt
type PartyLeaderWatchGamePrompt struct {
	dota_gcmessages_client_watch.CMsgPartyLeaderWatchGamePrompt
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartyLeaderWatchGamePrompt) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPartyLeaderWatchGamePrompt
}

// GetEventBody returns the event body.
func (e *PartyLeaderWatchGamePrompt) GetEventBody() proto.Message {
	return &e.CMsgPartyLeaderWatchGamePrompt
}

// GetEventName returns the event name.
func (e *PartyLeaderWatchGamePrompt) GetEventName() string {
	return "PartyLeaderWatchGamePrompt"
}

// PlayerInfo event.
// MessageID: k_EMsgGCPlayerInfo
type PlayerInfo struct {
	dota_gcmessages_client_fantasy.CMsgGCPlayerInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPlayerInfo
}

// GetEventBody returns the event body.
func (e *PlayerInfo) GetEventBody() proto.Message {
	return &e.CMsgGCPlayerInfo
}

// GetEventName returns the event name.
func (e *PlayerInfo) GetEventName() string {
	return "PlayerInfo"
}

// PlayerJoinedChannel event.
// MessageID: k_EMsgGCOtherJoinedChannel
type PlayerJoinedChannel struct {
	dota_gcmessages_client_chat.CMsgDOTAOtherJoinedChatChannel
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerJoinedChannel) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel
}

// GetEventBody returns the event body.
func (e *PlayerJoinedChannel) GetEventBody() proto.Message {
	return &e.CMsgDOTAOtherJoinedChatChannel
}

// GetEventName returns the event name.
func (e *PlayerJoinedChannel) GetEventName() string {
	return "PlayerJoinedChannel"
}

// PlayerLeftChannel event.
// MessageID: k_EMsgGCOtherLeftChannel
type PlayerLeftChannel struct {
	dota_gcmessages_client_chat.CMsgDOTAOtherLeftChatChannel
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerLeftChannel) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCOtherLeftChannel
}

// GetEventBody returns the event body.
func (e *PlayerLeftChannel) GetEventBody() proto.Message {
	return &e.CMsgDOTAOtherLeftChatChannel
}

// GetEventName returns the event name.
func (e *PlayerLeftChannel) GetEventName() string {
	return "PlayerLeftChannel"
}

// PlaytestStatus event.
// MessageID: k_EMsgGCToClientPlaytestStatus
type PlaytestStatus struct {
	dota_gcmessages_client.CMsgGCToClientPlaytestStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlaytestStatus) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientPlaytestStatus
}

// GetEventBody returns the event body.
func (e *PlaytestStatus) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPlaytestStatus
}

// GetEventName returns the event name.
func (e *PlaytestStatus) GetEventName() string {
	return "PlaytestStatus"
}

// Popup event.
// MessageID: k_EMsgGCPopup
type Popup struct {
	dota_gcmessages_client.CMsgDOTAPopup
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *Popup) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCPopup
}

// GetEventBody returns the event body.
func (e *Popup) GetEventBody() proto.Message {
	return &e.CMsgDOTAPopup
}

// GetEventName returns the event name.
func (e *Popup) GetEventName() string {
	return "Popup"
}

// ProcessFantasyScheduledEvent event.
// MessageID: k_EMsgGCProcessFantasyScheduledEvent
type ProcessFantasyScheduledEvent struct {
	dota_gcmessages_common.CMsgDOTAProcessFantasyScheduledEvent
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ProcessFantasyScheduledEvent) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCProcessFantasyScheduledEvent
}

// GetEventBody returns the event body.
func (e *ProcessFantasyScheduledEvent) GetEventBody() proto.Message {
	return &e.CMsgDOTAProcessFantasyScheduledEvent
}

// GetEventName returns the event name.
func (e *ProcessFantasyScheduledEvent) GetEventName() string {
	return "ProcessFantasyScheduledEvent"
}

// ProfileCardUpdated event.
// MessageID: k_EMsgGCToClientProfileCardUpdated
type ProfileCardUpdated struct {
	dota_gcmessages_common.CMsgDOTAProfileCard
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ProfileCardUpdated) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated
}

// GetEventBody returns the event body.
func (e *ProfileCardUpdated) GetEventBody() proto.Message {
	return &e.CMsgDOTAProfileCard
}

// GetEventName returns the event name.
func (e *ProfileCardUpdated) GetEventName() string {
	return "ProfileCardUpdated"
}

// QuestProgressUpdated event.
// MessageID: k_EMsgGCToClientQuestProgressUpdated
type QuestProgressUpdated struct {
	dota_gcmessages_client.CMsgGCToClientQuestProgressUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *QuestProgressUpdated) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientQuestProgressUpdated
}

// GetEventBody returns the event body.
func (e *QuestProgressUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientQuestProgressUpdated
}

// GetEventName returns the event name.
func (e *QuestProgressUpdated) GetEventName() string {
	return "QuestProgressUpdated"
}

// ReadyUpStatus event.
// MessageID: k_EMsgGCReadyUpStatus
type ReadyUpStatus struct {
	dota_gcmessages_client_match_management.CMsgReadyUpStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ReadyUpStatus) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCReadyUpStatus
}

// GetEventBody returns the event body.
func (e *ReadyUpStatus) GetEventBody() proto.Message {
	return &e.CMsgReadyUpStatus
}

// GetEventName returns the event name.
func (e *ReadyUpStatus) GetEventName() string {
	return "ReadyUpStatus"
}

// ResetMapLocations event.
// MessageID: k_EMsgGCResetMapLocations
type ResetMapLocations struct {
	dota_gcmessages_client.CMsgResetMapLocations
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ResetMapLocations) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCResetMapLocations
}

// GetEventBody returns the event body.
func (e *ResetMapLocations) GetEventBody() proto.Message {
	return &e.CMsgResetMapLocations
}

// GetEventName returns the event name.
func (e *ResetMapLocations) GetEventName() string {
	return "ResetMapLocations"
}

// SteamDatagramTicket event.
// MessageID: k_EMsgGCToClientSteamDatagramTicket
type SteamDatagramTicket struct {
	dota_gcmessages_client_match_management.CMsgGCToClientSteamDatagramTicket
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *SteamDatagramTicket) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientSteamDatagramTicket
}

// GetEventBody returns the event body.
func (e *SteamDatagramTicket) GetEventBody() proto.Message {
	return &e.CMsgGCToClientSteamDatagramTicket
}

// GetEventName returns the event name.
func (e *SteamDatagramTicket) GetEventName() string {
	return "SteamDatagramTicket"
}

// TeamInfo event.
// MessageID: k_EMsgGCToClientTeamInfo
type TeamInfo struct {
	dota_gcmessages_client_team.CMsgDOTATeamInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTeamInfo
}

// GetEventBody returns the event body.
func (e *TeamInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTATeamInfo
}

// GetEventName returns the event name.
func (e *TeamInfo) GetEventName() string {
	return "TeamInfo"
}

// TeamInviteGCImmediateResponseToInviter event.
// MessageID: k_EMsgGCTeamInvite_GCImmediateResponseToInviter
type TeamInviteGCImmediateResponseToInviter struct {
	dota_gcmessages_client_team.CMsgDOTATeamInvite_GCImmediateResponseToInviter
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteGCImmediateResponseToInviter) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter
}

// GetEventBody returns the event body.
func (e *TeamInviteGCImmediateResponseToInviter) GetEventBody() proto.Message {
	return &e.CMsgDOTATeamInvite_GCImmediateResponseToInviter
}

// GetEventName returns the event name.
func (e *TeamInviteGCImmediateResponseToInviter) GetEventName() string {
	return "TeamInviteGCImmediateResponseToInviter"
}

// TeamInviteReceived event.
// MessageID: k_EMsgGCTeamInvite_GCRequestToInvitee
type TeamInviteReceived struct {
	dota_gcmessages_client_team.CMsgDOTATeamInvite_GCRequestToInvitee
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteReceived) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee
}

// GetEventBody returns the event body.
func (e *TeamInviteReceived) GetEventBody() proto.Message {
	return &e.CMsgDOTATeamInvite_GCRequestToInvitee
}

// GetEventName returns the event name.
func (e *TeamInviteReceived) GetEventName() string {
	return "TeamInviteReceived"
}

// TeamInviteResponseReceived event.
// MessageID: k_EMsgGCTeamInvite_GCResponseToInviter
type TeamInviteResponseReceived struct {
	dota_gcmessages_client_team.CMsgDOTATeamInvite_GCResponseToInviter
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteResponseReceived) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter
}

// GetEventBody returns the event body.
func (e *TeamInviteResponseReceived) GetEventBody() proto.Message {
	return &e.CMsgDOTATeamInvite_GCResponseToInviter
}

// GetEventName returns the event name.
func (e *TeamInviteResponseReceived) GetEventName() string {
	return "TeamInviteResponseReceived"
}

// TeamsInfo event.
// MessageID: k_EMsgGCToClientTeamsInfo
type TeamsInfo struct {
	dota_gcmessages_client_team.CMsgDOTATeamsInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamsInfo) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo
}

// GetEventBody returns the event body.
func (e *TeamsInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTATeamsInfo
}

// GetEventName returns the event name.
func (e *TeamsInfo) GetEventName() string {
	return "TeamsInfo"
}

// TipNotification event.
// MessageID: k_EMsgGCToClientTipNotification
type TipNotification struct {
	dota_gcmessages_client.CMsgGCToClientTipNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TipNotification) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTipNotification
}

// GetEventBody returns the event body.
func (e *TipNotification) GetEventBody() proto.Message {
	return &e.CMsgGCToClientTipNotification
}

// GetEventName returns the event name.
func (e *TipNotification) GetEventName() string {
	return "TipNotification"
}

// TournamentItemDrop event.
// MessageID: k_EMsgGCToClientTournamentItemDrop
type TournamentItemDrop struct {
	dota_gcmessages_client.CMsgGCToClientTournamentItemDrop
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TournamentItemDrop) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTournamentItemDrop
}

// GetEventBody returns the event body.
func (e *TournamentItemDrop) GetEventBody() proto.Message {
	return &e.CMsgGCToClientTournamentItemDrop
}

// GetEventName returns the event name.
func (e *TournamentItemDrop) GetEventName() string {
	return "TournamentItemDrop"
}

// TrophyAwarded event.
// MessageID: k_EMsgGCToClientTrophyAwarded
type TrophyAwarded struct {
	dota_gcmessages_client.CMsgGCToClientTrophyAwarded
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TrophyAwarded) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientTrophyAwarded
}

// GetEventBody returns the event body.
func (e *TrophyAwarded) GetEventBody() proto.Message {
	return &e.CMsgGCToClientTrophyAwarded
}

// GetEventName returns the event name.
func (e *TrophyAwarded) GetEventName() string {
	return "TrophyAwarded"
}

// UpdateClientClippy event.
// MessageID: k_EMsgGCUpdateClientClippy
type UpdateClientClippy struct {
	dota_gcmessages_client.CMsgGCToClientUpdateClientClippy
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *UpdateClientClippy) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCUpdateClientClippy
}

// GetEventBody returns the event body.
func (e *UpdateClientClippy) GetEventBody() proto.Message {
	return &e.CMsgGCToClientUpdateClientClippy
}

// GetEventName returns the event name.
func (e *UpdateClientClippy) GetEventName() string {
	return "UpdateClientClippy"
}

// WageringUpdate event.
// MessageID: k_EMsgGCToClientWageringUpdate
type WageringUpdate struct {
	dota_gcmessages_client.CMsgGCToClientWageringUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *WageringUpdate) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCToClientWageringUpdate
}

// GetEventBody returns the event body.
func (e *WageringUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientWageringUpdate
}

// GetEventName returns the event name.
func (e *WageringUpdate) GetEventName() string {
	return "WageringUpdate"
}

// WatchDownloadedReplay event.
// MessageID: k_EMsgGCWatchDownloadedReplay
type WatchDownloadedReplay struct {
	dota_gcmessages_client.CMsgGCWatchDownloadedReplay
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *WatchDownloadedReplay) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCWatchDownloadedReplay
}

// GetEventBody returns the event body.
func (e *WatchDownloadedReplay) GetEventBody() proto.Message {
	return &e.CMsgGCWatchDownloadedReplay
}

// GetEventName returns the event name.
func (e *WatchDownloadedReplay) GetEventName() string {
	return "WatchDownloadedReplay"
}

// WatchGame event.
// MessageID: k_EMsgGCWatchGame
type WatchGame struct {
	dota_gcmessages_client_watch.CMsgWatchGame
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *WatchGame) GetDotaEventMsgID() dota_gcmessages_msgid.EDOTAGCMsg {
	return dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCWatchGame
}

// GetEventBody returns the event body.
func (e *WatchGame) GetEventBody() proto.Message {
	return &e.CMsgWatchGame
}

// GetEventName returns the event name.
func (e *WatchGame) GetEventName() string {
	return "WatchGame"
}
