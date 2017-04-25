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

// PopCountByCycle возвращает степень заполнения
// (количество установленных битов) значения x
func PopCountByCycle(x uint64) int {
	var res int
	for i := uint8(0); i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}

	return res
}
