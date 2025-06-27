package main

import (
	"os"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

type Cycle struct {
	Value int
	Tail  *Cycle
}

func TestDisplay(t *testing.T) {
	var i interface{} = 3
	var c Cycle
	// c = Cycle{42, &c}

	type args struct {
		name string
		x    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Mov",
			args: args{
				name: "strangelove",
				x: Movie{
					Title:    "Dr. Strangelove",
					Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
					Year:     1964,
					Color:    false,
					Actor: map[string]string{
						"Dr. Strangelove":            "Peter Sellers",
						"Grp. Capt. Lionel Mandrake": "Peter Sellers",
						"Pres. Merkin Muffley":       "Peter Sellers",
						"Gen. Buck Turgidson":        "George C. Scott",
						"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
						`Maj. T.J. "King" Kong`:      "Slim Pickens",
					},
					Oscars: []string{
						"Best Actor (Nomin.)",
						"Best Adapted Screenplay (Nomin.)",
						"Best Director (Nomin.)",
						"Best Picture (Nomin.)",
					},
				},
			},
		},
		{
			// ! Notice that even unexported fields are visible to reflection.
			name: "File",
			args: args{
				name: "file",
				x:    os.Stderr,
			},
		},
		{
			//! reflect.ValueOf always returns a Value of a concrete type since
			//! it extracts the contents of an interface value
			name: "Interface",
			args: args{
				name: "i",
				x:    i,
			},
		},
		{
			//? A Value obtained indirectly, like this one, may represent any value at all,
			//? including interfaces.
			name: "Interface PTR",
			args: args{
				name: "&i",
				x:    &i,
			},
		},
		{
			//? Display will never terminate if it encounters a cycle in
			//? the object graph, such as this linked list that eats its own tail:
			name: "Linked list - type var",
			args: args{
				name: "c",
				x:    c,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Display(tt.args.name, tt.args.x)
		})
	}
}

func TestMapKeys(t *testing.T) {
	s := map[struct{ x int }]int{
		{1}: 2,
		{2}: 3,
	}
	Display("s", s)
	// Output:
	// Display s (map[struct { x int }]int):
	// s[{x: 2}] = 3
	// s[{x: 1}] = 2

	a := map[[3]int]int{
		{1, 2, 3}: 3,
		{2, 3, 4}: 4,
	}
	Display("a", a)
	// Output:
	// Display a (map[[3]int]int):
	// a[1, 2, 3] = 3
	// a[2, 3, 4] = 4
}
