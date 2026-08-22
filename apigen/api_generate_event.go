package main

import (
	"fmt"
	"strings"

	gcm "github.com/paralin/go-dota2/protocol"
)

type generatedEventHandler struct {
	msgID     gcm.EDOTAGCMsg
	eventName string
	eventType *ProtoType
}

// buildGeneratedEventHandler builds a generated event handler.
func buildGeneratedEventHandler(
	msgID gcm.EDOTAGCMsg,
	protoMap map[string]*ProtoType,
	eventImports map[string]struct{},
) (*generatedEventHandler, error) {
	eventProtoType, err := LookupMessageProtoType(protoMap, msgID)
	if err != nil {
		return nil, err
	}
	eventImports[eventProtoType.Pak.Path()] = struct{}{}

	return &generatedEventHandler{
		msgID:     msgID,
		eventName: GetMessageEventName(msgID),
		eventType: eventProtoType,
	}, nil
}

func (g *generatedEventHandler) generateComment() string {
	if doc, ok := msgEventDocOverrides[g.msgID]; ok {
		return fmt.Sprintf(
			"// %s\n",
			strings.ReplaceAll(strings.TrimRight(doc, "\n"), "\n", "\n// "),
		)
	}
	return fmt.Sprintf(
		"// %s is an event delivered by the GC.\n//\n// Message: %s (%s).\n",
		g.eventName, g.msgID.String(), g.eventType.TypeName,
	)
}
