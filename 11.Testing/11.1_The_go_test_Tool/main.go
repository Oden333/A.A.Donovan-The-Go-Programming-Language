package thegotesttool

// The go test subcommand is a test driver for Go packages that are organized
// according to certain conventions.

//? In a package directory, files whose names end with _test.go are not part of the package
//? ordinarily built by go build but are a part of it when built by go test

//! Within *_test.go files, three kinds of functions are treated specially:
//! tests, benchmarks, and examples.

//? A test function, which is a function whose name begins
//? with Test, exercises some program logic for correct behavior
//
//* go test calls the test function and reports the result, which is either PASS or FAIL
//
//
//? A benchmark function has a name beginning with Benchmark
//? and measures the performance of some operation;
//
//* go test reports the mean execution time of the operation
//
//
//? An example function, whose name starts with Example,
//? provides machine-checked documentation

//! The go test tool scans the *_test.go files for these special functions,
//! generates a temporary main package that calls them all in the proper way,
//! builds and runs it, reports the results, and then cleans up
