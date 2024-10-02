// Within a Go program, every package is identified by a unique string called its import
// path.  These  are  the  strings  that  appear  in  an  import  declaration  like
// "gopl.io/ch2/tempconv".  The  language  specification  doesn’t  define  where
// these strings come from or what they mean;  it’s up to the tools to interpret them.
// When using the go tool (Chapter 10), an import path denotes a directory containing
// one or more Go source files that together make up the package.
// In addition to its import path, each package has a package name, which is the short
// (and  not  necessarily  unique)  name  that  appears  in  its  package  declaration.  By
// convention, a package’s name matches the last segment of its import path, making it
// easy  to  predict  that  the  package  name  of  gopl.io/ch2/tempconv  is tempconv

package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
	if len(os.Args) < 2 {
		var num int
		fmt.Print("Enter nums count: ")
		fmt.Scanf("%d", &num)

		nums := make([]float64, num)
		for i := 0; i < num; i++ {
			fmt.Scanln(&nums[i])
		}

		for _, i := range nums {
			f := tempconv.Fahrenheit(i)
			c := tempconv.Celsius(i)
			k := tempconv.Kelvin(i)
			fmt.Printf("%s = %s, %s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
		}
		// fmt.Println(nums)

	}

}
