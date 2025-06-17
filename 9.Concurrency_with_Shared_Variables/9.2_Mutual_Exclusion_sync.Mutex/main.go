package main

import (
	"sync"
)

// ! A binary semaphore - semaphore that counts only to 1
var (
	binSem         = make(chan struct{}, 1) // a binary semaphore guarding balance
	balance_binSem int
)

func Deposit_binSem(amount int) {
	binSem <- struct{}{} // acquire token
	balance = balance + amount
	<-binSem // release token
}
func Balance_binSem() int {
	binSem <- struct{}{} // acquire token
	b := balance
	<-binSem // release token
	return b
}

// This  pattern  of  mutual  exclusion  is  so  useful  that  it  is  supported  directly  by  the
// Mutex type from the sync package

var (
	mu      sync.Mutex // The mutex guards the shared variables
	balance int
	//? By convention, the variables guarded by a mutex are declared immediately
	//? after the declaration of the mutex itself. If you deviate from this, be sure to document it.
)

func deposit_Unexported(amount int) {
	//* Each time a goroutine accesses the variables of the bank,
	//* it must call the mutex’s Lock method to  acquire  an  exclusive lock
	mu.Lock()
	// If some other goroutine has acquired the lock, this operation will block until
	// the other goroutine calls Unlock and the lock becomes available again.
	balance = balance + amount
	mu.Unlock()
}
func nalance_Unexported() int {
	mu.Lock()
	//* The region of code between Lock and Unlock in which a goroutine is free to read
	//* and modify the shared variables is called a critical section.
	b := balance
	mu.Unlock()
	return b
}

func withdraw_Unexported(amount int) bool {
	deposit_Unexported(-amount)
	if nalance_Unexported() < 0 {
		deposit_Unexported(amount)
		return false // insufficient funds
	}
	return true
	//* The problem is that Withdraw is not atomic: it consists of a sequence of three separate operations,
	//* each of which acquires and then releases the mutex lock, but nothing locks the whole sequence
}

func Withdraw_Exported(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}
func Deposit_Exported(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}
func Balance_Exported() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }
