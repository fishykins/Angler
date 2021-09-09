package internal

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func processLink(index int, element *goquery.Selection) {
	link, found := element.Attr("href")
	if found {
		url, err := url.Parse(link)

		if err != nil {
			return
		}

		if url.Host != "" && strings.Contains(url.Host, "wiki") {
			AddUrl(link)
		}
	}
}
