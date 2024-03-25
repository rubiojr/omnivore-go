package omnivore_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rubiojr/omnivore-go"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	articles, err := client.Search(context.Background(), omnivore.SearchOpts{Query: "in:all"})
	assert.NoError(t, err, "Failed to search")
	for _, a := range articles {
		err = client.DeleteArticle(context.Background(), a.ID)
		assert.NoError(t, err, "Failed to delete article")
	}
}

func TestSaveUrl(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	wp := "https://en.wikipedia.org/wiki/Leet"
	err := client.SaveUrl(context.Background(), wp)
	assert.NoError(t, err, "Failed to save article")
}

func TestSearch(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	// https://docs.omnivore.app/using/search.html
	articles, err := client.Search(context.Background(), omnivore.SearchOpts{Query: "in:all -label:RSS"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 1)
	assert.Equal(t, articles[0].Title, "https://en.wikipedia.org/wiki/Leet")

	articles, err = client.Search(context.Background(), omnivore.SearchOpts{Query: "in:all label:RSS"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 0)

	articles, err = client.Search(context.Background(), omnivore.SearchOpts{Query: "title:Leet"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 1)
}

func TestSubcriptions(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	subscriptions, err := client.Subscriptions(context.Background())
	assert.NoError(t, err, "Failed to get subscriptions")
	assert.Equal(t, len(subscriptions), 1)
	assert.Equal(t, subscriptions[0].Name, "Omnivore Blog")
}

func TestApiKeys(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	keys, err := client.ApiKeys(context.Background())
	assert.NoError(t, err, "Failed to get api keys")
	assert.Equal(t, len(keys), 1)
	assert.Equal(t, keys[0].Name, "omnivore-go-github")
	assert.Equal(t, keys[0].ExpiresAt, "+275760-09-13T00:00:00.000Z")
	assert.False(t, keys[0].HasExpiry())
}

func TestLabels(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	labels, err := client.Labels(context.Background())
	assert.NoError(t, err, "Failed to get labels")
	assert.Equal(t, len(labels), 1)
	assert.Equal(t, labels[0].Name, "RSS")
	assert.Equal(t, labels[0].Color, "#F26522")
	assert.Equal(t, labels[0].Description, "")
}

func TestDeleteArticle(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	articles, err := client.Search(context.Background(), omnivore.SearchOpts{Query: fmt.Sprintf(`title:"%s"`, "https://en.wikipedia.org/wiki/Leet")})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 1)
	// Wait a bit, deletes right after saving are ignored otherwise
	err = client.DeleteArticle(context.Background(), articles[0].ID)
	assert.NoError(t, err, "Failed to delete article")
	articles, err = client.Search(context.Background(), omnivore.SearchOpts{Query: fmt.Sprintf(`title:"%s"`, "https://en.wikipedia.org/wiki/Leet")})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 0)
}
