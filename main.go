package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-spider/conf"
	"go-spider/httputil"
)


func main() {
	//日志配置
	conf.LogSetting()

	doc, err := httputil.Get("https://www.piaohua.com/html/dongzuo/list_2.html")
	if err != nil {
		fmt.Println("请求出错:", err)
	}
	//ret, err := doc.Html()
	//fmt.Println("页面...:",ret)
	find := doc.Find("div.m-film .ul-imgtxt2 .col-md-6")
	//fmt.Println("------:",find)
	find.Each(func(i int, s *goquery.Selection) {
		//下载页面地址
		if downPageUrl, exists := s.Find("div.pic").Find("a").Attr("href"); exists {
			fmt.Println("下载类容:", downPageUrl)
		}
		//小图片地址
		if picUrl, exists := s.Find("div.pic").Find("a").Find("img").Attr("src"); exists {
			fmt.Println("电影小图地址:", picUrl)
		}
		//电影名
		name := s.Find("div.txt h3 a b font").Text()
		fmt.Println("电影名:", name)
		//清晰渡
		definition := s.Find("div.txt h3 a em").Text()
		fmt.Println("清晰度:", definition)
		//简介
		introduce := s.Find("div.txt p").Text()
		fmt.Println("简介:", introduce)
	})
}
