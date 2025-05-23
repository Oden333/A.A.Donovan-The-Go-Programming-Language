package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//? The io.Writer type is one of the most widely used interfaces because it provides
//? an  abstraction  of  all  the  types  to  which  bytes  can  be  written,  which  includes
//* files,
//* memory buffers,
//* network connections,
//* HTTP clients,
//* archivers,
//* hashers, and so on

// ? A Reader represents any type from which you can read bytes
// ? A Closer is any value that you can close, such as a file or a network connection

type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type ReaddWriter interface {
	Read(p []byte) (n int, err error)
	io.Writer
}

type ReadWriteCloser interface {
	io.Reader
	io.Writer
	io.Closer
}

type StringReader struct {
	s string
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}

func main() {
	r := NewReader("dai manki")
	wrappedR := io.LimitReader(r, 3)
	scanner := bufio.NewScanner(wrappedR)

	for scanner.Scan() {
		fmt.Fprintf(os.Stdout, "%s", scanner.Bytes())
	}
}

type LReader struct {
	r io.Reader
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LReader{r: r, n: n}
}

func (wr *LReader) Read(dst []byte) (n int, err error) {
	if wr.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(dst)) > wr.n {
		dst = dst[0:wr.n]
	}

	n, err = wr.r.Read(dst)
	wr.n -= int64(n)

	return n, err
}
