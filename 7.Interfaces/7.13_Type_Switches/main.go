package main

import (
	"database/sql"
	"fmt"
)

//? Interfaces  are  used  in  two  distinct  styles.
//? In  the  first  style,  exemplified  by
//? io.Reader,  io.Writer,  fmt.Stringer,  sort.Interface, http.Handler, and error,

//! 1. An interface’s methods express the similarities of the concrete types
//! that satisfy the interface but hide the representation details and intrinsic
//! operations  of  those  concrete  types.

//! 2. The second style - discriminated unions interfaces - exploits the ability of an interface value
//! to hold values of a variety of concrete types and considers the interface to be the union of those types.
//? Type assertions are used to discriminate among these types dynamically and treat each case differently

//* In object-oriented programming, these two styles as subtype polymorphism and ad hoc polymorphism

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	result, err := db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	// ...
	//* Constructing queries this way helps avoid SQL injection attacks in which
	//* an adversary takes control of the query by exploiting improper quotation of input data.

	fmt.Fprintf(nil, "%v%v", result, err)
}

// A type switch enables a multi-way branch based on the interface value’s dynamic type
func x(x interface{}) {
	switch x.(type) {
	case nil: // ...
	case int, uint: // ...
	case bool: // ...
	case string: // ...
	default: // ...

	}

	//! As with an ordinary switch statement cases are considered in order and,
	//
	// when a match is found, the case’s body is executed. Case order becomes significant when
	// one or more case types are interfaces, since then there is a possibility of two cases matching.

	//! The position of the default case relative to the others is immaterial.

	//! No fallthrough is allowed.

	switch x := x.(type) { // binds the extracted value to a new variable within each case
	default:
		fmt.Println(x) /* ... */
	}
}

// ? Like  a  switch  statement,  a  type  switch  implicitly  creates  a
// ? lexical block, so the declaration of th/e new variable called x does not conflict with a
// ? variable  x  in  an  outer  block.
// ? Each  case  also  implicitly  creates  a  separate  lexical block
func sqlQuote(x interface{}) string {

	//! Although  the  type  of  x  is interface{},
	//! we  consider  it  a  discriminated  union  of  int,  uint,  bool,string, and nil.

	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x) // (not shown)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuoteString(x string) string { return "" }
