package main

import "time"

//! A variable or method of an object is said to be encapsulated
//! if it is inaccessible to clients of the object.
//? Encapsulation, sometimes called information  hiding,  is  a  key aspect of object-oriented programming

//? Go has only one mechanism to control the visibility of names:
//! capitalized identifiers  are exported from the package in which they are defined,
//! and uncapitalized names are not.
//
// The  same  mechanism  that  limits  access  to  members  of  a  package  also  limits
// access  to  the  fields  of  a  struct  or  the  methods  of  a  type.

//? The fields of a struct type are visible to all code within the same package.
//? Whether the code appears in a function or a method makes no difference.

// As  an  example,  consider  the  bytes.Buffer  type
// Since Buffer is a struct type, this space takes the form of an extra field of type [64]byte
// with an uncapitalized name.
//* When this field was added, because it was not exported, clients
//* of  Buffer  outside  the  bytes  package  were  unaware  of  any  change  except
//* improved performance.

type Buffer struct {
	buf     []byte
	initial [64]byte
	/* ... */
}

// Grow expands the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
//* func (b *Buffer) Grow(n int) {
//* 	if b.buf == nil {
//* 		b.buf = b.initial[:0] // use preallocated space initially
//* 	}
//* 	if len(b.buf)+n > cap(b.buf) {
//* 		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
//* 		copy(buf, b.buf)
//* 		b.buf = buf
//* 	}
//* }

//? The third benefit of encapsulation, and in many cases the most important, is that it
//! prevents  clients  from  setting  an  object’s  variables  arbitrarily.

// * For example, the Counter  type  below  permits  clients  to  increment  the  counter  or  to
// * reset it to zero, but not to set it to some arbitrary value:
type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

//? Functions that merely access or modify internal values of a type, such as the methods
//? of  the  Logger  type  from  log  package,  below,  are  called  getters  and  setters.

// Encapsulation is not always desirable. By revealing its representation as an int64
// number of nanoseconds, time.Duration lets us use all the usual arithmetic and
// comparison operations with durations, and even to define constants of this type:
// Click here to view code image
const day = 24 * time.Hour

//* fmt.Println(day.Seconds()) // "86400"
