package types

import (
	"time"

	"github.com/shurcooL/graphql"
)

type Label struct {
	Name        string
	Color       string
	Description string
	CreatedAt   time.Time
	ID          graphql.ID
	Source      string
	Internal    bool
	Position    int
}
