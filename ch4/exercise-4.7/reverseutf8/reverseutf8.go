package reverseutf8

import "unicode/utf8"

// ReverseUTF8 обращает последовательность символов среза s, который представляет
// строку в кодировке UTF-8 в обратном порядке
func ReverseUTF8(s *[]byte) {
	if len(string(*s)) < 2 {
		return
	}

	for i, j := 0, len(*s)-1; i < j; {
		// первая руна
		firstRune, firstRuneSize := utf8.DecodeRune((*s)[i:])

		// последняя руна
		lastRune, lastRuneSize := utf8.DecodeLastRune((*s)[:j+1])

		if i >= j-lastRuneSize+1 {
			break
		}
		// если первая и последняя руны имеют разный размер (количество байт),
		// то перед их перестановкой нужно сместить промежуточные байты

		// величина смещения (на сколько байт влево или вправо нужно сместить
		// байты между первой и последней рунами)
		offset := lastRuneSize - firstRuneSize
		from := i + firstRuneSize
		to := j - lastRuneSize

		if offset != 0 && from <= to {
			shift(s, from, to, offset)
		}
		putRune(s, lastRune, i)
		putRune(s, firstRune, j-firstRuneSize+1)

		i += lastRuneSize
		j -= firstRuneSize
	}
}

func shift(s *[]byte, from, to, offset int) {
	if from > to {
		panic("from must be less or equal than to")
	}

	if len(*s) < to+offset+1 {
		panic("out of range")
	}

	if from+offset < 0 {
		panic("out of range")
	}

	if offset > 0 {
		for i := to + offset; i >= from+offset; i-- {
			(*s)[i] = (*s)[i-offset]
		}
	} else {
		for i := from + offset; i < to+offset+1; i++ {
			(*s)[i] = (*s)[i-offset]
		}
	}
}

func putRune(s *[]byte, r rune, pos int) {
	if pos > len(*s)-len(string(r)) || pos < 0 {
		panic("out of range")
	}

	buf := make([]byte, 4)
	runeBytes := utf8.EncodeRune(buf, r)
	for i := pos; i < pos+runeBytes; i++ {
		(*s)[i] = buf[i-pos]
	}
}
