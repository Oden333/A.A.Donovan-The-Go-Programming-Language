package main

import (
	"fmt"
	"os"
)

//! A variadic function is one that can be called with varying numbers of arguments.

//? To  declare  a  variadic  function,  the  type  of  the  final  parameter  is  preceded  by  an
//? ellipsis, “...”, which indicates that the function may be called with any number of
//? arguments of this type.

//& Implicitly, the caller allocates an array, copies the arguments into it, and passes a slice
//& of the entire array to the function

// the ...int parameter behaves like a slice within the function body
func f(...int) {}
func g([]int)  {}

// Variadic functions are often used for string formatting. The errorf function below
// constructs a formatted error message with a line number at the beginning.

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
func errorPrint() {
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name) // "Line 12:undefined:count
}

func main() {
	ex()
}
