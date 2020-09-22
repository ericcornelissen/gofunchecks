package main

import "testing"

func TestCheckRecursive(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		originalPath := ""
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("short path", func(t *testing.T) {
		originalPath := "./"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("long path", func(t *testing.T) {
		originalPath := "path/to/directory"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != originalPath {
			t.Errorf("Expected adjusted path to equal original path (was '%s')", adjustedPath)
		}

		if recursive == true {
			t.Error("Expected recursive to be false")
		}
	})
	t.Run("package list wildcard", func(t *testing.T) {
		originalPath := "./..."
		expectedPath := "./"
		adjustedPath, recursive := checkRecursive(originalPath)

		if adjustedPath != expectedPath {
			t.Errorf("Expected adjusted path to '%s' (was '%s')", expectedPath, adjustedPath)
		}

		if recursive == false {
			t.Error("Expected recursive to be true")
		}
	})
}

func TestIncludes(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		result := includes([]string{}, "foobar")
		if result == true {
			t.Error("Expected first result to be false")
		}

		result = includes([]string{}, "")
		if result == true {
			t.Error("Expected second result to be false")
		}
	})
	t.Run("slice not containing element", func(t *testing.T) {
		result := includes([]string{"foo"}, "bar")
		if result == true {
			t.Error("Expected first result to be false")
		}

		result = includes([]string{"foo", "bar"}, "baz")
		if result == true {
			t.Error("Expected second result to be false")
		}
	})
	t.Run("slice containing element", func(t *testing.T) {
		result := includes([]string{"foo"}, "foo")
		if result == false {
			t.Error("Expected first result to be true")
		}

		result = includes([]string{"foo", "bar"}, "foo")
		if result == false {
			t.Error("Expected second result to be true")
		}

		result = includes([]string{"foo", "bar"}, "bar")
		if result == false {
			t.Error("Expected second result to be true")
		}
	})
}
