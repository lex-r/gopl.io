package reverseutf8

import (
	"bytes"
	"testing"
)

func TestReverseUTF8(t *testing.T) {
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
			data:     []byte("а"),
			expected: []byte("а"),
		},
		{
			data:     []byte("世"),
			expected: []byte("世"),
		},
		{
			data:     []byte("ab"),
			expected: []byte("ba"),
		},
		{
			data:     []byte("abc"),
			expected: []byte("cba"),
		},
		{
			data:     []byte("abcd"),
			expected: []byte("dcba"),
		},
		{
			data:     []byte("fб"),
			expected: []byte("бf"),
		},
		{
			data:     []byte("1fж"),
			expected: []byte("жf1"),
		},
		{
			data:     []byte("1fжё"),
			expected: []byte("ёжf1"),
		},
		{
			data:     []byte("1жj"),
			expected: []byte("jж1"),
		},
		{
			data:     []byte("строка"),
			expected: []byte("акортс"),
		},
		{
			data:     []byte("строка abc"),
			expected: []byte("cba акортс"),
		},
		{
			data:     []byte("abc строка"),
			expected: []byte("акортс cba"),
		},
		{
			data:     []byte("строка abc 世界"),
			expected: []byte("界世 cba акортс"),
		},
		{
			data:     []byte("世界 abc строка"),
			expected: []byte("акортс cba 界世"),
		},
		{
			data:     []byte("abc 世界"),
			expected: []byte("界世 cba"),
		},
	}

	for _, fxt := range fixtures {
		ReverseUTF8(&fxt.data)
		if !bytes.Equal(fxt.expected, fxt.data) {
			t.Errorf("Expected: %v, Actual: %v", fxt.expected, fxt.data)
		}
	}
}

func _TestShift_ZeroOffset(t *testing.T) {
	data := []byte{1, 2, 3}
	expected := []byte{1, 2, 3}
	shift(&data, 0, 1, 0)
	if !bytes.Equal(expected, data) {
		t.Errorf("ZeroOffset: Expected: %v, Actual: %v", expected, data)
	}
}

