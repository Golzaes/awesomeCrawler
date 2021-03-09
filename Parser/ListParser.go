package Parser

import (
	"github.com/payne/awesomeCrawler/Engine"
	"log"
	"regexp"
)

// RegexParseList
func RegexParseList(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	re := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)
	matches := re.FindAllSubmatch(content, -1)
	for _, m := range matches {
		href := string(m[1])
		title := string(m[2])
		log.Printf("Title: %s href: %s", title, href)
		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       string(m[1]),
			ParseFunc: Engine.NilParse,
		})
	}
	return result
}

// QueryParseList
//func QueryParseList(content []byte) {
//	fmt.Printf("QueryParseList:%s", content)
//}

// XpathParseList
//func XpathParseList(body []byte) Engine.ParseResult {
//	fmt.Printf("XpathParseList:%s", body)
//}
