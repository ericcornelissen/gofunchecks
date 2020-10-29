package main

import (
	"testing"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

func TestPrintLegal(t *testing.T) {
	p := utils.NoopLogger
	printLegal(p)
}
