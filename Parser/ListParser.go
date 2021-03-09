package Parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
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
func QueryParseList(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("NewDocumentFromReader err: %s", err)
	}
	// Find the review items
	doc.Find("#content .subject-list li.subject-item").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Find(`.info h2 a`).Attr(`title`)
		href, _ := s.Find(`.info h2 a`).Attr(`href`)
		log.Println(title, href)

		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       href,
			ParseFunc: Engine.NilParse,
		})
	})
	return result
}

// XpathParseList
func XpathParseList(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		fmt.Printf("htmlquery Parse Error: %s", err)
	}
	nodes := htmlquery.Find(doc, `//div[@id="subject_list"]/ul/li[@class="subject-item"]`)
	for _, node := range nodes {
		li := htmlquery.Find(node, `./*[@class="info"]/h2/a`)
		title := htmlquery.SelectAttr(li[0], `title`)
		href := htmlquery.SelectAttr(li[0], `href`)
		log.Printf(`Title: %s Href: %s`, title, href)
		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       href,
			ParseFunc: Engine.NilParse,
		})
	}
	return result
}
