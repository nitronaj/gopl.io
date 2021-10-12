// Issues prints a table of GitHub issues matching the search terms.

// repo:golang/go is:open json decoder

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mygopl.io/ch4/4.10-ex-struct-json-issues/github"
)

func main() {
	itemsByCategory := make(map[string][]*github.Issue)
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: len: %d\n", result.TotalCount, len(result.Items))

	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {

		switch {
		case item.CreatedAt.After(oneMonth):
			itemsByCategory["LESS_THAN_MONTH"] = append(itemsByCategory["LESS_THAN_MONTH"], item)

		case item.CreatedAt.After(oneYear):
			itemsByCategory["LESS_THAN_YEAR"] = append(itemsByCategory["LESS_THAN_YEAR"], item)

		default:
			itemsByCategory["MORE_THAN_YEAR"] = append(itemsByCategory["MORE_THAN_YEAR"], item)
		}

	}

	for cat, items := range itemsByCategory {
		fmt.Printf("Category: %s\tissues:%d\n", cat, len(items))
		fmt.Println("-----------------------------------------------------------")
		for _, item := range items {
			fmt.Printf("#%-5d | %.10s | %9.9s  | %.55s \n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		}
		fmt.Println("-----------------------------------------------------------")
	}
}
