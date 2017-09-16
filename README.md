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
 - [ ] SOCache tracking / state management
 - [ ] Lobby tracking / state management
    - [ ] Read lobby state correctly
    - [ ] Implement normal lobby operations
 - [ ] Party tracking / state management
    - [ ] Read party and invite state correctly
    - [ ] Implement normal party operations

... and others. This is the current short-term roadmap.

