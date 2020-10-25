// Package testdata exists for the purpose of testing gofunchecks.
//
// This file exists for the purpose of testing goparamcount. It is expected to
// have a private function with two parameters and a public function with two
// parameters.
package testdata

var hello = "world"

func localFunction(a string, b int) bool {
	return len(a) < b
}

// PublicFunction is ONLY for testing purposes.
func PublicFunction(a string, b int) bool {
	return len(a) > b
}
