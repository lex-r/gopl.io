package unique

import (
	"testing"
)

func TestUnique(t *testing.T) {
	type fixture struct {
		data     []string
		expected []string
	}

	fixtures := []fixture{
		{
			data:     []string{},
			expected: []string{},
		},
		{
			data:     []string{"a"},
			expected: []string{"a"},
		},
		{
			data:     []string{"a", "b"},
			expected: []string{"a", "b"},
		},
		{
			data:     []string{"a", "a"},
			expected: []string{"a", ""},
		},
		{
			data:     []string{"a", "a", "a"},
			expected: []string{"a", "", ""},
		},
		{
			data:     []string{"a", "b", "a"},
			expected: []string{"a", "b", "a"},
		},
		{
			data:     []string{"a", "a", "b"},
			expected: []string{"a", "b", ""},
		},
		{
			data:     []string{"a", "a", "b", "a"},
			expected: []string{"a", "b", "a", ""},
		},
		{
			data:     []string{"a", "a", "a"},
			expected: []string{"a", "", ""},
		},
		{
			data:     []string{"a", "a", "a", "b"},
			expected: []string{"a", "b", "", ""},
		},
		{
			data:     []string{"a", "b", "b", "a"},
			expected: []string{"a", "b", "a", ""},
		},
	}

	for _, fxt := range fixtures {
		unique(fxt.data)

		if !equal(fxt.data, fxt.expected) {
			t.Errorf("Expected: %v, Actual: %v", fxt.expected, fxt.data)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
