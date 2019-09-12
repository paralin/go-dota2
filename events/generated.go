package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

// ArcanaVotesUpdate event.
// MessageID: k_EMsgGCToClientArcanaVotesUpdate
type ArcanaVotesUpdate struct {
	protocol.CMsgGCToClientArcanaVotesUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ArcanaVotesUpdate) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientArcanaVotesUpdate
}

// GetEventBody returns the event body.
func (e *ArcanaVotesUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientArcanaVotesUpdate
}

// GetEventName returns the event name.
func (e *ArcanaVotesUpdate) GetEventName() string {
	return "ArcanaVotesUpdate"
}

// BattlePassRollupListRequest event.
// MessageID: k_EMsgGCToClientBattlePassRollupListRequest
type BattlePassRollupListRequest struct {
	protocol.CMsgGCToClientBattlePassRollupListRequest
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BattlePassRollupListRequest) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupListRequest
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
	protocol.CMsgGCToClientBattlePassRollupRequest
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BattlePassRollupRequest) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientBattlePassRollupRequest
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
	protocol.CMsgDOTABroadcastNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BroadcastNotification) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCBroadcastNotification
}

// GetEventBody returns the event body.
func (e *BroadcastNotification) GetEventBody() proto.Message {
	return &e.CMsgDOTABroadcastNotification
}

// GetEventName returns the event name.
func (e *BroadcastNotification) GetEventName() string {
	return "BroadcastNotification"
}

// CavernCrawlMapPathCompleted event.
// MessageID: k_EMsgGCToClientCavernCrawlMapPathCompleted
type CavernCrawlMapPathCompleted struct {
	protocol.CMsgGCToClientCavernCrawlMapPathCompleted
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CavernCrawlMapPathCompleted) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapPathCompleted
}

// GetEventBody returns the event body.
func (e *CavernCrawlMapPathCompleted) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCavernCrawlMapPathCompleted
}

// GetEventName returns the event name.
func (e *CavernCrawlMapPathCompleted) GetEventName() string {
	return "CavernCrawlMapPathCompleted"
}

// CavernCrawlMapUpdated event.
// MessageID: k_EMsgGCToClientCavernCrawlMapUpdated
type CavernCrawlMapUpdated struct {
	protocol.CMsgGCToClientCavernCrawlMapUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CavernCrawlMapUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCavernCrawlMapUpdated
}

// GetEventBody returns the event body.
func (e *CavernCrawlMapUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCavernCrawlMapUpdated
}

// GetEventName returns the event name.
func (e *CavernCrawlMapUpdated) GetEventName() string {
	return "CavernCrawlMapUpdated"
}

// ChatRegionsEnabled event.
// MessageID: k_EMsgGCToClientChatRegionsEnabled
type ChatRegionsEnabled struct {
	protocol.CMsgDOTAChatRegionsEnabled
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatRegionsEnabled) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientChatRegionsEnabled
}

// GetEventBody returns the event body.
func (e *ChatRegionsEnabled) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatRegionsEnabled
}

// GetEventName returns the event name.
func (e *ChatRegionsEnabled) GetEventName() string {
	return "ChatRegionsEnabled"
}

// ClaimEventActionUsingItemCompleted event.
// MessageID: k_EMsgGCToClientClaimEventActionUsingItemCompleted
type ClaimEventActionUsingItemCompleted struct {
	protocol.CMsgGCToClientClaimEventActionUsingItemCompleted
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClaimEventActionUsingItemCompleted) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientClaimEventActionUsingItemCompleted
}

// GetEventBody returns the event body.
func (e *ClaimEventActionUsingItemCompleted) GetEventBody() proto.Message {
	return &e.CMsgGCToClientClaimEventActionUsingItemCompleted
}

// GetEventName returns the event name.
func (e *ClaimEventActionUsingItemCompleted) GetEventName() string {
	return "ClaimEventActionUsingItemCompleted"
}

// ClearPracticeLobbyTeam event.
// MessageID: k_EMsgGCClearPracticeLobbyTeam
type ClearPracticeLobbyTeam struct {
	protocol.CMsgClearPracticeLobbyTeam
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClearPracticeLobbyTeam) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCClearPracticeLobbyTeam
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
	protocol.CMsgDOTAClientIgnoredUser
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClientIgnoredUser) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCClientIgnoredUser
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
	protocol.CMsgClientSuspended
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ClientSuspended) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCClientSuspended
}

