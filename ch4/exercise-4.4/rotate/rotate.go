package rotate

func rotate(s []int, shift int) []int {
	if len(s) <= 1 {
		return s
	}

	if shift < 0 {
		panic("shift must not be negative")
	}

	if shift >= len(s) {
		shift = shift % len(s)
	}

	if shift == 0 {
		return s
	}

	res := make([]int, 0)
	res = append(s[shift:], s[:shift]...)

	return res
}
