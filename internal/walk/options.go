package walk

type Options interface {
	Recursive() bool
	ExcludeTests() bool
	ExcludePatterns() []string
}
