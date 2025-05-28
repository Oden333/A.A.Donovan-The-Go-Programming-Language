package main

import (
	"errors"
	"fmt"
	"syscall"
)

func main() {
	// every call to New allocates a distinct error instance that is equal to no other.
	// We would not want a distinguished error such as io.EOF to compare
	// equal to one that merely happened to have the same message.
	fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"

	// the syscall package provides Go’s low-level system call API.
	var err error = syscall.Errno(2)
	fmt.Println(err.Error()) // "no such file or directory"
	fmt.Println(err)         // "no such file or directory"

}

// type Errno uintptr // operating system error code
// var errors = [...]string{
// 1:   "operation not permitted",   // EPERM
// 2:   "no such file or directory", // ENOENT
// 3:   "no such process",           // ESRCH
// ...
// }

// func (e Errno) Error() string {
// if 0 <= int(e) && int(e) < len(errors) {
// return errors[e]
// }
// return fmt.Sprintf("errno %d", e)
// }
