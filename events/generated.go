package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/paralin/go-dota2/protocol"
)

// AccountGuildEventDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientAccountGuildEventDataUpdated (CMsgGCToClientAccountGuildEventDataUpdated).
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

// ActiveGuildChallengeUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientActiveGuildChallengeUpdated (CMsgGCToClientActiveGuildChallengeUpdated).
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

// ActiveGuildContractsUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientActiveGuildContractsUpdated (CMsgGCToClientActiveGuildContractsUpdated).
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

// ArcanaVotesUpdate is an event delivered by the GC.
//
// Message: k_EMsgGCToClientArcanaVotesUpdate (CMsgGCToClientArcanaVotesUpdate).
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

// BattlePassRollupListRequest is an event delivered by the GC.
//
// Message: k_EMsgGCToClientBattlePassRollupListRequest (CMsgGCToClientBattlePassRollupListRequest).
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

// BattlePassRollupRequest is an event delivered by the GC.
//
// Message: k_EMsgGCToClientBattlePassRollupRequest (CMsgGCToClientBattlePassRollupRequest).
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

// BingoUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientBingoUserDataUpdated (CMsgGCToClientBingoUserDataUpdated).
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

// BroadcastNotification is an event delivered by the GC.
//
// Message: k_EMsgGCBroadcastNotification (CMsgDOTABroadcastNotification).
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

// CandyShopUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCandyShopUserDataUpdated (CMsgGCToClientCandyShopUserDataUpdated).
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

// CavernCrawlMapPathCompleted is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCavernCrawlMapPathCompleted (CMsgGCToClientCavernCrawlMapPathCompleted).
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

// CavernCrawlMapUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCavernCrawlMapUpdated (CMsgGCToClientCavernCrawlMapUpdated).
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

// ChatModeratorBan is an event delivered by the GC.
//
// Message: k_EMsgGCChatModeratorBan (CMsgDOTAChatModeratorBan).
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

// ChatRegionsEnabled is an event delivered by the GC.
//
// Message: k_EMsgGCToClientChatRegionsEnabled (CMsgDOTAChatRegionsEnabled).
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

// ClaimEventActionUsingItemCompleted is an event delivered by the GC.
//
// Message: k_EMsgGCToClientClaimEventActionUsingItemCompleted (CMsgGCToClientClaimEventActionUsingItemCompleted).
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

// ClientSuspended is emitted when the GC suspends new sessions, usually for
// a scheduled update or maintenance.
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

// CoachTeammateRatingsChanged is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCoachTeammateRatingsChanged (CMsgGCToClientCoachTeammateRatingsChanged).
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

// CommendNotification is emitted when another player commends your account.
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

// CompendiumRemoveAllSelections is an event delivered by the GC.
//
// Message: k_EMsgGCCompendiumRemoveAllSelections (CMsgDOTACompendiumRemoveAllSelections).
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

// CraftworksUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientCraftworksUserDataUpdated (CMsgGCToClientCraftworksUserDataUpdated).
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

// DOTAWeekendTourneySchedule is an event delivered by the GC.
//
// Message: k_EMsgDOTAWeekendTourneySchedule (CMsgWeekendTourneySchedule).
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

// EmoticonData is an event delivered by the GC.
//
// Message: k_EMsgGCToClientEmoticonData (CMsgGCToClientEmoticonData).
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

// FantasyFinalPlayerStats is an event delivered by the GC.
//
// Message: k_EMsgGCFantasyFinalPlayerStats (CMsgDOTAFantasyFinalPlayerStats).
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

// FightingGameChallenge is an event delivered by the GC.
//
// Message: k_EMsgGCToClientFightingGameChallenge (CMsgGCToClientFightingGameChallenge).
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

// FightingGameChallengeCanceled is an event delivered by the GC.
//
// Message: k_EMsgGCToClientFightingGameChallengeCanceled (CMsgGCToClientFightingGameChallengeCanceled).
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

// FightingGameStartMatch is an event delivered by the GC.
//
// Message: k_EMsgGCToClientFightingGameStartMatch (CMsgGCToClientFightingGameStartMatch).
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

// GuildDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientGuildDataUpdated (CMsgGCToClientGuildDataUpdated).
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

// GuildFeedUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientGuildFeedUpdated (CMsgGCToClientGuildFeedUpdated).
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

// GuildMembersDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientGuildMembersDataUpdated (CMsgGCToClientGuildMembersDataUpdated).
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

// GuildMembershipUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientGuildMembershipUpdated (CMsgGCToClientGuildMembershipUpdated).
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

// HeroStatueCreateResult is an event delivered by the GC.
//
// Message: k_EMsgGCToClientHeroStatueCreateResult (CMsgGCToClientHeroStatueCreateResult).
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

// InviteToDemoMode is an event delivered by the GC.
//
// Message: k_EMsgGCToClientInviteToDemoMode (CMsgGCToClientInviteToDemoMode).
type InviteToDemoMode struct {
	protocol.CMsgGCToClientInviteToDemoMode
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *InviteToDemoMode) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientInviteToDemoMode
}

// GetEventBody returns the event body.
func (e *InviteToDemoMode) GetEventBody() proto.Message {
	return &e.CMsgGCToClientInviteToDemoMode
}

// GetEventName returns the event name.
func (e *InviteToDemoMode) GetEventName() string {
	return "InviteToDemoMode"
}

// ItemBattlerUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientItemBattlerUserDataUpdated (CMsgGCToClientItemBattlerUserDataUpdated).
type ItemBattlerUserDataUpdated struct {
	protocol.CMsgGCToClientItemBattlerUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *ItemBattlerUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientItemBattlerUserDataUpdated
}

// GetEventBody returns the event body.
func (e *ItemBattlerUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientItemBattlerUserDataUpdated
}

// GetEventName returns the event name.
func (e *ItemBattlerUserDataUpdated) GetEventName() string {
	return "ItemBattlerUserDataUpdated"
}

// KickedFromMatchmakingQueue is emitted when the GC removes the account from
// a matchmaking queue it was waiting in.
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

// LeagueAdminList is an event delivered by the GC.
//
// Message: k_EMsgGCLeagueAdminList (CMsgLeagueAdminList).
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

// LobbyMVPAwarded is emitted when the MVP votes for a finished lobby match
// are awarded.
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

// LobbyUpdateBroadcastChannelInfo is an event delivered by the GC.
//
// Message: k_EMsgGCLobbyUpdateBroadcastChannelInfo (CMsgGCLobbyUpdateBroadcastChannelInfo).
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

// MatchGroupsVersion is an event delivered by the GC.
//
// Message: k_EMsgGCToClientMatchGroupsVersion (CMsgGCToClientMatchGroupsVersion).
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

// MatchSignedOut is an event delivered by the GC.
//
// Message: k_EMsgGCToClientMatchSignedOut (CMsgGCToClientMatchSignedOut).
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

// MergeGroupInviteReply is an event delivered by the GC.
//
// Message: k_EMsgGCToClientMergeGroupInviteReply (CMsgDOTAGroupMergeReply).
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

// MergePartyResponseReply is an event delivered by the GC.
//
// Message: k_EMsgGCToClientMergePartyResponseReply (CMsgDOTAGroupMergeReply).
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

// MonsterHunterUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientMonsterHunterUserDataUpdated (CMsgGCToClientMonsterHunterUserDataUpdated).
type MonsterHunterUserDataUpdated struct {
	protocol.CMsgGCToClientMonsterHunterUserDataUpdated
}

// GetDotaEventMsgID returns the dota message ID of the event.
func (e *MonsterHunterUserDataUpdated) GetDotaEventMsgID() protocol.EDOTAGCMsg {
	return protocol.EDOTAGCMsg_k_EMsgGCToClientMonsterHunterUserDataUpdated
}

// GetEventBody returns the event body.
func (e *MonsterHunterUserDataUpdated) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMonsterHunterUserDataUpdated
}

// GetEventName returns the event name.
func (e *MonsterHunterUserDataUpdated) GetEventName() string {
	return "MonsterHunterUserDataUpdated"
}

// NotificationsUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientNotificationsUpdated (CMsgGCNotificationsResponse).
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

// NotifyAccountFlagsChange is an event delivered by the GC.
//
// Message: k_EMsgGCNotifyAccountFlagsChange (CMsgDOTANotifyAccountFlagsChange).
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

// OverwatchCasesAvailable is an event delivered by the GC.
//
// Message: k_EMsgGCToClientOverwatchCasesAvailable (CMsgGCToClientOverwatchCasesAvailable).
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

// OverworldUserDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientOverworldUserDataUpdated (CMsgGCToClientOverworldUserDataUpdated).
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

