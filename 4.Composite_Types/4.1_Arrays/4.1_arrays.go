package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var shaPointer = flag.String("sha", "256", "sha number")

func arrays() {
	//! An array is a fixed-length sequence of zero or more elements of a particular type.
	// elements()
	// initialization()
	// comparing()
	// passing()
	// ex4_1()

	var input []string = os.Args[1:]
	fmt.Fprintf(os.Stdout, "user cmd input:%v\n", input)
	var s string
	fmt.Fprint(os.Stdout, "str: ")
	fmt.Fscanf(os.Stdin, "%s", &s)
	ex4_2(s)
}

func elements() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
func initialization() {
	//&	In an array literal, if an ellipsis “...” appears in place of the length, the array length
	//& is determined by the number of initializers.
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	//? The size must be a constant expression, that is, an expression whose value can be
	//? computed as the program is being compiled.
	//q := [3]int{1, 2, 3}
	//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	// The	specific form above is a list of values in order, but
	//! it is also possible to specify a list of index and value pairs, like this:
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{
		RMB: "¥", // 3 :  "¥"
		USD: "$", // 1 :  "$"
		EUR: "€", // 2 :  "€"
		GBP: "£", // 4 :  "£"
	}
	fmt.Println(RMB, symbol[RMB]) // "3 ¥"
	r := [...]int{99: -1}
	fmt.Println(len(r))
}

func comparing() {
	//& If an array’s element type is comparable then the array type is comparable too, so we
	// may directly compare two arrays of that type using the == operator, which reports
	// whether all corresponding elements are equal. The != operator is its negation.
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare
	// [2]int == [3]int

	//* The two inputs differ by only a single bit, but approximately half the bits are different in the digests
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func passing() {
	zero := func(ptr *[32]byte) {
		// for i := range ptr {
		// ptr[i] = 0
		// }
		*ptr = [32]byte{}
	}
	z := sha256.Sum256([]byte("x"))
	fmt.Println(z)
	zero(&z)
	fmt.Println(z)
	//?	Using a pointer to an array is efficient and allows the called function to mutate the
	//? caller’s variable, but arrays are still inherently inflexible because of their fixed size.
}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popCount returns the population count (number of set bits) of x.
func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func ex4_1() {
	var (
		c1      = sha256.Sum256([]byte("a"))
		c2      = sha256.Sum256([]byte("A"))
		counter int
		temp    byte
	)

	for idx := range 32 {
		temp = c1[idx] ^ c2[idx]
		counter += int(pc[temp])
		fmt.Printf("%8[1]b\n%8[2]b\t%[3]d\n", c1[idx], c2[idx], pc[temp])
	}

	fmt.Printf("\n%d", counter)
}

func ex4_2(s string) {
	flag.Parse()
	fmt.Fprintf(os.Stdout, "user input: %s\nsha: %s", s, *shaPointer)

	switch *shaPointer {
	case "256":
		fmt.Fprintf(os.Stdout, "| string: %s\t| hash: %x\n", *shaPointer, sha256.Sum256([]byte(s)))
	case "384":
		fmt.Println(sha512.Sum384([]byte(s)))
	case "512":
		fmt.Println(sha512.Sum512([]byte(s)))
	}
}
