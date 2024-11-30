package main

import (
	"fmt"
	"time"
)

func main() {
	//! A struct is an aggregate data type that groups together zero or more
	//! named values of arbitrary types as a single entity
	// typeInit()
	// trees()
	// struct_Literals()
	struct_Embedding_and_Anonymous_Fields()
}

func typeInit() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}

	var dilbert Employee
	//Because dilbert is a variable, its fields are variables too, so we may assign to a field
	dilbert.Salary -= 5000
	// or take its address and access it through a pointe
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert.Position)

	// The dot notation also works with a pointer to a struct
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player 0)"
	// The last statement is equivalent to
	(*employeeOfTheMonth).Position += " (proactive team player 1)"
	fmt.Println(dilbert.Position)

	//& Field order is significant to type identity. Had we also combined the declaration of the
	//& Position field (also a string), or interchanged Name and Address, we would be
	//& defining a different struct type. Typically we only combine the declarations of related
	//& fields.

	//? The name of a struct field is exported if it begins with a capital letter; this is Go’s
	//? main access control mechanism.
	//! A struct type may contain a mixture of exported andunexported fields.

	//! Every instance of a struct is just a part of memory with a particular size.
	//! When we get an OBJECT POINTER of a struct we get a pointer to the memory area.
	//! When we get an OBJECT FIELD POINTER we get the shifted pointer in the same memory area of this object.

	// 	If we have a []StructType slice and put an object into it consequently passing the slice into a function that changes
	// a field of an object and returns a pointer to that object, we won't change the original object, because while appending
	// or putting with slice[index] the object into []StructType, the slice will have the copy of the object.
	// 	! TO DEFEAT THIS WE SHOULD CREATE AN []*StructType SLICE TO STORE THE LINKS TO THE ORIGINAL OBJECTS. !
}

func trees() {
	//* A named struct type S can’t declare a field of the same type S: an aggregate value
	//* cannot contain itself. (An analogous restriction applies to arrays.) But S may declare a
	//* field of the pointer type *S, which lets us create recursive data structures like linked
	//* lists  and  trees.
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// The struct type with no fields is called the empty struct, written struct{}. It has
// size  zero  and  carries  no  information  but  may  be  useful  nonetheless.  Some  Go
// programmers use it instead of bool as the value type of a map that represents a set,
//! to emphasize that only the keys are significant, but the space saving is marginal and
// the syntax more cumbersome, so we generally avoid it.
