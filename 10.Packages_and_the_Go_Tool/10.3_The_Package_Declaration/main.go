package main

// A package declaration is required at the start of every Go source file.
// Its main purpose is to determine the default identifier for that package (called the package
// name) when it is imported by another package

//* Conventionally, the package name is the last segment of the import path, and as a
//* result, two packages may have the same name even though their import paths
//* necessarily differ.

//* There are three major exceptions to the “last segment” convention.
//
//! The first is that a package defining a command (an executable Go program)
//! always has the name main, regardless of the package’s import path.
//! This is a signal to go build that it must invoke the linker to make an executable file.
//
//! The second exception is that some files in the directory may have the suffix _test
//! on their package name if the file name ends with _test.go.
//? Such a directory may define two packages: the usual one, plus another one called an external test package.
//* The _test suffix signals to go test that it must build both packages, and it
//* indicates which files belong to each package.

//! The third exception is that some tools for dependency management append version
//! number suffixes to package import paths, such as "gopkg.in/yaml.v2".
// The package name excludes the suffix, so in this case it would be just yaml
