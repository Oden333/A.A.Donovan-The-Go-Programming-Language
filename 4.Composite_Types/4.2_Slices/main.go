package main

import "fmt"

//! Slices represent variable-length sequences whose elements all have the same type. A slice type is written []T, where the elements have type T

func main() {
	//! A  slice  has  three  components:  a  pointer,  a length,  and  a  capacity.
	//& The  pointer  points  to  the  first  element  of  the  array  that  is
	//& reachable  through  the  slice,  which  is  not  necessarily  the  array’s  first  element.  The
	//& length is the number of slice elements; it can’t exceed the capacity, which is usually
	//& the number of elements between the start of the slice and the end of the underlying
	//& array
	// months()
	// passing()
	// sliceInit()
	// theAppendFunction()
	// in_Place_Slice_Techniques()
	ex()
}

func months() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	// The slice operator s[i:j], where 0 ≤ i ≤ j ≤ cap(s), creates a new slice that
	// refers to elements i through j-1 of the sequence s, which may be an array variable,
	// a pointer to an array, or another slice.
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	//?	Slicing  beyond  cap(s)  causes  a  panic,
	//& but  slicing  beyond  len(s)  extends  the slice, so the result may be longer than the original
	// fmt.Println(summer[:20])    // panic: out of range
	endlessSummer := summer[:5] //! extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func passing() {
	a := [...]int{0, 1, 2, 3, 4, 5} // array
	reverse(a[:])                   //passing slice
	fmt.Println(a)
}

// A slice literal looks like an array literal, a sequence of values separated by commas
// and surrounded by braces, but the size is not given.
//! This implicitly creates an array variable of the right size and yields a slice that points to it.

//& Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contain the same elements.
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// Given how natural this “deep” equality test is, and that it is no more costly at run time
// than the == operator for arrays of strings, it may be puzzling that slice comparisons
// do  not  also  work  this  way.  There  are  two  reasons  why  deep  equivalence  is
// problematic. First, unlike array elements, the elements of a slice are indirect, making it
// possible for a slice to contain itself. Although there are ways to deal with such cases,
// none is simple, efficient, and most importantly, obvious.
// Second, because slice elements are indirect, a fixed slice value may contain different
// elements  at  different  times  as  the  contents  of  the  underlying  array  are  modified.
// Because a hash table such as Go’s map type makes only shallow copies of its keys, it
// requires  that  equality  for  each  key  remain  the  same  throughout  the  lifetime  of  the
// hash table. Deep equivalence would thus make slices unsuitable for use as map keys.
// For  reference  types  like  pointers  and  channels,  the  ==  operator  tests  reference
// identity,  that  is,  whether  the  two  entities  refer  to  the  same  thing.  An  analogous
// “shallow” equality test for slices could be useful, and it would solve the problem with
// maps, but the inconsistent treatment of slices and arrays by the == operator would be
// confusing. The safest choice is to disallow slice comparisons altogether.

func sliceInit() {
	//! The only legal slice comparison is against nil, as in
	// if summer == nil { /* ... */
	// }

	// 	The zero value of a slice type is nil. A nil slice has no underlying array. The nil slice has length and capacity zero, but
	// there are also non-nil slices of length and capacity zero, such as []int{} or make([]int, 3)[3:]

	// As  with  any  type  that  can have  nil  values,  the  nil  value  of a  particular  slice  type  can  be  written
	// using  a conversion expression such as []int(nil)
	var s []int    // len(s) == 0, s == nil
	s = nil        // len(s) == 0, s == nil
	s = []int(nil) // len(s) == 0, s == nil
	s = []int{}    // len(s) == 0, s != nil
	// z := make([]T, len)
	// v := make([]T, len, cap) // same as make([]T, cap)[:len]
	// o := make([]T, cap)[:len]
	fmt.Fprintf(nil, "%v", s)
	//? Under the hood, make creates an unnamed array variable and returns a slice of it; the
	//? array is accessible only through the returned slice. In the first form, the slice is a view
	//? of the entire array. In the second, the slice is a view of only the array’s first len
	//? elements, but its capacity includes the entire array. The additional elements are set
	//? aside for future growth
}
