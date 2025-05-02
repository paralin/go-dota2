package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

// AccountGuildEventDataUpdated event.
// MessageID: k_EMsgGCToClientAccountGuildEventDataUpdated
type AccountGuildEventDataUpdated struct {
	protocol.CMsgGCToClientAccountGuildEventDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *AccountGuildEventDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientAccountGuildEventDataUpdated
}

// GetEventBody returns the event body.
func (e *AccountGuildEventDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientAccountGuildEventDataUpdated
}

// GetEventName returns the event name.
func (e *AccountGuildEventDataUpdated) GetEventName() string {
	return "AccountGuildEventDataUpdated"
}

// ActiveGuildChallengeUpdated event.
// MessageID: k_EMsgGCToClientActiveGuildChallengeUpdated
type ActiveGuildChallengeUpdated struct {
	protocol.CMsgGCToClientActiveGuildChallengeUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ActiveGuildChallengeUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildChallengeUpdated
}

// GetEventBody returns the event body.
func (e *ActiveGuildChallengeUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientActiveGuildChallengeUpdated
}

// GetEventName returns the event name.
func (e *ActiveGuildChallengeUpdated) GetEventName() string {
	return "ActiveGuildChallengeUpdated"
}

// ActiveGuildContractsUpdated event.
// MessageID: k_EMsgGCToClientActiveGuildContractsUpdated
type ActiveGuildContractsUpdated struct {
	protocol.CMsgGCToClientActiveGuildContractsUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ActiveGuildContractsUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientActiveGuildContractsUpdated
}

// GetEventBody returns the event body.
func (e *ActiveGuildContractsUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientActiveGuildContractsUpdated
}

// GetEventName returns the event name.
func (e *ActiveGuildContractsUpdated) GetEventName() string {
	return "ActiveGuildContractsUpdated"
}

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

// BingoUserDataUpdated event.
// MessageID: k_EMsgGCToClientBingoUserDataUpdated
type BingoUserDataUpdated struct {
	protocol.CMsgGCToClientBingoUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *BingoUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientBingoUserDataUpdated
}

// GetEventBody returns the event body.
func (e *BingoUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientBingoUserDataUpdated
}

// GetEventName returns the event name.
func (e *BingoUserDataUpdated) GetEventName() string {
	return "BingoUserDataUpdated"
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

// CandyShopUserDataUpdated event.
// MessageID: k_EMsgGCToClientCandyShopUserDataUpdated
type CandyShopUserDataUpdated struct {
	protocol.CMsgGCToClientCandyShopUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CandyShopUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCandyShopUserDataUpdated
}

// GetEventBody returns the event body.
func (e *CandyShopUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCandyShopUserDataUpdated
}

// GetEventName returns the event name.
func (e *CandyShopUserDataUpdated) GetEventName() string {
	return "CandyShopUserDataUpdated"
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

// ChatModeratorBan event.
// MessageID: k_EMsgGCChatModeratorBan
type ChatModeratorBan struct {
	protocol.CMsgDOTAChatModeratorBan
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ChatModeratorBan) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCChatModeratorBan
}

// GetEventBody returns the event body.
func (e *ChatModeratorBan) GetEventBody() proto.Message {
	return &e.CMsgDOTAChatModeratorBan
}

// GetEventName returns the event name.
func (e *ChatModeratorBan) GetEventName() string {
	return "ChatModeratorBan"
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

// CompendiumRemoveAllSelections event.
// MessageID: k_EMsgGCCompendiumRemoveAllSelections
type CompendiumRemoveAllSelections struct {
	protocol.CMsgDOTACompendiumRemoveAllSelections
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CompendiumRemoveAllSelections) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCCompendiumRemoveAllSelections
}

// GetEventBody returns the event body.
func (e *CompendiumRemoveAllSelections) GetEventBody() proto.Message {
	return &e.CMsgDOTACompendiumRemoveAllSelections
}

// GetEventName returns the event name.
func (e *CompendiumRemoveAllSelections) GetEventName() string {
	return "CompendiumRemoveAllSelections"
}

// CraftworksUserDataUpdated event.
// MessageID: k_EMsgGCToClientCraftworksUserDataUpdated
type CraftworksUserDataUpdated struct {
	protocol.CMsgGCToClientCraftworksUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *CraftworksUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientCraftworksUserDataUpdated
}

// GetEventBody returns the event body.
func (e *CraftworksUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientCraftworksUserDataUpdated
}

// GetEventName returns the event name.
func (e *CraftworksUserDataUpdated) GetEventName() string {
	return "CraftworksUserDataUpdated"
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

// FantasyFinalPlayerStats event.
// MessageID: k_EMsgGCFantasyFinalPlayerStats
type FantasyFinalPlayerStats struct {
	protocol.CMsgDOTAFantasyFinalPlayerStats
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FantasyFinalPlayerStats) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCFantasyFinalPlayerStats
}

// GetEventBody returns the event body.
func (e *FantasyFinalPlayerStats) GetEventBody() proto.Message {
	return &e.CMsgDOTAFantasyFinalPlayerStats
}

// GetEventName returns the event name.
func (e *FantasyFinalPlayerStats) GetEventName() string {
	return "FantasyFinalPlayerStats"
}

// FightingGameChallenge event.
// MessageID: k_EMsgGCToClientFightingGameChallenge
type FightingGameChallenge struct {
	protocol.CMsgGCToClientFightingGameChallenge
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FightingGameChallenge) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameChallenge
}

