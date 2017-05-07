package anagram

import (
	"unicode"
)

func isAnagram(a, b string) bool {
	if a == b {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	runeCountA := runeCount(a)
	runeCountB := runeCount(b)
	for r, c := range runeCountA {
		if runeCountB[r] != c {
			return false
		}
	}

	return true
}

func runeCount(s string) map[rune]int {
	result := make(map[rune]int)
	for _, r := range s {
		result[unicode.ToLower(r)] += 1
	}

	return result
}
