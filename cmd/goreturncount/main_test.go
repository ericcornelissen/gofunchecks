package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestAnalyzeFile(t *testing.T) {
	t.Run("without issues", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 100,
			returnLimitPublic:  100,
		}

		issues, err := analyzeFile("../../testdata/src/return.go", options)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		issueCount := len(issues)
		if issueCount != 0 {
			t.Errorf("Expected no issues (got %d)", issueCount)
		}
	})
	t.Run("with issues", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 0,
			returnLimitPublic:  0,
		}

		issues, err := analyzeFile("../../testdata/src/return.go", options)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		issueCount := len(issues)
		if issueCount != 2 {
			t.Errorf("Expected two issues (got %d)", issueCount)
		}
	})
	t.Run("file does not exists", func(t *testing.T) {
		options := &options{}

		_, err := analyzeFile("this is definitely not a file!", options)
		if err == nil {
			t.Error("Expected an error but got none")
		}
	})
}

func TestCheckForReturnLimit(t *testing.T) {
	t.Run("no issues, below limit", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 2,
			returnLimitPublic:  2,
		}

		src := `
			package foo

			func localFunction(a, b int) int {
				return a + b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("no issues, at limit", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a, b int) int {
				return a + b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issue (got %d)", len(issues))
		}
	})
	t.Run("too many distinct return values", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a int, b string) (int, string) {
				return a + len(b), b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many distinct named return values", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a int, b string) (c int, d string) {
				return a + len(b), b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many return values of one type", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a, b, c int) (int, int) {
				return a, c
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many named return values of one type", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a, b, c int) (d int, f int) {
				return a, c
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many return values with shared type", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 1,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a int) (a, b int) {
				return a + 1, a + 2
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("separate limit for public and private functions", func(t *testing.T) {
		options := &options{
			returnLimitPrivate: 2,
			returnLimitPublic:  1,
		}

		src := `
			package foo

			func localFunctionFoo(a int, b uint) int {
				return a + int(b)
			}

			func localFunctionBar(a int) (bool, int, string) {
				return a > 3, a + 2, str(a)
			}

			func PublicFunctionFoo(a int) int {
				return a + 1
			}

			func PublicFunctionBar(a int, b string) (bool, string) {
				return len(b) > a, b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForReturnLimit(file, options)
		if len(issues) != 2 {
			t.Errorf("Expected two issues (got %d)", len(issues))
		}
	})
}