// GetEventBody returns the event body.
func (e *FightingGameChallenge) GetEventBody() proto.Message {
	return &e.CMsgGCToClientFightingGameChallenge
}

// GetEventName returns the event name.
func (e *FightingGameChallenge) GetEventName() string {
	return "FightingGameChallenge"
}

// FightingGameChallengeCanceled event.
// MessageID: k_EMsgGCToClientFightingGameChallengeCanceled
type FightingGameChallengeCanceled struct {
	protocol.CMsgGCToClientFightingGameChallengeCanceled
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FightingGameChallengeCanceled) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameChallengeCanceled
}

// GetEventBody returns the event body.
func (e *FightingGameChallengeCanceled) GetEventBody() proto.Message {
	return &e.CMsgGCToClientFightingGameChallengeCanceled
}

// GetEventName returns the event name.
func (e *FightingGameChallengeCanceled) GetEventName() string {
	return "FightingGameChallengeCanceled"
}

// FightingGameStartMatch event.
// MessageID: k_EMsgGCToClientFightingGameStartMatch
type FightingGameStartMatch struct {
	protocol.CMsgGCToClientFightingGameStartMatch
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *FightingGameStartMatch) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientFightingGameStartMatch
}

// GetEventBody returns the event body.
func (e *FightingGameStartMatch) GetEventBody() proto.Message {
	return &e.CMsgGCToClientFightingGameStartMatch
}

// GetEventName returns the event name.
func (e *FightingGameStartMatch) GetEventName() string {
	return "FightingGameStartMatch"
}

// GuildDataUpdated event.
// MessageID: k_EMsgGCToClientGuildDataUpdated
type GuildDataUpdated struct {
	protocol.CMsgGCToClientGuildDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientGuildDataUpdated
}

// GetEventBody returns the event body.
func (e *GuildDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGuildDataUpdated
}

// GetEventName returns the event name.
func (e *GuildDataUpdated) GetEventName() string {
	return "GuildDataUpdated"
}

// GuildFeedUpdated event.
// MessageID: k_EMsgGCToClientGuildFeedUpdated
type GuildFeedUpdated struct {
	protocol.CMsgGCToClientGuildFeedUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildFeedUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientGuildFeedUpdated
}

// GetEventBody returns the event body.
func (e *GuildFeedUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGuildFeedUpdated
}

// GetEventName returns the event name.
func (e *GuildFeedUpdated) GetEventName() string {
	return "GuildFeedUpdated"
}

// GuildMembersDataUpdated event.
// MessageID: k_EMsgGCToClientGuildMembersDataUpdated
type GuildMembersDataUpdated struct {
	protocol.CMsgGCToClientGuildMembersDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildMembersDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembersDataUpdated
}

// GetEventBody returns the event body.
func (e *GuildMembersDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGuildMembersDataUpdated
}

// GetEventName returns the event name.
func (e *GuildMembersDataUpdated) GetEventName() string {
	return "GuildMembersDataUpdated"
}

// GuildMembershipUpdated event.
// MessageID: k_EMsgGCToClientGuildMembershipUpdated
type GuildMembershipUpdated struct {
	protocol.CMsgGCToClientGuildMembershipUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *GuildMembershipUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientGuildMembershipUpdated
}

// GetEventBody returns the event body.
func (e *GuildMembershipUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGuildMembershipUpdated
}

