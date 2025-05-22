package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// ! Interface  types  express  generalizations  or  abstractions  about  the  behaviors  of  other types
//? Interfaces  let  us  write  functions  that  are  more  flexible  and
//? adaptable because they are not tied to the details of one particular implementation.

// ! Interfaces satisfied implicitly
// In other words, there’s no need to declare all the interfaces that a given concrete type satisfies;
// simply possessing the necessary methods is enough.

// ? A  concrete  type specifies the exact representation of its values and
// ? exposes the intrinsic operations of that  representation,
// ? such  as  arithmetic  for  numbers, or  indexing,  append,  and range for slices.
type v int

// ? A concrete type may also provide additional behaviors through its methods.
func (v v) f() {}

// ! There  is  another  kind  of  type  in  Go  called  an  interface  type. An  interface  is  an abstract type.
type zzz interface{}

//* It doesn’t expose the representation or internal structure of its values or  the  set  of  basic  operations  they  support;
//? When you have a value of an interface type, you know nothing about what it is;
//? you know  only what  behaviors  are  provided  by  its methods.

// package fmt
//
// func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
//
// func Printf(format string, args ...interface{}) (int, error) {
//     return Fprintf(os.Stdout, format, args...)
// }
// func Sprintf(format string, args ...interface{}) string {
//     var buf bytes.Buffer
//     Fprintf(&buf, format, args...)
//     return buf.String()
// }
//? The F  prefix  of  Fprintf  stands  for  file  and  indicates  that  the  formatted  output
//? should be written to the file provided as the first argument

// Writer is the interface that wraps the basic Write method.
type Writer interface {
	// Write writes len(p) bytes from p to the underlying data stream.
	// It returns the number of bytes written from p (0 <= n <= len(p))
	// and any error encountered that caused the write to stop early.
	// Write must return a non-nil error if it returns n < len(p).
	// Write must not modify the slice data, even temporarily.
	//
	// Implementations must not retain p.
	Write(p []byte) (n int, err error)
}

//! The io.Writer interface defines the contract between Fprintf and its callers.

//* Because fmt.Fprintf assumes nothing about the representation of the value and
//* relies only on the behaviors guaranteed by the io.Writer contract, we can safely
//* pass a value of any concrete type that satisfies io.Writer as the first argument to
//* fmt.Fprintf. This freedom to substitute one type for another that satisfies the
//* same  interface  is  called  substitutability,  and  is  a  hallmark  of  object-oriented
//* programming.

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	// var word string
	for scanner.Scan() {
		// word = scanner.Text()
		// fmt.Println(word)
		*c += WordsCounter(1)
	}
	return int(*c), nil
}

type LinesCounter int

func (c *LinesCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	// var word string
	for scanner.Scan() {
		// word = scanner.Text()
		// fmt.Println(word)
		*c += LinesCounter(1)
	}
	return int(*c), nil
}

func ex7_1() {
	var b WordsCounter
	fmt.Println(b.Write([]byte("zzzz das dassad dasdasdas  d ")))
	var a LinesCounter
	fmt.Println(a.Write([]byte("zzzz das \n dassad dasdasdas \n   d ")))
}

type countingWriter struct {
	count  int64
	writer io.Writer
}

func (z *countingWriter) Write(p []byte) (n int, err error) {
	// defer func() { z.count += int64(n) }()
	// не верно,иногда n может быть 0 при ошибке, и тогда count не изменится, что правильно.
	n, err = z.writer.Write(p)
	z.count += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrapped := new(countingWriter)
	wrapped.writer = w
	return wrapped, &wrapped.count
}

func ex7_2() {
	wr, counter := CountingWriter(os.Stdout)
	wr.Write([]byte("s string stringssss"))
	fmt.Fprintf(wr, "\nbytes:%d\n", *counter)
	fmt.Fprintf(wr, "\nbytes:%d\n", *counter)
}

func main() {
}
