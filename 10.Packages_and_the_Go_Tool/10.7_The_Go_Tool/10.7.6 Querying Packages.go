package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

//! The go list tool reports information about available packages

//? An argument to go list may contain the “...” wildcard, which matches any
//? substring of a package’s import path

// The -json flag causes go list to print the entire record of each package in JSON format

// The -f flag lets users customize the output format using the template language of package text/template
// go list -f '{{join .Deps " "}}'

// this command prints the direct imports of each package in the compress subtree of the standard library
// * go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
// compress/bzip2 -> bufio cmp io slices
// compress/flate -> bufio errors fmt io math math/bits sort strconv sync
// compress/gzip -> bufio compress/flate encoding/binary errors fmt hash/crc32 io time
// compress/lzw -> bufio errors fmt io
// compress/zlib -> bufio compress/flate encoding/binary errors fmt hash hash/adler32 io

func logCommandError(context string, err error) {
	ee, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("%s: %s", context, err)
	}
	log.Printf("%s: %s", context, err)
	os.Stderr.Write(ee.Stderr)
	os.Exit(1)
}

// packages returns a slice of package import paths corresponding to slice of
// package patterns.
// See 'go help packages' for different ways of specifying packages.
func packages(patterns []string) []string {
	args := []string{"list", "-f={{.ImportPath}}"}
	for _, pkg := range patterns {
		args = append(args, pkg)
	}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("resolve packages", err)
	}
	return strings.Fields(string(out))
}

func ancestors(packageNames []string) []string {
	targets := make(map[string]bool)
	for _, pkg := range packageNames {
		targets[pkg] = true
	}

	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		logCommandError("find ancestors", err)
	}
	var pkgs []string
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if targets[dep] {
				pkgs = append(pkgs, pkg)
				break
			}
		}
	}
	return pkgs
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	pkgs := ancestors(packages(os.Args[1:]))
	sort.StringSlice(pkgs).Sort()
	for _, pkg := range pkgs {
		fmt.Println(pkg)
	}
}
