package main

import "fmt"

func struct_Embedding_and_Anonymous_Fields() {
	wheelInit()
}

type point struct {
	X, Y int
}

// type Circle struct {
// 	Center Point
// 	Radius int
// }
// type Wheel struct {
// 	Circle Circle
// 	Spokes int
// }

// func wheelInit() {
// 	var w Wheel
// 	w.Circle.Center.X = 8
// 	w.Circle.Center.Y = 8
// 	w.Circle.Radius = 5
// 	w.Spokes = 20
// }

// Go lets us declare a field with a type but no name; such fields are called anonymous fields.
// The  type  of  the  field  must  be  a  named  type  or  a  pointer  to  a  named  type

// & We say that a Point is embedded within Circle, and a Circle is embedded within Wheel.
type Circle struct {
	point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func wheelInit() {
	var w Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	//* w = Wheel{8, 8, 5, 20}                       //  compile error: unknown fields
	//* w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} //  compile error: unknown fields

	// The struct literal must follow the shape of the type declaration, so we must use one of
	// the two forms below, which are equivalent to each other
	w = Wheel{
		Circle: Circle{point{8, 8}, 5},
	}
	w = Wheel{
		Circle: Circle{
			point:  point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	// #  adverb  causes  Printf’s  %v  verb  to  display  values  in  a  form similar to Go syntax.
	fmt.Printf("%#v\n", w)
	fmt.Printf("%+v\n", w)
}

//& Because “anonymous” fields do have implicit names, you can’t have two anonymous
//& fields of the same type since their names would conflict.

// But why would you want to embed a type that has no subfields?
// The answer has to do with methods. The shorthand notation used for selecting the
// fields  of  an  embedded  type  works  for  selecting  its  methods  as  well.
// ! In  effect,  the outer struct type gains not just the fields of the embedded type but its methods too.
// This mechanism is the main way that complex object behaviors are composed from
// simpler ones.
//? Composition is central to object-oriented programming in Go, and we’ll explore it further in Section 6.3.
