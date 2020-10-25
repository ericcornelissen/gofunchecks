package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

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
	funcReturnCount := 3
	funcDecl := &funcdecl{
		name:        funcName,
		returnCount: funcReturnCount,
		pos:         file.Decls[0].Pos(),
	}

	result := constructMessage(fileSet, funcDecl)
	if !strings.Contains(result, funcName) {
		t.Error("Expected message to contain the function name")
	}

	if !strings.Contains(result, fmt.Sprintf("(%d)", funcReturnCount)) {
		t.Error("Expected message to contain the actual return values count")
	}
}
