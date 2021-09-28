package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseByte(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// reverse reverses a slice of ints in place.
func reverse(s []byte) {

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		reverseByte(s[i : i+size])
		fmt.Printf("%d\t%q\t%d\n", i, r, size)
		i += size

	}
	reverseByte(s)
}

func main() {
	s := []byte("Hسello س")
	reverse(s)
	fmt.Printf("%s", s)

}
