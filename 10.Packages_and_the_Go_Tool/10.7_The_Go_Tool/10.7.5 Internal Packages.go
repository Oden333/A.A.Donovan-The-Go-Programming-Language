package main

//! The package is the most important mechanism for encapsulation in Go programs.
// Unexported identifiers are visible only within the same package, and exported
// identifiers are visible to the world.

// Sometimes, though, a middle ground would be helpful, a way to define identifiers that
// are visible to a small set of trusted packages, but not to everyone
//! The go build tool treats a package specially if its import path contains a path segment named internal.

//? An internal package may be imported only by another package that is
//? inside the tree rooted at the parent of the internal directory.
