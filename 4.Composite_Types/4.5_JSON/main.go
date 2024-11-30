package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//! JavaScript Object Notation (JSON) is a standard notation for sending and receiving structured information.
	//? XML (§7.14),  ASN.1,  Google’s  Protocol  Buffers
	//? serve  similar  purposes  and  each  has  its  niche,  but
	//? because of its simplicity, readability, and universal support, JSON is the most widely used.

	//? Go has excellent support for encoding and decoding these formats, provided by the standard  library  packages
	//? encoding/json,  encoding/xml, encoding/asn1

	//! JSON is an encoding of JavaScript values—strings, numbers, booleans, arrays, and objects as Unicode text.

	//* The  basic  JSON  types  are
	// numbers  (in  decimal  or  scientific  notation),
	// booleans (true  or  false),
	// strings,  which  are  sequences  of  Unicode  code  points enclosed  in  double  quotes,
	// with  backslash  escapes  using  a  similar  notation  to  Go,
	// though JSON’s \Uhhhh numeric escapes denote UTF-16 codes, not runes.

	// These basic types may be combined recursively using JSON arrays and objects.
	// A JSON  array  is  an  ordered  sequence  of  values,  written  as  a  comma-separated  list
	// enclosed in square brackets; JSON arrays are used to encode Go arrays and slices.

	//? A JSON  object  is  a  mapping  from  strings  to  values,  written  as  a  sequence  of
	//? name:value pairs separated by commas and surrounded by braces;
	//& JSON objects are used to encode Go maps (with string keys) and structs. For example

	// boolean            true
	// number             -273.15
	// string             "She said \"Hello,  \""
	// array              ["gold", "silver", "bronze"]
	// object             {"year": 1980,
	//                     "event": "archery",
	//                     "medals": ["gold", "silver", "bronze"]}
	reccomendationsApp()
}

func reccomendationsApp() {

	type Movie struct {
		Title string
		//! A field tag is a string of metadata associated at compile time with the field of a struct
		Year  int  `json:"released"`
		Color bool `json:"color,omitempty"`

		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}

	// Data structures like this are an excellent fit for JSON, and it’s easy to convert in both directions.
	//& Converting a Go  data  structure  like  movies to JSON is called marshaling.

	// Marshaling is done by json.Marshal:
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	// fmt.Printf("%+v\n", movies)
	//? Marshal produces a byte slice containing a very long string with no extraneous white space;

	// json.MarshalIndent produces neatly indented output (аккуратный вывод с отступами)
	data1, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data1)
	//?	Marshaling uses the Go struct field names as the field names for the JSON objects through reflection
	//& Only exported fields are marshaled, which is why we chose capitalized names for all the Go field names

	//! A field tag is a string of metadata associated at compile time with the field of a struct:
	// Интерпретируется как разделенный пробелами разделенный пробелами список пар «ключ: значение»
	//! A field tag may be any literal string, but it is conventionally interpreted as
	//! a space-separated list of key:"value" pairs
	// The  json  key  controls  the behavior of the encoding/json package,
	// and other encoding/... packages follow this convention.

	//! "omitempty" indicates that no JSON output should be produced if the field has the zero value

	// The  inverse  operation  to  marshaling,  decoding  JSON  and  populating  a  Go  data structure, is called unmarshaling
	var titles []struct {
		Title string
	}
	// By  defining  suitable  Go  data  structures  in  this  way,
	// we  can  select  which parts of the JSON input to decode and which to discard.
	if err := json.Unmarshal(data1, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} 	{Bullitt}]"

	// Many web services provide a JSON interface

}
