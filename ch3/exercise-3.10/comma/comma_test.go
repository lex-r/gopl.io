package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	fixtures := map[string]string{
		"":        "",
		"1":       "1",
		"123":     "123",
		"1234":    "1,234",
		"123456":  "123,456",
		"1234567": "1,234,567",
	}

	for data, expected := range fixtures {
		actual := comma(data)
		if expected != actual {
			t.Errorf("Expected: %s, actual: %s", expected, actual)
		}
	}
}
