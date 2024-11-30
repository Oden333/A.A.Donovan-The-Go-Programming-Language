package main

import "fmt"

//! If  all  the  fields  of  a  struct  are  comparable,  the  struct  itself  is  comparable
type Point struct{ X, Y int }

func comparing_Structs() {
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)

	//? Comparable struct types, like other comparable types, may be used as the key type of a map.
	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
}
