package main

// In  most  operating  systems  and  programming  languages  that  support  multithreading,
// the current thread has a distinct identity that can be easily obtained as an ordinary
// value, typically an integer or pointer. This makes it easy to build an abstraction called
// thread-local  storage, which is essentially a global map keyed by thread identity, so
// that each thread can store and retrieve values independent of other threads.

//? Goroutines have no notion of identity that is accessible to the programmer.
//? This is by design, since thread-local storage tends to be abused.
