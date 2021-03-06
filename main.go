package main

import (
	"github.com/payne/awesomeCrawler/Engine"
	"github.com/payne/awesomeCrawler/Parser"
)

func main() {
	Engine.Run(Engine.Request{
		Method: "GET",
		URL:    "https://book.douban.com/tag/?view=cloud",
		//ParseFunc: Parser.CssParseTag,
		ParseFunc: Parser.XpathParseTag,
	})
}
