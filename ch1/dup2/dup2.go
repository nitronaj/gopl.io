// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
func countFileLines(f *os.File, counts map[string]int, filename string, fileNames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNames[input.Text()] += " " + filename
	}
	// NOTE: ignoring potential errors from input.Err()
}

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countFileLines(f, counts, arg, fileNames)
			f.Close()
		}
	}

	uniqueFilename := make(map[string]int)

	var fileNamesArr [10]string
	for line, n := range counts {
		if n > 1 {
			fileList := strings.Split(fileNames[line], " ")[1:]
			for _, fileArg := range fileList {
				uniqueFilename[fileArg]++;
			}
			for file, _ := range uniqueFilename {
				append(fileNamesArr, file)
			}

			fmt.Printf("%d\t%s\t%s\n", n, fileNamesArr.Join(" "), line)
		}
	}

}
