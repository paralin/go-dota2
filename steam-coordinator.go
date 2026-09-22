package dota2

import (
	"github.com/paralin/go-steam"
	"github.com/paralin/go-steam/steamid"
)

// SteamCoordinator combines GC transport with its authenticated Steam identity.
type SteamCoordinator struct {
	// GameCoordinator carries GC messages on this client's connection.
	*steam.GameCoordinator
	// client supplies the current identity after Steam authentication.
	client *steam.Client
}

// NewSteamCoordinator binds the GC transport and identity to the same client.
func NewSteamCoordinator(client *steam.Client) *SteamCoordinator {
	return &SteamCoordinator{GameCoordinator: client.GC, client: client}
}

// SteamID returns the Steam identity established by the current connection.
func (s *SteamCoordinator) SteamID() steamid.SteamId { return s.client.SteamId() }

var _ Coordinator = (*SteamCoordinator)(nil)
