package main

import (
	"fmt"
	"math"
	"time"
)

// ! Constants  are  expressions  whose  value  is  known  to  the  compiler  and  whose
// ! evaluation is guaranteed  to occur  at compile  time, not  at run  time.
const (
// arr [3]int = [3]int{1, 2, 3} - invalid constant type
)

func temp() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	fmt.Println(FlagUp, FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 	false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 	true"
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool {
	return v&FlagUp ==
		FlagUp
}
func TurnDown(v *Flags) {
	*v &^= FlagUp
}
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool {
	return v&
		(FlagBroadcast|FlagMulticast) != 0
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776              (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424     (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

// В Go константы (объявленные через const) представляют собой безразмерные числа ("untyped" numbers),
// пока не приведены к определенному типу.
// Это позволяет компилятору работать с константами произвольной точности во время компиляции,
// вплоть до максимального значения типа float или complex.
// Таким образом, константы могут содержать значения,
// превышающие 64-битное целое число, так как Go оптимизирует их на этапе компиляции.

func main() {
	// temp()
	// fmt.Println(GiB, 1<<30)

	// fmt.Println(string(YiB)) //& Error
	// fmt.Println(YiB / ZiB)
	// ex3_13()
	// var f float64 = 212
	// fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	// fmt.Println(5 / 9 * (f - 32))     // "0";   5/9 is an untyped integer, 0
	// fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	//! Only constants can be untyped
	var f float64 = 3 + 0i // untyped complex -> float64
	//? Equal to: var f float64 = float64(3 + 0i)
	f = 2 // untyped integer -> float64
	//? Equal to: f = float64(2)
	f = 1e123 // untyped floating-point -> float64
	//? Equal to: f = float64(1e123)
	f = 'a' // untyped rune -> float64
	//? Equal to: f = float64('a')
	fmt.Println(f)

	const (
		deadbeef = 0xdeadbeef // untyped int with value 3735928559

		// d        = int32(deadbeef)   // compile error: constant overflows int32
		// e        = float64(1e309)    // compile error: constant overflows float64
		// ff        = uint(-1)          // compile error: constant underflows uint
	)
	fmt.Printf("%[1]d 10\n%[1]b 2\n%[1]x 16\n", deadbeef)
	fmt.Println(uint32(deadbeef))  // uint32 with value 3735928559
	fmt.Println(float32(deadbeef)) // float32 with value 3735928576 (rounded up)
	fmt.Println(float64(deadbeef)) // float64 with value 3735928559 (exact)

}

func ex3_13() {
	const (
		_      = 1e1
		KB     = 1e3
		MB     = 1e6
		GB     = 1e9
		TB     = 1e12
		PB     = 1e15
		EB     = 1e18
		ZB     = 1e21
		YB     = 1e24
		zalupa = 10000000000000000000000000000900000.0568335868356
	)
	fmt.Println(KB, MB, GB, TB, PB)
	// var x float64 = zalupa
	fmt.Println(float32(math.Pi))
}
