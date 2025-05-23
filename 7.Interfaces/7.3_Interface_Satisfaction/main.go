package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//! A type satisfies an interface if it possesses all the methods the interface requires.

// The  assignability  rule for  interfaces  is  very  simple:
//! an  expression  may  be assigned to an interface only if its type satisfies the interface.
// * var w io.Writer
// * w = os.Stdout         // OK: *os.File has Write method
// * w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
// * w = time.Second       // compile error: time.Duration lacks Write method

//* var rwc io.ReadWriteCloser
//* rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
//* rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

//? Because  ReadWriter  and  ReadWriteCloser  include  all  the  methods  of Writer,
//? any type that satisfies ReadWriter or ReadWriteCloser necessarily satisfies Writer.
//* w = rwc                 // OK: io.ReadWriteCloser has Write method
//* rwc = w                 // compile error: io.Writer lacks Close method

type IntSet struct { /* ... */
}

func (*IntSet) String() string { return "" }

func m() {
	var s IntSet
	//? However, since only *IntSet has a String method, only *IntSet satisfies the fmt.Stringer interface:
	var _ fmt.Stringer = &s
	//* var _ fmt.Stringer = s // compile error: IntSet lacks String method

	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method
}

// The type interface{}
// which  is  called  the  empty  interface  type,  is  indispensable.  Because  the  empty
// interface type places no demands on the types that satisfy it, we can assign any value to the empty interface

func mm() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Fprintf(os.Stdout, "%v", any)
}

//? Since interface satisfaction depends only on the methods of the two types involved,
//? there is no need to declare the relationship between a concrete type and the interfaces
//? it satisfies.

// ! The declaration below asserts  at  compile  time  that  a  value  of  type  *bytes.Buffer  satisfies io.Writer:
// *bytes.Buffer must satisfy io.Writer
var _ io.Writer = new(bytes.Buffer)

//? interfaces are but one useful way to group related concrete types together and express the facets they share in common

//? Each grouping of concrete types based on their shared behaviors can be expressed as an interface type.
