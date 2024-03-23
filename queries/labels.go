package queries

import (
	"time"

	"github.com/shurcooL/graphql"
)

var Labels struct {
	Labels struct {
		LabelsSuccess struct {
			Labels []struct {
				ID          graphql.ID
				Name        string
				Color       string
				CreatedAt   time.Time
				Source      string
				Internal    bool
				Position    int
				Description string
			}
		} `graphql:"... on LabelsSuccess"`
	}
}
