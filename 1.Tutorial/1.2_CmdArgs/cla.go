// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	t1 := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "\n"
	}

	fmt.Println("The name of command that invoked os.Args -", fmt.Sprintf("%s", s), fmt.Sprintf("\nTime elapsed %.10f", time.Since(t1).Seconds()))
	t1 = time.Now()
	fmt.Println("The name of command that invoked os.Args -", strings.Join(os.Args, "\n"), fmt.Sprintf("\nTime elapsed %.10f", time.Since(t1).Seconds()))

}
