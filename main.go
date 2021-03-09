package main

import (
	"github.com/payne/awesomeCrawler/Engine"
	"github.com/payne/awesomeCrawler/Parser"
)

func main() {
	Engine.Run(Engine.Request{
		Method: "GET",
		URL:    `https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4`,
		//ParseFunc: Parser.CssParseTag,
		ParseFunc: Parser.RegexParseList,
	})
	//a := `123`
	//b := "123"
	//c := "321"
	////fmt.Printf("Type:%T", a, reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(b))
	//fmt.Println(reflect.TypeOf(c))
	////fmt.Print(strings.Compare(a, c))
	//fmt.Println(a + b)

}
