package main

import (
	"fmt"
	"io"
	"os"
)

//! Conceptually, a value of an interface type, or interface value, has two components,
//! a concrete type and a value of that type.
//
//! These are called the interface’s dynamic type and dynamic value.

//? In Go, types are a compile-time concept, so a type is not a value.

// A set of values called type descriptors provide information about each type, such as its name and methods.
// ? In an interface value, the type component is represented by the appropriate type descriptor.
var w io.Writer

//! The zero value for an interface has both its type and value components set to nil
// An interface value is described as nil or non-nil based on its dynamic type, so this is a nil interface value

func main() {
	w = os.Stdout
	fmt.Println(w == nil) // false
	w = nil
	fmt.Println(w == nil) // true
	// w.Write([]byte{}) // panic

	w = os.Stdout //? The interface value’s dynamic type is set to the type descriptor for the pointer type *os.File
	//? and its dynamic value holds a copy of os.Stdout which
	//? is a pointer to the os.File variable representing the standard output of the process

}
