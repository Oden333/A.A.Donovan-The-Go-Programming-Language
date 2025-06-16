package main

import (
	"fmt"
	"image"
	"time"
)

//! A type is concurrency-safe if all its accessible methods and operations are concurrency-safe.
//! - if it continues to work correctly even when called concurrently, that is, from two or more goroutines with no additional synchronization.

//? We avoid concurrent access to most variables either
//? by confining them to a single goroutine or
//? by maintaining a higher-level invariant of mutual exclusion.

//* In  contrast,  exported  package-level  functions  are  generally  expected  to  be
//* concurrency-safe.  Since  package-level  variables  cannot  be  confined  to  a  single
//* goroutine, functions that modify them must enforce mutual exclusion.

// A race condition is a situation in which the program does not give the correct result
// for some interleavings of the operations of multiple goroutines.

//! A data race occurs whenever two goroutines access the same variable concurrently
//! and at least one of the accesses is a write.

var balance int

func Deposit(amount int) {
	balance = balance + amount
}
func Balance() int {
	return balance
}

func maind() {
	// Alice:
	go func() {
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()
	// Bob:
	go Deposit(100) // B

	<-time.Tick(time.Microsecond * 400)
	fmt.Println(balance)

	var x []int
	go func() { x = make([]int, 1000000) }()
	go func() { x = make([]int, 10) }()
	<-time.Tick(time.Microsecond * 400)
	x[999999] = 1
}

var icons = make(map[string]image.Image)

func loadIcon(name string) image.Image {
	return nil
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon = loadIcon(name)
		icons[name] = icon
	}
	return icon
}

// ? If instead we initialize the map with all necessary entries before creating additional
// ? goroutines and never modify it again, then any number of goroutines may safely call
// ? Icon concurrently since each only reads the map
var icons1 = map[string]image.Image{
	"spades.png":   loadIcon("spades.png"),
	"hearts.png":   loadIcon("hearts.png"),
	"diamonds.png": loadIcon("diamonds.png"),
	"clubs.png":    loadIcon("clubs.png"),
}

// ! 1 Data structures that are never modified or are immutable are inherently concurrency-safe and need no synchronization.
// Concurrency-safe.
func Icon1l(name string) image.Image {
	return icons[name]
}

// ! 2 The second way to avoid a data race is to avoid accessing the variable from multiple goroutines

// “Do not communicate by sharing memory; instead, share memory by communicating.”
//* A goroutine that brokers access to a confined variable
//* using channel requests is called a monitor goroutine for that variable

// Package bank provides a concurrency-safe bank with one account.
var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit1(amount int) {
	deposits <- amount
}

func Balance1() int {
	return <-balances
}

func Withdraw(amount int) bool {
	b0 := Balance1()
	Deposit1(-amount)
	return b0 == Balance1()
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
func main() {
	// Alice:
	go func() {
		Deposit(200)                // A1
		fmt.Println("=", Balance()) // A2
	}()
	// Bob:
	go Deposit(100) // B

	<-time.Tick(time.Microsecond * 400)
	fmt.Println(balance)
}

//* Even when a variable cannot be confined to a single goroutine for its entire lifetime,
//* confinement may still be a solution to the problem of concurrent access

//! Serial confinement -  If  each  stage  of  the  pipeline refrains from accessing the variable after sending it
//! to the next stage, then all accesses to the variable are sequential.

type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}
func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}
}

//! The third way to avoid a data race is to allow many goroutines to access the variable,
//! but only one at a time.
