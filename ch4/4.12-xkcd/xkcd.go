package main

import (
	"fmt"
	"log"

	"mygopl.io/ch4/4.12-xkcd/comics"
)

func main() {
	// var comicNumber string
	// args := os.Args[1:]

	// if len(args) > 0 {
	// 	comicNumber = args[0]
	// }

	comics, err := comics.FetchAllComics()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, comic := range comics {

		fmt.Printf("#%-5d | %.20s | %2.55s \n", comic.Num, comic.Title, comic.Transcript)
		fmt.Println("-----------------------------------------------------------")
	}

}
