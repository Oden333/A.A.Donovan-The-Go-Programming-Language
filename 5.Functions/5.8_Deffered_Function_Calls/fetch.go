package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = f.ReadFrom(resp.Body)
	// n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return local, n, err
}

//* The deferred call to resp.Body.Close tempting to use a second deferred call, to f.Close, to close the local file
//* but this would be subtly wrong because os.Create opens a file for writing, creating it as needed.

//* On many file systems, notsably NFS, write errors are not reported immediately but may
//* be postponed until the file is closed. Failure to check the result of the close operation
//* could  cause  serious  data  loss  to  go  unnoticed.

//* However, if both io.Copy and f.Close fail, we should prefer to report the error from io.Copy
//* since it occurred first and is more likely to tell us the root cause.
