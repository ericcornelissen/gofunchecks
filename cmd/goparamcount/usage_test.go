package main

import (
	"testing"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

func TestPrintUsage(t *testing.T) {
	p := utils.NoopLogger
	printUsage(p)
}
