package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

func Exs() {

	// fmt.Println(ex3_11("213332111000"))
	fmt.Println(ex3_12("zalupa", "zzzzalupa"))

}

// comma inserts commas in a non-negative decimal integer string.
func ex3_10(s string) string {
	res := bytes.Buffer{}
	offset := len(s) % 3
	if offset == 0 {
		offset = 3
	}
	var j = offset - 3
	if j < 0 {
		j = 0
	}
	for i := offset; i < len(s); i += 3 {
		res.WriteString(fmt.Sprintf("%s,", s[j:i]))
		j = i
	}

	res.WriteString(s[len(s)-3:])
	return res.String()
}

func ex3_11(s string) string {
	idx := strings.Index(s, ".")
	if idx == -1 {
		return ex3_10(s)
	}
	res := bytes.Buffer{}
	res.WriteString(ex3_10(s[:idx]))
	res.WriteRune('.')
	res.WriteString(ex3_10(s[idx+1:]))
	return res.String()
}

func ex3_12(s1 string, s2 string) bool {
	start := time.Now()
	// time.Sleep(1 * (time.Millisecond))
	defer func() { fmt.Println(time.Since(start).Nanoseconds()) }()
	chars := make(map[rune]bool)
	for _, char := range s1 {
		if !chars[char] {
			chars[char] = true
		} else {
			continue
		}
		if !strings.Contains(s2, string(char)) {
			return false
		}
	}
	for _, char := range s2 {
		if !chars[char] {
			chars[char] = true
		} else {
			continue
		}
		if !strings.Contains(s1, string(char)) {
			return false
		}
	}
	return true
}
