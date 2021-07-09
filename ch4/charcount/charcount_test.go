// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharCount(t *testing.T) {
	values := []struct {
		value   string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
		letters int
		spaces  int
		symbols int
	}{{
		"this is a test", map[rune]int{'t': 3, 'h': 1, 'i': 2, 's': 3, 'a': 1, 'e': 1, ' ': 3}, [utf8.UTFMax + 1]int{0, 14, 0, 0}, 0, 11, 3, 0},
		{"João", map[rune]int{'J': 1, 'o': 2, 'ã': 1}, [utf8.UTFMax + 1]int{0, 3, 1, 0}, 0, 4, 0, 0}}
	for _, a := range values {
		counts, utflen, invalid, letters, spaces, symbols := CharCount(strings.NewReader(a.value))
		if len(counts) != len(a.counts) {
			t.Errorf("CharCount(%q).counts = %v", a.value, counts)
		}
		for r, v := range counts {
			if v != a.counts[r] {
				t.Errorf("CharCount(%q).counts[%q] = %v", a.value, r, v)
			}
		}
		for i, v := range utflen {
			if v != a.utflen[i] {
				t.Errorf("CharCount(%q).utflen[%d] = %d", a.value, i, v)
			}
		}
		if invalid != a.invalid {
			t.Errorf("CharCount(%q).invalid = %v", a.value, invalid)
		}
		if letters != a.letters {
			t.Errorf("CharCount(%q).letters = %v", a.value, letters)
		}
		if spaces != a.spaces {
			t.Errorf("CharCount(%q).spaces = %v", a.value, spaces)
		}
		if symbols != a.symbols {
			t.Errorf("CharCount(%q).symbols = %v", a.value, symbols)
		}
	}
}

//!-
