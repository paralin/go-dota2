package main

import (
	"fmt"
	"github.com/urfave/cli"
)

func init() {
	commands = append(commands, cli.Command{
		Name:  "print-type-list",
		Usage: "prints the protobuf type names",
		Action: func(c *cli.Context) error {
			packageMap, typeMap, err := BuildProtoTypeMap()
			if err != nil {
				return err
			}
			_ = packageMap

			for typeName := range typeMap {
				fmt.Printf("%s\n", typeName)
			}

			return nil
		},
	})
}
