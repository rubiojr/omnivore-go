# Omnivore Go

[![GoDoc](https://godoc.org/github.com/rubiojr/omnivore-go?status.svg)](https://godoc.org/github.com/rubiojr/omnivore-go)

This is a Go client library for the [Omnivore GraphQL API](https://github.com/omnivore-app/omnivore).

## Work in Progress

Supported read queries:

- [x] ApiKeys
- [ ] Feeds
- [ ] Filters
- [ ] Groups
- [ ] Integrations
- [x] Labels
- [x] NewsletterEmails
- [ ] Rules
- [x] Search
- [x] Subscriptions
- [ ] Users
- [ ] Webhooks

Supported mutation queries:

- [ ] AddPopularRead
- [ ] BulkAction
- [ ] CreateArticle
- [ ] CreateGroup
- [ ] DeleteAccount
- [x] DeleteArticle (SetBookmarkArticle)
- [ ] DeleteFilter
- [ ] DeleteIntegration
- [ ] DeleteLabel
- [ ] DeleteRule
- [ ] DeleteWebhook
- [x] SaveUrl
- [ ] EmptyTrash

## Usage

Get an API key from https://docs.omnivore.app/integrations/api.html.

See the [examples](examples) directory.
