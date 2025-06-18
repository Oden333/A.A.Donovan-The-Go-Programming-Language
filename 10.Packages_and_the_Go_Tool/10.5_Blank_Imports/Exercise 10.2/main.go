package main

import (
	_ "archive/tar"
	_ "archive/zip"
	"bufio"
	"flag"
	"fmt"
	"os"
)

var file = flag.String("f", "", "defines an archive file for reading")

func main() {
	flag.Parse()
	var fName string = *file

	if fName == "" {
		fmt.Println("You haven't typed filename as flag. Type it below:")
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if fName = scanner.Text(); fName != "" {
				fmt.Println("Oppening", fName)
			}
			break
		}
	}

	file, err := os.Open(fName)
	if err != nil {
		fmt.Println("file open:", err)
		os.Exit(1)
	}
	defer file.Close()

}
