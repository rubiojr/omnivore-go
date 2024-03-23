package omnivore

import (
	"net/http"
	"strings"
)

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

type headerTransport struct {
	Token string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", t.Token)
	req.Header.Set("User-Agent", "omnivore-go/1.0")
	return http.DefaultTransport.RoundTrip(req)
}
