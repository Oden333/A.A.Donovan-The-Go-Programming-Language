package main

import (
	"fmt"

	"gopl.io/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.FToC(tempconv.CToF(tempconv.BoilingC)))
}
