package bench

import "testing"

func BenchmarkCloneList(b *testing.B) {
	b.ReportAllocs()
	input := []string{"abb", "bbbb", "cbbbb"}

	for i := 0; i < b.N; i++ {
		CloneList(input)
	}
}
