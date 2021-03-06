package Engine

type Request struct {
	Method    string
	URL       string
	ParseFunc func([]byte) ParseResult
	//Header map[string]string
}

type ParseResult struct {
	Request []Request
	Item    []interface{}
}

//func NewParseResult(request []Request, item []interface{}) *ParseResult {
//	return &ParseResult{Request: request, Item: item}
//}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
