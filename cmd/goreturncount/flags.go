package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	flagExcludes = flag.String(
		"excludes",
		"",
		"Comma separated list of file patterns to exclude",
	)
	flagExcludesAlias = flag.String(
		"e",
		"",
		"(alias of -excludes)",
	)

	flagLegal = flag.Bool(
		"legal",
		false,
		"Show legal information about the program and exit",
	)

	flagMax = flag.Int(
		"max",
		math.MaxInt32,
		fmt.Sprintf("Maximum number of return values (default: %d)", defaultReturnLimit),
	)
	flagMaxAlias = flag.Int(
		"m",
		math.MaxInt32,
		"(alias of -max)",
	)

	flagPrivateMax = flag.Int(
		"private-max",
		math.MaxInt32,
		"Maximum number of return values for private functions",
	)
	flagPublicMax = flag.Int(
		"public-max",
		math.MaxInt32,
		"Maximum number of return values for public functions",
	)

	flagSetExitStatus = flag.Bool(
		"set_exit_status",
		false,
		fmt.Sprintf("Set exit status to %d if any issues are found", setExitStatusExitCode),
	)
	flagSetExitStatusAlias = flag.Bool(
		"S",
		false,
		"(alias of -set_exit_status)",
	)

	flagTests = flag.Bool(
		"tests",
		false,
		"Include test files in analysis",
	)
	flagTestsAlias = flag.Bool(
		"t",
		false,
		"(alias of -tests)",
	)

	flagVerbose = flag.Bool(
		"verbose",
		false,
		"Enable debug logging",
	)
	flagVerboseAlias = flag.Bool(
		"v",
		false,
		"(alias of -verbose)",
	)

	flagVersion = flag.Bool(
		"version",
		false,
		"Show the program version and exit",
	)
)

func noLimitIsSet(flags ...int) bool {
	for _, flag := range flags {
		if flag != math.MaxInt32 {
			return false
		}
	}

	return true
}
