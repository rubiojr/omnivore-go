package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"context"

	"github.com/rubiojr/omnivore-go"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "omnivore",
		Usage: "Omnivore API client",
		Flags: []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List available articles",
				Action: func(cCtx *cli.Context) error {
					return listSaved(cCtx)
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name: "long",
					},
				},
			},
			{
				Name:  "add",
				Usage: "Save a URL",
				Action: func(cCtx *cli.Context) error {
					return saveUrl(cCtx)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func client() (*omnivore.Omnivore, error) {
	token := getAPIToken()
	if token == "" {
		return nil, fmt.Errorf("OMNIVORE_API_TOKEN is required")
	}

	return omnivore.NewClient(omnivore.Opts{Token: token}), nil
}

func saveUrl(ctx *cli.Context) error {
	client, err := client()
	if err != nil {
		return err
	}
	url := ctx.Args().First()
	if url == "" {
		return fmt.Errorf("URL is required")
	}
	return client.SaveUrl(context.Background(), url)
}

func listSaved(ctx *cli.Context) error {
	client, err := client()
	if err != nil {
		return err
	}
	longFormat := ctx.Bool("long")
	// https://docs.omnivore.app/using/search.html
	a, err := client.Search(context.Background(), omnivore.SearchOpts{Query: "in:all sort:saved"})
	if err != nil {
		return err
	}
	for _, searchItem := range a {
		if longFormat {
			formatLong(&searchItem)
		} else {
			formatShort(&searchItem)
		}
	}
	fmt.Println("Total items:", len(a))

	return nil
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