// GetEventBody returns the event body.
func (e *ClientSuspended) GetEventBody() proto.Message {
	return &e.CMsgClientSuspended
}

// GetEventName returns the event name.
func (e *ClientSuspended) GetEventName() string {
	return "ClientSuspended"
}

// CoachTeammateRatingsChanged event.
// MessageID: k_EMsgGCToClientCoachTeammateRatingsChanged
type CoachTeammateRatingsChanged struct {
	protocol.CMsgGCToClientCoachTeammateRatingsChanged
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CoachTeammateRatingsChanged) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCoachTeammateRatingsChanged
}

// GetEventBody returns the event body.
func (e *CoachTeammateRatingsChanged) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCoachTeammateRatingsChanged
}

// GetEventName returns the event name.
func (e *CoachTeammateRatingsChanged) GetEventName() string {
	return "CoachTeammateRatingsChanged"
}

// CommendNotification event.
// MessageID: k_EMsgGCToClientCommendNotification
type CommendNotification struct {
	protocol.CMsgGCToClientCommendNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CommendNotification) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCommendNotification
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
	protocol.CMsgDOTACompendiumData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CompendiumDataChanged) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCCompendiumDataChanged
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
	protocol.CMsgDOTALiveLeagueGameUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DOTALiveLeagueGameUpdate) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgDOTALiveLeagueGameUpdate
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
	protocol.CMsgWeekendTourneySchedule
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *DOTAWeekendTourneySchedule) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgDOTAWeekendTourneySchedule
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
	protocol.CMsgGCToClientEmoticonData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *EmoticonData) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientEmoticonData
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
	protocol.CMsgGCToClientEventStatusChanged
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *EventStatusChanged) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientEventStatusChanged
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
	protocol.CMsgDOTAFantasyPlayerHisoricalStatsResponse_PlayerStats
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyFinalPlayerStats) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats
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
	protocol.CMsgDOTAFantasyLeagueDraftStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyLeagueDraftStatus) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyLeagueDraftStatus
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
	protocol.CMsgDOTAFantasyLeagueInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyLeagueInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyLeagueInfo
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
	protocol.CMsgDOTAFantasyMessageAdd
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyMessageAdd) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyMessageAdd
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
	protocol.CMsgDOTAFantasyRemoveOwner
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyRemoveOwner) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyRemoveOwner
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
	protocol.CMsgDOTAFantasyTeamInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyTeamInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyTeamInfo
}

// GetEventBody returns the event body.
func (e *FantasyTeamInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyTeamInfo
}

// GetEventName returns the event name.
func (e *FantasyTeamInfo) GetEventName() string {
	return "FantasyTeamInfo"
}

// GuildInviteData event.
// MessageID: k_EMsgGCGuildInviteData
type GuildInviteData struct {
	protocol.CMsgDOTAGuildInviteData
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildInviteData) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCGuildInviteData
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
	protocol.CMsgDOTAGuildUpdateMessage
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildUpdateMessage) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCGuildUpdateMessage
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
	protocol.CMsgDOTAHallOfFame
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *HallOfFame) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCHallOfFame
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
	protocol.CMsgGCToClientHeroStatueCreateResult
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *HeroStatueCreateResult) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientHeroStatueCreateResult
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
	protocol.CMsgDOTAKickedFromMatchmakingQueue
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *KickedFromMatchmakingQueue) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCKickedFromMatchmakingQueue
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
	protocol.CMsgDOTALastHitChallengeHighScorePost
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LastHitChallengeHighScorePost) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCLastHitChallengeHighScorePost
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
	protocol.CMsgLeagueAdminList
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LeagueAdminList) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCLeagueAdminList
}

// GetEventBody returns the event body.
func (e *LeagueAdminList) GetEventBody() proto.Message {
	return &e.CMsgLeagueAdminList
}

// GetEventName returns the event name.
func (e *LeagueAdminList) GetEventName() string {
	return "LeagueAdminList"
}

