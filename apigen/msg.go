package main

import (
	"sort"

	dota_gcmessages_msgid "github.com/paralin/go-dota2/protocol"
)

// IsValidMsg checks if the message is valid.
func IsValidMsg(msg dota_gcmessages_msgid.EDOTAGCMsg) bool {
	_, ok := dota_gcmessages_msgid.EDOTAGCMsg_name[int32(msg)]
	return ok && msg > dota_gcmessages_msgid.EDOTAGCMsg_k_EMsgGCDOTABase
}

func getSortedMsgIDs() []dota_gcmessages_msgid.EDOTAGCMsg {
	var msgIds []dota_gcmessages_msgid.EDOTAGCMsg
	for msgIDNum := range dota_gcmessages_msgid.EDOTAGCMsg_name {
		msgID := dota_gcmessages_msgid.EDOTAGCMsg(msgIDNum)
		msgIds = append(msgIds, msgID)
	}

	sort.Slice(msgIds, func(i int, j int) bool {
		return msgIds[i] < msgIds[j]
	})
	return msgIds
}
