// Signed numbers are represented in 2’s-complement form, in which the high-order bit
// is reserved for the sign of the number and the range of values of an n-bit number is
// from −2n−1 to 2n−1−1. Unsigned integers use the full range of bits for non-negative
// values and thus have the range 0 to 2n−1. For instance, the range of int8 is −128 to
// 127, whereas the range of uint8 is 0 to 255.

// Например:int8 (8 бит):
// 01111111(2)= 127 - максимальное число
// 10000000(2)=-128 — минимальное отрицательное число

// Пример 2's complement:
// Прямой код для 5: 00000101.
// Инвертируем все биты: 11111010.
// Добавляем 1: 11111011.
// Так, в памяти -5 будет представлено как 11111011 в системе 2's complement.

package main

import "fmt"

func main() {
	printTrick()
	// bitOpers()

}

// Usually a Printf format string containing multiple % verbs would require the same
// number of extra operands, but the [1] “adverbs” after %  tell  Printf
// to  use  the  first  operand  over  and  over  again.

// Second,  the  #adverb for %o or %x or %X tells Printf
// to emit a 0 or 0x or 0X prefix respectively.
func printTrick() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	ascii := 'a'
	unicode := ' '
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269   ''"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

}
func bitOpers() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b,%d\n", x, int(x)) // "00100010", the set {1,5}
	fmt.Printf("%08b,%d\n", y, int(y)) // "00000110", the set {1,2}
	fmt.Printf("%08b\n", x&y)          // "00000010", theintersection {1}
	fmt.Printf("%08b\n", x|y)          // "00100110", the union	{1, 2, 5}
	fmt.Printf("%08b\n", x^y)          // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y)         // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2,6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0,4}

	// Минимальное значение для int64
	minInt64 := int64(-1 << 63)
	// Мы используем выражение -1 << 63, чтобы сдвинуть единицу на 63 бита влево и
	// установить старший бит в 1, что даёт минимальное значение

	// Максимальное значение для int64
	maxInt64 := int64(^uint64(0) >> 1)
	// Мы используем побитовое отрицание нуля ^uint64(0), чтобы получить число, у
	// которого все биты установлены в 1, то есть максимальное значение для uint64:

	fmt.Printf("%d\n", minInt64)
	fmt.Printf("%d\n", maxInt64)
}
