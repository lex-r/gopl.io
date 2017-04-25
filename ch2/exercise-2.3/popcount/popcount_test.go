package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountByCycle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByCycle(uint64(i))
	}
}
