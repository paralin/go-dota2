# Dota2

[![GoDoc Widget]][GoDoc] [![Go Report Card Widget]][Go Report Card]

[GoDoc]: https://godoc.org/github.com/paralin/go-dota2
[GoDoc Widget]: https://godoc.org/github.com/paralin/go-dota2
[Go Report Card Widget]: https://goreportcard.com/badge/github.com/paralin/go-dota2
[Go Report Card]: https://goreportcard.com/report/github.com/paralin/go-dota2

## Introduction 
**go-dota2** is a DOTA 2 client plugin for go-steam. The intent is to replicate as much of the client functionality as possible.

Also see: the [C#/SteamKit Implementation](https://github.com/paralin/Dota2).

## Implementation Progress

Currently, the following client features have been implemented in this library:

 - [x] GC session state management
 - [x] Player profile fetching / call tracking
 - [x] SOCache tracking / state management
 - [x] Basic chat interaction
 - [~] Lobby tracking / state management
    - [x] Read lobby state correctly
    - [ ] Implement normal lobby operations
 - [~] Party tracking / state management
    - [x] Read party and invite state correctly
    - [ ] Implement normal party operations

... and others. This is the current short-term roadmap.

## SOCache Mechanism

The caching mechanism makes it easy to watch for changes to common objects, like `Lobby, LobbyInvite, Party, PartyInvite`.

This mechanism is used everywhere, these objects are not exposed in their own events.

```go
import (
	gcmm "github.com/paralin/go-dota2/protocol/dota_gcmessages_common_match_management"
	"github.com/paralin/go-dota2/cso"
)

eventCh, eventCancel, err := dota.GetCache().SubscribeType(cso.Lobby)
if err != nil {
    return err
}

defer eventCancel()

lobbyEvent := <-eventCh
lobby := lobbyEvent.Object.(*gcmm.CSODOTALobby)
```

Events for the object type are emitted on the eventCh. Be sure to call `eventCancel` once you are done with the channel to prevent resource leaks.

The cache object also adds interfaces to get and list the current objects in the cache.

## go-steam Dependency

This library depends on `go-steam`. Currently we are using the [FACEIT fork](https://github.com/faceit/go-steam).

