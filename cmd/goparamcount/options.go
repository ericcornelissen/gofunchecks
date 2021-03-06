package main

type options struct {
	excludePatterns   []string
	excludeTests      bool
	paramLimitPrivate int
	paramLimitPublic  int
	recursive         bool
}

func (o options) Recursive() bool {
	return o.recursive
}

func (o options) ExcludeTests() bool {
	return o.excludeTests
}

func (o options) ExcludePatterns() []string {
	return o.excludePatterns
}
