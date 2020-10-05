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
	"strings"
	"unicode"
)

type funcdecl struct {
	name       string
	paramCount int
	pos        token.Pos
}

type options struct {
	excludePatterns   []string
	excludeTests      bool
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

	issues, err := run(args, logger)
	if err != nil {
		os.Exit(invalidArgumentExitCode)
	} else if len(issues) > 0 {
		printAll(logger, issues)

		if *flagSetExitStatus {
			os.Exit(setExitStatusExitCode)
		}
	}
}

func run(paths []string, logger *log.Logger) (issues []string, err error) {
	if !(*flagVerbose) {
		logger = noopLogger
	}

	if noLimitIsSet(*flagMax, *flagPrivateMax, *flagPublicMax) {
		*flagMax = defaultParamLimit
	}

	excludePatterns := strings.Split(*flagExcludes, ",")
	if err := checkPatterns(excludePatterns); err != nil {
		logger.Printf("invalid pattern(s): %s", err)
	}

	baseOptions := &options{
		excludePatterns:   excludePatterns,
		excludeTests:      !(*flagTests),
		paramLimitPrivate: min(*flagMax, *flagPrivateMax),
		paramLimitPublic:  min(*flagMax, *flagPublicMax),
	}

	return runWith(paths, baseOptions, logger)
}

func runWith(
	paths []string,
	baseOptions *options,
	logger *log.Logger,
) (issues []string, err error) {
	for _, path := range paths {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return []string{}, fmt.Errorf("invalid path %s", path)
		}

		root, recursive := checkRecursive(absPath)
		baseOptions.recursive = recursive

		pathIssues := analyzeWith(root, baseOptions, logger)
		issues = append(issues, pathIssues...)
	}

	return issues, nil
}

func analyzeWith(
	path string,
	options *options,
	logger *log.Logger,
) (issues []string) {
	for _, filePath := range getFiles(path, options) {
		logger.Printf("analyzing %s", filePath)
		fileIssues, err := analyzeFile(filePath, options)
		if err != nil {
			logger.Printf("error parsing %s", filePath)
		} else {
			issues = append(issues, fileIssues...)
		}
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
