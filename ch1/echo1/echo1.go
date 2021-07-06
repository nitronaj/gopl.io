// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func JoinArgs(args []string) string  {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func main()  {
	start := time.Now()

	var s string
	s = JoinArgs(os.Args)
	fmt.Println(s)

	// sub time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}
