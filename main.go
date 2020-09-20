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
	missingArgumentExitCode = iota
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
		os.Exit(missingArgumentExitCode)
	}

	anyIssues := run(args)
	if anyIssues && *flagSetExitStatus {
		os.Exit(setExistStatusExitCode)
	}
}

func run(args []string) bool {
	lintFailed := false

	paramLimit := defaultParamLimit
	for _, path := range args {
		root, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error finding absolute path: %s", err)
			os.Exit(invalidArgumentExitCode)
		}

		if walkPath(root, paramLimit) {
			lintFailed = true
		}
	}

	return lintFailed
}

func walkPath(root string, paramLimit int) bool {
	lintFailed := false
	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			return nil
		}

		if fi.IsDir() {
			if path != root && (filepath.Base(path) == "testdata" ||
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

	return checkForParamLimit(file, paramLimit), nil
}

func checkForParamLimit(file *ast.File, paramLimit int) []string {
	var issues []string
	for _, decl := range file.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		params := decl.Type.Params.List
		paramCount := len(params)
		if paramCount > paramLimit {
			issues = append(issues, fmt.Sprintf("to many parameters in %s, %d > %d", decl.Name, paramCount, paramLimit))
		}
	}

	return issues
}
