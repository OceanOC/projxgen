package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func main() {
	fv := FlagVars{}

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
						Usage:    "A minimal C project with GNU Make build system",
						Category: "C",
						Aliases:  []string{"c", "c-make-min"},
						Flags:    cFlags(&fv),
						Action: func(ctx context.Context, cmd *cli.Command) error {
							return cMakeMini(cmd.Args().Get(0), strings.Join(fv.cflags, " --")+" --cstd="+fv.cver, strings.Join(fv.ldflags, " "), fv.compiler)
						},
					},
					{
						Name:     "C-Minimal",
						Usage:    "A minimal C project with no build system",
						Category: "C",
						Aliases:  []string{"c-min"},
						Action: func(ctx context.Context, cmd *cli.Command) error {
							return cMini(cmd.Args().Get(0))
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
