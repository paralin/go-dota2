package events

import (
	gccm "github.com/paralin/go-dota2/protocol/dota_gcmessages_client"
)

// MatchSignedOut is emitted when the match signout event arrives from the GC.
type MatchSignedOut struct {
	gccm.CMsgGCToClientMatchSignedOut
}
