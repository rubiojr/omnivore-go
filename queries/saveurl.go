package queries

import "github.com/shurcooL/graphql"

var SaveUrl struct {
	SaveUrl struct {
		SaveSuccess struct {
			ClientRequestID graphql.ID
			Url             string
		} `graphql:"... on SaveSuccess"`
		SaveError struct {
			ErrorCodes string
		} `graphql:"... on SaveError"`
	} `graphql:"saveUrl(input: $input)"`
}

type SaveUrlInput struct {
	Url             graphql.String `json:"url"`
	ClientRequestId graphql.ID     `json:"clientRequestId"`
	Source          graphql.String `json:"source"`
}
