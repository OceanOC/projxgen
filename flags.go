package main

import "github.com/urfave/cli/v3"

type FlagVars struct {
	ldflags, cflags []string
	cver, compiler  string
	wextra, wall    bool
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
			Name:        "C-Standard",
			Aliases:     []string{"cstd"},
			Value:       "c99",
			Usage:       "Version of C Standard to use",
			Destination: &fv.cver,
		},
		&cli.BoolFlag{
			Name:        "Extra-Warnings",
			Aliases:     []string{"Wextra", "extra-warnings"},
			Value:       false,
			Usage:       "Version of C Standard to use",
			Destination: &fv.wextra,
		},
		&cli.BoolFlag{
			Name:        "All-Warnings",
			Aliases:     []string{"Wall", "all-warnings"},
			Value:       false,
			Usage:       "Version of C Standard to use",
			Destination: &fv.wall,
		},
		&cli.StringSliceFlag{
			Name:        "C-Flag",
			Aliases:     []string{"cflag"},
			Usage:       "Add multiple C flags",
			Destination: &fv.cflags,
		},
		&cli.StringSliceFlag{
			Name:        "ld-Flag",
			Aliases:     []string{"ld"},
			Usage:       "Add multiple Linker flags",
			Destination: &fv.ldflags,
		},
	}
}
