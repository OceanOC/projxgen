package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func cMini(name string) error {
	namec := filepath.Clean(name)
	os.MkdirAll(filepath.Join(namec, "src"), 0700)
	os.MkdirAll(filepath.Join(namec, "build"), 0700)

	cf, err := os.Create(filepath.Join(namec, "src", "main.c"))
	if err != nil {
		fmt.Println("Could not create main.c", err.Error())
		return err
	}
	mcs := getTemplate(C_Mini)
	mcs = strings.ReplaceAll(mcs, ".$PROJNAME", name)
	cf.WriteString(mcs)

	return nil
}

func cMakeMini(name, cflags, ldflags, cc string) error {
	namec := filepath.Clean(name)
	cMini(namec)

	mf, err := os.Create(filepath.Join(namec, "Makefile"))
	if err != nil {
		fmt.Println("Could not create Makefile", err.Error())
		return err
	}

	mfs := getTemplate(Makefile_C)
	mfs = strings.ReplaceAll(mfs, ".$PROJNAME", name)
	mfs = strings.ReplaceAll(mfs, ".$CFLAGS", cflags)
	mfs = strings.ReplaceAll(mfs, ".$LDFLAGS", ldflags)
	mfs = strings.ReplaceAll(mfs, ".$CC", cc)
	mf.WriteString(mfs)

	return nil
}
