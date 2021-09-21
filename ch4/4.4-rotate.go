package main

import "fmt"



func rotate(s []int, i int) {
	var swap = make([]int, i);
	copy(swap, s[:i])
	// fmt.Printf("%v",swap)
	copy(s, s[i:])
	// fmt.Printf("%v",s)
	copy(s[len(s) - i:], swap)
}


func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate(s,2);
	fmt.Printf("%v",s)
}

