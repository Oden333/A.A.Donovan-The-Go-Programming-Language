package main

import (
	"flag"
	"fmt"
	"time"

	t "gopl.io/ch7/tempconv"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

//? The flag.Duration function creates a flag variable
//? of type time.Duration and allows the user to specify the duration in a variety of
//? user-friendly formats, including the same notation printed by the String  method.

func q() {
	flag.Parse()
	fmt.Println(period)
	fmt.Printf("Sleeping for %s...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ t.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = t.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = t.FToC(t.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = t.Celsius(kToC(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func kToC(f float64) float64 {
	return (f - 273.15)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value t.Celsius, usage string) *t.Celsius {
	f := celsiusFlag{value}
	//? The call to Var adds the flag to the application’s set of command-line flags, the global variable
	flag.CommandLine.Var(&f, name, usage)

	//! The  call  to  Var assigns  a  *celsiusFlag  argument  to  a  flag.Value  parameter,
	//! causing  the compiler to check that *celsiusFlag has the necessary methods

	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

// Exercise 7.7: Explain why the help message contains °C when the default value of 20.0 does not
// bcs String implementation formats the value in this way
// func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
