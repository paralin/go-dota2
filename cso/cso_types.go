package cso

import (
	"github.com/golang/protobuf/proto"
	bgcm "github.com/paralin/go-dota2/protocol"
	gcclm "github.com/paralin/go-dota2/protocol"
	gccm "github.com/paralin/go-dota2/protocol"
	gcmm "github.com/paralin/go-dota2/protocol"
	"github.com/pkg/errors"
)

// CSOType is a shared object type identifier.
//go:generate stringer -type=CSOType
type CSOType int32

const (
	// EconItem is an economy item.
	EconItem CSOType = 1
	// ItemRecipe is an item recipe.
	ItemRecipe = 5
	// EconGameAccountClient is a economy game account client..
	EconGameAccountClient = 7
	// SelectedItemPreset is a selected item preset.
	SelectedItemPreset = 35
	// ItemPresetInstance is a instance of an item preset.
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
)

// csoTypeCtors links type IDs to constructors.
var csoTypeCtors = map[CSOType]func() proto.Message{
	EconItem: func() proto.Message {
		return &bgcm.CSOEconItem{}
	},
	GameAccountClient: func() proto.Message {
		return &gccm.CSODOTAGameAccountClient{}
	},
	Party: func() proto.Message {
		return &gcmm.CSODOTAParty{}
	},
	Lobby: func() proto.Message {
		return &gcmm.CSODOTALobby{}
	},
	PartyInvite: func() proto.Message {
		return &gcmm.CSODOTAPartyInvite{}
	},
	GameHeroFavorites: func() proto.Message {
		return &gcclm.CSODOTAGameHeroFavorites{}
	},
	MapLocationState: func() proto.Message {
		return &gccm.CSODOTAMapLocationState{}
	},
	PlayerChallenge: func() proto.Message {
		return &gccm.CSODOTAPlayerChallenge{}
	},
	LobbyInvite: func() proto.Message {
		return &gcmm.CSODOTALobbyInvite{}
	},
	DropRateBonus: func() proto.Message {
		return &bgcm.CSOEconItemDropRateBonus{}
	},
}

// NewSharedObject builds a new shared object from a type ID.
func NewSharedObject(typ CSOType) (proto.Message, error) {
	ctor, ok := csoTypeCtors[typ]
	if !ok {
		return nil, errors.Errorf("unknown shared object type id: %d", typ)
	}

	return ctor(), nil
}
