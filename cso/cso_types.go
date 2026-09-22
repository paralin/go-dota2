package cso

import (
	protobuf "github.com/aperturerobotics/protobuf-go-lite"
	bgcm "github.com/paralin/go-dota2/protocol"
	"github.com/pkg/errors"
)

// CSOType is a shared object type identifier.
//
//go:generate stringer -type=CSOType
type CSOType int32

const (
	// EconItem is an economy item.
	EconItem CSOType = 1
	// ItemRecipe is an item recipe.
	ItemRecipe = 5
	// EconGameAccountClient is an economy game account client.
	EconGameAccountClient = 7
	// SelectedItemPreset is a selected item preset.
	SelectedItemPreset = 35
	// ItemPresetInstance is an instance of an item preset.
	ItemPresetInstance = 36
	// DropRateBonus is an active drop rate bonus.
	DropRateBonus = 38
	// EventTicket is a ticket to an event.
	EventTicket = 40
	// ItemTournamentPassport is an item representing a tournament passport.
	ItemTournamentPassport = 42
	// GameAccountClient is the DOTA game account for a client.
	GameAccountClient = 2002
	// Party is a Dota 2 party.
	Party = 2003
	// Lobby is a Dota 2 lobby.
	Lobby = 2004
	// PartyInvite is an invite to a party.
	PartyInvite = 2006
	// GameHeroFavorites are game hero favorites.
	GameHeroFavorites = 2007
	// MapLocationState is the minimap location state.
	MapLocationState = 2008
	// Tournament represents a tournament.
	Tournament = 2009
	// PlayerChallenge represents a player challenge.
	PlayerChallenge = 2010
	// LobbyInvite is an invitation to a lobby.
	LobbyInvite = 2011
	// GameAccountPlus is the Dota Plus account.
	GameAccountPlus = 2012
)

// csoTypeCtors links type IDs to constructors.
var csoTypeCtors = map[CSOType]func() protobuf.Message{
	EconItem: func() protobuf.Message {
		return &bgcm.CSOEconItem{}
	},
	GameAccountClient: func() protobuf.Message {
		return &bgcm.CSODOTAGameAccountClient{}
	},
	Party: func() protobuf.Message {
		return &bgcm.CSODOTAParty{}
	},
	Lobby: func() protobuf.Message {
		return &bgcm.CSODOTALobby{}
	},
	PartyInvite: func() protobuf.Message {
		return &bgcm.CSODOTAPartyInvite{}
	},
	GameHeroFavorites: func() protobuf.Message {
		return &bgcm.CSODOTAGameHeroFavorites{}
	},
	MapLocationState: func() protobuf.Message {
		return &bgcm.CSODOTAMapLocationState{}
	},
	PlayerChallenge: func() protobuf.Message {
		return &bgcm.CSODOTAPlayerChallenge{}
	},
	LobbyInvite: func() protobuf.Message {
		return &bgcm.CSODOTALobbyInvite{}
	},
	DropRateBonus: func() protobuf.Message {
		return &bgcm.CSOEconItemDropRateBonus{}
	},
	GameAccountPlus: func() protobuf.Message {
		return &bgcm.CSODOTAGameAccountPlus{}
	},
}

// NewSharedObject builds a new shared object from a type ID.
func NewSharedObject(typ CSOType) (protobuf.Message, error) {
	ctor, ok := csoTypeCtors[typ]
	if !ok {
		return nil, errors.Errorf("unknown shared object type id: %d", typ)
	}

	return ctor(), nil
}
