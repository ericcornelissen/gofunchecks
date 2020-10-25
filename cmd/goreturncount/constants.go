package main

// The default return value limit used by goreturncount.
const defaultReturnLimit = 2

// The non-zero exit codes used by goreturncount.
const (
	missingArgumentExitCode = iota + 1
	setExitStatusExitCode
	invalidArgumentExitCode
)
