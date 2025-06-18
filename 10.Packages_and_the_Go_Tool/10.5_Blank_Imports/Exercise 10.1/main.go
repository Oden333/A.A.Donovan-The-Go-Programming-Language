package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

var format = flag.String("f", "jpeg", "defines output image file format")

func main() {
	flag.Parse()

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "img %q: %v\n", kind, err)
		os.Exit(1)
	}

	var out *os.File
	if out, err = os.Create(fmt.Sprintf("out.%s", *format)); err != nil {
		fmt.Fprintf(os.Stderr, "file create: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	switch *format {
	case "jpeg":
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: 95}); err != nil {
			fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
			os.Exit(1)
		}
	case "png":
		if err := png.Encode(out, img); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	case "gif":
		if err := gif.Encode(out, img, nil); err != nil {
			fmt.Fprintf(os.Stderr, "png: %v\n", err)
			os.Exit(1)
		}
	}

}
