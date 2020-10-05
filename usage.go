package main

import "fmt"

const usageDocTemplate = `goparamcount: find functions with too many parameters

Usage:

  goparamcount <flags> <directory> [<directory>...]

Flags:

  -max              Maximum number of parameters (default: %d).
  -public-max       Maximum number of parameters for public functions. Defaults
                      to the value of -max. Public functions are not analyzed if
                      only -private-max is set.
  -private-max      Maximum number of parameters for private functions. Defaults
                      to the value of -max. Private functions are not analyzed
                      if only -public-max is set.
  -excludes         Comma separated list of file patterns to exclude.
  -set_exit_status  Set exit status to %d if any issues are found.
  -verbose          Enable debug logging.

Examples:

  goparamcount ./...
  goparamcount -max 3 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -public-max 2 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -set_exit_status $GOPATH/src/github.com/cockroachdb/cockroach
`

func printUsage(p printer) {
	usageDoc := fmt.Sprintf(usageDocTemplate, defaultParamLimit, setExitStatusExitCode)
	p.Print(usageDoc)
}
