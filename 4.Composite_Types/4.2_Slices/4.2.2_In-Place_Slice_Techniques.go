package main

import "fmt"

func in_Place_Slice_Techniques() {
	// q := []string{"asd", "", "zalupa"}
	// zal := nonempty2(q)
	// fmt.Println(zal, len(zal), q)
	// fmt.Println(zal[:3], q)
	// s := []int{5, 6, 7, 8, 9}
	// fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	// fmt.Println(s)

}

// Nonempty is an example of an in-place slice algorithm.
// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]    // zero-length slice of	original
	fmt.Println(cap(out)) // inhetits the capacity of parent slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
