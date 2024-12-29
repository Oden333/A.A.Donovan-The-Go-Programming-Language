package main

// Although there is no universally accepted definition of object-oriented programming, for  our  purposes,
//! An  object  is  simply  a  value  or  variable  that  has  methods,  and
//! A method is a function associated with a particular type.

//? An object-oriented program is one that uses methods to express the properties and operations of each data structure
//? so that clients need not access the object’s representation directly.

// A method is declared with a variant of the ordinary function declaration
// in which an extra  parameter  appears  before  the  function  name.

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
