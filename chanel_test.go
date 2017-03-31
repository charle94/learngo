package learn

import (
	"testing"
)

func Test_mychan(t *testing.T) {
	my()
}
func Benchmark_mycha(b *testing.B) {
	my()
}
func Benchmark_myloop(b *testing.B) {
	myloop()
}
