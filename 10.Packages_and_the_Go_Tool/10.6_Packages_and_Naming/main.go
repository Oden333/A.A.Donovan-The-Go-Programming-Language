package main

//? Naming packages and their members
//
// Avoid  choosing  package  names  that  are  commonly  used  for  related  local
// variables, or you may compel the package’s clients to use renaming imports, as with the path package.
//
// Avoid  package  names  that  already  have  other  connotations.
//
//
//
//
// We can identify some common naming patterns. The strings package provides a number of independent functions for manipulating strings

// package strings
func Index(needle, haystack string) int

type Replacer struct { /* ... */
}

func NewReplacer(oldnew ...string) *Replacer

type Reader struct{}

func NewReader(s string) *Reader

//* The word string does not appear in any of their names.
//* Clients refer to them as strings.Index, strings.Replacer, and so on
