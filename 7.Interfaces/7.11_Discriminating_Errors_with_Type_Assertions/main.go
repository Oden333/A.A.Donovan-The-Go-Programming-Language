package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

//? When you are working with errors handling, A reliable  approach  is  to  represent  structured  error  values  using  a  dedicated type.

// The  os  package  defines  a  type  called  PathError  to  describe  failures
// involving  an  operation  on  a  file  path,  like  Open or Delete,  and  a  variant  called
// LinkError  to  describe  failures  of  operations  involving  two  file  paths,  like
// Symlink and Rename. Here’s os.PathError:

//? Most clients are oblivious to PathError and deal with all errors in a uniform way
//? by calling their Error methods.

func main() {
	_, err := os.Open("/no/such/file")
	fmt.Println(err) // "open /no/such/file: No such file or directory"
	fmt.Printf("%#v\n", err)
	// Output:
	// &os.PathError{Op:"open", Path:"/no/such/file", Err:0x2}

	_, err = os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err)) // "true"
}

var ErrNotExist = errors.New("file does not exist")

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

// Error  discrimination  must usually be done immediately after the failing operation,
// before an error is propagated to the caller.
