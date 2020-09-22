package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCheckForParamLimit(t *testing.T) {
	t.Run("no issues", func(t *testing.T) {
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

		issues := checkForParamLimit(file, 2)
		if len(issues) != 0 {
			t.Errorf("Expected zero issues (got %d)", len(issues))
		}
	})
	t.Run("no issues with variadic function", func(t *testing.T) {
		src := `
			package foo

			func localFunction(a ...int) int {
				return a + b + c
			}
		`

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, "", src, 0)
		if err != nil {
			t.Fatal("Test file could not be parsed")
		}

		issues := checkForParamLimit(file, 1)
		if len(issues) != 0 {
			t.Errorf("Expected zero issue (got %d)", len(issues))
		}
	})
	t.Run("too many distinct parameters", func(t *testing.T) {
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

		issues := checkForParamLimit(file, 1)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("too many paramters of one type", func(t *testing.T) {
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

		issues := checkForParamLimit(file, 1)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
	t.Run("variadic function with too many parameters", func(t *testing.T) {
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

		issues := checkForParamLimit(file, 1)
		if len(issues) != 1 {
			t.Errorf("Expected one issue (got %d)", len(issues))
		}
	})
}