// LobbyMVPAwarded event.
// MessageID: k_EMsgGCToClientLobbyMVPAwarded
type LobbyMVPAwarded struct {
	protocol.CMsgDOTALobbyMVPAwarded
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyMVPAwarded) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPAwarded
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
	protocol.CMsgDOTALobbyMVPNotifyRecipient
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyMVPNotifyRecipient) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientLobbyMVPNotifyRecipient
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
	protocol.CMsgGCLobbyUpdateBroadcastChannelInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *LobbyUpdateBroadcastChannelInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCLobbyUpdateBroadcastChannelInfo
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
	protocol.CMsgMakeOffering
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MakeOffering) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCMakeOffering
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
	protocol.CMsgGCToClientMatchGroupsVersion
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MatchGroupsVersion) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientMatchGroupsVersion
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
	protocol.CMsgGCToClientMatchSignedOut
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MatchSignedOut) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientMatchSignedOut
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
	protocol.CMsgDOTAGroupMergeReply
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MergeGroupInviteReply) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientMergeGroupInviteReply
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
	protocol.CMsgDOTAGroupMergeReply
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MergePartyResponseReply) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientMergePartyResponseReply
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
	protocol.CMsgGCNotificationsResponse_Notification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NewNotificationAdded) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientNewNotificationAdded
}

// GetEventBody returns the event body.
func (e *NewNotificationAdded) GetEventBody() proto.Message {
	return &e.CMsgGCNotificationsResponse_Notification
}

// GetEventName returns the event name.
func (e *NewNotificationAdded) GetEventName() string {
	return "NewNotificationAdded"
}

// NotifyAccountFlagsChange event.
// MessageID: k_EMsgGCNotifyAccountFlagsChange
type NotifyAccountFlagsChange struct {
	protocol.CMsgDOTANotifyAccountFlagsChange
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NotifyAccountFlagsChange) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCNotifyAccountFlagsChange
}

// GetEventBody returns the event body.
func (e *NotifyAccountFlagsChange) GetEventBody() proto.Message {
	return &e.CMsgDOTANotifyAccountFlagsChange
}

// GetEventName returns the event name.
func (e *NotifyAccountFlagsChange) GetEventName() string {
	return "NotifyAccountFlagsChange"
}

// PartyBeaconUpdate event.
// MessageID: k_EMsgGCToClientPartyBeaconUpdate
type PartyBeaconUpdate struct {
	protocol.CMsgGCToClientPartyBeaconUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartyBeaconUpdate) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPartyBeaconUpdate
}

// GetEventBody returns the event body.
func (e *PartyBeaconUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartyBeaconUpdate
}

// GetEventName returns the event name.
func (e *PartyBeaconUpdate) GetEventName() string {
	return "PartyBeaconUpdate"
}

// PartyLeaderWatchGamePrompt event.
// MessageID: k_EMsgGCPartyLeaderWatchGamePrompt
type PartyLeaderWatchGamePrompt struct {
	protocol.CMsgPartyLeaderWatchGamePrompt
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartyLeaderWatchGamePrompt) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCPartyLeaderWatchGamePrompt
}

// GetEventBody returns the event body.
func (e *PartyLeaderWatchGamePrompt) GetEventBody() proto.Message {
	return &e.CMsgPartyLeaderWatchGamePrompt
}

// GetEventName returns the event name.
func (e *PartyLeaderWatchGamePrompt) GetEventName() string {
	return "PartyLeaderWatchGamePrompt"
}

// PartySearchInvite event.
// MessageID: k_EMsgGCToClientPartySearchInvite
type PartySearchInvite struct {
	protocol.CMsgGCToClientPartySearchInvite
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartySearchInvite) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvite
}

// GetEventBody returns the event body.
func (e *PartySearchInvite) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartySearchInvite
}

// GetEventName returns the event name.
func (e *PartySearchInvite) GetEventName() string {
	return "PartySearchInvite"
}

// PartySearchInvites event.
// MessageID: k_EMsgGCToClientPartySearchInvites
type PartySearchInvites struct {
	protocol.CMsgGCToClientPartySearchInvites
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PartySearchInvites) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPartySearchInvites
}

// GetEventBody returns the event body.
func (e *PartySearchInvites) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPartySearchInvites
}

// GetEventName returns the event name.
func (e *PartySearchInvites) GetEventName() string {
	return "PartySearchInvites"
}

// PlayerBeaconState event.
// MessageID: k_EMsgGCToClientPlayerBeaconState
type PlayerBeaconState struct {
	protocol.CMsgGCToClientPlayerBeaconState
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerBeaconState) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPlayerBeaconState
}

// GetEventBody returns the event body.
func (e *PlayerBeaconState) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPlayerBeaconState
}

// GetEventName returns the event name.
func (e *PlayerBeaconState) GetEventName() string {
	return "PlayerBeaconState"
}

// PlayerInfo event.
// MessageID: k_EMsgGCPlayerInfo
type PlayerInfo struct {
	protocol.CMsgDOTAPlayerInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCPlayerInfo
}

