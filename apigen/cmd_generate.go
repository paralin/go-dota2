package main

import (
	"bytes"
	"context"
	"github.com/urfave/cli"
	gofmt "go/format"
	"io/ioutil"
)

var clientOutFile string
var eventsOutFile string

func init() {
	commands = append(commands, cli.Command{
		Name:  "generate-api",
		Usage: "generates the API code",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "client-output",
				Usage:       "client output file",
				Value:       "../client_generated.go",
				Destination: &clientOutFile,
			},
			cli.StringFlag{
				Name:        "events-output",
				Usage:       "events output file",
				Value:       "../events/generated.go",
				Destination: &eventsOutFile,
			},
		},
		Action: func(c *cli.Context) error {
			clientFileBuf := &bytes.Buffer{}
			eventsFileBuf := &bytes.Buffer{}

			if err := GenerateAPI(context.Background(), clientFileBuf, eventsFileBuf); err != nil {
				return err
			}

			if err := ioutil.WriteFile(clientOutFile, clientFileBuf.Bytes(), 0655); err != nil {
				return err
			}

			clientFileDat, err := gofmt.Source(clientFileBuf.Bytes())
			if err != nil {
				clientFileDat = clientFileBuf.Bytes()
			}

			if err := ioutil.WriteFile(clientOutFile, clientFileDat, 0655); err != nil {
				return err
			}

			if err != nil {
				return err
			}

			eventsFileDat, err := gofmt.Source(eventsFileBuf.Bytes())
			if err != nil {
				eventsFileDat = eventsFileBuf.Bytes()
			}

			if err := ioutil.WriteFile(eventsOutFile, eventsFileDat, 0655); err != nil {
				return err
			}

			return err
		},
	})
}
