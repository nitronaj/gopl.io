package main

import (
	"fmt"
)

// func rotate(s []string, i int) {
// 	var swap = make([]string, i);
// 	copy(swap, s[:i])
// 	// fmt.Printf("%v",swap)
// 	copy(s, s[i:])
// 	// fmt.Printf("%v",s)
// 	copy(s[len(s) - i:], swap)
// }


func set(s []string) []string {
	i := 0
	for index, value := range s {
		if (index + 1) < len(s) &&  value == s[index + 1] {
			continue
		}
		s[i] = value
		i++
	}
	return s[:i]
}

func removeAdjacentDuplicates(s []string) []string {
    i := 0
    for _, m := range s[1:] {
        if s[i] != m {
            i++
            s[i] = m
        }
    }
    return s[:i+1]
}

func filteDup(s []string) []string {
    if len(s) == 0 {
        return s
    }

    i := 0
    for j := 1; j < len(s); j++ {
        if s[j] != s[i] {
			i++
			fmt.Printf("%d %v\n", i, s[:i])
            s[i] = s[j]
        }
    }
    return s[:i+1]
}

func main() {
	s := []string {"Hello", "Hello", "Hi", "No", "No", "No", "MMMM", "Hello"}
	// s2 := set(s)
	// fmt.Printf("%v", s2)


	s3 := filteDup(s)
	fmt.Printf("%v", s3)
}

