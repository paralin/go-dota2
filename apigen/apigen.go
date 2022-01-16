package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var commands []cli.Command

// Generates a Go API for the Dota GC automatically with static analysis.
func main() {
	app := cli.NewApp()
	app.Name = "apigen"
	app.Usage = "generates a Go API client for the DOTA API using analysis and heuristics"
	app.HideVersion = true
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		os.Exit(1)
	}
}
