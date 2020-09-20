package main

import (
	"flag"
	"fmt"
)

var (
	flagIgnore = flag.Int(
		"max",
		defaultParamLimit,
		fmt.Sprintf("Maximum number of function parameters (default: %d)", defaultParamLimit),
	)

	flagSetExitStatus = flag.Bool(
		"set_exit_status",
		false,
		fmt.Sprintf("Set exit status to %d if any issues are found", setExistStatusExitCode),
	)
)
