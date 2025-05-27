package main

import (
	"bytes"
	"io"
)

//! A nil interface value, which contains no value at all, is not the same as an interface
//! value containing a pointer that happens to be nil.

const debug = true

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	//? When main calls f, it assigns a nil pointer of type *bytes.Buffer to the out parameter,
	//?  so  the  dynamic  value  of  out  is  nil.
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	//? However,  its  dynamic  type  is *bytes.Buffer, meaning that out is a non-nil interface containing a nil pointer value
	//! so the defensive check out != nil is still true.
	if out != nil {
		out.Write([]byte("done!\n")) // panic: nil pointer dereference
	}
}

// The  problem  is  that  although  a  nil  *bytes.Buffer  pointer  has  the  methods
// needed to satisfy the interface, it doesn’t satisfy the behavioral requirements of the
// interface.

// The solution is to change the type of buf in main to io.Writer,
// thereby  avoiding  the  assignment  of  the  dysfunctional  value  to  the
// interface in the first place

func fix() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // OK
}
