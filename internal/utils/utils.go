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

var NoopLogger = log.New(bytes.NewBuffer([]byte{}), "", 0)

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

func CheckRecursive(path string) (adjustedPath string, recursive bool) {
	pathLen := len(path)
	if pathLen >= 5 && path[pathLen-3:] == "..." {
		recursive = true
		path = path[:pathLen-3]
	}

	return path, recursive
}

func IsPublicFunc(decl *ast.FuncDecl) bool {
	name := []rune(decl.Name.String())
	return unicode.IsUpper(name[0])
}

func Min(v0 int, vs ...int) int {
	m := v0
	for _, c := range vs {
		if c < m {
			m = c
		}
	}

	return m
}

func PrintAll(p Printer, issues []string) {
	for _, issue := range issues {
		p.Print(issue)
	}
}