// GetEventName returns the event name.
func (e *GuildMembershipUpdated) GetEventName() string {
	return "GuildMembershipUpdated"
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

// NotificationsUpdated event.
// MessageID: k_EMsgGCToClientNotificationsUpdated
type NotificationsUpdated struct {
	protocol.CMsgGCNotificationsResponse
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *NotificationsUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientNotificationsUpdated
}

// GetEventBody returns the event body.
func (e *NotificationsUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCNotificationsResponse
}

// GetEventName returns the event name.
func (e *NotificationsUpdated) GetEventName() string {
	return "NotificationsUpdated"
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

// OverwatchCasesAvailable event.
// MessageID: k_EMsgGCToClientOverwatchCasesAvailable
type OverwatchCasesAvailable struct {
	protocol.CMsgGCToClientOverwatchCasesAvailable
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *OverwatchCasesAvailable) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientOverwatchCasesAvailable
}

// GetEventBody returns the event body.
func (e *OverwatchCasesAvailable) GetEventBody() proto.Message {
	return &e.CMsgGCToClientOverwatchCasesAvailable
}

// GetEventName returns the event name.
func (e *OverwatchCasesAvailable) GetEventName() string {
	return "OverwatchCasesAvailable"
}

// OverworldUserDataUpdated event.
// MessageID: k_EMsgGCToClientOverworldUserDataUpdated
type OverworldUserDataUpdated struct {
	protocol.CMsgGCToClientOverworldUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *OverworldUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientOverworldUserDataUpdated
}

// GetEventBody returns the event body.
func (e *OverworldUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientOverworldUserDataUpdated
}

// GetEventName returns the event name.
func (e *OverworldUserDataUpdated) GetEventName() string {
	return "OverworldUserDataUpdated"
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

// PrivateCoachingSessionUpdated event.
// MessageID: k_EMsgGCToClientPrivateCoachingSessionUpdated
type PrivateCoachingSessionUpdated struct {
	protocol.CMsgGCToClientPrivateCoachingSessionUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *PrivateCoachingSessionUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientPrivateCoachingSessionUpdated
}

// GetEventBody returns the event body.
func (e *PrivateCoachingSessionUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPrivateCoachingSessionUpdated
}

// GetEventName returns the event name.
func (e *PrivateCoachingSessionUpdated) GetEventName() string {
	return "PrivateCoachingSessionUpdated"
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

// RankUpdate event.
// MessageID: k_EMsgGCToClientRankUpdate
type RankUpdate struct {
	protocol.CMsgGCToClientRankUpdate
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *RankUpdate) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientRankUpdate
}

// GetEventBody returns the event body.
func (e *RankUpdate) GetEventBody() proto.Message {
	return &e.CMsgGCToClientRankUpdate
}

// GetEventName returns the event name.
func (e *RankUpdate) GetEventName() string {
	return "RankUpdate"
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

// RequestMMInfo event.
// MessageID: k_EMsgGCToClientRequestMMInfo
type RequestMMInfo struct {
	protocol.CMsgGCToClientRequestMMInfo
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *RequestMMInfo) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientRequestMMInfo
}

// GetEventBody returns the event body.
func (e *RequestMMInfo) GetEventBody() proto.Message {
	return &e.CMsgGCToClientRequestMMInfo
}

// GetEventName returns the event name.
func (e *RequestMMInfo) GetEventName() string {
	return "RequestMMInfo"
}

// RoadToTIQuestDataUpdated event.
// MessageID: k_EMsgGCToClientRoadToTIQuestDataUpdated
type RoadToTIQuestDataUpdated struct {
	protocol.CMsgGCToClientRoadToTIQuestDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *RoadToTIQuestDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientRoadToTIQuestDataUpdated
}

// GetEventBody returns the event body.
func (e *RoadToTIQuestDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientRoadToTIQuestDataUpdated
}

// GetEventName returns the event name.
func (e *RoadToTIQuestDataUpdated) GetEventName() string {
	return "RoadToTIQuestDataUpdated"
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

// UnderDraftGoldUpdated event.
// MessageID: k_EMsgGCToClientUnderDraftGoldUpdated
type UnderDraftGoldUpdated struct {
	protocol.CMsgGCToClientGuildUnderDraftGoldUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *UnderDraftGoldUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientUnderDraftGoldUpdated
}

// GetEventBody returns the event body.
func (e *UnderDraftGoldUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGuildUnderDraftGoldUpdated
}

// GetEventName returns the event name.
func (e *UnderDraftGoldUpdated) GetEventName() string {
	return "UnderDraftGoldUpdated"
}

// VACReminder event.
// MessageID: k_EMsgGCToClientVACReminder
type VACReminder struct {
	protocol.CMsgGCToClientVACReminder
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *VACReminder) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientVACReminder
}

// GetEventBody returns the event body.
func (e *VACReminder) GetEventBody() proto.Message {
	return &e.CMsgGCToClientVACReminder
}

// GetEventName returns the event name.
func (e *VACReminder) GetEventName() string {
	return "VACReminder"
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
