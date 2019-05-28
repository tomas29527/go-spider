package httputil

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

var httpClient = &http.Client{}

//get请求地址,返回document对象
func Get(url string) (doc *goquery.Document, err error) {
	reqest, e := http.NewRequest("GET", url, nil)
	if e != nil {
		err = e
		return
	}
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36")
	resp, e := httpClient.Do(reqest)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()

	document, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		err = e
		return
	}
	doc = document
	return
}
