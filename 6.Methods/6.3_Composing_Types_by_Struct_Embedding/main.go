package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	//? We could have defined ColoredPoint as a struct of three fields, but instead we
	//? embedded  a  Point  to  provide  the  X  and  Y  fields.
	Point

	Color color.RGBA
}

var (
	red  = color.RGBA{255, 0, 0, 255}
	blue = color.RGBA{0, 0, 255, 255}
)

// we can select the fields of  ColoredPoint  that  were
// contributed  by  the  embedded  Point  without mentioning Point:
func q() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	fmt.Println(cp.Y) // "2"
	// A similar mechanism applies to the methods of Point. We can call methods of the
	// embedded  Point  field  using  a  receiver  of  type  ColoredPoint,  even  though
	// ColoredPoint has no declared methods
}

//! The  methods  of  Point  have  been  promoted  to  ColoredPoint.  In  this  way,
//! embedding  allows  complex  types  with  many  methods  to  be  built  up  by  the
//! composition of several fields, each providing a few methods.

//? The type of an anonymous field may be a pointer to a named type, in which case
//? fields  and  methods  are  promoted  indirectly  from  the  pointed-to  object.

func qq() {
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}
	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share 	the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}

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

func qqq() {
	type ColoredPoint struct {
		Point
		color.RGBA
	}
	//! When the compiler resolves a selector such as p.ScaleBy to a method,
	//!
	//! it first looks for a	directly declared method named ScaleBy,
	//! then for methods promoted once from ColoredPoint’s  embedded  fields,
	//! then  for  methods  promoted  twice  from embedded fields within Point and RGBA, and so on

}

//? it’s possible and sometimes useful for unnamed
//? struct types to have methods too.

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

//? The new variable gives more expressive names to the variables related to the cache,
//? and because the sync.Mutex field is embedded within it,
//! its Lock and Unlock methods are promoted to the unnamed struct type, allowing us to lock the cache
//? with a self-explanatory syntax.
