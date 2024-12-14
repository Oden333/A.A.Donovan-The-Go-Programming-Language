package main

// This panic is a sure sign of a bug in the calling code and should never
// happen in a well-written program.

// Errors  are   an  important  part  of  a  package’s  API  or  an  application’s  user interface,
// and failure is just one of several expected behaviors.

//& This is the approach Go takes to error handling.

//? A  function  for  which  failure  is  an  expected  behavior  returns  an  additional  result,
//? conventionally the last one.

//? If the failure has only one possible cause, the result is a boolean, usually called ok
//* value, ok := cache.Lookup(key)
//* if !ok {
// ()...cache[key] does not exist...()
//* }

//? More often, and especially for I/O, the failure may have a variety of causes for which the caller will need an explanation.
//? In such cases, the type of the additional result is error.

//! The built-in type error is an interface type
// error may be nil or non-nil, that nil implies success and non-nil implies failure, and
// that a non-nil error has an error message string which we can obtain by calling its
// Error  method  or  print  by  calling  fmt.Println(err)  or fmt.Printf("%v", err)

//& Usually when a function returns a non-nil error, its other results are undefined and should be ignored
// However, a few functions may return partial results in error cases.
// For example, if an error occurs while reading from a file, a call to Read returns the
// number of bytes it was able to read and an error value describing the problem

//? Go’s approach sets it apart from many other languages
//? in which failures are reported using  exceptions,  not  ordinary  values.
