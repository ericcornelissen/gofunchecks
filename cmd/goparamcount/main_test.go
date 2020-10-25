package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestAnalyzeFile(t *testing.T) {
	t.Run("without issues", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 100,
			paramLimitPublic:  100,
		}

		issues, err := analyzeFile("../../testdata/src/param.go", options)
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
			paramLimitPrivate: 0,
			paramLimitPublic:  0,
		}

		issues, err := analyzeFile("../../testdata/src/param.go", options)
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

func TestCheckForParamLimitNoIssues(t *testing.T) {
	t.Run("below limit", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 2,
			paramLimitPublic:  2,
		}

		src := `
			package foo

			func localFunction(a int) int {
				return a + 1
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("at limit", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 2,
			paramLimitPublic:  2,
		}

		src := `
			package foo

			func localFunction(a string, b int) int {
				return len(a) + b
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("below limit, shared type", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 3,
			paramLimitPublic:  3,
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

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("at limit, shared type", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 2,
			paramLimitPublic:  2,
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

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("below limit, variadic function", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 2,
			paramLimitPublic:  2,
		}

		src := `
			package foo

			func localFunction(a ...int) int {
				return len(a)
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issue (got %d)", len(issues))
		}
	})
	t.Run("at limit, variadic function", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 1,
			paramLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a ...int) int {
				return len(a)
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 0 {
			t.Errorf("Expected zero issue (got %d)", len(issues))
		}
	})
}

func TestCheckForParamLimit(t *testing.T) {
	t.Run("too many distinct parameters", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 1,
			paramLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a int, b string) int {
				return a + len(b)
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many parameters of one type", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 1,
			paramLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a, b, c int) int {
				return a + b + c
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("variadic function with too many parameters", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 1,
			paramLimitPublic:  1,
		}

		src := `
			package foo

			func localFunction(a int, b ...int) int {
				return a + b + c
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("separate limit for public and private functions", func(t *testing.T) {
		options := &options{
			paramLimitPrivate: 2,
			paramLimitPublic:  1,
		}

		src := `
			package foo

			func localFunctionFoo(a int, b uint) int {
				return a + int(b)
			}

			func localFunctionBar(a int, b uint, c string) bool {
				return len(c) > localFunctionFoo(a, b)
			}

			func PublicFunctionFoo(a int) int {
				return a + 1
			}

			func PublicFunctionBar(a int, b string) bool {
				return len(b) > a
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, options)
		if len(issues) != 2 {
			t.Errorf("Expected two issues (got %d)", len(issues))
		}
	})
}
