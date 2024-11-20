package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func ex() {
	// ex4_8()
	ex4_9()
}

func ex4_8() {
	// countCharsUnicode()
	fmt.Printf("%[1]c char \n\b%8[1]d unicode code point, \n\b%08[2]b binary view", unicode.MaxRune, []byte(string(unicode.MaxRune)))
}

func ex4_9() {
	wordfreq()
}

func countCharsUnicode() {
	letters := make(map[rune]int)
	digits := make(map[rune]int)
	numbers := make(map[rune]int)
	symbol := make(map[rune]int)
	var utflen [4]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		unicode.IsDigit(r)
		switch {
		case unicode.IsLetter(r):
			letters[r]++
		case unicode.IsDigit(r):
			digits[r]++
		case unicode.IsNumber(r):
			numbers[r]++
		case unicode.IsSymbol(r):
			symbol[r]++
		}

		utflen[n]++
	}
	fmt.Printf("\nletter\tcount\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\ndigits\tcount\n")
	for c, n := range digits {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nnumbers\tcount\n")
	for c, n := range numbers {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nsymbol\tcount\n")
	for c, n := range symbol {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func wordfreq() {
	var (
		words = make(map[string]int)
		token string
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		token = scanner.Text()
		if token == "end" {
			break
		}
		words[token]++
		if err := scanner.Err(); err != nil {
			fmt.Println("Scan error:", err)
		}
	}
	fmt.Fprintf(os.Stdout, "%v", words)
}
