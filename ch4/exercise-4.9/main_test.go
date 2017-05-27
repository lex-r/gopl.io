package main

import (
	"bytes"
	"io"
	"testing"
)

func TestWordFreq(t *testing.T) {
	strWithBreakLines := `a b
   c
  a b`
	fixtures := []struct {
		data     io.Reader
		expected map[string]int
	}{
		{
			data:     bytes.NewBufferString(""),
			expected: make(map[string]int),
		},
		{
			data:     bytes.NewBufferString("word"),
			expected: map[string]int{"word": 1},
		},
		{
			data:     bytes.NewBufferString("word word"),
			expected: map[string]int{"word": 2},
		},
		{
			data:     bytes.NewBufferString("word word2"),
			expected: map[string]int{"word": 1, "word2": 1},
		},
		{
			data:     bytes.NewBufferString(`a b c c b a`),
			expected: map[string]int{"a": 2, "b": 2, "c": 2},
		},
		{
			data:     bytes.NewBufferString(strWithBreakLines),
			expected: map[string]int{"a": 2, "b": 2, "c": 1},
		},
	}

	for _, fxt := range fixtures {
		actual := wordfreq(fxt.data)
		if !mapequal(fxt.expected, actual) {
			t.Errorf("TestWordFreq: expected %v, actual %v", fxt.expected, actual)
		}
	}
}

func mapequal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bv, ok := b[k]; !ok || bv != v {
			return false
		}
	}

	return true
}
