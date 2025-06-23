package testfunctions

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// ! Each test file must import the testing package.
// ! Test functions have the following signature
func TestName(t *testing.T) {
	// ...
}

//* The t parameter provides methods for reporting test failures and
//* logging additional information.

//? A package named main ordinarily produces an executable program,
//? but it can be imported as a library too

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout // modified during testing

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}

//? Notice that the test code is in the same package as the production code. Although the
//? package name is main and it defines a main function, during testing this package
//? acts as a library that exposes the function TestEcho to the test driver;
//? its main function is ignored.
