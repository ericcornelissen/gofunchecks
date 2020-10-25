package main

import (
	"fmt"

	"github.com/ericcornelissen/gofunchecks/internal/utils"
)

const usageDocTemplate = `find functions with too many parameters

Usage:

  goparamcount <flags> <directory> [<directory>...]

Flags:

  -m, -max              Maximum number of parameters (default: %d).
      -public-max       Maximum number of parameters for public functions.
                          Defaults to the value of -max. The analysis won't
                          cover public functions if only -private-max is set.
      -private-max      Maximum number of parameters for private functions.
                          Defaults to the value of -max. The analysis won't
                          cover private functions if only -public-max is set.
  -e, -excludes         Comma separated list of file patterns to exclude.
  -S, -set_exit_status  Set exit status to %d if any issues are found.
  -t, -tests            Include test files in analysis.
  -v, -verbose          Enable debug logging.

Examples:

  goparamcount ./...
  goparamcount -max 3 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -public-max 2 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -set_exit_status $GOPATH/src/github.com/cockroachdb/cockroach
`

func printUsage(p utils.Printer) {
	usageDoc := fmt.Sprintf(usageDocTemplate, defaultParamLimit, setExitStatusExitCode)
	p.Print(usageDoc)
}
