package main

import (
	"fmt"
	"runtime"
)

// The go build command compiles each argument package.
//? If the package is a library, the result is discarded; this merely checks that the package is free of compile errors.
//
//? If the package is named main, go build invokes the linker to create an executable in the current directory;
// the name of the executable is taken from the last segment of the package’s import path

//! go build command builds the requested package and all its
//! dependencies, then throws away all the compiled code except the final executable, if any.

// Since each directory contains one package, each executable program, or command in
// Unix terminology, requires its own directory

// Packages may also be specified as a list of file names,
// though this tends to be used only for small programs and one-off experiments.
//? If the package name is main, the executable name comes from the basename of the first .go file

// * The go install command is very similar to go build, except that it saves the
// * compiled code for each package and command instead of throwing it away

func goos_goarch() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

// Some packages may need to compile different versions of the code for certain
// platforms or processors, to deal with low-level portability issues or to provide
// optimized versions of important routines, for instance.
//
//! If a file name includes an operating system or processor architecture name like
//* net_linux.go or asm_amd64.s,
//! then the go tool will compile the file only when building for that target.

//! Special comments called build tags give more fine-grained control.
// For example, if a file contains this comment:
//* go:build linux || darwin - will compile it only when building for Linux or Mac OS X,
//* +build linux darwin

//* +build ignore - never to compile the file
