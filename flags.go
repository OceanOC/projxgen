package main

import "github.com/urfave/cli/v3"

type FlagVars struct {
	cver, compiler string
}

func cFlags(fv *FlagVars) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "compiler",
			Aliases:     []string{"cc"},
			Value:       "gcc",
			Usage:       "command to use for the compiler",
			Destination: &fv.compiler,
		},
		&cli.StringFlag{
			Name:        "C Standard",
			Aliases:     []string{"cstd"},
			Value:       "c99",
			Usage:       "Version of C Standard to use",
			Destination: &fv.cver,
		},
	}
}
