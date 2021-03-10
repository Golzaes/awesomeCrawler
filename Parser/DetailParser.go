package Parser

type BookDetail struct {
	Title            string // 标题
	Author           string // 作者
	PublicationHouse string // 出版社
	Pages            string // 页数
	Price            string // 价格
	Binding          string // 装帧
	Series           string // 系列
	ISBN             string // 书号
	Score            string // 评分
	Translator       string // 译者
	Introduction     string // 简介
	Catalog          string // 目录
}

// QueryParseDetail
//func QueryParseDetail(content []byte) BookDetail {
//	bookdetail := BookDetail{}
//	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(content))
//	if err != nil {
//		log.Printf("NewDocumentFromReader err: %s", err)
//	}
//	//doc.Find()
//	return bookdetail
//}

// XpathParseDetail
//func XpathParseDetail(content []byte) BookDetail {
//	bookdetail := BookDetail{}
//	doc, err := htmlquery.Parse(bytes.NewBuffer(content))
//	if err != nil {
//		fmt.Printf("htmlquery Parse Error: %s", err)
//	}
//	nodes := htmlquery.Find(doc, ``)
//
//}
