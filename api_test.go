package omnivore_test

import (
	"os"
	"testing"

	"github.com/rubiojr/omnivore-go"
	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	// https://docs.omnivore.app/using/search.html
	articles, err := client.Search(omnivore.SearchOpts{Query: "in:all -label:RSS"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 3)
	assert.Equal(t, articles[0].Title, "Organize your Omnivore library with labels")

	articles, err = client.Search(omnivore.SearchOpts{Query: "in:all label:RSS"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 1)
	assert.Equal(t, articles[0].Title, "Web changes for an improved experience")
}

func TestSubcriptions(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	subscriptions, err := client.Subscriptions()
	assert.NoError(t, err, "Failed to get subscriptions")
	assert.Equal(t, len(subscriptions), 1)
	assert.Equal(t, subscriptions[0].Name, "Omnivore Blog")
}
