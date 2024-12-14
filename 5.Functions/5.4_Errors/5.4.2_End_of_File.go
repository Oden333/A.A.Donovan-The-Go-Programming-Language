package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Usually, the variety of errors that a function may return is interesting to the end user
// but not to the intervening program logic

// A program must take different  actions  depending  on  the  kind  of  error  that  has  occurred

// Consider  an attempt to read n bytes of data from a file.
// If n is chosen to be the length of the file, any error represents a failure.

//& io package guarantees that any read failure caused by an end-of-file condition
//& is always reported by a distinguished error, io.EOF

func read() error {
	in := bufio.NewReader(os.Stdin)
	for {
		_, _, err := in.ReadRune()
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		// ...use r...
	}
	return nil
}

// For other errors, we may need to report both the quality and quantity of the error,
// so to speak, so a fixed error value will not do.
