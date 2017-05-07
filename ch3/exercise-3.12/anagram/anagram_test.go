package anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	fixtures := []struct {
		a, b     string
		expected bool
	}{
		{"a", "b", false},
		{"", "", true},
		{"a", "a", true},
		{"ab", "ba", true},
		{"abcd", "dcba", true},
		{"abcd", "bacd", true},
		{"abcd", "abca", false},
		{"aba", "bab", false},
		{"The Eyes", "They See", true},
	}

	for _, fixture := range fixtures {
		actual := isAnagram(fixture.a, fixture.b)
		if fixture.expected != actual {
			t.Errorf("Str1: %s, Str2: %s, Expected: %t, Actual: %t", fixture.a, fixture.b, fixture.expected, actual)
		}
	}
}
