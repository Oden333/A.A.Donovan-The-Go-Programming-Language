package main

// net package provides  the  components  for  building  networked  client  and  server  programs  that
// communicate  over  TCP,  UDP,  or  Unix  domain  sockets.
// Clock1 is a TCP server that periodically writes the
import (
	"flag"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func sequential() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//? The listener’s Accept method blocks until an incoming connection request is made,
		//? then returns a net.Conn object representing the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		log.Printf("Accepted:%s", conn.RemoteAddr())
		handleConn(conn) // handle one connection at a time
	}
}

func concurrent() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections 		concurrently
	}
}

// handleConn  function  handles  one  complete  client  connection
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(tz).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var (
	sep    = flag.String("port", "8080", "vairiable defines server port")
	tzFlag = flag.String("TZ", "Russia/Moscow", "vairiable defines time zone")

	tz *time.Location
)

func ex8_1() {
	flag.Parse()
	if *sep == "" {
		log.Printf("Port flag is not set, will use defaul 8080 port")
	}
	if *tzFlag == "" {
		log.Printf("Port tz is not set, will use defaul MSC tz")
	}

	var err error
	tz, err = time.LoadLocation(*tzFlag)
	if err != nil {
		log.Fatalf("error reading tzone %q", *tzFlag)
	}

	listener, err := net.Listen("tcp", strings.Join([]string{"localhost", *sep}, ":"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections 		concurrently
	}

}

// func main() {
// concurrent()
// ex8_1()
// }
