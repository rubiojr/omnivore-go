package omnivore

import (
	"time"

	"github.com/shurcooL/graphql"
)

var searchQuery struct {
	Search struct {
		SearchSuccess struct {
			Edges []struct {
				Cursor string
				Node   struct {
					Author        string
					Title         string
					Content       string
					IsArchived    bool
					PublishedAt   time.Time
					Description   string
					SavedAt       time.Time
					ID            graphql.ID
					ReadAt        time.Time
					Url           string
					PageType      string
					CreatedAt     time.Time
					ContentReader string
					WordsCount    int
					FeedContent   string
					Folder        string
					Labels        []struct {
						Name        string
						Color       string
						CreatedAt   time.Time
						Description string
					}
				}
			}
			PageInfo struct {
				TotalCount      int
				EndCursor       string
				HasNextPage     bool
				HasPreviousPage bool
				StartCursor     string
			}
		} `graphql:"... on SearchSuccess"`
		SearchError struct {
			errorCodes string
		} `graphql:"... on SearchError"`
	} `graphql:"search(after: $after, format: $format, query: $query, includeContent: $includeContent)"`
}
