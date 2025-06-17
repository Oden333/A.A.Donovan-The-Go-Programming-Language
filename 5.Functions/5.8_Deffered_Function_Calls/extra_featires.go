package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// the world
// the sun is shining
// tasty bebra
func DirectPassing() {
	s := "bebra"

	defer func(s string) {
		fmt.Println("tasty", s)
	}(s)

	s = "the sun"

	defer func(s string) {
		fmt.Println(s, "is shining")
	}(s)

	s = "the world"
	fmt.Println(s)
}

func ClosurePassing() {
	s := "bebra"

	defer func() {
		fmt.Println("tasty", s)
	}()

	s = "the sun"

	defer func() {
		fmt.Println(s, "is shining")
	}()

	s = "the world"
	fmt.Println(s)
}

func UnnamedReturn() int {
	num := 100

	defer func(num int) {
		fmt.Printf("num value in the 1st defer `before`: %d\n", num)
		num = 777
		fmt.Printf("num value in the 1st after `after`: %d\n", num)
	}(num)

	num = 200

	defer func() {
		fmt.Printf("num value in the 2nd defer `before`: %d\n", num)
		num = 300
		defer func() { fmt.Printf("num value in the 2nd defer `after`: %d\n", num) }()
	}()

	fmt.Printf("num value before function exiting: %d\n", num)
	return num
}

// Defer with closure 	close phuong-secrets.txt: file already closed
// Func result: 		close phuong-secrets.txt: file already closed
func NamedReturn() (err error) {
	const (
		filePath = "fetch.go"
	)

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("opening the file %q: %w", filePath, err)
	}
	defer func() {
		err = errors.Join(err, f.Close())
		fmt.Println("Defer with closure", err)
	}()

	sr := bufio.NewScanner(f)
	for sr.Scan() {
		_ = sr.Text()
	}

	// It's made deliberately to show the changes in defer statement
	f.Close()

	return
}

func main() {
	fmt.Println(fn())
	fmt.Println(fn1())
	fmt.Println(fn2())
}

func fn() int {
	var i = 34
	defer func() { i = 13 }()
	return i
}

func fn1() (i int) {
	i = 34
	defer func() { i = 1 }()
	return
}

func fn2() (i int) {
	i = 34
	defer func(i int) { i = 1 }(i)
	return
}
