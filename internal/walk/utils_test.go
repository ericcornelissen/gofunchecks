package walk

import "testing"

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
