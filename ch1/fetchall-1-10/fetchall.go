// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nBytes, url)
}

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.Create("times.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot wite file %v", err);
	}

	w := bufio.NewWriter(f)
	for range os.Args[1:] {
		fmt.Fprintln(w, <-ch)  // receive from channel ch
	}


	fmt.Fprintf(w, "%.2fs elapsed\n", time.Since(start).Seconds())
    w.Flush()
}


