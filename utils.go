package main

import (
	"fmt"
	"go/token"
)

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

func printAll(p printer, issues []string) {
	for _, issue := range issues {
		p.Print(issue)
	}
}
