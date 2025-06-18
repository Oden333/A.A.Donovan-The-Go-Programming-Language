package main

// Each package is identified by a unique string called its import path.
//? Import paths are the strings that appear in import declarations
import (
	_ "encoding/json"
	_ "fmt"
	_ "math/rand"
	// _ "github.com/go-sql-driver/mysql"
	// _ "golang.org/x/net/html"
)

//! The Go language specification doesn’t define the meaning of these strings or
//! how to determine a package’s import path, but leaves these issues to the tools.

//! For packages you intend to share or publish, import paths should be globally unique

//? To avoid conflicts, the import paths of all packages other than those from the
//? standard library should start with the Internet domain name of the organization that
//? owns or hosts the package; this also makes it possible to find packages.
