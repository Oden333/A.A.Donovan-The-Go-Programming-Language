package main

import (
	_ "archive/tar"
	_ "archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Define a generic archive file-reading function capable of reading ZIP
// files  (archive/zip)  and  POSIX  tar  files  (archive/tar).  Use  a  registration
// mechanism similar to the one described above so that support for each file format can be plugged in using blank imports.

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

	// reader,err := Open(file)
	// ...
}

type format struct {
	name, magic string
	magicOffset int
	reader      NewReader
}

type NewReader func(*os.File) (io.Reader, error)

var formats []format

// We could probably just try opening readers instead of checking magic numbers.
func RegisterFormat(name, magic string, magicOffset int, f NewReader) {
	formats = append(formats, format{name, magic, magicOffset, f})
}

func Open(file *os.File) (io.Reader, error) {
	var found *format
	r := bufio.NewReader(file)
	for _, f := range formats {
		p, err := r.Peek(f.magicOffset + len(f.magic))
		if err != nil {
			continue
		}
		if string(p[f.magicOffset:]) == f.magic {
			found = &f
			break
		}
	}
	if found == nil {
		return nil, fmt.Errorf("open archive: can't determine format")
	}
	_, err := file.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, fmt.Errorf("open archive: %s", err)
	}
	return found.reader(file)
}
