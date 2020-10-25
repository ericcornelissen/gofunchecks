package main

import "fmt"

const usageDocTemplate = `find functions with too many return values

Usage:

  goreturncount <flags> <directory> [<directory>...]

Flags:

  -m, -max              Maximum number of return values (default: %d).
      -public-max       Maximum number of return values for public functions.
                          Defaults to the value of -max. The analysis won't
                          cover public functions if only -private-max is set.
      -private-max      Maximum number of return values for private functions.
                          Defaults to the value of -max. The analysis won't
                          cover private functions if only -public-max is set.
  -e, -excludes         Comma separated list of file patterns to exclude.
  -S, -set_exit_status  Set exit status to %d if any issues are found.
  -t, -tests            Include test files in analysis.
  -v, -verbose          Enable debug logging.

Examples:

  goreturncount ./...
  goreturncount -max 3 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goreturncount -public-max 2 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goreturncount -set_exit_status $GOPATH/src/github.com/cockroachdb/cockroach
`

func printUsage(p printer) {
	usageDoc := fmt.Sprintf(usageDocTemplate, defaultReturnLimit, setExitStatusExitCode)
	p.Print(usageDoc)
}
