package comma

import (
	"bytes"
)

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	var result bytes.Buffer
	sb := bytes.NewBufferString(s)

	firstRead := sb.Len() % 3
	result.Write(sb.Next(firstRead))

	for sb.Len() >= 3 {
		if result.Len() > 0 {
			result.Write([]byte(","))
		}
		result.Write(sb.Next(3))
	}

	return result.String()
}
