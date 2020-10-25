// Package testdata exists for the purpose of testing gofunchecks.
//
// This file exists for the purpose of testing goreturncount. It is expected to
// have a private function with two return values and a public function with two
// return values.
package testdata

var foo = "bar"

func localFunction() (bool, string) {
	return false, foo
}

// PublicFunction is ONLY for testing purposes.
func PublicFunction(a string, b int) (c bool, d string) {
	return len(a) > 1, a + string(b)
}
