package main

import (
	"strings"

	"github.com/fatih/camelcase"
	dota_gcmessages_msgid "github.com/paralin/go-dota2/protocol"
)

// GetMessageFuncName determines what function name we should assign a message.
func GetMessageFuncName(msg dota_gcmessages_msgid.EDOTAGCMsg) string {
	if str, ok := msgMethodNameOverrides[msg]; ok {
		return str
	}

	msgName := msg.String()
	msgName = strings.TrimPrefix(msgName, "k_EMsg")

	if strings.HasPrefix(msgName, "DOTA") {
		msgName = msgName[4:]
	}

	switch {
	case strings.HasPrefix(msgName, "ClientToGC"):
		msgName = strings.TrimPrefix(msgName, "ClientToGC")
	case strings.HasPrefix(msgName, "GC"):
		msgName = strings.TrimPrefix(msgName, "GC")
	}

	switch {
	case strings.HasSuffix(msgName, "Request"):
		msgName = "Request" + strings.TrimSuffix(msgName, "Request")
	case strings.HasSuffix(msgName, "Query"):
		msgName = "Query" + strings.TrimSuffix(msgName, "Query")
	}

	msgName = strings.Replace(msgName, "PracticeLobby", "Lobby", -1)
	words := camelcase.Split(msgName)
	switch {
	case len(words) < 2:
		return msgName
	case len(words) == 2:
		if !IsWordVerb(words[0]) && IsWordVerb(words[1]) {
			words[0], words[1] = words[1], words[0]
		}
		break
	case IsWordVerb(words[0]):
		break
	}

	switch {
	case words[0] == "List" && !strings.HasSuffix(words[len(words)-1], "s"):
		if words[1] == "Lobby" {
			words[1] = "Lobbies"
		} else {
			words[1] += "s"
		}
	case words[0] == "Request" && IsWordVerb(words[len(words)-1]):
		words[0] = words[len(words)-1]
		words = words[:len(words)-1]
	}

	anyVerbs := false
	var verbi int
	for i, w := range words {
		verbi = i
		if IsWordVerb(w) {
			anyVerbs = true
			break
		}
	}

	if !anyVerbs {
		words = append([]string{"Send"}, words...)
	} else if !IsWordVerb(words[0]) {
		words[0], words[verbi] = words[verbi], words[0]
	}

	msgName = strings.Join(words, "")
	return msgName
}
