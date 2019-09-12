package state

import (
	gcmm "github.com/paralin/go-dota2/protocol"
	gcsdkm "github.com/paralin/go-dota2/protocol"
)

// Dota2State is a snapshot of the client state at a point in time.
type Dota2State struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus gcsdkm.GCConnectionStatus
	// Lobby is the current lobby object.
	Lobby *gcmm.CSODOTALobby
	// Party is the current party object.
	Party *gcmm.CSODOTAParty
	// PartyInvite is the active incoming party invite.
	PartyInvite *gcmm.CSODOTAPartyInvite
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *gcsdkm.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *Dota2State) ClearState() {
	*s = Dota2State{ConnectionStatus: gcsdkm.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *Dota2State) IsReady() bool {
	return s.ConnectionStatus == gcsdkm.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}
