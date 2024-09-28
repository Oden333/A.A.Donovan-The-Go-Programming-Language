package main

import (
	"fmt"
)

func main() {

	var x, y int
	_, _ = fmt.Scanf("%d %d", &x, &y)
	fmt.Println(gcd(x, y))
}

// greatest common divisor (GCD) of two integers:
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// the n-th Fibonacci number iteratively:
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
