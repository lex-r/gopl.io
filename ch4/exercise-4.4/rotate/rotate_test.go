package rotate

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type fixture struct {
		data     []int
		shift    int
		expected []int
	}

	fixtures := []fixture{
		{
			data:     []int{},
			shift:    0,
			expected: []int{},
		},
		{
			data:     []int{},
			shift:    1,
			expected: []int{},
		},
		{
			data:     []int{1},
			shift:    0,
			expected: []int{1},
		},
		{
			data:     []int{1},
			shift:    1,
			expected: []int{1},
		},
		{
			data:     []int{1},
			shift:    2,
			expected: []int{1},
		},
		{
			data:     []int{1, 2},
			shift:    1,
			expected: []int{2, 1},
		},
		{
			data:     []int{1, 2},
			shift:    2,
			expected: []int{1, 2},
		},
		{
			data:     []int{1, 2},
			shift:    3,
			expected: []int{2, 1},
		},
		{
			data:     []int{1, 2},
			shift:    4,
			expected: []int{1, 2},
		},
		{
			data:     []int{1, 2, 3},
			shift:    1,
			expected: []int{2, 3, 1},
		},
		{
			data:     []int{1, 2, 3},
			shift:    3,
			expected: []int{1, 2, 3},
		},
		{
			data:     []int{1, 2, 3},
			shift:    2,
			expected: []int{3, 1, 2},
		},
		{
			data:     []int{1, 2, 3, 4, 5, 6, 7},
			shift:    4,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
	}

	for _, fxt := range fixtures {
		actual := rotate(fxt.data, fxt.shift)

		if !equal(actual, fxt.expected) {
			t.Errorf("Expected: %v, Actual: %v, Shift: %d", fxt.expected, actual, fxt.shift)
		}
	}
}

func TestRotate_NegativeShift(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("panic expected while shift is negative")
		}
	}()
	s := []int{1, 2, 3}
	rotate(s, -1)

	t.Errorf("panic expected while shift is negative")
}

func equal(a, b []int) bool {
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
