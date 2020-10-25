package main

import (
	"fmt"
	"go/token"
)

func constructMessage(fileSet *token.FileSet, funcdecl *funcdecl) string {
	return fmt.Sprintf("%s:%d - %d parameters in function '%s' is too many",
		fileSet.Position(funcdecl.pos).Filename,
		fileSet.Position(funcdecl.pos).Line,
		funcdecl.paramCount,
		funcdecl.name,
	)
}