// GetEventBody returns the event body.
func (e *PlayerInfo) GetEventBody() proto.Message {
	return &e.CMsgDOTAPlayerInfo
}

// GetEventName returns the event name.
func (e *PlayerInfo) GetEventName() string {
	return "PlayerInfo"
}

// PlayerJoinedChannel event.
// MessageID: k_EMsgGCOtherJoinedChannel
type PlayerJoinedChannel struct {
	protocol.CMsgDOTAOtherJoinedChatChannel
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerJoinedChannel) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCOtherJoinedChannel
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
	protocol.CMsgDOTAOtherLeftChatChannel
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlayerLeftChannel) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCOtherLeftChannel
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
	protocol.CMsgGCToClientPlaytestStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PlaytestStatus) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPlaytestStatus
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
	protocol.CMsgDOTAPopup
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *Popup) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCPopup
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
	protocol.CMsgDOTAProcessFantasyScheduledEvent
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ProcessFantasyScheduledEvent) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCProcessFantasyScheduledEvent
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
	protocol.CMsgDOTAProfileCard
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ProfileCardUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientProfileCardUpdated
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
	protocol.CMsgGCToClientQuestProgressUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *QuestProgressUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientQuestProgressUpdated
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
	protocol.CMsgReadyUpStatus
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ReadyUpStatus) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCReadyUpStatus
}

// GetEventBody returns the event body.
func (e *ReadyUpStatus) GetEventBody() proto.Message {
	return &e.CMsgReadyUpStatus
}

// GetEventName returns the event name.
func (e *ReadyUpStatus) GetEventName() string {
	return "ReadyUpStatus"
}

// RequestLaneSelection event.
// MessageID: k_EMsgGCToClientRequestLaneSelection
type RequestLaneSelection struct {
	protocol.CMsgGCToClientRequestLaneSelection
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *RequestLaneSelection) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientRequestLaneSelection
}

// GetEventBody returns the event body.
func (e *RequestLaneSelection) GetEventBody() proto.Message {
	return &e.CMsgGCToClientRequestLaneSelection
}

// GetEventName returns the event name.
func (e *RequestLaneSelection) GetEventName() string {
	return "RequestLaneSelection"
}

// ResetMapLocations event.
// MessageID: k_EMsgGCResetMapLocations
type ResetMapLocations struct {
	protocol.CMsgResetMapLocations
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ResetMapLocations) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCResetMapLocations
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
	protocol.CMsgGCToClientSteamDatagramTicket
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *SteamDatagramTicket) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientSteamDatagramTicket
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
	protocol.CMsgDOTATeamInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientTeamInfo
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
	protocol.CMsgDOTATeamInvite_GCImmediateResponseToInviter
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteGCImmediateResponseToInviter) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCImmediateResponseToInviter
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
	protocol.CMsgDOTATeamInvite_GCRequestToInvitee
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteReceived) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee
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
	protocol.CMsgDOTATeamInvite_GCResponseToInviter
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamInviteResponseReceived) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCTeamInvite_GCResponseToInviter
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
	protocol.CMsgDOTATeamsInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TeamsInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientTeamsInfo
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
	protocol.CMsgGCToClientTipNotification
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TipNotification) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientTipNotification
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
	protocol.CMsgGCToClientTournamentItemDrop
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TournamentItemDrop) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientTournamentItemDrop
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
	protocol.CMsgGCToClientTrophyAwarded
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *TrophyAwarded) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientTrophyAwarded
}

// GetEventBody returns the event body.
func (e *TrophyAwarded) GetEventBody() proto.Message {
	return &e.CMsgGCToClientTrophyAwarded
}

// GetEventName returns the event name.
func (e *TrophyAwarded) GetEventName() string {
	return "TrophyAwarded"
}

// WageringUpdate event.
// MessageID: k_EMsgGCToClientWageringUpdate
type WageringUpdate struct {
	protocol.CMsgGCToClientWageringUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *WageringUpdate) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientWageringUpdate
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
	protocol.CMsgGCWatchDownloadedReplay
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *WatchDownloadedReplay) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCWatchDownloadedReplay
}

// GetEventBody returns the event body.
func (e *WatchDownloadedReplay) GetEventBody() proto.Message {
	return &e.CMsgGCWatchDownloadedReplay
}

// GetEventName returns the event name.
func (e *WatchDownloadedReplay) GetEventName() string {
	return "WatchDownloadedReplay"
}
