package Engine

import (
	"fmt"
	"github.com/payne/awesomeCrawler/Fetcher"
	"log"
)

func Run(seed ...Request) {
	var requests []Request
	for _, e := range seed {
		requests = append(requests, e)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		log.Printf("Fetch URL:  %s\n", request.URL)
		body, err := Fetcher.Fetch("GET", request.URL)
		if err != nil {
			fmt.Println("Fetch Error", request.URL)
		}
		ParseResult := request.ParseFunc(body)
		requests = append(requests, ParseResult.Request...)
		for _, item := range ParseResult.Item {
			log.Printf("Getting Item %s\n", item)
		}
	}
}
