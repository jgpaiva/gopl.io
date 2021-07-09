// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 222.

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(loc *time.Location, c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var (
	tz   = flag.String("tz", "", "timezone")
	port = flag.Uint("p", 8011, "port")
)

func main() {
	flag.Parse()
	loc, err := time.LoadLocation(*tz)
	if err != nil {
		log.Fatalf("error loading timezone: %v", err)
	}
	fmt.Println("Running server for timezone", loc, "at port", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("main: %v", err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(loc, conn) // handle connections concurrently
	}
	//!-
}
