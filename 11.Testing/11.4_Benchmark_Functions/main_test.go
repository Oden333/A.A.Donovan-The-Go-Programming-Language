package benchmarkfunctions

import (
	"testing"
)

// We run it with the command below. Unlike tests, by default no benchmarks are run.
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// Benchmarks like this tell us the absolute time required for a given operation, but in
// many settings the interesting performance questions are about the relative timings of
// two different operations.

// If a function takes 1ms to process 1,000 elements, how long will it take to process 10,000 or a million?
// What is the best size for an I/O buffer?
// Which algorithm performs best for a given job?

func benchmark(b *testing.B, size int) {
	for i := 0; i < size; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}
func Benchmark10(b *testing.B)   { benchmark(b, 10) }
func Benchmark100(b *testing.B)  { benchmark(b, 100) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }
