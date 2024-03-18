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

type Label struct {
	Name        string
	Color       string
	Description string
	CreatedAt   time.Time
}

type SearchItem struct {
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
	FeedContent   string
	Folder        string
	Labels        []Label
}

type SearchOpts struct {
	Query          string
	IncludeContent bool
	Format         string
}

func (c *Omnivore) Search(opts SearchOpts) ([]SearchItem, error) {
	afterCursor := ""

	a := []SearchItem{}

	for {
		variables := map[string]any{
			"query":          graphql.String(opts.Query),
			"includeContent": graphql.Boolean(opts.IncludeContent),
			"after":          graphql.String(afterCursor),
			"format":         graphql.String(opts.Format),
		}

		err := c.graphql.Query(context.Background(), &searchQuery, variables)
		if err != nil {
			return nil, err
		}

		for _, edge := range searchQuery.Search.SearchSuccess.Edges {
			si := SearchItem{
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
				FeedContent:   edge.Node.FeedContent,
				Folder:        edge.Node.Folder,
			}
			labels := []Label{}
			for _, label := range edge.Node.Labels {
				labels = append(labels, Label{
					Name:        label.Name,
					Color:       label.Color,
					Description: label.Description,
					CreatedAt:   label.CreatedAt,
				})
			}
			si.Labels = labels
			a = append(a, si)

		}

		if !searchQuery.Search.SearchSuccess.PageInfo.HasNextPage {
			break
		}

		afterCursor = searchQuery.Search.SearchSuccess.PageInfo.EndCursor
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
