package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	fv := FlagVars{}
	var compiler, linkerflags, langflags string

	cmd := &cli.Command{
		Name:  "projxgen",
		Usage: "Simply make a project template",
		Commands: []*cli.Command{
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Make a new project",
				Commands: []*cli.Command{
					{
						Name:     "C-GNUMake-Minimal",
						Category: "C",
						Aliases:  []string{"c", "c-make-min"},
						Flags:    cFlags(&fv),
						Action: func(ctx context.Context, cmd *cli.Command) error {
							cMakeMini(cmd.Args().Get(0), langflags, linkerflags, compiler)
							return nil
						},
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
