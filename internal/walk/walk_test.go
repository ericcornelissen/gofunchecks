package walk

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
	excludeTests := true
	var noExcludePatterns []string

	t.Run("non-.go file", func(t *testing.T) {
		result := skipFile("file.txt", excludeTests, noExcludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
	t.Run(".go file", func(t *testing.T) {
		result := skipFile("file.go", excludeTests, noExcludePatterns)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("_test.go file", func(t *testing.T) {
		result := skipFile("file_test.go", true, noExcludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}

		result = skipFile("file_test.go", false, noExcludePatterns)
		if result == true {
			t.Error("Expected result to be false")
		}
	})
	t.Run("custom exclude pattern", func(t *testing.T) {
		excludePatterns := []string{"foo*.go"}
		result := skipFile("foobar.go", excludeTests, excludePatterns)
		if result == false {
			t.Error("Expected result to be true")
		}
	})
}
