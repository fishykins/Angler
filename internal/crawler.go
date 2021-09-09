package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var visited []string
var visiting []string
var potentialNouns []string

func GetPotentialNouns() []string {
	ret := make([]string, len(potentialNouns))
	copy(ret, potentialNouns)
	return ret
}

func AddUrl(url string) bool {
	canAppend := true
	for _, w := range visited {
		if w == url {
			canAppend = false
			break
		}
	}
	for _, w := range visiting {
		if w == url {
			canAppend = false
			break
		}
	}
	if canAppend {
		visiting = append(visiting, url)
	}
	return canAppend
}

func Crawl(recursive bool) bool {
	if len(visiting) == 0 {
		return false
	}

	url := visiting[len(visiting)-1]
	visiting = append(make([]string, 0), visiting[:len(visiting)-1]...)
	visited = append(visited, url)
	totalFound := len(potentialNouns)

	// Make HTTP request
	response, err := http.Get(url)
	if err != nil {
		return true
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("p").Each(processParagraph)

	if recursive {
		document.Find("a").Each(processLink)
	}

	found := len(potentialNouns) - totalFound
	println("found " + fmt.Sprint(found) + " potential nouns at '" + url + "' (" + fmt.Sprint(len(visiting)) + " urls left)")
	return len(visiting) > 0
}
