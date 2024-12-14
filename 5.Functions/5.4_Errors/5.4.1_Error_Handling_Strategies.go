package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// ! First, and most common, is to propagate the error, so that a failure in a subroutine
// ! becomes a failure of the calling routine.
// * resp, err := http.Get(url)
// * if err != nil {
// *     return nil, err
// * }
// The fmt.Errorf  function  formats  an  error  message  using  fmt.Sprintf  and
// returns  a  new  error  value
// * doc, err := html.Parse(resp.Body)
// * resp.Body.Close()
// * if err != nil {
// *     return nil, fmt.Errorf("parsing %s as HTML: %v",
// * url, err)
// * }
// In general, the call f(x) is responsible for reporting the attempted operation f and
// the  argument  value  x  as  they  relate  to  the  context  of  the  error.
// & The  caller  is responsible for adding further information that it has but the call f(x) does not, such
// & as the URL in the call to html.Parse above.
//

// ! For errors that represent transient or unpredictable problems, it may make sense to retry the failed operation,
// & possibly  with  a  delay  between  tries,  and  perhaps  with  a  limit  on  the  number  of
// & attempts or the time spent trying before giving up entirely.
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// ! Third, if progress is impossible, the caller can print the error and stop the program
// & gracefully, but this course of action should generally be reserved for the main package of a program.
// & Library functions should usually propagate errors to the caller, unless
// & the error is a sign of an internal inconsistency—that is, a bug.
func errHandle() {
	// url := "https://google.com"
	url := os.Args[1]
	// if err := WaitForServer(url); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	// 	os.Exit(1)
	// }
	// A more convenient way to achieve the same effect
	if err := WaitForServer(url); err != nil {
		log.Fatalf("Site is down: %v\n", err)
	}
	// 2006/01/02 15:04:05 Site is down: no such domain: bad.gopl.io
	// 	For a more attractive output, we can set the prefix used by the log package to the
	// name of the command, and suppress the display of the date and time:
	log.SetPrefix("wait: ")
	log.SetFlags(0)
	// All log functions append a newline if one is not already present.
	// if err := Ping(); err != nil {
	// 	log.Printf("ping failed: %v; networking disabled",
	// 		err)
	// }
}

func ignore() error {
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v",
			err)
	}
	//! And fifth and finally, in rare cases we can safely ignore an error entirely
	// ...use temp dir...
	os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned	periodically
	return nil
}

// The  call  to  os.RemoveAll  may  fail,  but  the  program  ignores  it  because  the
// operating  system  periodically  cleans  out  the  temporary  directory.  In  this  case,
// discarding the error was intentional

//& Error  handling  in  Go  has  a  particular  rhythm.  After  checking  an  error,  failure  is
//& usually dealt with before success.
//& If failure causes the function to return, the logic for success is not indented
//& within an else block but follows at the outer level.

//& Function stend  to  exhibit  a  common  structure,  with  a  series  of  initial  checks  to  reject  errors,
//& followed by the substance of the function at the end, minimally indented.
