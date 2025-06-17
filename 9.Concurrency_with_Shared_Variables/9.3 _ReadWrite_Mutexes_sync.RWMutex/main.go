package main

import "sync"

// ? A special kind of lock that allows read-only operations to proceed in parallel with each other,
// ? but write operations to have fully exclusive access is provided by sync.RWMutex:

var mu sync.RWMutex
var balance int

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}

//? The RLock and RUnlock methods to acquire and release a readers or shared lock
//
//? The mu.Lock and mu.Unlock methods to acquire and release a writer or exclusive lock.
//

//* RLock can be used only if there are no writes to shared variables in the critical section.

//* It’s only profitable to use an RWMutex when most of the goroutines that acquire the
//* lock are readers, and the lock is under contention,
// That is, goroutines routinely have to wait to acquire it.
// An RWMutex requires more complex internal bookkeeping, making it slower than a regular mutex for uncontended locks.
