package main

import (
	"bytes"
	"context"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/tools/imports"
)

var (
	// clientOutFile selects the generated client implementation.
	clientOutFile string
	// eventsOutFile selects the generated event types.
	eventsOutFile string
)

// init registers the deterministic API generation command.
func init() {
	commands = append(commands, &cli.Command{
		Name:  "generate-api",
		Usage: "generates the API code",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "client-output",
				Usage:       "client output file",
				Value:       "../client_generated.go",
				Destination: &clientOutFile,
			},
			&cli.StringFlag{
				Name:        "events-output",
				Usage:       "events output file",
				Value:       "../events/generated.go",
				Destination: &eventsOutFile,
			},
		},
		Action: func(c *cli.Context) error {
			// Generate both boundaries from the same schema snapshot.
			clientFileBuf := &bytes.Buffer{}
			eventsFileBuf := &bytes.Buffer{}

			if err := GenerateAPI(context.Background(), clientFileBuf, eventsFileBuf); err != nil {
				return err
			}

			// Format both outputs before replacing either generated file.
			clientFileDat, err := imports.Process(clientOutFile, clientFileBuf.Bytes(), nil)
			if err != nil {
				return err
			}
			eventsFileDat, err := imports.Process(eventsOutFile, eventsFileBuf.Bytes(), nil)
			if err != nil {
				return err
			}

			// Retain generated source as ordinary non-executable files.
			if err := os.WriteFile(clientOutFile, clientFileDat, 0o644); err != nil {
				return err
			}
			return os.WriteFile(eventsOutFile, eventsFileDat, 0o644)
		},
	})
}
