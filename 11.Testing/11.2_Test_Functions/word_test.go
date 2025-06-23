package testfunctions

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

// The -v flag prints the name and execution time of each test in the package

//		-run flag, whose argument is a regular expression, causes go test to run
// 					only those tests whose function name matches the pattern
//* go test -v -run="French|Canal"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// The output of a failing test does not include the entire stack trace at the moment of
// the call to t.Errorf. Nor does t.Errorf cause a panic or stop the execution of
// the test, unlike assertion failures in many test frameworks for other languages.
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.

// ? 11.2.1_Randomized_Testing
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		if r%10 == 0 {
			r = 0x20
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}
func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	t.Logf("Random seed: %d", seed)

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		t.Logf("Random pal: %q", p)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	t.Logf("Random seed: %d", seed)

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		p = strings.Join([]string{p, "asda"}, "")
		t.Logf("Random pal: %q", p)

		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
