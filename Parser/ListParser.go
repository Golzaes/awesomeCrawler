package Parser

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/payne/awesomeCrawler/Engine"
	"log"
	"net/url"
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
func QueryParseList(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("NewDocumentFromReader err: %s", err)
	}
	// Find the review items
	doc.Find("#content .subject-list li.subject-item").Each(func(i int, s *goquery.Selection) {
		title := s.Find(`.info h2 a`).Text()
		href := s.Find(`.info h2 a`).Attr(`href`)
		log.Println(tag)
		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       "https://book.douban.com/tag/" + url.QueryEscape(tag),
			ParseFunc: Engine.NilParse,
		})
	})
	return result
}

// XpathParseList
//func XpathParseList(body []byte) Engine.ParseResult {
//	fmt.Printf("XpathParseList:%s", body)
//}
