package main

import (
	"bytes"
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

func as() {
	w = os.Stdout
	fmt.Println(w == nil) // false
	w = nil
	fmt.Println(w == nil) // true
	// w.Write([]byte{}) // panic

	//? The interface value’s dynamic type is set to the type descriptor for the pointer type *os.File
	w = os.Stdout
	//? and its dynamic value holds a copy of os.Stdout which is a pointer
	//? to the os.File variable representing the standard output of the process

	// Calling the Write method on an interface value containing an *os.File pointer
	// causes the (*os.File).Write method to be called. The call prints "hello".
	w.Write([]byte("das\n"))

	//* we cannot know at compile time what the dynamic type of an interface
	//* value will be, so a call through an interface must use dynamic dispatch.

	//! Instead of a direct  call,  the  compiler  must  generate  code  to  obtain  the  address  of  the  method
	//! named Write from the type descriptor, then make an indirect call to that address.
	//! The  receiver  argument  for  the  call  is  a  copy  of  the  interface’s  dynamic  value,
	//! os.Stdout. The effect is as if we had made this call directly

	w = new(bytes.Buffer)
	//? The dynamic type is now *bytes.Buffer and
	//? the dynamic value is a pointer to the newly allocated buffer

	w.Write([]byte("hello")) // writes "hello" to the bytes.Buffer
	//? This  time,  the  type  descriptor  is  *bytes.Buffer,  so  the (*bytes.Buffer).Write method is called,
	//? with the address of the buffer as the value of the receiver parameter. The call appends "hello" to the buffer

	// An  interface  value  can  hold  arbitrarily  large  dynamic  values
	// var x interface{} = time.Now()
}

//! Interface values may be compared using == and !=.
//! Two interface values are equal if both are nil, or if their dynamic types are identical
//! and their dynamic values are equal according to the usual behavior of == for that type

//? Because interface values are comparable, they may be used as the keys of a map or as the operand of a switch statement
// However, if two interface values are compared and have the same dynamic type, but
// that type is not comparable (a slice, for instance), then the comparison fails with a  panic

func panicc() {
	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x) // panic: comparing uncomparable type []int
}
