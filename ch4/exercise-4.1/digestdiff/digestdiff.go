package digestdiff

// pc[i] - количество единичных битов в i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func digestdiff(digest1, digest2 [32]byte) int {
	var diffs int
	for i, b1 := range digest1 {
		b2 := digest2[i]
		d := b1 ^ b2
		diffs += int(pc[d])
	}
	return diffs
}
