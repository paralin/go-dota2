package main

import (
	"fmt"
	"reflect"
	"strings"

	gcm "github.com/paralin/go-dota2/protocol/dota_gcmessages_msgid"
	"github.com/pkg/errors"
)

// GetResponseMessageID returns the response message ID for the request.
// Error is returned if the request ID indicates there should be a response, but there is none.
func GetResponseMessageID(reqID gcm.EDOTAGCMsg) (gcm.EDOTAGCMsg, error) {
	if override, ok := msgResponseOverrides[reqID]; ok {
		return override, nil
	}

	msgID := reqID
	msgStr := msgID.String()
	msgStr = strings.TrimPrefix(msgStr, "k_EMsg")

	responseImplied := strings.HasSuffix(msgStr, "Request")
	msgStr = strings.TrimSuffix(msgStr, "Request")

	clientToGC := strings.HasPrefix(msgStr, "ClientToGC")
	msgStr = strings.TrimPrefix(msgStr, "ClientToGC")

	queryRespStr := func(respStr string) (gcm.EDOTAGCMsg, bool) {
		val, ok := gcm.EDOTAGCMsg_value["k_EMsg"+respStr]
		if ok {
			return gcm.EDOTAGCMsg(val), true
		}

		if clientToGC {
			val, ok = gcm.EDOTAGCMsg_value["k_EMsgGCToClient"+respStr]
			if ok {
				return gcm.EDOTAGCMsg(val), true
			}

			val, ok = gcm.EDOTAGCMsg_value["k_EMsgClientToGC"+respStr]
			if ok {
				return gcm.EDOTAGCMsg(val), true
			}
		}

		return gcm.EDOTAGCMsg(0), false
	}

	if respID, ok := queryRespStr(msgStr + "Response"); ok {
		return respID, nil
	}

	if responseImplied {
		queryStrs := []string{
			msgStr,
			msgStr + "RequestResponse",
		}
		for _, mstr := range queryStrs {
			if respID, ok := queryRespStr(mstr); ok {
				return respID, nil
			}
		}

		return gcm.EDOTAGCMsg(0), errors.Errorf("response was implied by request %v but no response type found", msgID.String())
	}

	return gcm.EDOTAGCMsg(0), nil
}

// LookupMessageProtoType lookup proto from message ID.
func LookupMessageProtoType(protoMap map[string]*ProtoType, msgID gcm.EDOTAGCMsg) (*ProtoType, error) {
	var protoName string
	if overrideMsg, ok := msgProtoTypeOverrides[msgID]; ok {
		protoName = reflect.TypeOf(overrideMsg).Elem().Name()

		pt, ok := protoMap[protoName]
		if !ok {
			return nil, errors.Errorf("proto not found: %s", protoName)
		}

		return pt, nil
	}

	msgStr := msgID.String()
	msgStr = strings.TrimPrefix(msgStr, "k_EMsg")

	withoutGC := strings.TrimPrefix(msgStr, "GC")
	withoutResponse := strings.TrimSuffix(msgStr, "Response")
	withoutGCOrResponse := strings.TrimPrefix(withoutResponse, "GC")
	withoutDota := strings.Replace(msgStr, "DOTA", "", -1)
	withoutToFrom := strings.Replace(msgStr, "GCToClient", "", -1)
	withoutToFrom = strings.Replace(withoutToFrom, "ClientToGC", "", -1)
	toAttempt := []string{
		msgStr,
		withoutDota,
		withoutToFrom,
		"ClientToGC" + withoutDota,
		"ClientToGC" + withoutDota + "Response",
		"GCToClient" + withoutDota,
		withoutGCOrResponse,
		"DOTA" + msgStr,
		withoutGC,
		"DOTA" + withoutGC,
		withoutResponse,
		"DOTA" + withoutResponse,
		withoutResponse + "Result",
		"DOTA" + withoutResponse + "Result",
		withoutGC + "Result",
		"DOTA" + withoutGC + "Result",
		"DOTA" + withoutGC + "Response",
		withoutGC + "Response",
		"DOTA" + withoutGCOrResponse + "Result",
		"DOTA" + withoutGCOrResponse + "Response",
		withoutGCOrResponse + "Response",
		withoutGCOrResponse + "Result",
	}

	for _, att := range toAttempt {
		if pt, ok := protoMap["CMsg"+att]; ok {
			fmt.Printf("Request: %v matched to type: %v\n", msgID.String(), pt.TypeName)
			return pt, nil
		}
	}

	return nil, errors.Errorf("unable to find proto for: %s", msgID.String())
}
