package main

//! It is an error to import a package into a file but not refer to the name it defines within that file.

//? However, on occasion we must import a package merely for the side effects of doing so:
//? evaluation of the initializer expressions of its package-level variables and
//? execution of its init functions

// import _ "image/png" // register PNG decoder //* blank import
//? It is most often used to implement a compile-time mechanism whereby
//? the main program can enable optional features by blank-importing additional packages.
import (
	"fmt"
	"image"
	"image/jpeg"

	// _ "image/png" // register PNG decoder
	"io"
	"os"
)

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	// The standard library’s image package exports a Decode function that reads bytes
	// from an io.Reader, figures out which image format was used to encode the data,
	// invokes the appropriate decoder, then returns the resulting image.Image
	if err != nil {
		fmt.Fprintln(os.Stderr, img, kind)
		return err
	}
	fmt.Fprintln(os.Stderr, img, kind)

	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

//! Notice the blank import of image/png. Without that line, the program compiles and
//! links as usual but can no longer recognize or decode input in PNG format:

//? The standard library provides decoders for GIF, PNG, and JPEG, and users may provide others,
//? but to keep executables small, decoders are not included in an application unless explicitly requested
//
//? The image.Decode function consults a table of supported formats
// Each entry in the table specifies four things: the name of the format;
// a string that is a prefix of all images encoded this way, used to detect the encoding;
// a function Decode that decodes an encoded image;
// and another function DecodeConfig that decodes only the image metadata, such as its size and color space.

//! An entry is added to the table by calling image.RegisterFormat,
//! typically from within the package initializer of the supporting package for each format, like this one in image/png

// SEE INTERNAL PKG /snap/go/10888/src/image/png/reader.go:1052
//* package png // image/png

//* func Decode(r io.Reader) (image.Image, error)
//* func DecodeConfig(r io.Reader) (image.Config, error)
//* func init() {
//* 	const pngHeader = "\x89PNG\r\n\x1a\n"
//* 	image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
//* }

//? The database/sql package uses a similar mechanism to let users install just the database drivers they need
// import (
//		 "database/mysql"
//		 _ "github.com/lib/pq"              // enable support for Postgres
//		 _ "github.com/go-sql-driver/mysql" // enable support for MySQL
// )
