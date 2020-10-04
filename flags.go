package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	flagMax = flag.Int(
		"max",
		math.MaxInt32,
		fmt.Sprintf("Maximum number of parameters (default: %d)", defaultParamLimit),
	)
	flagPrivateMax = flag.Int(
		"private-max",
		math.MaxInt32,
		"Maximum number of parameters for private functions",
	)
	flagPublicMax = flag.Int(
		"public-max",
		math.MaxInt32,
		"Maximum number of parameters for public functions",
	)

	flagSetExitStatus = flag.Bool(
		"set_exit_status",
		false,
		fmt.Sprintf("Set exit status to %d if any issues are found", setExitStatusExitCode),
	)
)

func noLimitIsSet() bool {
	return *flagMax == math.MaxInt32 &&
		*flagPrivateMax == math.MaxInt32 &&
		*flagPublicMax == math.MaxInt32
}
