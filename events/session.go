package events

import (
	sprotocol "github.com/faceit/go-steam/protocol/gamecoordinator"
	gcsdkm "github.com/paralin/go-dota2/protocol"
)

// GCConnectionStatusChanged is emitted when the client connection state is updated.
type GCConnectionStatusChanged struct {
	// OldState contains the old connection status.
	OldState gcsdkm.GCConnectionStatus
	// NewState contains the new connection status.
	NewState gcsdkm.GCConnectionStatus
	// Update contains the message from the server that triggered this change, may be nil.
	Update *gcsdkm.CMsgConnectionStatus
}

// ClientWelcomed is emitted when the client receives the GC welcome
type ClientWelcomed struct {
	// Welcome is the welcome message from the GC.
	Welcome *gcsdkm.CMsgClientWelcome
}

// UnhandledGCPacket is called when the client ignores an unhandled packet.
type UnhandledGCPacket struct {
	// Packet is the unhandled packet.
	Packet *sprotocol.GCPacket
}
