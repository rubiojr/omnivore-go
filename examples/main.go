package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rubiojr/omnivore-go"
)

func main() {
	client := omnivore.NewClient(omnivore.Opts{Token: getAPIToken()})
	a, err := client.GetArticles(omnivore.SearchOpts{})
	if err != nil {
		log.Fatalf("Failed to fetch articles: %v", err)
	}
	for _, article := range a {
		fmt.Println(article.ID, article.PageType, article.Title, article.PublishedAt, article.Url)
	}
}

func getAPIToken() string {
	return os.Getenv("OMNIVORE_API_TOKEN")
}
