package internal

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML paragraph found
func processParagraph(index int, element *goquery.Selection) {
	words := strings.Split(element.Text(), " ")
	for _, w := range words {
		checkWord(w)
	}
}
