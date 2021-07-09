// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func CharCount(in io.RuneReader) (counts map[rune]int, utflen [utf8.UTFMax + 1]int, invalid int, letters int, spaces int, symbols int) {
	counts = map[rune]int{}

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsSpace(r) {
			spaces++
		}
		if unicode.IsSymbol(r) {
			symbols++
		}
		counts[r]++
		utflen[n]++
	}
	return
}

func main() {
	// here's a string to get unicode chars ããããããããêêñ
	f, err := os.Open("main.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v", err)
		os.Exit(-1)
	}

	in := bufio.NewReader(f)
	counts, utflen, invalid, letters, spaces, symbols := CharCount(in)

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if letters > 0 {
		fmt.Printf("letters: %d\n", letters)
	}
	if spaces > 0 {
		fmt.Printf("spaces: %d\n", spaces)
	}
	if symbols > 0 {
		fmt.Printf("symbols: %d\n", symbols)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
