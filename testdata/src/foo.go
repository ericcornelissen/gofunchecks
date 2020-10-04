// Package foo exists for the purpose of testing goparamcount. This file is
// expected to have a private function with two parameters and a public function
// with two parameters.
package foo

var hello = "world"

func localFunction(a string, b int) bool {
	return len(a) < b
}

// PublicFunction is ONLY for testing purposes.
func PublicFunction(a string, b int) bool {
	return len(a) > b
}
