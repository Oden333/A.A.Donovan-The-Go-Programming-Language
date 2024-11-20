package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	//! Maps is an unordered  collection  of  key/value  pairs  in  which  all  the  keys  are  distinct,  and  the
	//! value  associated  with  a  given  key  can  be  retrieved,  updated,  or  removed  using  a
	//! constant number of key comparisons on the average, no matter how large the hash
	//! table.

	//! In Go, a map is a reference to a hash table

	//! The key type must be comparable using ==
}

func init() {
	ages0 := make(map[string]int)
	ages0["alice"] = 31
	ages0["charlie"] = 34

	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	ages1["zz"] = 13

	//& But a map element is not a variable, and we cannot take its address:
	//* _ = &ages["bob"] // compile error: cannot take address of map element
	// 	One reason that we can’t take the address of a map element is that growing a map
	// might cause rehashing of existing elements into new storage locations, thus potentially
	// invalidating the address.
	// fmt.Println(len(ages1))
	// for name, age := range ages1 {
	// 	//? The order of map iteration is unspecified, and different implementations might use a
	// 	//? different  hash  function,  leading  to  a  different  ordering.  In  practice,  the  order  is
	// 	//? random,  varying  from  one  execution  to  the  next.
	// 	fmt.Println(name, age)
	// }

	//! 	But a map element is not a variable, and we cannot take its address:
	//* _ = &ages["bob"] // compile error: cannot take address of map element
	//& One reason that we can’t take the address of a map element is that growing a map
	//& might cause rehashing of existing elements into new storage locations, thus potentially
	//& invalidating the address.

	//! The order of map iteration is unspecified, and different implementations might use a
	//! different  hash  function,  leading  to  a  different  ordering.
	// In  practice,  the  order  is random,  varying  from  one  execution  to  the  next.

	//The zero value for a map type is nil, that is, a reference to no hash table at all.
	// var ages map[string]int
	// fmt.Println(ages == nil)    // "true"
	// fmt.Println(len(ages) == 0) // "true"

	// delete, len, and range loops, are safe  to  perform  on  a  nil  map  reference,
	// since  it  behaves  like  an  empty  map.  But storing to a nil map causes a panic:
	// ages["carol"] = 21 // panic: assignment to entry in nil map

	// fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
	// dedup()

	//? Sometimes we need a map or set whose keys are slices, but because a map’s keys
	//? must be comparable, this cannot be expressed directly. However, it can be done in
	//? two steps.
	// compateUncomparable()
	// countChars()
	// graphFunc()
	ex()
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func dedup() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func compateUncomparable() {
	// First we define a helper function k that maps each key to a string, with the
	// property that k(x) == k(y) if and only if we consider x and y equivalent.

}

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

// And  the  type  of  k(x) needn’t be a string;
// any comparable type with the desired equivalence property will
// do, such as integers, arrays, or structs.

func countChars() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("\nrune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func graphFunc() {

}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
