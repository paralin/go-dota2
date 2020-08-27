package main

import (
	"fmt"
	"sort"

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

			names := make([]string, 0, len(typeMap))
			for typeName := range typeMap {
				names = append(names, typeName)
			}
			sort.Strings(names)
			for _, typeName := range names {
				fmt.Printf("%s\n", typeName)
			}

			return nil
		},
	})
}
