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

func TestMarshall(t *testing.T) {
	var i interface{} = 3

	type args struct {
		name string
		x    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Movit Struct",
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := Marshal(tt.args.x)
			if err != nil {
				t.Log(err)
			}
			t.Logf("\n-------------------\nMarshalling %s:\n%s\n", tt.name, v)
		})
	}
}
