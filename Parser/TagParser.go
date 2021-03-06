package Parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/payne/awesomeCrawler/Engine"
	"log"
	"net/url"
	"regexp"
)

func RegexParseTag(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	re := regexp.MustCompile(`<td><a href="(.*?)">`)
	matches := re.FindAllSubmatch(content, -1)
	for _, m := range matches {
		result.Item = append(result.Item, m[1])
		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       "https://book.douban.com/" + url.QueryEscape(string(m[1])),
			ParseFunc: Engine.NilParse,
		})
	}
	return result
}

//CssParseTag use go query to parse txt
func CssParseTag(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
	if err != nil {
		log.Printf("NewDocumentFromReader err: %s", err)
	}
	// Find the review items
	doc.Find("table tr td").Each(func(i int, s *goquery.Selection) {
		tag := s.Find("a").Text()
		log.Println(tag)
		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       "https://book.douban.com/tag/" + url.QueryEscape(tag),
			ParseFunc: Engine.NilParse,
		})
	})
	return result
}

// XpathParseTag use html query to parse txt
func XpathParseTag(content []byte) Engine.ParseResult {
	result := Engine.ParseResult{}
	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		fmt.Printf("htmlquery Parse Error: %s", err)
	}
	nodes := htmlquery.Find(doc, `//table//tr/td`)
	for _, node := range nodes {
		tagHref := htmlquery.Find(node, `./a/@href`)
		tagTxt := htmlquery.Find(node, `./a/text()`)
		href := htmlquery.SelectAttr(tagHref[0], `href`)
		txt := htmlquery.InnerText(tagTxt[0])

		log.Printf("Href:%s, Txt:%s \n", href, txt)
		result.Item = append(result.Item, txt)

		result.Request = append(result.Request, Engine.Request{
			Method:    "GET",
			URL:       "https://book.douban.com/" + url.QueryEscape(href),
			ParseFunc: Engine.NilParse,
		})
	}
	return result
}
