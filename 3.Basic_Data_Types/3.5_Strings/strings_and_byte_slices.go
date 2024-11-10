package main

import (
	"bytes"
	"fmt"
	"strings"
)

//* Because strings are immutable, building up strings incrementally can involve a lot of allocation and copying.
//* In such cases, it’s more efficient to use the bytes.Buffer type

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
// func basename(s string) string {
// 	// Discard last '/' and everything before.
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '/' {
// 			s = s[i+1:]
// 			break
// 		}
// 	}
// 	// Preserve everything before last '.'.
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '.' {
// 			s = s[:i]
// 			break
// 		}
// 	}
// 	return s
// }

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
func Ch3_5_4() {
	//& Strings can be converted to byte slices and back again:
	// s := "abc"
	//& Conceptually, the []byte(s) conversion allocates a new byte array holding a copy
	//& of  the  bytes  of  s,  and  yields  a  slice  that  references  the  entirety  of  that  array.
	// b := []byte(s)
	//& The conversion from byte slice back to	string with string(b) also makes a copy,
	//& to ensure immutability of the resulting	string s2
	// s2 := string(b)
	// fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	Exs()
}
