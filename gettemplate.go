package main

import (
	_ "embed"
	"io"
	"net/http"
)

type TemplateType int

const (
	Makefile_C TemplateType = iota
	C_Mini
)

func getTemplate(tt TemplateType) string {
	embedtemplate, templateURL := getTemplateTypeString(&tt)

	body, err := http.Get(templateURL)
	if err != nil || body.StatusCode != http.StatusOK {
		return embedtemplate
	} else {
		bodyBytes, err := io.ReadAll(body.Body)
		if err != nil {
			return embedtemplate
		} else {
			return string(bodyBytes)
		}
	}
}

func getTemplateTypeString(tt *TemplateType) (string, string) {
	var embedtemplate string
	var templateURL string
	switch *tt {
	case Makefile_C:
		embedtemplate = makefilec
		templateURL = "https://raw.githubusercontent.com/OceanOC/projxgen/refs/heads/main/templates/Makefile-C"
	case C_Mini:
		embedtemplate = cmini
		templateURL = "https://raw.githubusercontent.com/OceanOC/projxgen/refs/heads/main/templates/code/C-Mini.c"

	}

	return embedtemplate, templateURL
}

//go:embed templates/Makefile-C
var makefilec string

//go:embed templates/code/C-Mini.c
var cmini string
