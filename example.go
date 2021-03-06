package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"log"
	"net/http"
)

// ScrapeExample
func ScrapeExample() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://book.douban.com/tag/?view=cloud", nil)
	if err != nil {
		log.Printf("NewRequest Error:%#v", err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Client Error:%#v", err)
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}

// CssParseExample
func CssParseExample(body []byte) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}
	// Find the review items
	doc.Find("table tr td").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		tagText := s.Find("a").Text()
		tagHref, _ := s.Find("a").Attr("href")
		fmt.Printf("[tagText]:%s, [tagHref]:%s \n", tagText, tagHref)
	})
}

// XpathParseExample
func XpathParseExample(body []byte) {
	selector, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		log.Printf("htmlquery Parse Error: %s", err)
	}
	for _, node := range htmlquery.Find(selector, `//table//tr/td`) {
		tagHref := htmlquery.Find(node, `./a/@href`)
		tagTxt := htmlquery.Find(node, `./a/text()`)
		href := htmlquery.SelectAttr(tagHref[0], `href`)
		txt := htmlquery.InnerText(tagTxt[0])
		//fmt.Printf("[tagText]:%#v, [tagHref]:%#v \n", tagHref, tagTxt)
		//exa := htmlquery.InnerText(item[2])
		log.Printf("Href:%s, Txt:%s \n", href, txt)
	}
}

// func main() {
// 	body, _ := ScrapeExample()
// 	//CssParseExample(body)
// 	XpathParseExample(body)
// }
