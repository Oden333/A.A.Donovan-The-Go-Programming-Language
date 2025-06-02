package main

import "io"

// encoding/xml package provides a similar to encoding/json API.
// This approach is convenient when we want to construct a representation of the document tree,
// but that’s unnecessary for  many  programs.
//? The  encoding/xml  package  also  provides  a  lower-level token-based API for decoding XML.

// In the token-based style, the parser consumes
// the  input  and  produces  a  stream  of  tokens,  primarily  of  four  kinds
// StartElement,  EndElement,  CharData,  and  Comment—each  being
// a concrete  type  in  the  encoding/xml  package

// Each  call  to (*xml.Decoder).Token returns a token
// package xml
type Name struct {
	Local string // e.g., "Title" or "id"
}
type Attr struct { // e.g., name="value"
	Name  Name
	Value string
}

// A Token includes StartElement, EndElement, CharData,
// and Comment, plus a few esoteric types (not shown).
// ? The Token interface, which has no methods, is also an example of a discriminated union
// ! The purpose of a traditional interface like io.Reader is to hide details of the
// ! concrete  types  that  satisfy  it  so  that  new  implementations  can  be  created;
// ! each concrete type is treated uniformly.
//
// ! By contrast, the set of concrete types that satisfy a discriminated  union  is
// ! fixed  by  the  design  and  exposed,  not  hidden.  Discriminated
// ! union types have few methods; functions that operate on them are expressed as a set
// ! of cases using a type switch, with different logic in each case.
type Token interface{}

type StartElement struct { // e.g., <name>
	Name Name
	Attr []Attr
}
type EndElement struct{ Name Name } // e.g., </name>
type CharData []byte                // e.g., <p>CharData</p>
type Comment []byte                 // e.g., <!-- Comment -->
type Decoder struct {               /* ... */
}

func NewDecoder(io.Reader) *Decoder
func (*Decoder) Token() (Token, error) // returns next Token in sequence
