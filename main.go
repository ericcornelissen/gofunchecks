package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

type funcdecl struct {
	name       string
	paramCount int
	pos        token.Pos
}

type options struct {
	paramLimitPrivate int
	paramLimitPublic  int
	recursive         bool
}

type printer interface {
	Print(msgs ...interface{})
}

func main() {
	logger := log.New(os.Stdout, "goparamcount: ", 0)

	flag.Usage = func() { printUsage(logger) }
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		printUsage(logger)
		os.Exit(missingArgumentExitCode)
	}

	anyIssues, err := run(args, logger)
	if err != nil {
		os.Exit(invalidArgumentExitCode)
	} else if anyIssues && *flagSetExitStatus {
		os.Exit(setExitStatusExitCode)
	}
}

func run(paths []string, logger *log.Logger) (anyIssues bool, err error) {
	if noLimitIsSet(*flagMax, *flagPrivateMax, *flagPublicMax) {
		*flagMax = defaultParamLimit
	}

	var issues []string
	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return false, fmt.Errorf("invalid path %s", path)
		}

		pathIssues := analyze(absPath)
		issues = append(issues, pathIssues...)
	}

	printAll(logger, issues)
	return len(issues) > 0, nil
}

func analyze(path string) (issues []string) {
	root, recursive := checkRecursive(path)
	options := &options{
		paramLimitPrivate: min(*flagMax, *flagPrivateMax),
		paramLimitPublic:  min(*flagMax, *flagPublicMax),
		recursive:         recursive,
	}

	return analyzeWith(root, options)
}

func analyzeWith(path string, options *options) (issues []string) {
	for _, file := range getFiles(path, options) {
		fileIssues, err := analyzeFile(file, options)
		if err != nil {
			fileIssues = []string{err.Error()}
		}

		issues = append(issues, fileIssues...)
	}

	return issues
}

func analyzeFile(path string, options *options) (issues []string, err error) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, path, nil, 0)
	if err != nil {
		return nil, err
	}

	for _, decl := range checkForParamLimit(file, options) {
		issues = append(issues, constructMessage(fileSet, decl))
	}

	return issues, nil
}

func checkForParamLimit(file *ast.File, options *options) (issues []*funcdecl) {
	for _, decl := range file.Decls {
		issue := checkDecl(decl, options)
		if issue != nil {
			issues = append(issues, issue)
		}
	}

	return issues
}

func checkDecl(d ast.Decl, options *options) *funcdecl {
	decl, ok := d.(*ast.FuncDecl)
	if !ok {
		return nil
	}

	paramLimit := options.paramLimitPrivate
	if isPublicFunc(decl) {
		paramLimit = options.paramLimitPublic
	}

	paramCount := getParamCount(decl)
	if paramCount <= paramLimit {
		return nil
	}

	return &funcdecl{
		name:       decl.Name.String(),
		paramCount: paramCount,
		pos:        decl.Pos(),
	}
}

func isPublicFunc(decl *ast.FuncDecl) bool {
	name := []rune(decl.Name.String())
	return unicode.IsUpper(name[0])
}

func getParamCount(decl *ast.FuncDecl) int {
	paramCount := 0
	for _, param := range decl.Type.Params.List {
		// Multiple parameters with the same type, as in `func(a, b int)` will
		// appear as a single "param" in the `Params.List`. Therefore, we instead
		// count the number of `Names` per "param".
		paramCount += len(param.Names)
	}

	return paramCount
}
