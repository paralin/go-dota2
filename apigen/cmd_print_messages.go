package main

import (
	"bytes"
	"fmt"

	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:  "print-messages",
		Usage: "prints message IDs and derived information",
		Action: func(c *cli.Context) error {
			m := DumpMessageIDDebug()
			fmt.Printf("%s", m)
			return nil
		},
	})
}

// DumpMessageIDDebug dumps debug information about the messages in the system.
func DumpMessageIDDebug() string {
	var buf bytes.Buffer
	msgIds := getSortedMsgIDs()

	var allMethods []string
	for _, msgID := range msgIds {
		sender := GetMessageSender(msgID)
		fmt.Fprintf(&buf, "%d", msgID)

		switch sender {
		case MsgSenderNone:
			buf.WriteString("\tN/A   ")
		case MsgSenderClient:
			buf.WriteString("\tCLIENT")
		case MsgSenderGC:
			buf.WriteString("\tGC    ")
		}

		fmt.Fprintf(&buf, "\t%s", msgID.String())

		if sender == MsgSenderClient {
			methodName := GetMessageFuncName(msgID)
			fmt.Fprintf(&buf, "\t%s", methodName)
			allMethods = append(allMethods, methodName)
		}

		(&buf).WriteString("\n")
	}

	(&buf).WriteString(fmt.Sprintf("\nComputed client methods: %v\n", allMethods))

	return (&buf).String()
}
