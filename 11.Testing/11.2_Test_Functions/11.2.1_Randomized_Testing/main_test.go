package randomizedtesting

import (
	"math/rand/v2"
	"strings"
	"testing"
	"time"
)

// The output of a failing test does not include the entire stack trace at the moment of
// the call to t.Errorf. Nor does t.Errorf cause a panic or stop the execution of
// the test, unlike assertion failures in many test frameworks for other languages.
// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.

func randomPalindrome(rng *rand.Rand) string {
	n := rng.IntN(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.IntN(0x1000)) // random rune up to '\u0999'
		if r%10 == 0 {
			r = 0x20
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

type ss int64

func (s ss) Uint64() uint64 {
	return uint64(s)
}

func TestRandomPalindromes(t *testing.T) {
	seed := ss(time.Now().UTC().UnixNano())
	rng := rand.New(seed)

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
	rng := rand.New(ss(seed))

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