// PartyBeaconUpdate is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPartyBeaconUpdate (CMsgGCToClientPartyBeaconUpdate).
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

// PartyLeaderWatchGamePrompt is an event delivered by the GC.
//
// Message: k_EMsgGCPartyLeaderWatchGamePrompt (CMsgPartyLeaderWatchGamePrompt).
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

// PartySearchInvite is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPartySearchInvite (CMsgGCToClientPartySearchInvite).
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

// PartySearchInvites is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPartySearchInvites (CMsgGCToClientPartySearchInvites).
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

// PlayerBeaconState is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPlayerBeaconState (CMsgGCToClientPlayerBeaconState).
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

// OtherJoinedChannel is emitted when another member joins a chat channel you
// are in.
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

// OtherLeftChannel is emitted when another member leaves a chat channel you
// are in.
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

// PlaytestStatus is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPlaytestStatus (CMsgGCToClientPlaytestStatus).
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

// Popup is an event delivered by the GC.
//
// Message: k_EMsgGCPopup (CMsgDOTAPopup).
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

// PrivateCoachingSessionUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientPrivateCoachingSessionUpdated (CMsgGCToClientPrivateCoachingSessionUpdated).
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

// ProfileCardUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientProfileCardUpdated (CMsgDOTAProfileCard).
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

// QuestProgressUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientQuestProgressUpdated (CMsgGCToClientQuestProgressUpdated).
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

// RankUpdate is an event delivered by the GC.
//
// Message: k_EMsgGCToClientRankUpdate (CMsgGCToClientRankUpdate).
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

// ReadyUpStatus is an event delivered by the GC.
//
// Message: k_EMsgGCReadyUpStatus (CMsgReadyUpStatus).
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

// RequestLaneSelection is an event delivered by the GC.
//
// Message: k_EMsgGCToClientRequestLaneSelection (CMsgGCToClientRequestLaneSelection).
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

// RequestMMInfo is an event delivered by the GC.
//
// Message: k_EMsgGCToClientRequestMMInfo (CMsgGCToClientRequestMMInfo).
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

// RoadToTIQuestDataUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientRoadToTIQuestDataUpdated (CMsgGCToClientRoadToTIQuestDataUpdated).
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

// SteamDatagramTicket is an event delivered by the GC.
//
// Message: k_EMsgGCToClientSteamDatagramTicket (CMsgGCToClientSteamDatagramTicket).
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

// TeamInfo is an event delivered by the GC.
//
// Message: k_EMsgGCToClientTeamInfo (CMsgDOTATeamInfo).
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

// TeamInviteGCImmediateResponseToInviter is an event delivered by the GC.
//
// Message: k_EMsgGCTeamInvite_GCImmediateResponseToInviter (CMsgDOTATeamInvite_GCImmediateResponseToInviter).
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

// TeamInviteReceived is an event delivered by the GC.
//
// Message: k_EMsgGCTeamInvite_GCRequestToInvitee (CMsgDOTATeamInvite_GCRequestToInvitee).
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

// TeamInviteResponseReceived is an event delivered by the GC.
//
// Message: k_EMsgGCTeamInvite_GCResponseToInviter (CMsgDOTATeamInvite_GCResponseToInviter).
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

// TeamsInfo is an event delivered by the GC.
//
// Message: k_EMsgGCToClientTeamsInfo (CMsgDOTATeamsInfo).
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

// TournamentItemDrop is an event delivered by the GC.
//
// Message: k_EMsgGCToClientTournamentItemDrop (CMsgGCToClientTournamentItemDrop).
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

// TrophyAwarded is an event delivered by the GC.
//
// Message: k_EMsgGCToClientTrophyAwarded (CMsgGCToClientTrophyAwarded).
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

// UnderDraftGoldUpdated is an event delivered by the GC.
//
// Message: k_EMsgGCToClientUnderDraftGoldUpdated (CMsgGCToClientGuildUnderDraftGoldUpdated).
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

// VACReminder is an event delivered by the GC.
//
// Message: k_EMsgGCToClientVACReminder (CMsgGCToClientVACReminder).
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

// WageringUpdate is an event delivered by the GC.
//
// Message: k_EMsgGCToClientWageringUpdate (CMsgGCToClientWageringUpdate).
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

// WatchDownloadedReplay is an event delivered by the GC.
//
// Message: k_EMsgGCWatchDownloadedReplay (CMsgGCWatchDownloadedReplay).
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
