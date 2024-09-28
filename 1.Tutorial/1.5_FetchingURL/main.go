// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		// Ex 1.8
		if !strings.Contains(url, "http://") || !strings.Contains(url, "https://") {
			url = strings.Join([]string{"https://", url}, "")
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)

		// Ex 1.7
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: copying %s: %v\n", url, err)
			os.Exit(1)
		}

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		os.Stdout.Write([]byte(url))
		os.Stdout.Sync()
		// Ex 1.9
		os.Stdout.Write([]byte(resp.Status))

	}
}
