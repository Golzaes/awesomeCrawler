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

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
