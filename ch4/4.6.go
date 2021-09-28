package main

import (
	"fmt"
	"unicode"
)


func squashSpace(s []byte) []byte {
	i := 0;
	if(unicode.IsSpace(rune(s[i])) ) {
		s[i] = 32
	}

	for _, v := range s[1:] {
		if !unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(v))  {
			i++
			s[i] =  32
		} else {
			i++
			s[i] = v
		}
	}

	return s[:i]
}





func main() {
	s := []byte("\tHello\t, friend s\n")

	s1 := squashSpace(s)

	fmt.Printf("%c", s1)
}
