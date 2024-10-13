package main

import (
	"fmt"
	"unicode"
)

func f() {}

var g = "g"

func main() {
	main1()
}
func main0() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", unicode.ToUpper(rune(x)))
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}

func main1() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}
