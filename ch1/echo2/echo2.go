// Echo2 “prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main()  {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
