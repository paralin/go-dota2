package main

import (
	"strings"
)

var verbs = make(map[string]struct{})

func init() {
	verbList := []string{
		"Find",
		"Publish",
		"Retrieve",
		"Cancel",
		"Autograph",
		"Destroy",
		"Award",
		"Close",
		"Send",
		"Kick",
		"Create",
		"Edit",
		"Flip",
		"Get",
		"Redeem",
		"Record",
		"Recycle",
		"Refresh",
		"Track",
		"Upgrade",
		"Vote",
		"Reroll",
		"Promote",
		"Demote",
		"Give",
		"Grant",
		"Join",
		"List",
		"Launch",
		"Purchase",
		"Abandon",
		"Leave",
		"Open",
		"Query",
		"Release",
		"Report",
		"Request",
		"Reserve",
		"Select",
		"Set",
		"Spectate",
		"Start",
		"Submit",
		"Swap",
		"Toggle",
		"Transfer",
	}
	for _, v := range verbList {
		verbs[strings.ToLower(v)] = struct{}{}
	}
}

// IsWordVerb checks if a word is a suspected verb, from a hardcoded dictionary of common verbs.
func IsWordVerb(word string) bool {
	_, ok := verbs[strings.ToLower(word)]
	return ok
}
