package queries

import "time"

var NewsletterEmail struct {
	NewsletterEmails struct {
		NewsletterEmailsSuccess struct {
			NewsletterEmails []struct {
				Address           string
				CreatedAt         time.Time
				Name              string
				SubscriptionCount int
				Folder            string
			}
		} `graphql:"... on NewsletterEmailsSuccess"`
	}
}
