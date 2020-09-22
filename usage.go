package main

import "fmt"

const usageDocTemplate = `goparamcount: find functions with too many parameters

Usage:

  goparamcount <flags> <directory> [<directory>...]

Flags:

  -max               Maximum number of parameters (default: %d)
  -set_exit_status   Set exit status to %d if any issues are found

Examples:

  goparamcount ./...
  goparamcount -max 3 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -set_exit_status $GOPATH/src/github.com/cockroachdb/cockroach
`

func printUsage(p printer) {
	usageDoc := fmt.Sprintf(usageDocTemplate, defaultParamLimit, setExitStatusExitCode)
	p.Print(usageDoc)
}
