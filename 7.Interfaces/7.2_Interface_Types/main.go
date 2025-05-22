package main

import "io"

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
