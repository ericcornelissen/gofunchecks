package main

// The default parameter limit used by goparamcount.
const defaultParamLimit = 2

// The non-zero exit codes used by goparamcount.
const (
	missingArgumentExitCode = iota + 1
	setExitStatusExitCode
	invalidArgumentExitCode
)
