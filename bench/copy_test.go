package bench_test

import (
	"testing"

	"github.com/bingoohuang/gogotcha/bench"
)

func BenchmarkCloneList(b *testing.B) {
	b.ReportAllocs()

	input := []string{"abb", "bbbb", "cbbbb"}

	for i := 0; i < b.N; i++ {
		bench.CloneList(input)
	}
}
