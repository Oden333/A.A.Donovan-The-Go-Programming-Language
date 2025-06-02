package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//! A type assertion is an operation applied to an interface value.
//? Syntactically, it looks like x.(T), where x is an expression of an interface type
//? and T is a type, called the “asserted” type

// There are two possibilities.
//! First, if the asserted type T is a concrete type, then the
//! type  assertion  checks  whether  x’s  dynamic  type  is  identical  to  T.
//
//? If  this  check succeeds,  the  result  of  the  type  assertion  is  x’s  dynamic  value,  whose  type  is  of course T.
//
//! In other words, a type assertion to a concrete type extracts the concrete value from its operand.
//? If the check fails, then the operation panics.

func t() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer

	println(f, c)
}

// ! Second, if instead the asserted type T  is  an  interface  type,  then  the  type  assertion
// ! checks whether x’s dynamic type satisfies T
// ? If  this  check  succeeds,  the  dynamic value  is  not  extracted;
// ? the  result  is  still  an  interface  value  with  the  same  type  and
// ? value  components,  but  the  result  has  the  interface  type  T

// ! In  other  words,  a  type assertion to an interface type changes the type of the expression,
// making a different (and usually larger) set of methods accessible, but it preserves the dynamic type
// and value components inside the interface value
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (a int, b error) { return }

func main() {
	var w io.Writer
	fmt.Printf("w: %T, %[1]v,\n", w)
	w = os.Stdout

	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("w: %T, %[1]v,\t rw: %T, %[2]v\n", w, rw)
	//! both w and rw hold os.Stdout so each has a dynamic type of *os.File

	//? but w, an io.Writer, exposes only the file’s Write method, whereas
	//? rw exposes its Read method too

	w = new(ByteCounter)
	fmt.Printf("w: %T, %[1]v,\t rw: %T, %[2]v\n", w, rw)
	// rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method

	w = rw             // io.ReadWriter is assignable to io.Writer
	w = rw.(io.Writer) // fails only if rw == nil

	// If the type assertion appears in an assignment in which two results are expected,
	// such as the following declarations, the operation does not  panic  on  failure  but  instead  returns
	// an  additional  second  result,  a  boolean indicating success:
}

func b() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File) // success:  ok, f == os.Stdout
	fmt.Println(ok, f == os.Stdout)
	b, ok := w.(*bytes.Buffer) // failure: !ok, b == nil
	fmt.Println(ok, b == &bytes.Buffer{})
}
