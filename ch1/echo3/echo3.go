package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func JoinArgs(args []string) string  {
	return strings.Join(args[1:], " ")
}

func main()  {
	start := time.Now()
	fmt.Println(JoinArgs(os.Args))

	// sub time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}
