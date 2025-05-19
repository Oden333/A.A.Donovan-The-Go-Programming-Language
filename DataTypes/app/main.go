package main

import (
	. "DataTypes/bit_Vector"
	"fmt"
)

func bitset() {
	var x IntSet
	x.Add(0)
	x.Add(1)
	x.Add(4)
	x.Add(144)
	// fmt.Println(x.String()) // "{1 9 144}"
	// fmt.Fprintf(os.Stdout, "%b", x)
	// fmt.Println(x.Len(), x.String())
	x.Remove(4)

	y := x.Copy()

	x.Remove(144)

	x.Add(10)
	y.Add(100)

	fmt.Println("x: len", x.Len(), x.String())
	fmt.Println("y: len", y.Len(), y.String())

	c := x.Copy()
	c.IntersectWith(y)

	d := x.Copy()
	d.DifferenceWith(y)

	fmt.Println("intersection:\t\t", c)
	fmt.Println("defference:\t\t", d)
	fmt.Println("symmetric difference:\t", x.SymmetricDifference(y))

	fmt.Println("x int slice:\t\t", x.Elems())
	fmt.Println("y int slice:\t\t", y.Elems())

}
