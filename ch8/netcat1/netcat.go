// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Target struct {
	city string
	host string
}

func connect(t Target) {
	conn, err := net.Dial("tcp", t.host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn, t.city)
}

func main() {
	targets := []Target{}
	for _, arg := range os.Args[1:] {
		args := strings.SplitN(arg, "=", 2)
		t := Target{
			city: args[0],
			host: args[1],
		}
		targets = append(targets, t)
	}
	for _, t := range targets {
		go connect(t)
	}
	time.Sleep(1 * time.Minute)
}

func mustCopy(dst io.Writer, src io.Reader, city string) {
	input := bufio.NewScanner(src)
	for input.Scan() {
		fmt.Println(city, "-", input.Text())
	}
}

//!-
