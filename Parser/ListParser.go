package Parser

import "fmt"

// RegexParseList
func RegexParseList(body []byte) {
	fmt.Printf("RegexParseList:%s", body)
}

// QueryParseList
func QueryParseList(body []byte) {
	fmt.Printf("QueryParseList:%s", body)
}

// XpathParseList
func XpathParseList(body []byte) {
	fmt.Printf("XpathParseList:%s", body)
}
