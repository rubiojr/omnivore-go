package omnivore_test

import (
	"os"
	"testing"

	"github.com/rubiojr/omnivore-go"
	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	client := omnivore.NewClient(omnivore.Opts{Token: os.Getenv("OMNIVORE_API_TOKEN")})
	// https://docs.omnivore.app/using/search.html
	articles, err := client.Search(omnivore.SearchOpts{Query: "in:all"})
	assert.NoError(t, err, "Failed to search")
	assert.Equal(t, len(articles), 3)
	assert.Equal(t, articles[0].Title, "Organize your Omnivore library with labels")
}
