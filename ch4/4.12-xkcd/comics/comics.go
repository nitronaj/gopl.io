package comics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func getComicURL(comicNumber int) string {
	if comicNumber == 0 {
		return "https://xkcd.com/info.0.json"

	}
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
}

func GetComic(comicNumber int) (*Comic, error) {
	comicUrl := getComicURL(comicNumber)

	resp, err := http.Get(comicUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("comic not found: %s", resp.Status)
	}

	var result Comic

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil

}

func FetchAllComics() ([]*Comic, error) {
	var comics []*Comic
	lastComic, err := GetComic(0)
	if err != nil {
		return nil, err
	}

	numberOfComics := lastComic.Num

	f, err := os.Create("times.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot write file %v", err)
	}
	w := bufio.NewWriter(f)

	fmt.Printf("Get comics ... %d\n", numberOfComics)
	end := numberOfComics - 2525

	fmt.Fprintln(w, "[")
	for i := 1; i <= end; i++ {
		comic, err := GetComic(i)
		if err != nil {
			fmt.Printf("\n%s\n", err)
			continue
		}

		fmt.Printf("%d ", i)
		if i%20 == 0 {
			fmt.Println()
		}
		comics = append(comics, comic)
		json.NewEncoder(w).Encode(comic)
		if i < end {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprintln(w, "]")
	w.Flush()

	return comics, err
}