func _TestShift_PositiveOffset(t *testing.T) {
	type fixture struct {
		data             []byte
		from, to, offset int
		expected         []byte
	}

	fixtures := []fixture{
		{
			data:     []byte{1, 2},
			from:     0,
			to:       0,
			offset:   1,
			expected: []byte{1, 1},
		},
		{
			data:     []byte{1, 2, 3},
			from:     0,
			to:       0,
			offset:   1,
			expected: []byte{1, 1, 3},
		},
		{
			data:     []byte{1, 2, 3},
			from:     0,
			to:       0,
			offset:   2,
			expected: []byte{1, 2, 1},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       1,
			offset:   1,
			expected: []byte{1, 1, 2, 4, 5, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       1,
			offset:   2,
			expected: []byte{1, 2, 1, 2, 5, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       1,
			offset:   3,
			expected: []byte{1, 2, 3, 1, 2, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       1,
			offset:   4,
			expected: []byte{1, 2, 3, 4, 1, 2},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       0,
			offset:   4,
			expected: []byte{1, 2, 3, 4, 1, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       0,
			offset:   5,
			expected: []byte{1, 2, 3, 4, 5, 1},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       3,
			offset:   1,
			expected: []byte{1, 1, 2, 3, 4, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     0,
			to:       3,
			offset:   2,
			expected: []byte{1, 2, 1, 2, 3, 4},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     4,
			to:       4,
			offset:   1,
			expected: []byte{1, 2, 3, 4, 5, 5},
		},
	}

	for i, fxt := range fixtures {
		shift(&fxt.data, fxt.from, fxt.to, fxt.offset)
		if !bytes.Equal(fxt.expected, fxt.data) {
			t.Errorf("TestShift_PositiveOffset(%d). Expected: %v, Actual: %v", i, fxt.expected, fxt.data)
		}
	}
}

func _TestShift_NegativeOffset(t *testing.T) {
	type fixture struct {
		data             []byte
		from, to, offset int
		expected         []byte
	}

	fixtures := []fixture{
		{
			data:     []byte{1, 2},
			from:     1,
			to:       1,
			offset:   -1,
			expected: []byte{2, 2},
		},
		{
			data:     []byte{1, 2, 3},
			from:     1,
			to:       1,
			offset:   -1,
			expected: []byte{2, 2, 3},
		},
		{
			data:     []byte{1, 2, 3},
			from:     2,
			to:       2,
			offset:   -1,
			expected: []byte{1, 3, 3},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     5,
			to:       5,
			offset:   -1,
			expected: []byte{1, 2, 3, 4, 6, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     4,
			to:       5,
			offset:   -1,
			expected: []byte{1, 2, 3, 5, 6, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     1,
			to:       5,
			offset:   -1,
			expected: []byte{2, 3, 4, 5, 6, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     5,
			to:       5,
			offset:   -2,
			expected: []byte{1, 2, 3, 6, 5, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     2,
			to:       5,
			offset:   -2,
			expected: []byte{3, 4, 5, 6, 5, 6},
		},
		{
			data:     []byte{1, 2, 3, 4, 5, 6},
			from:     3,
			to:       5,
			offset:   -2,
			expected: []byte{1, 4, 5, 6, 5, 6},
		},
	}

	for i, fxt := range fixtures {
		shift(&fxt.data, fxt.from, fxt.to, fxt.offset)
		if !bytes.Equal(fxt.expected, fxt.data) {
			t.Errorf("TestShift_NegativeOffset(%d). Expected: %v, Actual: %v", i, fxt.expected, fxt.data)
		}
	}
}

func _TestShift_OutOfTheRange(t *testing.T) {
	type fixture struct {
		data             []byte
		from, to, offset int
	}

	fixtures := []fixture{
		{
			data:   []byte{1, 2},
			from:   -1,
			to:     1,
			offset: 0,
		},
		{
			data:   []byte{1, 2},
			from:   0,
			to:     1,
			offset: -1,
		},
		{
			data:   []byte{1, 2, 3, 4, 5},
			from:   1,
			to:     1,
			offset: -2,
		},
		{
			data:   []byte{1, 2, 3, 4, 5},
			from:   2,
			to:     2,
			offset: -3,
		},
		{
			data:   []byte{1, 2},
			from:   0,
			to:     2,
			offset: 0,
		},
		{
			data:   []byte{1, 2},
			from:   0,
			to:     1,
			offset: 1,
		},
		{
			data:   []byte{1, 2},
			from:   0,
			to:     0,
			offset: 2,
		},
		{
			data:   []byte{1, 2},
			from:   0,
			to:     0,
			offset: 3,
		},
		{
			data:   []byte{1, 2, 3, 4, 5},
			from:   0,
			to:     4,
			offset: 1,
		},
		{
			data:   []byte{1, 2, 3, 4, 5},
			from:   0,
			to:     3,
			offset: 2,
		},
	}

	for _, fxt := range fixtures {
		shouldPanic(t, func() {
			shift(&fxt.data, fxt.from, fxt.to, fxt.offset)
		}, "out of range")
	}
}

func _TestShift_FromMustNotGreaterThanTo(t *testing.T) {
	shouldPanic(t, func() {
		data := []byte{1, 2, 3}
		shift(&data, 1, 0, 1)
	}, "from must be less or equal than to")
}

func _TestPutRune(t *testing.T) {
	fixtures := []struct {
		data     []byte
		r        rune
		pos      int
		expected []byte
	}{
		{
			data:     []byte("abc"),
			r:        'd',
			pos:      0,
			expected: []byte("dbc"),
		},
		{
			data:     []byte("abc"),
			r:        'd',
			pos:      1,
			expected: []byte("adc"),
		},
		{
			data:     []byte("abc"),
			r:        'd',
			pos:      2,
			expected: []byte("abd"),
		},
		{
			data:     []byte("abc"),
			r:        'б',
			pos:      0,
			expected: []byte("бc"),
		},
		{
			data:     []byte("abc"),
			r:        'б',
			pos:      1,
			expected: []byte("aб"),
		},
		{
			data:     []byte("abc"),
			r:        '‰',
			pos:      0,
			expected: []byte("‰"),
		},
	}

	for _, fxt := range fixtures {
		putRune(&fxt.data, fxt.r, fxt.pos)

		if !bytes.Equal(fxt.expected, fxt.data) {
			t.Errorf("TestPutRune: expected: %v, actual: %v", fxt.expected, fxt.data)
		}
	}

}

func shouldPanic(t *testing.T, f func(), panicMsg string) {
	defer func() {
		r := recover()
		if err := r; err == nil || err != panicMsg {
			t.Errorf("func should throw panic with message `%s`, actual message is `%s`", panicMsg, err)
		}
	}()

	f()
}
