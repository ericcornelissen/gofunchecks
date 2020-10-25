package utils

import (
	"go/ast"
	"testing"
)

func TestCheckPatterns(t *testing.T) {
	t.Run("no patterns", func(t *testing.T) {
		err := CheckPatterns([]string{})
		if err != nil {
			t.Errorf("Unexpected error (got '%s')", err)
		}
	})
	t.Run("no invalid patterns", func(t *testing.T) {
		err := CheckPatterns([]string{"valid pattern"})
		if err != nil {
			t.Errorf("Unexpected error (got '%s')", err)
		}
	})
	t.Run("some invalid patterns", func(t *testing.T) {
		err := CheckPatterns([]string{"invalid[pattern"})
		if err == nil {
			t.Error("Expected an error but got none")
		}
	})
}

func TestCheckRecursive(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		originalPath := ""
		adjustedPath, recursive := CheckRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("short path", func(t *testing.T) {
		originalPath := "./"
		adjustedPath, recursive := CheckRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("long path", func(t *testing.T) {
		originalPath := "path/to/directory"
		adjustedPath, recursive := CheckRecursive(originalPath)

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
		adjustedPath, recursive := CheckRecursive(originalPath)

		if adjustedPath != expectedPath {
			t.Errorf("Expected adjusted path to '%s' (was '%s')", expectedPath, adjustedPath)
		}

		if recursive == false {
			t.Error("Expected recursive to be true")
		}
	})
}

func TestIsPublicFunc(t *testing.T) {
	t.Run("private function", func(t *testing.T) {
		decl := &ast.FuncDecl{
			Name: ast.NewIdent("localFunction"),
		}

		result := IsPublicFunc(decl)
		if result == true {
			t.Error("The function declaration is not public")
		}
	})
	t.Run("public function", func(t *testing.T) {
		decl := &ast.FuncDecl{
			Name: ast.NewIdent("PublicFunction"),
		}

		result := IsPublicFunc(decl)
		if result == false {
			t.Error("The function declaration is public")
		}
	})
	t.Run("unconventional function name", func(t *testing.T) {
		decl := &ast.FuncDecl{
			Name: ast.NewIdent("_localFunction"),
		}

		result := IsPublicFunc(decl)
		if result == true {
			t.Error("The function declaration is not public")
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("a < b", func(t *testing.T) {
		a, b := 1, 2
		if !(a < b) {
			t.Fatal("For this test a must be less than b")
		}

		result := Min(a, b)
		if result != a {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("a > b", func(t *testing.T) {
		a, b := 2, 1
		if !(a > b) {
			t.Fatal("For this test a must be greater than b")
		}

		result := Min(a, b)
		if result != b {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("a == b", func(t *testing.T) {
		a, b := 2, 2
		if !(a == b) {
			t.Fatal("For this test a must be equal to b")
		}

		result := Min(a, b)
		if result != a {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("only one value", func(t *testing.T) {
		v := 4
		result := Min(v)
		if result != v {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("more than two values", func(t *testing.T) {
		a, b, c := 1, 2, 3
		if !(a < b) || !(a < c) {
			t.Fatal("The value of a must be the lowest for this test")
		}

		result := Min(a, b, c)
		if result != a {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("only one value", func(t *testing.T) {
		v := 4
		result := Min(v)
		if result != v {
			t.Errorf("Unexpected result (got %d)", result)
		}
	})
	t.Run("more than two values", func(t *testing.T) {
		a, b, c := 1, 2, 3
		if !(a < b) || !(a < c) {
			t.Fatal("The value of a must be the lowest for this test")
		}

		result := Min(a, b, c)
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
		PrintAll(p, noIssues)

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
		PrintAll(p, issues)

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
		PrintAll(p, issues)

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
