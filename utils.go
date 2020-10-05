package main

import (
	"bytes"
	"fmt"
	"go/token"
	"log"
)

var noopLogger = log.New(bytes.NewBuffer([]byte{}), "", 0)

func checkRecursive(path string) (adjustedPath string, recursive bool) {
	pathLen := len(path)
	if pathLen >= 5 && path[pathLen-3:] == "..." {
		recursive = true
		path = path[:pathLen-3]
	}

	return path, recursive
}

func constructMessage(fileSet *token.FileSet, funcdecl *funcdecl) string {
	return fmt.Sprintf("%s:%d - %d parameters in function '%s' is too many",
		fileSet.Position(funcdecl.pos).Filename,
		fileSet.Position(funcdecl.pos).Line,
		funcdecl.paramCount,
		funcdecl.name,
	)
}

func includes(vs []string, x string) bool {
	for _, v := range vs {
		if v == x {
			return true
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func printAll(p printer, issues []string) {
	for _, issue := range issues {
		p.Print(issue)
	}
}
