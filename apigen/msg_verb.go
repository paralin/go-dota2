package main

import (
	"strings"
)

var verbs = make(map[string]struct{})

func init() {
	verbList := []string{
		"Anchor",
		"Abandon",
		"Apply",
		"Autograph",
		"Award",
		"Cancel",
		"Check",
		"Close",
		"Create",
		"Demote",
		"Destroy",
		"Edit",
		"Find",
		"Flip",
		"Get",
		"Give",
		"Grant",
		"Join",
		"Kick",
		"Launch",
		"Leave",
		"List",
		"Open",
		"Promote",
		"Publish",
		"Purchase",
		"Query",
		"Record",
		"Recycle",
		"Redeem",
		"Refresh",
		"Release",
		"Report",
		"Revoke",
		"Request",
		"Reroll",
		"Reserve",
		"Retrieve",
		"Select",
		"Send",
		"Set",
		"Spectate",
		"Start",
		"Stop",
		"Submit",
		"Swap",
		"Toggle",
		"Track",
		"Transfer",
		"Upgrade",
		"Vote",
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
