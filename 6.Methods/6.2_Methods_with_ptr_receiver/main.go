package main

import (
	"fmt"
	"math"
	"net/url"
)

//? Because calling a function makes a copy of each argument value, if a function needs
//? to update a variable, or if an argument is so large that we wish to avoid copying it, we
//? must pass the address of the variable using a pointer

//! Convention dictates that if any method of Point has a pointer
//! receiver, then all methods of Point should have a pointer receiver

	type Point struct{ X, Y float64 }

	// traditional function
	func Distance(p, q Point) float64 {
		return math.Hypot(q.X-p.X, q.Y-p.Y)
	}

	// same thing, but as a method of the Point type
	func (p *Point) Distance(q Point) float64 {
		return math.Hypot(q.X-p.X, q.Y-p.Y)
	}

	func (p *Point) ScaleBy(factor float64) {
		p.X *= factor
		p.Y *= factor
	}

//! method declarations are not permitted
//! on named types that are themselves pointer types

//* type P *int
//* func (P) f() { /* ... */ } // compile error: invalid receiver type

func t() {
	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	p.ScaleBy(2)
	fmt.Println(p) // "{2, 4}"

	//? If the receiver p is a variable of type Point but the method requires
	//? a *Point receiver, we can use this shorthand
	p.ScaleBy(2)
	//? and  the  compiler  will  perform  an  implicit  &p on the variable

	//? We cannot call a *Point method on a non-addressable Point receiver, because
	//? there’s no way to obtain the address of a temporary value.
	// Point{1, 2}.ScaleBy(2) // compile error: can't take

	//? But we can call a Point method like Point.Distance with a *Point receiver
	pptr.Distance(p)
	(*pptr).Distance(p)
	//? The compiler inserts an implicit * operation for us.

}

// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func q() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)
	m = nil
	fmt.Println(m.Get("item")) // ""
	//? In  the  final  call  to  Get,  the  nil  receiver  behaves  like  an  empty  map.
	//? We  could equivalently  have  written  it  as
	url.Values(nil).Get("item")
	//? but nil.Get("item")  will  not  compile  because  the  type  of  nil
	//? has  not  been determined

	m.Add("item", "3") // panic: assignment to entry in nil map
}

func main() {
	// q()
}
