package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

//? Reflection is provided by the reflect package. It defines Type and Value.

//? A Type represents a Go type.
//& It is an interface with many methods for discriminating among types and inspecting
//& their components, like the fields of a struct or the parameters of a function.

//* The sole implementation of reflect.Type is the type descriptor,
//* the same entity that identifies the dynamic type of an interface value.

// * The reflect.TypeOf function accepts any interface{} and returns its
// * dynamic type as a reflect.Type:
func t1() {
	// The TypeOf(3) call assigns the value 3 to the interface{} parameter
	t := reflect.TypeOf(3)  // a reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

func t2() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	//* fmt.Printf provides a shorthand, %T, that uses reflect.TypeOf internally:
	fmt.Printf("%T\n", 3) // "int"
}

//? A reflect.Value can  hold  a  value  of  any  type.
//& The  reflect.ValueOf  function  accepts  any interface{} and returns a reflect.Value
//& containing the interface’s dynamic value

func v1() {
	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) //* NOTE: "<int Value>"
	// unless the Value holds a string, the result of the String method reveals only the type

	//* Use  the  fmt  package’s  %v  verb,  which  treats  reflect.Values specially.
	t := v.Type()           // a reflect.Type
	fmt.Println(t.String()) // "int"
}

//? The  inverse  operation  to  reflect.ValueOf  is  the reflect.Value.Interface  method.
//? It  returns  an  interface{}  holding the same concrete value as the reflect.Value:

func StoI() {
	v := reflect.ValueOf(3) // a reflect.Value
	x := v.Interface()      // an interface{}
	i := x.(int)            // an int
	fmt.Printf("%d\n", i)   // "3"
}

//? There are infinitely many types, there are only a finite number of kinds of type:
//
//? - the basic types Bool, String, and all the numbers;
//? - the aggregate types Array and Struct;
//? - the reference types Chan, Func, Ptr,  Slice,  and Map;
//? - Interface types;
//? - and finally Invalid, meaning no value at all. (The zero value of a reflect.Value has kind Invalid.)
//

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64,
		reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct,	reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}
