package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountByReset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByReset(uint64(i))
	}
}
