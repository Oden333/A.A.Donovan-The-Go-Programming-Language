package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	Ex1_4()
}

// Чтобы прекратить ввод в процессе чтения данных с консоли, можно использовать комбинацию клавиш, которая сигнализирует о конце ввода (EOF). В разных операционных системах эти комбинации могут отличаться:

// На Linux/macOS: комбинация клавиш Ctrl+D.
// На Windows: комбинация клавиш Ctrl+Z, а затем Enter.
// После ввода этой комбинации, метод input.Scan() завершит чтение ввода, и программа продолжит выполнение оставшегося кода.

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from
	input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",
					err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) (ex bool) {
	input := bufio.NewScanner(f)
	mergeMap := make(map[string]int)
	for input.Scan() {
		mergeMap[input.Text()]++
		if mergeMap[input.Text()] > 1 {
			ex = true
		}
	}
	for k, v := range mergeMap {
		counts[k] += v
	}
	// NOTE: ignoring potential errors from
	input.Err()
	return
}

// ReadFile returns a byte slice that must be converted into a string so it can be
// split  by  strings.Split.  We  will  discuss  strings  and  byte  slices  at  length  in
// Section 3.5.4.
// Under  the  covers,  bufio.Scanner,  ioutil.ReadFile,  and
// ioutil.WriteFile use the Read and Write methods of *os.File, but it’s
// rare that most programmers need to access those lower-level routines directly. The
// higher-level functions like those from bufio and io/ioutil are easier to use.
// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated
// line occurs.
func Dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			str := strings.Split(line, " ")
			for _, s := range str {
				counts[s]++
			}
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Ex1_4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		fmt.Println(files)
		var dupsExtended bool
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",
					err)
				continue
			}
			dupsExtended = countLines(f, counts)
			if dupsExtended {
				fmt.Print(f.Name(), " ")
			}
			f.Close()
		}
	}
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
