package examplefunctions

import (
	"unicode"
)

// The third kind of function treated specially by go test is an example function, one
// whose name starts with Example.
//
// func ExampleIsPalindrome() {
//
// It has neither parameters nor results. Here’s an example function for IsPalindrome:

//? The primary one is documentation: a good example can be a more succinct or intuitive way to
//? convey the behavior of a library function than its prose description, especially
//? when used as a reminder or quick reference

// ? The second purpose is that examples are executable tests run by go test
// ? If the example function contains a final // Output: comment like the  one  above,  the
// ? test  driver  will  execute  the  function  and  check  that  what  it  printed  to  its  standard
// ? output matches the text within the comment.

// ? The third purpose of an example is hands-on experimentation

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
