package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates/Makefile-C
var mfc string

func cMakeMini(name, cflags, ldflags, cc string) {
	namec := filepath.Clean(name)
	os.MkdirAll(filepath.Join(namec, "src"), 0700)
	os.MkdirAll(filepath.Join(namec, "build"), 0700)

	mf, err := os.Create(filepath.Join(namec, "Makefile"))
	if err != nil {
		fmt.Println("Could not create Makefile", err.Error())
	}

	var mfs string
	body, err := http.Get("https://raw.githubusercontent.com/OceanOC/projxgen/refs/heads/main/templates/Makefile-C")
	if err != nil || body.StatusCode != http.StatusOK {
		// Incase of any errors use embedded Makefile-C
		mfs = mfc
	} else {
		bodyBytes, err := io.ReadAll(body.Body)
		if err != nil {
			mfs = mfc
		} else {
			mfs = string(bodyBytes)
		}
	}
	mfs = strings.ReplaceAll(mfs, ".$PROJNAME", name)
	mfs = strings.ReplaceAll(mfs, ".$CFLAGS", cflags)
	mfs = strings.ReplaceAll(mfs, ".$LDFLAGS", ldflags)
	mfs = strings.ReplaceAll(mfs, ".$CC", cc)

	mf.WriteString(mfs)
}
