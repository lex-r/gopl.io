package unique

func unique(s []string) {
	if len(s) <= 1 {
		return
	}

	cmp := 0
	for i := 1; i < len(s); i++ {
		if s[cmp] == s[i] {
			continue
		}

		cmp += 1
		s[cmp] = s[i]
	}

	for i := cmp + 1; i < len(s); i++ {
		s[i] = ""
	}
}
