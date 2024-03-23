package queries

import (
	"time"

	"github.com/shurcooL/graphql"
)

var ApiKeys struct {
	ApiKeys struct {
		ApiKeysSuccess struct {
			ApiKeys []struct {
				ID        graphql.ID
				Key       string
				CreatedAt time.Time
				UsedAt    time.Time
				ExpiresAt string
				Scopes    []string
				Name      string
			}
		} `graphql:"... on ApiKeysSuccess"`
	}
}
