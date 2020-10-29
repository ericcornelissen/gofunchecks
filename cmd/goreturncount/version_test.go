package main

import (
	"testing"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

func TestPrintVersion(t *testing.T) {
	p := utils.NoopLogger
	printVersion(p)
}
