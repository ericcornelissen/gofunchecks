package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const (
	missingArgumentExitCode = iota + 1
	setExistStatusExitCode
	invalidArgumentExitCode
)

const defaultParamLimit = 2

func main() {
	flag.Usage = printUsage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		if *flagHelp {
			os.Exit(0)
		} else {
			os.Exit(missingArgumentExitCode)
		}
	}

	anyIssues := run(args)
	if anyIssues && *flagSetExitStatus {
		os.Exit(setExistStatusExitCode)
	}
}

func run(args []string) bool {
	lintFailed := false
	for _, path := range args {
		root, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error finding absolute path: %s", err)
			os.Exit(invalidArgumentExitCode)
		}

		if walkPath(root, *flagMax) {
			lintFailed = true
		}
	}

	return lintFailed
}

func walkPath(root string, paramLimit int) bool {
	lintFailed := false

	pathLen := len(root)
	recursive := false
	if pathLen >= 5 && root[pathLen-3:] == "..." {
		recursive = true
		root = root[:pathLen-3]
	}

	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			return nil
		}

		if fi.IsDir() {
			if path != root && (!recursive ||
				filepath.Base(path) == "testdata" ||
				filepath.Base(path) == "vendor" ||
				filepath.Base(path) == ".git") {
				return filepath.SkipDir
			}

			return nil
		}

		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		issues, err := checkFile(path, paramLimit)
		for _, issue := range issues {
			fmt.Println(issue)
			lintFailed = true
		}

		return nil
	})
	return lintFailed
}

func checkFile(path string, paramLimit int) ([]string, error) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, path, nil, 0)
	if err != nil {
		return nil, err
	}

	return checkForParamLimit(fileSet, file, paramLimit), nil
}

func checkForParamLimit(fileSet *token.FileSet, file *ast.File, paramLimit int) []string {
	var issues []string
	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		paramCount := 0
		for _, param := range decl.Type.Params.List {
			// Multiple parameters with the same type, as in `func(a, b int)` will
			// appear as a single "param" in the `Params.List`. Therefore, we instead
			// count the number of `Names` per "param".
			paramCount += len(param.Names)
		}

		if paramCount > paramLimit {
			issue := fmt.Sprintf("%s - too many parameters in %s (%d > %d)",
				fileSet.Position(decl.Pos()),
				decl.Name,
				paramCount,
				paramLimit,
			)
			issues = append(issues, issue)
		}
	}

	return issues
}
