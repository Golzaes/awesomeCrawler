package main

import (
	"fmt"
	"github.com/payne/awesomeCrawler/Engine"
	"github.com/payne/awesomeCrawler/Parser"
)

func main() {
	const method = `GET`
	const tagURL = `https://book.douban.com/tag/`
	const ListURL = `https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4`
	fmt.Println(method, tagURL, ListURL)
	//weakRun(method, tagURL, Parser.RegexParseTag)
	slimRun(100, method, tagURL, Parser.RegexParseTag)

}
func weakRun(startMethod, startURL string, startParse func([]byte) Engine.ParseResult) {
	Engine.SimpleRun(Engine.Request{
		Method:    startMethod,
		URL:       startURL,
		ParseFunc: startParse,
	})
}

func slimRun(WorkCount int, method, URL string, Parse func([]byte) Engine.ParseResult) {
	e := Engine.ConcurrentEngine{
		Scheduler: &Engine.SimpleScheduler{},
		WorkCount: WorkCount,
	}
	e.ConcurrentRun(Engine.Request{
		Method:    method,
		URL:       URL,
		ParseFunc: Parse,
	})
}
