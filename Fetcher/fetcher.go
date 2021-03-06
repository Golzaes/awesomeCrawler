package Fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(method, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("NewRequest Error:%#v", err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Client Error:%#v", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}
	// Detection encode，eg: UTF8 GBK...
	bodyReader := bufio.NewReader(res.Body)
	PageEncode := DetectionEncode(bodyReader)
	// recoding
	encodedReader := transform.NewReader(bodyReader, PageEncode.NewDecoder())
	return ioutil.ReadAll(encodedReader)
}

// DetectionEncode 检测网页编码，实现自动解码
func DetectionEncode(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("DetectionEncode Error:%#v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
