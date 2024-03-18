package omnivore

import (
	"context"
	"net/http"
	"time"

	"github.com/rubiojr/omnivore-go/queries"
	"github.com/shurcooL/graphql"
)

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

type Omnivore struct {
	graphql *graphql.Client
}

type Opts struct {
	Token string
}

type Subscription struct {
	AutoAddToLibrary bool
	Count            int
	CreatedAt        time.Time
	Description      string
	FailedAt         time.Time
	FetchContent     bool
	Folder           string
	Icon             string
	IsPrivate        bool
	LastFetchedAt    time.Time
	Name             string
	NewsletterEmail  string
	RefreshedAt      time.Time
	Url              string
}

type Label struct {
	Name        string
	Color       string
	Description string
	CreatedAt   time.Time
}

type NewsletterEmail struct {
	Address           string
	CreatedAt         time.Time
	Name              string
	SubscriptionCount int
	Folder            string
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

		err := c.graphql.Query(context.Background(), &queries.Search, variables)
		if err != nil {
			return nil, err
		}

		results := queries.Search.Search.SearchSuccess

		for _, edge := range results.Edges {
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

		if !results.PageInfo.HasNextPage {
			break
		}

		afterCursor = results.PageInfo.EndCursor
	}

	return a, nil
}

func (c *Omnivore) NewsletterEmails() ([]NewsletterEmail, error) {
	a := []NewsletterEmail{}

	err := c.graphql.Query(context.Background(), &queries.NewsletterEmail, nil)
	if err != nil {
		return nil, err
	}

	result := queries.NewsletterEmail.NewsletterEmails.NewsletterEmailsSuccess
	for _, email := range result.NewsletterEmails {
		a = append(a, NewsletterEmail{
			Address:           email.Address,
			CreatedAt:         email.CreatedAt,
			Name:              email.Name,
			SubscriptionCount: email.SubscriptionCount,
			Folder:            email.Folder,
		})
	}

	return a, nil
}

func (c *Omnivore) Subscriptions() ([]Subscription, error) {
	a := []Subscription{}

	err := c.graphql.Query(context.Background(), &queries.Subscriptions, nil)
	if err != nil {
		return nil, err
	}

	result := queries.Subscriptions.Subscriptions.SubscriptionsSuccess
	for _, sub := range result.Subscriptions {
		a = append(a, Subscription{
			AutoAddToLibrary: sub.AutoAddToLibrary,
			Count:            sub.Count,
			CreatedAt:        sub.CreatedAt,
			Description:      sub.Description,
			FailedAt:         sub.FailedAt,
			FetchContent:     sub.FetchContent,
			Folder:           sub.Folder,
			Icon:             sub.Icon,
			IsPrivate:        sub.IsPrivate,
			LastFetchedAt:    sub.LastFetchedAt,
			Name:             sub.Name,
			NewsletterEmail:  sub.NewsletterEmail,
			RefreshedAt:      sub.RefreshedAt,
			Url:              sub.Url,
		})
	}

	return a, nil
}
