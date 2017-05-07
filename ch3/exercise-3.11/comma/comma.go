package comma

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	var sign bool
	var integer *bytes.Buffer
	var fraction *bytes.Buffer

	if len(s) > 1 && s[0] == '-' {
		sign = true
		s = s[1:]
	}

	if pos := strings.Index(s, "."); pos > 0 {
		integer = bytes.NewBufferString(s[:pos])
		fraction = bytes.NewBufferString(s[pos+1:])
	} else {
		integer = bytes.NewBufferString(s)
	}

	if integer.Len() <= 3 {
		return repr(sign, integer, fraction)
	}

	integerResult := commaDelim(integer)

	return repr(sign, integerResult, fraction)
}

func commaDelim(s *bytes.Buffer) *bytes.Buffer {
	if s.Len() <= 3 {
		return s
	}
	var result bytes.Buffer

	firstRead := s.Len() % 3
	result.Write(s.Next(firstRead))

	for s.Len() >= 3 {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write(s.Next(3))
	}

	return &result
}

func repr(sign bool, integer *bytes.Buffer, fraction *bytes.Buffer) string {
	var result string
	if sign {
		result += "-"
	}

	result += integer.String()

	if fraction != nil && fraction.Len() > 0 {
		result += "." + fraction.String()
	}

	return result
}
