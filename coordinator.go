package dota2

import (
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/protocol/gamecoordinator"
	"github.com/paralin/go-steam/steamid"
)

// Coordinator carries Dota2 messages over an authenticated Steam connection.
// Implementations must preserve message order and serialize writes safely.
type Coordinator interface {
	// SteamID returns the authenticated Steam account for lobby authority checks.
	SteamID() steamid.SteamId

	// Write queues a message; callers must not modify it afterward.
	Write(gamecoordinator.IGCMsg)

	// SetGamesPlayed announces the applications running on this connection.
	SetGamesPlayed(...uint64)

	// RegisterPacketHandler attaches a handler before receiving packets.
	RegisterPacketHandler(steam.GCPacketHandler)
}
