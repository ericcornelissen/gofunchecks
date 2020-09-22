package main

import (
	"path/filepath"
	"testing"
)

func TestSkipDir(t *testing.T) {
	t.Run("recursive disabled", func(t *testing.T) {
		result := skipDir("./directory", false)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}
	})
	t.Run("random directory (recursive)", func(t *testing.T) {
		result := skipDir("./directory", true)
		if result != nil {
			t.Error("Expected result to be `nil`")
		}
	})
	t.Run("directories always excluded (recursive)", func(t *testing.T) {
		result := skipDir("./.git", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}

		result = skipDir("./testdata", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}

		result = skipDir("./vendor", true)
		if result != filepath.SkipDir {
			t.Error("Expected result to be `filepath.SkipDir`")
		}
	})
}

func TestSkipFile(t *testing.T) {
	t.Run("non-.go file", func(t *testing.T) {
		result := skipFile("file.txt")
		if result == false {
			t.Error("Expected result to be true")
		}
	})
	t.Run(".go file", func(t *testing.T) {
		result := skipFile("file.go")
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("_test.go file", func(t *testing.T) {
		result := skipFile("file_test.go")
		if result == false {
			t.Error("Expected result to be true")
		}
	})
}
