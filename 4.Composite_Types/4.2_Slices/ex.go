package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

func ex() {
	// ex4_3()
	// ex4_4()
	// ex4_5()
	// ex4_6()
	ex4_7()
}

func ex4_3() {
	zal := []int{0, 1, 2, 3, 4, 5, 6}
	reverseArray(&zal)
	fmt.Println(zal)
}

// To get an array value at a particular indices with a pointer we should use the following form:
// ! (*name)[index]
func reverseArray(s *[]int) {
	length := len(*s) - 1
	for i, j := 0, length; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func ex4_4() {
	zal := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(rotate(zal, 6))
}

func rotate(s []int, idx int) []int {
	return append(s[idx:], s[:idx]...)
}

func ex4_5() {
	zal := []string{"zalupa", "zalupa", "zalupas", "zalupaj", "zalupaj", "zalupad"}
	cleanZ := removeAdjDups(zal)
	fmt.Println(cleanZ)
}

func removeAdjDups(str []string) []string {
	res := str[:0]
	for idx, s := range str {
		if s == str[idx+1] {
			continue
		}
		res = append(res, s)
		if idx == len(str)-2 && str[idx+1] != s {
			res = append(res, str[idx+1])
			break
		}
	}
	return res
}

func ex4_6() {
	fmt.Println(string(removeAdjSpaces([]byte("zal \t \n upa   123  das     w   "))))
}

func removeAdjSpaces(str []byte) []byte {
	var (
		res  = []byte(str)
		cur  rune
		step int
		skip int
	)

	for i := 0; i < len(res); i++ {
		cur, step = utf8.DecodeRune(res[i:])
		if unicode.IsSpace(cur) {
			res[i] = 32 // common space ' '
			for spaceIdx := i + step; spaceIdx < len(res); spaceIdx += step {
				cur, step = utf8.DecodeRune(res[spaceIdx:])
				if !unicode.IsSpace(cur) {
					break
				}
				skip = spaceIdx
			}
			res = append(res[:i], res[skip:]...)
		}
	}
	return res
}

func ex4_7() {
	zal := []byte("ZalupA")
	fmt.Println(unsafe.Pointer(&zal[0]))
	fmt.Println(unsafe.Pointer(&[]rune(string(zal))[0]))

	zal = reverseString([]byte(zal))
	fmt.Println(string(zal))
	fmt.Println(unsafe.Pointer(&zal[0]))
}

func reverseString(s []byte) []byte {
	// var (
	// 	stepL, stepR int
	// 	runeL, runeR rune
	// )
	// for l, r := 0, len(s); l < r; l, r = l+stepL, r-stepR {
	// 	runeR, stepL = utf8.DecodeRune(s[l:])
	// 	runeL, stepR = utf8.DecodeLastRune(s[:r])
	// 	utf8.AppendRune(s[:l], runeR)
	// 	utf8.AppendRune(s[r:], runeL)
	// }

	runes := []rune(string(s))
	for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
		runes[l], runes[r] = runes[r], runes[l]
	}
	return []byte(string(runes))
}
