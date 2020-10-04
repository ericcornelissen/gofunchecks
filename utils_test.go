package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func TestCheckRecursive(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		originalPath := ""
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("short path", func(t *testing.T) {
		originalPath := "./"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("long path", func(t *testing.T) {
		originalPath := "path/to/directory"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("package list wildcard", func(t *testing.T) {
		originalPath := "./..."
		expectedPath := "./"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != expectedPath {
			t.Errorf("Expected adjusted path to '%s' (was '%s')", expectedPath, adjustedPath)
		}

		if recursive == false {
			t.Error("Expected recursive to be true")
		}
	})
}

func TestConstructMessage(t *testing.T) {
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

	if len(file.Decls) < 1 {
		t.Fatal("Test file must contains at least one declaration")
	}

	funcName := "foobar"
	funcParamCount := 3
	funcDecl := &funcdecl{
		name:       funcName,
		paramCount: funcParamCount,
		pos:        file.Decls[0].Pos(),
	}

	result := constructMessage(fileSet, funcDecl)
	if !strings.Contains(result, funcName) {
		t.Error("Expected message to contain the function name")
	}

	if !strings.Contains(result, fmt.Sprintf("%d parameters", funcParamCount)) {
		t.Error("Expected message to contain the actual parameter count")
	}
}

func TestIncludes(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := includes([]string{}, "foobar")
		if result == true {
			t.Error("Expected first result to be false")
		}

		result = includes([]string{}, "")
		if result == true {
			t.Error("Expected second result to be false")
		}
	})
	t.Run("slice not containing element", func(t *testing.T) {
		result := includes([]string{"foo"}, "bar")
		if result == true {
			t.Error("Expected first result to be false")
		}

		result = includes([]string{"foo", "bar"}, "baz")
		if result == true {
			t.Error("Expected second result to be false")
		}
	})
	t.Run("slice containing element", func(t *testing.T) {
		result := includes([]string{"foo"}, "foo")
		if result == false {
			t.Error("Expected first result to be true")
		}

		result = includes([]string{"foo", "bar"}, "foo")
		if result == false {
			t.Error("Expected second result to be true")
		}

		result = includes([]string{"foo", "bar"}, "bar")
		if result == false {
			t.Error("Expected second result to be true")
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("a < b", func(t *testing.T) {
		a, b := 1, 2
		if !(a < b) {
			t.Fatal("For this test a must be less than b")
		}

		result := min(a, b)
		if result != a {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("a > b", func(t *testing.T) {
		a, b := 2, 1
		if !(a > b) {
			t.Fatal("For this test a must be greater than b")
		}

		result := min(a, b)
		if result != b {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("a == b", func(t *testing.T) {
		a, b := 2, 2
		if !(a == b) {
			t.Fatal("For this test a must be equal to b")
		}

		result := min(a, b)
		if result != a {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
}

func TestPrintAll(t *testing.T) {
	t.Run("no issues", func(t *testing.T) {
		var callCount uint
		p := mockPrinter{callCount: &callCount}

		noIssues := []string{}
		printAll(p, noIssues)

		if callCount > 0 {
			t.Errorf("Expected printer not to be called (called %d times)", callCount)
		}
	})
	t.Run("one issue", func(t *testing.T) {
		var calls [][]interface{}
		var callCount uint
		p := mockPrinter{
			callCount: &callCount,
			calls:     &calls,
		}

		issues := []string{"Hello world!"}
		printAll(p, issues)

		if callCount != 1 {
			t.Errorf("Expected printer to be called once (called %d times)", callCount)
		}

		actualCallValue := calls[0][0].(string)
		if actualCallValue != issues[0] {
			t.Errorf("Unexpected arguments (got '%s')", actualCallValue)
		}
	})
	t.Run("multiple issue", func(t *testing.T) {
		var calls [][]interface{}
		var callCount uint
		p := mockPrinter{
			callCount: &callCount,
			calls:     &calls,
		}

		issues := []string{"foo", "bar", "baz"}
		printAll(p, issues)

		if callCount != 3 {
			t.Errorf("Expected printer to be called thrice (called %d times)", callCount)
		}

		for i, expected := range issues {
			actual := calls[i][0].(string)
			if actual != expected {
				t.Errorf("Unexpected arguments in call %d (got '%s')", i, actual)
			}
		}
	})
}
