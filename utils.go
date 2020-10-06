package main

import (
	"bytes"
	"fmt"
	"go/token"
	"log"
	"path/filepath"
	"strings"
)

var noopLogger = log.New(bytes.NewBuffer([]byte{}), "", 0)

func checkPatterns(patterns []string) error {
	var invalidPatterns []string
	for _, pattern := range patterns {
		_, err := filepath.Match(pattern, pattern)
		if err != nil {
			invalidPatterns = append(invalidPatterns, pattern)
		}
	}

	if len(invalidPatterns) > 0 {
		return fmt.Errorf("'%s'", strings.Join(invalidPatterns, "', '"))
	}

	return nil
}

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

func min(v0 int, vs ...int) int {
	m := v0
	for _, c := range vs {
		if c < m {
			m = c
		}
	}

	return m
}

func printAll(p printer, issues []string) {
	for _, issue := range issues {
		p.Print(issue)
	}
}
