package main

import (
	"fmt"
	"go/token"
)

func constructMessage(fileSet *token.FileSet, funcdecl *funcdecl) string {
	return fmt.Sprintf("%s:%d - function '%s' has too many (%d) return values",
		fileSet.Position(funcdecl.pos).Filename,
		fileSet.Position(funcdecl.pos).Line,
		funcdecl.name,
		funcdecl.returnCount,
	)
}
