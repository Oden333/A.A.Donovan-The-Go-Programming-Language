package main

import (
	"fmt"
	"os"
	"runtime"
)

// Go’s type system catches many mistakes at compile time, but others, like an out-of-
// bounds array access or nil pointer dereference, require checks at run time.
// When the Go runtime detects these mistakes, it panics.

//! During  a  typical  panic,  normal  execution  stops,  all  deferred  function  calls  in  that
//! goroutine  are  executed,  and  the  program  crashes  with  a  log  message

//* This  log message includes the panic  value, which is usually an error message of some sort, and, for each goroutine,
//* a stack trace showing the stack of function calls that were active  at  the  time  of  the  panic

// Not all panics come from the runtime. The built-in panic function may be called directly
//* func testPanic() {
//* 	switch s := suit(drawCard()); s {
//* 	case "Spades": // ...
//* 	case "Hearts": // ...
//* 	case "Diamonds": // ...
//* 	case "Clubs": // ...
//* 	default:
//* 		panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
//* 	}
//* }

// func Reset(x *bufio) {
// 	if x == nil {
// 		panic("x is nil") // unnecessary!
// 	}
// 	x.elements = nil
// }

//! Since a panic causes the program to crash, it is generally used for grave errors,
//! such as a logical inconsistency in the program;
//! diligent programmers consider any crash to be proof of a bug in their code.

//& In  a  robust  program,  “expected”  errors,  the  kind  that  arise  from  incorrect  input,
//& misconfiguration, or failing I/O, should be handled gracefully; they are best dealt with
//& using error values.

// ! When  a  panic  occurs,  all  deferred  functions  are  run  in  reverse  order,  starting  with
// ! those  of  the  topmost  function  on  the  stack  and  proceeding  up  to  main
// func main() {
// f(3)
// }
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

//? For diagnostic purposes, the runtime package lets the programmer dump the stack
//? using the same machinery. By deferring a call to printStack in main,

func main() {
	defer printStack()
	f(3)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//! Go’s  panic  mechanism  runs  the  deferred  functions  before  it unwinds the stack.
