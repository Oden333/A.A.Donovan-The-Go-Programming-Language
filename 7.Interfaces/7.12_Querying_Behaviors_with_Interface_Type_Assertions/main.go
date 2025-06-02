package main

import "io"

// The io.Writer w represents the HTTP response;
// the bytes written to it are ultimately sent to someone’s web browser.
func writeHeader(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type: ")); err != nil {
		return err
	}
	//* This conversion allocates memory and  makes  a  copy,
	//* but  the  copy  is  thrown  away  almost  immediately  after.
	if _, err := w.Write([]byte(contentType)); err !=
		nil {
		return err
	}
	// ...
	return nil
}

// The  io.Writer  interface  tells  us  only  one  fact  about  the  concrete  type  that  w holds:
//  that  bytes  may  be  written  to  it.

//* If  we  look  behind  the  curtains  of  the
//* net/http package, we see that the dynamic type that w holds in this program also
//* has  a  WriteString  method  that  allows  strings  to  be  efficiently  written  to  it,
//? avoiding the need to allocate a temporary copy

// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}

func writeHeader1(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err !=
		nil {
		return err
	}
	// ...
	return nil
}

// The technique above relies on the assumption that
// ! if a type satisfies the interface  below,
// ! then  WriteString(s)  must  have  the  same  effect  as Write([]byte(s))
type f interface {
	io.Writer
	WriteString(s string) (n int, err error)
}

//? Defining a method of a  particular  type  is  taken  as  an  implicit  assent  for  a  certain  behavioral  contract.
