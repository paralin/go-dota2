package main

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/fatih/camelcase"
	gcm "github.com/paralin/go-dota2/protocol"
	"github.com/pkg/errors"
)

// generatedRequestFunc is a auto-generated request.
type generatedRequestFunc struct {
	methodName string
	reqType    types.Object
	reqMsgID   gcm.EDOTAGCMsg
	respType   types.Object
	respMsgID  gcm.EDOTAGCMsg
}

// buildGeneratedRequestFunc builds a generated request function.
func buildGeneratedRequestFunc(
	msgID gcm.EDOTAGCMsg,
	protoMap map[string]*ProtoType,
	clientImports map[string]struct{},
) (*generatedRequestFunc, error) {
	genReqFunc := &generatedRequestFunc{
		methodName: GetMessageFuncName(msgID),
	}

	rMsgID, err := GetResponseMessageID(msgID)
	if err != nil {
		return nil, err
	}

	genReqFunc.reqMsgID = msgID
	reqProtoType, err := LookupMessageProtoType(protoMap, msgID)
	if err != nil {
		return nil, err
	}
	genReqFunc.reqType = reqProtoType.Obj
	clientImports[reqProtoType.Pak.Path()] = struct{}{}

	genReqFunc.respMsgID = rMsgID
	if rMsgID != 0 {
		sender, ok := msgSenderOverrides[rMsgID]
		if !ok || sender == MsgSenderGC {
			fmt.Printf("Request: %v matched to response: %v\n", msgID.String(), rMsgID.String())
			respProtoType, err := LookupMessageProtoType(protoMap, rMsgID)
			if err != nil {
				return nil, err
			}

			genReqFunc.respType = respProtoType.Obj
			clientImports[respProtoType.Pak.Path()] = struct{}{}
			if genReqFunc.respType == genReqFunc.reqType {
				return nil, errors.Errorf("request type cannot be the same as response type: %s %s", msgID.String(), respProtoType.TypeStr)
			}
		}

		if genReqFunc.respType == nil {
			genReqFunc.respMsgID = 0
		}

		if genReqFunc.respMsgID != 0 && genReqFunc.respMsgID == msgID {
			return nil, errors.Errorf("request message cannot be the same as the response message: %s", msgID.String())
		}
	}

	return genReqFunc, nil
}

// generateComment generates a comment for the method.
func (f *generatedRequestFunc) generateComment() string {
	purpose := "is undocumented."

	words := camelcase.Split(f.methodName)
	if IsWordVerb(words[0]) {
		w0 := words[0]
		if w0 == "Query" {
			w0 = "Queries"
		} else {
			w0 = w0 + "s"
		}
		w0 = strings.ToLower(w0)

		if IsWordVerb(words[1]) || w0 == "queries" {
			purpose = fmt.Sprintf(
				"%s to check if the target %s.",
				w0,
				strings.ToLower(strings.Join(words[1:], " ")),
			)
		} else {
			a := " a"
			if f.methodName[len(f.methodName)-1] == 's' || words[1] == "All" {
				a = ""
			}

			purpose = fmt.Sprintf(
				"%s%s %s.",
				w0,
				a,
				strings.ToLower(strings.Join(words[1:], " ")),
			)
		}
	} else if IsWordVerb(words[len(words)-1]) {
		f := " for/from"

		purpose = fmt.Sprintf(
			"%ss%s a %s.",
			strings.ToLower(words[len(words)-1]),
			f,
			strings.ToLower(strings.Join(words[0:len(words)-1], " ")),
		)
	}

	acA := "// Request ID: " + f.reqMsgID.String()
	if f.respMsgID != 0 {
		acA += "\n// Response ID: " + f.respMsgID.String()
	}

	acB := "// Request type: " + f.reqType.Name()
	if f.respMsgID != 0 {
		acB += "\n// Response type: " + f.respType.Name()
	}

	return fmt.Sprintf("%s %s\n%s\n%s", f.methodName, purpose, acA, acB)
}
