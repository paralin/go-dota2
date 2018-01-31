package main

import (
	"strings"

	"github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
)

// GetMessageEventName returns the event name for the message.
func GetMessageEventName(msg dota_gcmessages_msgid.EDOTAGCMsg) string {
	msgName := msg.String()
	msgName = strings.TrimPrefix(msgName, "k_EMsg")
	msgName = strings.TrimPrefix(msgName, "GC")
	msgName = strings.TrimPrefix(msgName, "ToClient")
	return msgName
}
