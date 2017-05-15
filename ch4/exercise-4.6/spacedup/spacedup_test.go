package spacedup

import (
	"bytes"
	"testing"
)

func TestSpacedup(t *testing.T) {
	type fixture struct {
		data     []byte
		expected []byte
	}

	fixtures := []fixture{
		{
			data:     []byte(""),
			expected: []byte(""),
		},
		{
			data:     []byte("1"),
			expected: []byte("1"),
		},
		{
			data:     []byte(" "),
			expected: []byte(" "),
		},
		{
			data:     []byte("  "),
			expected: []byte(" "),
		},
		{
			data:     []byte(" 1 "),
			expected: []byte(" 1 "),
		},
		{
			data:     []byte("  1"),
			expected: []byte(" 1"),
		},
		{
			data:     []byte("1  "),
			expected: []byte("1 "),
		},
		{
			data:     []byte("  1  "),
			expected: []byte(" 1 "),
		},
		{
			data:     []byte("1 1"),
			expected: []byte("1 1"),
		},
		{
			data:     []byte("1  1"),
			expected: []byte("1 1"),
		},
		{
			data:     []byte("1  1 1"),
			expected: []byte("1 1 1"),
		},
		{
			data:     []byte("1  1  1"),
			expected: []byte("1 1 1"),
		},
		{
			data:     []byte("  1   1   1  "),
			expected: []byte(" 1 1 1 "),
		},
		{
			data:     []byte("世界"),
			expected: []byte("世界"),
		},
		{
			data:     []byte("世 界"),
			expected: []byte("世 界"),
		},
		{
			data:     []byte("世  界"),
			expected: []byte("世 界"),
		},
		{
			data:     []byte(" 世界"),
			expected: []byte(" 世界"),
		},
		{
			data:     []byte("  世界"),
			expected: []byte(" 世界"),
		},
		{
			data:     []byte("世界  "),
			expected: []byte("世界 "),
		},
		{
			data:     []byte("  世 界  "),
			expected: []byte(" 世 界 "),
		},
		{
			data:     []byte("   世   界   "),
			expected: []byte(" 世 界 "),
		},
	}

	for i, fxt := range fixtures {
		before := clone(fxt.data)
		var data *[]byte = &fxt.data
		var expected *[]byte = &fxt.expected
		spacedup(data)

		if !bytes.Equal(*data, *expected) {
			t.Errorf("Expected: %v, Actual: %v, Before: %v, ind: %d", expected, data, before, i)
		}
	}
}

func clone(src []byte) []byte {
	dst := make([]byte, len(src))
	for i, v := range src {
		dst[i] = v
	}

	return dst
}
