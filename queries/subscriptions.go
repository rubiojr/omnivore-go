package queries

import "time"

var Subscriptions struct {
	Subscriptions struct {
		SubscriptionsSuccess struct {
			Subscriptions []struct {
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
		} `graphql:"... on SubscriptionsSuccess"`
	}
}
