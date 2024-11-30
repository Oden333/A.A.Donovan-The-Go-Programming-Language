package main

func struct_Literals() {
	type Point struct {
		X, Y int
	}
	// This form requires a value to be specified for every field, in the right order
	// p := Point{1, 2}
	// It burdens the writer (and reader) with remembering exactly what the fields are,
	// and it makes the code fragile should the set of fields later grow or be reordered

	// usage examples of such form
	// image.Point{x,  y}
	// color.RGBA{red, green, blue, alpha}
}
