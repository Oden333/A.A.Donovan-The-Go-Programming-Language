package externaltestpackages

//* one of the tests in net/url is an example demonstrating the interaction between
//* URLs and the HTTP client library

//* We resolve the "cycle import" problem by declaring the test function in an external test package, that
//* is, in a file in the net/url directory whose package declaration reads package url_test.

//? The extra suffix _test is a signal to go test that it should build an
//? additional package containing just these files and run its tests.

// ? By avoiding import cycles, external test packages allow tests, especially
// ? integration tests (which test the interaction of several components),
// ? to import other packages freely, exactly as an application would
//
// Sometimes an external test package may need privileged access to the internals of the
// package under test, if for example a white-box test must live in a separate package to
// avoid an import cycle
// ! In such cases, we use a trick: we add declarations to an in-package _test.go
// ! file to expose the necessary internals to the external test.

//& An application that often fails when it encounters new but valid inputs is called buggy;
//
//& A test that spuriously fails when a sound change was made to the program is called brittle
