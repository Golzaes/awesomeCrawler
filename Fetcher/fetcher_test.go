package Fetcher

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFetch(t *testing.T) {
	html, _ := Fetch("GET", "https://ssr1.scrape.center/page/2")
	//html, _ := Fetch("GET", "https://book.douban.com/tag/?view=cloud")
	fmt.Print(string(html))
}

func BenchmarkFetch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		html, _ := Fetch("GET", "https://ssr1.scrape.center/page/2")
		fmt.Print(string(html))
	}
}

func benchmarkFetch(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		html, _ := Fetch("GET", "https://ssr1.scrape.center/page/"+strconv.Itoa(n))
		fmt.Print(string(html))
	}
}
func BenchmarkFetch1(b *testing.B)  { benchmarkFetch(b, 1) }
func BenchmarkFetch2(b *testing.B)  { benchmarkFetch(b, 2) }
func BenchmarkFetch5(b *testing.B)  { benchmarkFetch(b, 5) }
func BenchmarkFetch10(b *testing.B) { benchmarkFetch(b, 10) }
