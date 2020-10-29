package main

import (
	"fmt"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

const version = "1.0.1"

const versionMessageTemplate = "version %s"

func printVersion(p utils.Printer) {
	versionMessage := fmt.Sprintf(versionMessageTemplate, version)
	p.Print(versionMessage)
}
