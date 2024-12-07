package main

// func name(parameter-list) (result-list) {
//     body
// }
//
// The parameter list specifies the names and types of the function’s parameters, which
// are  the  local  variables  whose  values  or  arguments  are  supplied  by  the  caller
//
// The result list specifies the types of the values that the function returns. If the function
// returns one unnamed result or no results at all, parentheses are optional and usually
// omitted.

//? Like  parameters,  results  may  be  named.  In  that  case,  each  name  declares  a  local
//? variable initialized to the zero value for its type.

// A sequence of parameters or results of the same type can be factored  so  that  the  type  itself  is  written  only  once.

// func f(i, j, k int, s, t string){ /*	... */ }

//? Here are four ways to declare a function with two parameters and one result, all of
//? type int. The blank identifier can be used to emphasize that a parameter is unused.
func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

//& The  type  of  a  function  is  sometimes  called  its  signature.
//& Two  functions  have  the same type or signature if they have the same sequence of parameter types and the same sequence of result types.

//* The names of parameters and results don’t affect the type, nor does whether or not they were declared using the factored form.

//* Every  function  call  must  provide  an  argument  for  each  parameter,
//* in  the  order  in which the parameters were declared.

// Parameters are local variables within the body of the function, with their initial values
// set to the arguments supplied by the caller.

//! Arguments are passed by  value, so the function receives a copy of each argument; modifications to the copy do not affect the caller.
// However, if the argument contains some kind of reference,
//* like a pointer, slice, map, function, or channel,
// then the caller may  be  affected  by  any  modifications  the  function  makes  to  variables
// indirectly referred to by the argument.

func main() {

}

// You may occasionally encounter
//! A function declaration without a body, indicating that
//! the function is implemented in a language other than Go.

//& Such a declaration defines the function signature
func Sin(x float64) float64 // implemented in assembly language
