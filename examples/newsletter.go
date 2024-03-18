package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rubiojr/omnivore-go"
)

func main() {
	client := omnivore.NewClient(omnivore.Opts{Token: getAPIToken()})
	// https://docs.omnivore.app/using/search.html
	emails, err := client.NewsletterEmails()
	if err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	for _, email := range emails {
		fmt.Println("* " + email.Address)
		fmt.Println("    Name:", email.Name)
		fmt.Println("    Created at:", email.CreatedAt.Format("2006-01-02"))
		fmt.Println("    Subscription count:", email.SubscriptionCount)
		fmt.Println("    Folder:", email.Folder)
	}
	fmt.Println("Total items:", len(emails))
}

func getAPIToken() string {
	return os.Getenv("OMNIVORE_API_TOKEN")
}
