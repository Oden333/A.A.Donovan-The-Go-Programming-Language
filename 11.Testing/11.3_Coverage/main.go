package main

// The degree to which a test suite exercises the package under test is called the test’s coverage

// Coverage can’t be quantified directly—the dynamics of all but the most
// trivial programs are beyond precise measurement—but there are heuristics that can
// help us direct our testing efforts to where they are more likely to be useful.

//& Go’s cover tool - integrated into go test -
//& to measure statement coverage and help identify obvious gaps in the tests.
//! The statement coverage of a test suite is the fraction of source statements
//! that are executed at least once during the test.

//? The go tool command runs one of the executables from the Go toolchain.
//? These programs live in the directory $GOROOT/pkg/tool/${GOOS}_${GOARCH}.
//? Thanks to go build, we rarely need to invoke them directly

//* Achieving  100%  statement  coverage  sounds  like  a  noble  goal,  but  it  is  not  usually
//* feasible in practice, nor is it likely to be a good use of effort. Just because a statement
//* is executed does not mean it is bug-free;
// statements containing complex expressions must  be  executed
//  many  times  with  different  inputs  to  cover  the  interesting  cases.
// Some statements, like the panic statements above, can never be reached.

// Coverage  tools  can  help  identify  the  weakest  spots,  but  devising  good  test  cases
// demands the same rigorous thinking as programming in general.
