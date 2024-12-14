package main

import (
	"fmt"
	"sort"
)

// Named functions can be declared only at the package level, but
// we can use a function literal to denote a function value within any expression.

//& A function literal is written like a function declaration, but without a name following the func keyword.
//& It is an expression, and its value is called an anonymous function.

//! Functions  defined  in  this  way  have  access  to  the  entire  lexical
//! environment, so the inner function can refer to variables from the enclosing function

// squares returns a function that returns the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main0() {
	f := squares()
	//? A call to squares creates a local variable x and returns an anonymous function that, each
	//? time  it  is  called,  increments  x  and  returns  its  square.
	fmt.Println(f()) // "1"

	//? A  second  call  to  squares would  create  a  second  variable  x  and
	//? return  a  new  anonymous  function  which increments that variable
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	//& The squares example demonstrates that function values are not just code but can have state
	//& The anonymous inner function can access and update the local variables of the enclosing function squares

	//! These hidden variable references are why we classify  functions  as  reference  types
	//! and  why  function  values  are  not  comparable

	//& The variable x exists after squares has returned within main, even though x is hidden inside f
}

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main1() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	//& When an anonymous function requires recursion, as in this example, we must  first
	//& declare a variable, and then assign the anonymous function to that variable.
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
