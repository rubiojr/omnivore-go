package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rubiojr/omnivore-go"
)

func main() {
	client := omnivore.NewClient(omnivore.Opts{Token: getAPIToken()})
	subs, err := client.Subscriptions()
	if err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	for _, sub := range subs {
		fmt.Println("* " + sub.Name)
		fmt.Println("    Folder:", sub.Folder)
		fmt.Println("    Created at:", sub.CreatedAt.Format("2006-01-02"))
		fmt.Println("    Last fetched at:", sub.LastFetchedAt.Format("2006-01-02"))
		fmt.Println("    Description:", sub.Description)
		fmt.Println("    Newsletter email:", sub.NewsletterEmail)
		fmt.Println("    Refreshed at:", sub.RefreshedAt.Format("2006-01-02"))
		fmt.Println("    Count:", sub.Count)
		fmt.Println("    Icon:", sub.Icon)
		fmt.Println("    Is private:", sub.IsPrivate)
		fmt.Println("    Auto add to library:", sub.AutoAddToLibrary)
		fmt.Println("    Fetch content:", sub.FetchContent)
		fmt.Println("    Failed at:", sub.FailedAt.Format("2006-01-02"))
		fmt.Println("    URL:", sub.Url)
	}
	fmt.Println("Total items:", len(subs))
}

func getAPIToken() string {
	return os.Getenv("OMNIVORE_API_TOKEN")
}
