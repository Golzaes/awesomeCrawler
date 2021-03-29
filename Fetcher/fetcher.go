package Fetcher

import (
	"bufio"
	"fmt"
	"github.com/payne/awesomeCrawler/Tools"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func Fetch(method, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf(`NewRequest Error:%#v`, err)
	}
	req.Header.Add(`User-Agent`, Tools.RandomUa())
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf(` Client Error:%#v`, err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(` Get invalid status code %s while scraping %s`, res.Status, url)
	}
	// Detection encode，eg: UTF8 GBK...
	bodyReader := bufio.NewReader(res.Body)
	PageEncode := Tools.DetectionEncode(bodyReader)
	// recoding
	encodedReader := transform.NewReader(bodyReader, PageEncode.NewDecoder())
	return ioutil.ReadAll(encodedReader)
}
