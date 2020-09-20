package main

import "fmt"

const usageDoc = `goparamcount: find functions with to many parameters

Usage:

  goparamcount <flags> <directory> [<directory>...]

Flags:

  -max               Maximum number of parameters (default: 2)
  -set_exit_status   Set exit status to 2 if any issues are found

Examples:

  goparamcount ./...
  goparamcount -max 3 $GOPATH/src/github.com/cockroachdb/cockroach/...
  goparamcount -set_exit_status $GOPATH/src/github.com/cockroachdb/cockroach
`

func printUsage() {
	fmt.Print(usageDoc)
}
