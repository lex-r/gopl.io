package popcount

// pc[i] - количество единичных битов в i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения
// (количество установленных битов) значения x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountByShifting возвращает степень заполнения
// (количество установленных битов) значения x
func PopCountByShifting(x uint64) int {
	var res int = int(x & 1)
	for i := 0; i < 64; i++ {
		x = x >> 1
		res += int(x & 1)
	}

	return res
}
