package walk

// Options is an interface for the options when calling walk.GetFiles.
type Options interface {
	// Should the walk be recursive. I.e. should (sub)directories be explored.
	Recursive() bool

	// Should tests be excluded or not.
	ExcludeTests() bool

	// The pattern(s) of files to be excluded.
	ExcludePatterns() []string
}
