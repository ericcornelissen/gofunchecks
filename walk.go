package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getFiles(root string, options *options) (paths []string) {
	filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error during filesystem walk: %v\n", err)
			return nil
		}

		if path == root {
			return nil
		}

		if fi.IsDir() {
			return skipDir(path, options.recursive)
		}

		if skipFile(path) {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	return paths
}

func skipDir(path string, recursive bool) error {
	if !recursive {
		return filepath.SkipDir
	}

	ignoreDirs := []string{".git", "testdata", "vendor"}
	if includes(ignoreDirs, filepath.Base(path)) {
		return filepath.SkipDir
	}

	return nil
}

func skipFile(path string) bool {
	if !strings.HasSuffix(path, ".go") {
		return true
	}

	if strings.HasSuffix(path, "_test.go") {
		return true
	}

	return false
}
