package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type fixture struct {
		data     *[]int
		expected *[]int
	}

	fixtures := []fixture{
		{
			data:     link([]int{}),
			expected: link([]int{}),
		},
		{
			data:     link([]int{1}),
			expected: link([]int{1}),
		},
		{
			data:     link([]int{1, 2}),
			expected: link([]int{2, 1}),
		},
		{
			data:     link([]int{1, 2, 3}),
			expected: link([]int{3, 2, 1}),
		},
		{
			data:     link([]int{1, 2, 3, 4}),
			expected: link([]int{4, 3, 2, 1}),
		},
		{
			data:     link([]int{3, 2, 1}),
			expected: link([]int{1, 2, 3}),
		},
		{
			data:     link([]int{2, 2, 2}),
			expected: link([]int{2, 2, 2}),
		},
		{
			data:     link([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}),
			expected: link([]int{16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}),
		},
	}

	for _, fxt := range fixtures {
		reverse(fxt.data)

		if !equal(fxt.data, fxt.expected) {
			t.Errorf("Expected: %v, Actual: %v", fxt.expected, fxt.data)
		}
	}
}

func equal(a, b *[]int) bool {
	if len(*a) != len(*b) {
		return false
	}

	for i := range *a {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}

	return true
}

func link(s []int) *[]int {
	return &s
}
