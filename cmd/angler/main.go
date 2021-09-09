package main

import (
	"fmt"

	"example.com/fishy/angler/internal"
)

func main() {
	internal.AddUrl("https://en.wikipedia.org/wiki/Proper_and_common_nouns")

	crawling := true

	for i := 0; i < 100 && crawling; i++ {
		crawling = internal.Crawl(true)
	}

	for i, pn := range internal.GetPotentialNouns() {
		println(fmt.Sprint(i) + ": " + pn)
	}
}
