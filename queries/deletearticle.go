package queries

import "github.com/shurcooL/graphql"

var DeleteArticle struct {
	SetBookmarkArticle struct {
		SetBookmarkArticleSuccess struct {
			BookmarkedArticle struct {
				ID     graphql.ID
				LinkId graphql.ID
			}
		} `graphql:"... on SetBookmarkArticleSuccess"`
		SetBookmarkArticleError struct {
			ErrorCodes string
		} `graphql:"... on SetBookmarkArticleError"`
	} `graphql:"setBookmarkArticle(input: $input)"`
}

type SetBookmarkArticleInput struct {
	ArticleId graphql.ID      `json:"articleID"`
	Bookmark  graphql.Boolean `json:"bookmark"`
}
