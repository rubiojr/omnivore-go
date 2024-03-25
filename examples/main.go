package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rubiojr/omnivore-go"
)

func main() {
	client := omnivore.NewClient(omnivore.Opts{Token: getAPIToken()})
	// https://docs.omnivore.app/using/search.html
	a, err := client.Search(context.Background(), omnivore.SearchOpts{Query: "in:all sort:saved"})
	if err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	for _, searchItem := range a {
		fmt.Println("* " + searchItem.Title)
		fmt.Println("    Labels:", labelsToString(searchItem.Labels))
		fmt.Println("    Folder:", searchItem.Folder)
		fmt.Println("    URL:", searchItem.Url)
		fmt.Println("    Published at:", searchItem.PublishedAt.Format("2006-01-02"))
		fmt.Println("    Saved at:", searchItem.SavedAt.Format("2006-01-02"))
	}
	fmt.Println("Total items:", len(a))
}

func labelsToString(labels []omnivore.Label) string {
	l := []string{}
	for _, label := range labels {
		l = append(l, label.Name)
	}
	return fmt.Sprintf("[%s]", strings.Join(l, ", "))
}
func getAPIToken() string {
	return os.Getenv("OMNIVORE_API_TOKEN")
}
