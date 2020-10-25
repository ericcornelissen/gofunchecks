package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"log"
	"path/filepath"
	"strings"
	"unicode"
)

// NoopLogger is a logger that won't log anything when called.s
var NoopLogger = log.New(bytes.NewBuffer([]byte{}), "", 0)

// CheckPatterns validates -exclude patterns.
func CheckPatterns(patterns []string) error {
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

// CheckRecursive checks if a provided path must be analyzed recursively. I.e.
// if it ends with the substring "...". If so, the first return value will be
// the path with "..." removed (else it will be exactly the provided input).
func CheckRecursive(path string) (adjustedPath string, recursive bool) {
	pathLen := len(path)
	if pathLen >= 5 && path[pathLen-3:] == "..." {
		recursive = true
		path = path[:pathLen-3]
	}

	return path, recursive
}

// IsPublicFunc checks whether a function deceleration is for a public function.
func IsPublicFunc(decl *ast.FuncDecl) bool {
	name := []rune(decl.Name.String())
	return unicode.IsUpper(name[0])
}

// Min returns the minimum value out of n integers.
func Min(v0 int, vs ...int) int {
	m := v0
	for _, c := range vs {
		if c < m {
			m = c
		}
	}

	return m
}

// PrintAll will print all messages using the provided Printer.
func PrintAll(p Printer, issues []string) {
	for _, issue := range issues {
		p.Print(issue)
	}
}
