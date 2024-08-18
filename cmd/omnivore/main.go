package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rubiojr/omnivore-go"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "omnivore",
		Usage: "Omnivore API client",
		Action: func(*cli.Context) error {
			return nil
		},
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List available articles",
				Action: func(cCtx *cli.Context) error {
					listSaved(cCtx)
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "long",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func listSaved(ctx *cli.Context) {
	longFormat := ctx.Bool("long")
	client := omnivore.NewClient(omnivore.Opts{Token: getAPIToken()})
	// https://docs.omnivore.app/using/search.html
	a, err := client.Search(omnivore.SearchOpts{Query: "in:all sort:saved"})
	if err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	for _, searchItem := range a {
		if longFormat {
			formatLong(&searchItem)
		} else {
			formatShort(&searchItem)
		}
	}
	fmt.Println("Total items:", len(a))
}

func formatLong(searchItem *omnivore.SearchItem) {
	fmt.Println("* " + searchItem.Title)
	fmt.Println("    Labels:", labelsToString(searchItem.Labels))
	fmt.Println("    Folder:", searchItem.Folder)
	fmt.Println("    URL:", searchItem.Url)
	fmt.Println("    Published at:", searchItem.PublishedAt.Format("2006-01-02"))
	fmt.Println("    Saved at:", searchItem.SavedAt.Format("2006-01-02"))
}

func formatShort(searchItem *omnivore.SearchItem) {
	fmt.Printf("* [%s] %s %s (%s)\n", searchItem.SavedAt.Format("2006-01-02"), searchItem.Title, labelsToString(searchItem.Labels), searchItem.Url)
}

func labelsToString(labels []omnivore.Label) string {
	l := []string{}
	for _, label := range labels {
		l = append(l, label.Name)
	}
	if len(l) == 0 {
		return ""
	}
	return fmt.Sprintf("[%s]", strings.Join(l, ", "))
}
func getAPIToken() string {
	return os.Getenv("OMNIVORE_API_TOKEN")
}
