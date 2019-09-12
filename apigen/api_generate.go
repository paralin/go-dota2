package main

import (
	"context"
	"fmt"
	"go/types"
	"io"
	"sort"
	"strings"
	"unicode"

	gcm "github.com/paralin/go-dota2/protocol"
	"github.com/pkg/errors"
	"github.com/serenize/snaker"
)

// GenerateAPI uses protobuf reflection to generate an API client for Dota.
func GenerateAPI(ctx context.Context, clientOutput, eventsOutput io.Writer) error {
	senderMap := make(map[gcm.EDOTAGCMsg]MsgSender)

	packageMap, protoMap, err := BuildProtoTypeMap()
	if err != nil {
		return err
	}

	fmt.Fprintf(clientOutput, "package dota2\n\nimport (\n")

	msgIds := getSortedMsgIDs()
	var requestFuncs []*generatedRequestFunc

	eventsImports := make(map[string]struct{})
	eventsImports["github.com/paralin/go-dota2/protocol"] = struct{}{}
	eventsImports["github.com/golang/protobuf/proto"] = struct{}{}

	clientImports := make(map[string]struct{})
	clientImports["context"] = struct{}{}
	clientImports["github.com/paralin/go-dota2/protocol"] = struct{}{}
	clientImports["github.com/faceit/go-steam/steamid"] = struct{}{}
	clientImports["github.com/paralin/go-dota2/events"] = struct{}{}

	// responseMsgs are messages that are known to be responses.
	responseMsgs := make(map[gcm.EDOTAGCMsg]struct{})
	eventHandlers := make(map[gcm.EDOTAGCMsg]*generatedEventHandler)
	var eventHandlersOrdered []*generatedEventHandler

	for _, msgID := range msgIds {
		sender := GetMessageSender(msgID)
		if sender == MsgSenderNone {
			continue
		}

		senderMap[msgID] = sender
		if sender == MsgSenderClient {
			genReqFunc, err := buildGeneratedRequestFunc(msgID, protoMap, clientImports)
			if err != nil {
				return err
			}

			if genReqFunc.respMsgID != 0 {
				responseMsgs[genReqFunc.respMsgID] = struct{}{}
			}

			requestFuncs = append(requestFuncs, genReqFunc)
		} else if sender == MsgSenderGC {
			msgStr := msgID.String()
			if strings.HasSuffix(msgStr, "Response") {
				continue
			}

			eventHandler, err := buildGeneratedEventHandler(msgID, protoMap, eventsImports)
			if err != nil {
				return err
			}

			eventHandlers[msgID] = eventHandler
			eventHandlersOrdered = append(eventHandlersOrdered, eventHandler)
		}
	}

	for pakPath := range clientImports {
		fmt.Fprintf(clientOutput, "\t\"%s\"\n", pakPath)
	}
	fmt.Fprintf(clientOutput, ")\n")

	sort.Slice(requestFuncs, func(i int, j int) bool {
		return requestFuncs[i].methodName < requestFuncs[j].methodName
	})

	steamIDFieldOverrides := make(map[string]string)
	for _, f := range requestFuncs {
		fmt.Fprintf(clientOutput, "\n// %s\n", f.generateComment())
		fmt.Fprintf(clientOutput, "func (d *Dota2) %s(\n", f.methodName)

		respTyp := f.respType
		if respTyp != nil {
			fmt.Fprintf(clientOutput, "\tctx context.Context,\n")
		}

		reqTyp := f.reqType.Type()
		reqTypUnderlying := reqTyp.Underlying()
		reqDs := reqTypUnderlying.(*types.Struct)
		reqFields := make(map[string]string)
		var reqFieldsOrdered []string

		reqObjAsArgument := reqDs.NumFields() > 8
		if v, ok := msgArgAsParameterOverrides[f.reqMsgID]; ok {
			reqObjAsArgument = v
		}

		if reqObjAsArgument {
			fmt.Fprintf(clientOutput, "\treq *")
			if err := printFieldType(clientOutput, reqTyp); err != nil {
				return err
			}
			fmt.Fprintf(clientOutput, ",\n")
		} else {
			for i := 0; i < reqDs.NumFields(); i++ {
				reqField := reqDs.Field(i)
				if !reqField.Exported() {
					continue
				}

				reqFieldName := reqField.Name()
				if strings.HasPrefix(reqFieldName, "XXX_") {
					continue
				}

				reqFieldNameUpper := reqFieldName
				reqFieldName = snaker.SnakeToCamel(reqFieldName)
				reqFieldName = string(append([]rune{unicode.ToLower([]rune(reqFieldName)[0])}, []rune(reqFieldName[1:])...))
				switch reqFieldName {
				case "map":
					reqFieldName = "gameMap"
				case "clientVersion":
					continue
				}

				reqFieldName = strings.Replace(reqFieldName, "Id", "ID", -1)

				_, isSlice := reqField.Type().(*types.Slice)
				if isSlice {
					if _, ok := reqFields[reqFieldName]; !ok {
						reqFieldsOrdered = append(reqFieldsOrdered, reqFieldName)
					}
					reqFields[reqFieldName] = reqFieldNameUpper
				} else {
					if _, ok := reqFields["&"+reqFieldName]; !ok {
						reqFieldsOrdered = append(reqFieldsOrdered, "&"+reqFieldName)
					}
					reqFields["&"+reqFieldName] = reqFieldNameUpper
				}

				reqFieldType := reqField.Type()
				reqFieldTypePtr, isPtr := reqFieldType.(*types.Pointer)
				if isPtr {
					reqFieldType = reqFieldTypePtr.Elem()
				}

				if !isSlice &&
					strings.Contains(strings.ToLower(reqFieldName), "steamid") &&
					strings.Contains(reqFieldType.String(), "uint64") {
					reqFieldType = types.NewNamed(
						types.NewTypeName(
							0,
							packageMap["github.com/faceit/go-steam/steamid"],
							"SteamId",
							reqFieldType,
						),
						reqFieldType,
						nil,
					)
					reqFieldNameAfter := reqFieldName + "U64"
					if _, ok := reqFields[reqFieldNameAfter]; !ok {
						reqFieldsOrdered = append(reqFieldsOrdered, reqFieldNameAfter)
					}
					steamIDFieldOverrides[reqFieldNameAfter] = reqFieldName
					delete(reqFields, "&"+reqFieldName)
					reqFields[reqFieldNameAfter] = reqFieldNameUpper
				}

				fmt.Fprintf(clientOutput, "\t%s ", reqFieldName)
				if err := printFieldType(clientOutput, reqFieldType); err != nil {
					return err
				}
				fmt.Fprintf(clientOutput, ",\n")
			}
		}

		fmt.Fprintf(clientOutput, ") ")

		if respTyp != nil {
			fmt.Fprintf(clientOutput, "(*")
			if err := printFieldType(clientOutput, respTyp.Type()); err != nil {
				return err
			}
			fmt.Fprintf(clientOutput, ", error) ")
		}
		fmt.Fprintf(clientOutput, "{\n")

		// transform steam IDs
		for _, fieldName := range reqFieldsOrdered {
			if _, ok := reqFields[fieldName]; !ok {
				continue
			}

			if sidOverride, ok := steamIDFieldOverrides[fieldName]; ok {
				fmt.Fprintf(clientOutput, "\t%sVal := uint64(%s)\n", fieldName, sidOverride)
				fmt.Fprintf(clientOutput, "\t%s := &%sVal\n", fieldName, fieldName)
			}
		}

		if !reqObjAsArgument {
			fmt.Fprintf(clientOutput, "\treq := &")
			if err := printFieldType(clientOutput, reqTyp); err != nil {
				return err
			}
			fmt.Fprintf(clientOutput, "{")
			for _, fieldName := range reqFieldsOrdered {
				fieldNameUpper, ok := reqFields[fieldName]
				if !ok {
					continue
				}

				fmt.Fprintf(clientOutput, "\n\t\t%s: %s,", fieldNameUpper, fieldName)
			}
			if len(reqFields) != 0 {
				fmt.Fprintf(clientOutput, "\n\t")
			}
			fmt.Fprintf(clientOutput, "}\n")
		}

		if respTyp != nil {
			fmt.Fprintf(clientOutput, "\tresp := &")
			if err := printFieldType(clientOutput, respTyp.Type()); err != nil {
				return err
			}
			fmt.Fprintf(clientOutput, "{}\n\n\treturn resp, d.MakeRequest(\n")
			fmt.Fprintf(clientOutput, "\t\tctx,\n\t\tuint32(protocol.EDOTAGCMsg_%s),\n", f.reqMsgID.String())
			fmt.Fprintf(clientOutput, "\t\treq,\n\t\tuint32(protocol.EDOTAGCMsg_%s),\n", f.respMsgID.String())
			fmt.Fprintf(clientOutput, "\t\tresp,\n\t)\n")
		} else {
			fmt.Fprintf(clientOutput, "\td.write(uint32(protocol.EDOTAGCMsg_%s), req)\n", f.reqMsgID.String())
		}

		fmt.Fprintf(clientOutput, "}\n")
	}

	sort.Slice(eventHandlersOrdered, func(i int, j int) bool {
		return eventHandlersOrdered[i].eventName < eventHandlersOrdered[j].eventName
	})

	fmt.Fprintf(clientOutput, "\n// registerGeneratedHandlers registers the auto-generated event handlers.\n")
	fmt.Fprintf(clientOutput, "func (d *Dota2) registerGeneratedHandlers() {\n")

	fmt.Fprintf(eventsOutput, "package events\n\nimport (\n")
	for pakPath := range eventsImports {
		fmt.Fprintf(eventsOutput, "\t\"%s\"\n", pakPath)
	}
	fmt.Fprintf(eventsOutput, ")\n")

	for _, eventHandler := range eventHandlersOrdered {
		fmt.Fprintf(eventsOutput, "\n")
		fmt.Fprintf(eventsOutput, eventHandler.generateComment())
		fmt.Fprintf(eventsOutput, "type %s struct {\n", eventHandler.eventName)
		fmt.Fprintf(eventsOutput, "\t")
		if err := printFieldType(eventsOutput, eventHandler.eventType.Obj.Type()); err != nil {
			return err
		}
		fmt.Fprintf(eventsOutput, "\n}\n")

		fmt.Fprintf(eventsOutput, "\n// GetDotaEventMsgID returns the dota message ID of the event.\n")
		fmt.Fprintf(eventsOutput, "func (e *%s) GetDotaEventMsgID() protocol.EDOTAGCMsg {\n", eventHandler.eventName)
		fmt.Fprintf(eventsOutput, "\treturn protocol.EDOTAGCMsg_%s\n", eventHandler.msgID.String())
		fmt.Fprintf(eventsOutput, "}\n")

		fmt.Fprintf(eventsOutput, "\n// GetEventBody returns the event body.\n")
		fmt.Fprintf(eventsOutput, "func (e *%s) GetEventBody() proto.Message {\n", eventHandler.eventName)
		fmt.Fprintf(eventsOutput, "\treturn &e.%s\n", eventHandler.eventType.TypeName)
		fmt.Fprintf(eventsOutput, "}\n")

		fmt.Fprintf(eventsOutput, "\n// GetEventName returns the event name.\n")
		fmt.Fprintf(eventsOutput, "func (e *%s) GetEventName() string {\n", eventHandler.eventName)
		fmt.Fprintf(eventsOutput, "\treturn \"%s\"\n", eventHandler.eventName)
		fmt.Fprintf(eventsOutput, "}\n")

		fmt.Fprintf(
			clientOutput,
			"\td.handlers[uint32(protocol.EDOTAGCMsg_%s)] = d.getEventEmitter(func() events.Event {\n",
			eventHandler.msgID.String(),
		)
		fmt.Fprintf(clientOutput, "\t\treturn &events.%s{}\n\t})\n", eventHandler.eventName)
	}

	fmt.Fprintf(clientOutput, "}\n")
	return nil
}

func printFieldType(output io.Writer, reqFieldType types.Type) error {
	switch rft := reqFieldType.(type) {
	case *types.Basic:
		fmt.Fprintf(output, "%s", rft.Name())
	case *types.Pointer:
		fmt.Fprintf(output, "*")
		return printFieldType(output, rft.Elem())
	case *types.Named:
		pkgName := rft.Obj().Pkg().Name()
		typName := rft.Obj().Name()
		fmt.Fprintf(output, "%s.%s", pkgName, typName)
	case *types.Slice:
		elemType := rft.Elem()
		fmt.Fprintf(output, "[]")
		return printFieldType(output, elemType)
	default:
		return errors.Errorf("type unhandled: %#v", reqFieldType)
	}

	return nil
}
