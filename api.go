package omnivore

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/shurcooL/graphql"
)

type Omnivore struct {
	graphql *graphql.Client
}

type Opts struct {
	Token string
}

func NewClient(opts Opts) *Omnivore {
	return NewClientFor("https://api-prod.omnivore.app/api/graphql", opts)
}

func NewClientFor(url string, opts Opts) *Omnivore {
	httpClient := &http.Client{
		Transport: &headerTransport{
			Token: opts.Token,
		},
	}

	client := graphql.NewClient(url, httpClient)
	return &Omnivore{graphql: client}
}

type Article struct {
	Title         string
	Content       string
	Author        string
	Description   string
	IsArchived    bool
	PublishedAt   time.Time
	SavedAt       time.Time
	ID            string
	ReadAt        time.Time
	Url           string
	PageType      PageType
	CreatedAt     time.Time
	ContentReader string
	Words         int
}

type SearchOpts struct {
	Query          string
	IncludeContent bool
}

func (c *Omnivore) GetArticles(opts SearchOpts) ([]Article, error) {
	var query struct {
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
					}
				}
				PageInfo struct {
					TotalCount int
				}
			} `graphql:"... on SearchSuccess"`
			SearchError struct {
				errorCodes string
			} `graphql:"... on SearchError"`
		} `graphql:"search(query: $query, includeContent: $includeContent)"`
	}

	variables := map[string]any{
		"query":          graphql.String(opts.Query),
		"includeContent": graphql.Boolean(opts.IncludeContent),
	}

	err := c.graphql.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}

	a := []Article{}
	for _, edge := range query.Search.SearchSuccess.Edges {
		a = append(a, Article{
			Title:         edge.Node.Title,
			PublishedAt:   edge.Node.PublishedAt,
			Content:       edge.Node.Content,
			Description:   edge.Node.Description,
			IsArchived:    edge.Node.IsArchived,
			SavedAt:       edge.Node.SavedAt,
			ID:            edge.Node.ID.(string),
			ReadAt:        edge.Node.ReadAt,
			Url:           edge.Node.Url,
			PageType:      pageTypeToName(edge.Node.PageType),
			CreatedAt:     edge.Node.CreatedAt,
			ContentReader: edge.Node.ContentReader,
			Author:        edge.Node.Author,
			Words:         edge.Node.WordsCount,
		})
	}

	return a, nil
}

type headerTransport struct {
	Token string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.Token)
	return http.DefaultTransport.RoundTrip(req)
}

const (
	PageTypeArticle = "article"
	PageTypeVideo   = "video"
	PageTypeTweet   = "tweet"
	PageTypeFile    = "file"
	PageTypeBook    = "book"
	PageTypeImage   = "image"
	PageTypeUnknown = "unknown"
	PageTypeWebsite = "website"
)

type PageType string

func pageTypeToName(pageType string) PageType {
	switch strings.ToLower(pageType) {
	case "article":
		return PageTypeArticle
	case "video":
		return PageTypeVideo
	case "website":
		return PageTypeWebsite
	case "tweet":
		return PageTypeTweet
	case "file":
		return PageTypeFile
	case "book":
		return PageTypeBook
	case "image":
		return PageTypeImage
	default:
		return PageTypeUnknown
	}
}
