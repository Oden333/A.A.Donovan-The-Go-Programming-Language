package main

//! A Go source file may contain zero or more import declarations immediately after
//! the package declaration and before the first non-import declaration.

// Imported packages may be grouped by introducing blank lines;
//? such groupings usually indicate different domains
import (
	_ "fmt"
	_ "html/template"
	_ "os"
	//
	// "golang.org/x/net/html"
	// "golang.org/x/net/ipv4"
	// If we need to import two packages whose names are the same, like math/rand and crypto/rand
	//
	// "crypto/rand"
	// mrand "math/rand" // alternative name mrand avoids conflict
	//
	//* The  alternative  name  affects  only  the  importing  file.  Other  files,  even  ones  in  the
	//* same package, may import the package using its default name, or a different name
)
