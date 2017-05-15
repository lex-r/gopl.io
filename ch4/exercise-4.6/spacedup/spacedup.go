package spacedup

import (
	"unicode"
	"unicode/utf8"
)

func spacedup(s *[]byte) {
	lastRuneIsSpace := false
	for i := 0; i < len(*s); {
		r, sz := utf8.DecodeRune((*s)[i:])
		if unicode.IsSpace(r) {
			if lastRuneIsSpace {
				shift(s, i, sz)
				*s = (*s)[:len(*s)-sz]
			} else {
				lastRuneIsSpace = true
				i += sz
			}
		} else {
			lastRuneIsSpace = false
			i += sz
		}
	}
}

func shift(s *[]byte, from, offset int) {
	for i := from; i < len(*s)-offset; i++ {
		(*s)[i] = (*s)[i+1]
	}
}
