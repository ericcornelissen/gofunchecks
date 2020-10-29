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

	"github.com/ericcornelissen/gofunchecks/internal/utils"
	"github.com/ericcornelissen/gofunchecks/internal/walk"
)

type funcdecl struct {
	name        string
	returnCount int
	pos         token.Pos
}

type printer interface {
	Print(msgs ...interface{})
}

func main() {
	logger := log.New(os.Stdout, "goreturncount: ", 0)

	flag.Usage = func() { printUsage(logger) }
	flag.Parse()

	if *flagLegal {
		printLegal(logger)
		os.Exit(0)
	}

	if *flagVersion {
		printVersion(logger)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		printUsage(logger)
		os.Exit(missingArgumentExitCode)
	}

	issues, err := run(args, logger)
	if err != nil {
		os.Exit(invalidArgumentExitCode)
	} else if len(issues) > 0 {
		utils.PrintAll(logger, issues)

		if *flagSetExitStatus || *flagSetExitStatusAlias {
			os.Exit(setExitStatusExitCode)
		}
	}
}

func run(paths []string, logger *log.Logger) (issues []string, err error) {
	if !(*flagVerbose || *flagVerboseAlias) {
		logger = utils.NoopLogger
	}

	if noLimitIsSet(*flagMax, *flagMaxAlias, *flagPrivateMax, *flagPublicMax) {
		*flagMax = defaultReturnLimit
	}

	*flagExcludes += *flagExcludesAlias
	excludePatterns := strings.Split(*flagExcludes, ",")
	if err := utils.CheckPatterns(excludePatterns); err != nil {
		logger.Printf("invalid pattern(s): %s", err)
	}

	baseOptions := &options{
		excludePatterns:    excludePatterns,
		excludeTests:       !(*flagTests || *flagTestsAlias),
		returnLimitPrivate: utils.Min(*flagMax, *flagMaxAlias, *flagPrivateMax),
		returnLimitPublic:  utils.Min(*flagMax, *flagMaxAlias, *flagPublicMax),
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

		root, recursive := utils.CheckRecursive(absPath)
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
	for _, filePath := range walk.GetFiles(path, options) {
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

	for _, decl := range checkForReturnLimit(file, options) {
		issues = append(issues, constructMessage(fileSet, decl))
	}

	return issues, nil
}

func checkForReturnLimit(file *ast.File, options *options) (issues []*funcdecl) {
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

	returnLimit := options.returnLimitPrivate
	if utils.IsPublicFunc(decl) {
		returnLimit = options.returnLimitPublic
	}

	returnCount := getReturnCount(decl)
	if returnCount <= returnLimit {
		return nil
	}

	return &funcdecl{
		name:        decl.Name.String(),
		returnCount: returnCount,
		pos:         decl.Pos(),
	}
}

func getReturnCount(decl *ast.FuncDecl) int {
	if decl.Type.Results != nil {
		return decl.Type.Results.NumFields()
	}

	return 0
}
