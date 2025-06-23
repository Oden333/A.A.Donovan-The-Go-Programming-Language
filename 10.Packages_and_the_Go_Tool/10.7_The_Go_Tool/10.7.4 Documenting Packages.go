package main

import "io"

// Go style strongly encourages good documentation of package APIs.
// Each declaration of an exported package member and the package declaration itself
// should be immediately preceded by a comment explaining its purpose and usage.

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, format string, a ...interface{}) (int, error)

//! Good documentation need not be extensive, and documentation is no substitute for simplicity
