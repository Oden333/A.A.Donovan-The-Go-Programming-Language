package storage

//? One way of categorizing tests is by the level of knowledge they require of the internal
//? workings of the package under test

//* A black-box test assumes nothing about the package other than what is exposed
//* by its API and specified by its documentation; the package’s internals are opaque.

//* In contrast, a white-box test has privileged access to the internal functions and data structures
//* of the package and can make observations and changes that an ordinary client cannot.

// For example, a white-box test can check that the invariants of the package’s data types
// are maintained after every operation.
//
// TestIsPalindrome calls only the exported function IsPalindrome and is thus a black-box test.
// TestEcho  calls the  echo  function  and  updates  the  global  variable  out,  both  of  which  are
// unexported, making it a white-box test
