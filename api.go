package omnivore

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	// Omnivore API token.
	Token string
}

type Label struct {
	// Label name.
	Name string
	// Label color (hex).
	Color string
	// Label description.
	Description string
	// Label creation date.
	CreatedAt time.Time
	// Label ID.
	ID graphql.ID
	// FIXME: no clue
	Source string
	// FIXME: no clue
	Internal bool
	// Label position.
	Position int
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

type NewsletterEmail struct {
	Address           string
	CreatedAt         time.Time
	Name              string
	SubscriptionCount int
	Folder            string
}

type SearchItem struct {
	// Title of the item.
	Title string
	// Content of the item.
	Content string
	// Author of the item.
	Author string
	//Description of the item.
	Description string
	// IsArchived is true if the item is archived.
	IsArchived bool
	// PublishedAt is the date the item was published.
	PublishedAt time.Time
	// SavedAt is the date the item was saved.
	SavedAt time.Time
	// ID is the item ID.
	ID string
	// ReadAt is the date the item was read.
	ReadAt time.Time
	// Url is the URL of the item.
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
	// Query is the search query.
	// Documented here: https://docs.omnivore.app/using/search.html
	Query string
	// IncludeContent will include the content of the articles in the search results.
	IncludeContent bool
}

type ApiKey struct {
	ID graphql.ID
	// Key is the API key.
	Key string
	// CreatedAt is the date the API key was created.
	CreatedAt time.Time
	// ExpiresAt is the date the API key expires.
	ExpiresAt string
	// Name is the name of the API key.
	Name   string
	Scopes []string
	// UsedAt is the date the API key was last used.
	UsedAt time.Time
}

// NewClient creates a new Omnivore client.
//
// Opts.Token is required, or the function will panic.
func NewClient(opts Opts) *Omnivore {
	return NewClientFor("https://api-prod.omnivore.app/api/graphql", opts)
}

// NewClientFor creates a new Omnivore client for a specific URL.
//
// Opts.Token is required, or the function will panic.
func NewClientFor(url string, opts Opts) *Omnivore {
	if opts.Token == "" {
		panic("token is missing")
	}

	httpClient := &http.Client{
		Transport: &headerTransport{
			Token: opts.Token,
		},
	}

	client := graphql.NewClient(url, httpClient)
	return &Omnivore{graphql: client}
}

// Search searches for items in Omnivore.
//
// The search is paginated, and it will fetch all the results by default.
func (c *Omnivore) Search(ctx context.Context, opts SearchOpts) ([]SearchItem, error) {
	afterCursor := ""

	a := []SearchItem{}

	for {
		variables := map[string]any{
			"query":          graphql.String(opts.Query),
			"includeContent": graphql.Boolean(opts.IncludeContent),
			"after":          graphql.String(afterCursor),
		}

		err := c.graphql.Query(ctx, &queries.Search, variables)
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

// NewsletterEmails returns the newsletter emails of the user.
func (c *Omnivore) NewsletterEmails(ctx context.Context) ([]NewsletterEmail, error) {
	a := []NewsletterEmail{}

	err := c.graphql.Query(ctx, &queries.NewsletterEmail, nil)
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

// Subscriptions returns the subscriptions of the user.
func (c *Omnivore) Subscriptions(ctx context.Context) ([]Subscription, error) {
	a := []Subscription{}

	err := c.graphql.Query(ctx, &queries.Subscriptions, nil)
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

// ApiKeys returns all the API keys in Omnivore.
func (c *Omnivore) ApiKeys(ctx context.Context) ([]ApiKey, error) {
	a := []ApiKey{}

	err := c.graphql.Query(ctx, &queries.ApiKeys, nil)
	if err != nil {
		return nil, err
	}

	result := queries.ApiKeys.ApiKeys.ApiKeysSuccess
	for _, apiKey := range result.ApiKeys {
		a = append(a, ApiKey{
			ID:        apiKey.ID,
			Key:       apiKey.Key,
			CreatedAt: apiKey.CreatedAt,
			ExpiresAt: apiKey.ExpiresAt,
			Name:      apiKey.Name,
			Scopes:    apiKey.Scopes,
			UsedAt:    apiKey.UsedAt,
		})
	}

	return a, nil
}

// HasExpiry returns true if the API key has an expiry date.
func (a *ApiKey) HasExpiry() bool {
	return a.ExpiresAt != "+275760-09-13T00:00:00.000Z"
}

// Labels returns all the labels in Omnivore.
func (c *Omnivore) Labels(ctx context.Context) ([]*Label, error) {
	a := []*Label{}

	err := c.graphql.Query(ctx, &queries.Labels, nil)
	if err != nil {
		return nil, err
	}

	result := queries.Labels.Labels.LabelsSuccess
	for _, label := range result.Labels {
		a = append(a, &Label{
			ID:    label.ID,
			Name:  label.Name,
			Color: label.Color,
		})
	}

	return a, nil
}

// IsUnread returns true if the item is unread.
func (c *SearchItem) IsUnread() bool {
	return c.ReadAt.IsZero()
}

func (c *Omnivore) SaveUrl(ctx context.Context, url string) error {
	input := queries.SaveUrlInput{
		Url:             graphql.String(url),
		ClientRequestId: graphql.ID(uuid.New().String()),
		Source:          graphql.String("api"),
	}

	variables := map[string]any{
		"input": input,
	}

	return c.graphql.Mutate(ctx, &queries.SaveUrl, variables)
}

// DeleteArticle deletes an article from the library.
//
// The query name is weird, see https://github.com/omnivore-app/omnivore/issues/2380
func (c *Omnivore) DeleteArticle(ctx context.Context, id string) error {
	input := queries.SetBookmarkArticleInput{
		ArticleId: graphql.ID(id),
		Bookmark:  graphql.Boolean(false),
	}

	variables := map[string]any{
		"input": input,
	}

	err := c.graphql.Mutate(ctx, &queries.DeleteArticle, variables)
	return err
}
