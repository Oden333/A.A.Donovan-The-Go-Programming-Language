package benchmarkfunctions

import "unicode"

// Benchmarking is the practice of measuring the performance of a program on a fixed workload.

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

//* go test -bench=.
// goos: linux
// goarch: amd64
// pkg: word
// cpu: 12th Gen Intel(R) Core(TM) i5-1235U
// BenchmarkIsPalindrome-12         5193930               196.3 ns/op
// PASS
// ok      word    1.259s

//* The benchmark name’s numeric suffix, 12 here, indicates the value of GOMAXPROCS,
//* which is important for concurrent benchmarks

//? The  -benchmem  command-line  flag  will  include  memory allocation statistics in its report.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
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
