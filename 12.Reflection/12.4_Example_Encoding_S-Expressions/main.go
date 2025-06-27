package main

// Go’s standard library supports a variety of formats, including JSON, XML, and ASN.1.
// Another notation that is still widely used is S-expressions, the syntax of Lisp

//* Package encodes arbitrary Go objects using an S-expression notation that supports the following constructs
//* 42             integer
//* "hello"        string (with Go-style quotation)
//* foo            symbol (an unquoted name)
//* (1 2 3)        list   (zero or more items enclosed in parentheses)

// We’ll encode the types of Go using S-expressions as follows.
// Integers and strings are encoded in the obvious way.
// Nil values are encoded as the symbol nil.
// Arrays and slices are encoded using list notation

//
//& Traditionally, S-expressions represent lists of key/value pairs
//& using a single cons cell (key . value) for each pair, rather than a two-element
//& list, but to simplify the decoding we’ll ignore dotted list notation.
