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

func TestApiKeys(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	keys, err := client.ApiKeys()
	assert.NoError(t, err, "Failed to get api keys")
	assert.Equal(t, len(keys), 1)
	assert.Equal(t, keys[0].Name, "omnivore-go-github")
	assert.Equal(t, keys[0].ExpiresAt, "+275760-09-13T00:00:00.000Z")
	assert.False(t, keys[0].HasExpiry())
}

func TestLabels(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	labels, err := client.Labels()
	assert.NoError(t, err, "Failed to get labels")
	assert.Equal(t, len(labels), 1)
	assert.Equal(t, labels[0].Name, "RSS")
	assert.Equal(t, labels[0].Color, "#F26522")
	assert.Equal(t, labels[0].Description, "")
}
