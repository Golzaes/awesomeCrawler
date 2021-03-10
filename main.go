package main

import (
	"fmt"
)

func main() {
	// tagURL := `https://book.douban.com/tag/`
	// ListURL := `https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4`
	// fmt.Println(tagURL, ListURL)
	// Engine.Run(Engine.Request{
	// 	Method: "GET",
	// 	URL:    ListURL,
	// 	//URL:    `https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4`,
	// 	//ParseFunc: Parser.CssParseTag,
	// 	//ParseFunc: Parser.QueryParseList,
	// 	ParseFunc: Parser.XpathParseList,
	// })
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
