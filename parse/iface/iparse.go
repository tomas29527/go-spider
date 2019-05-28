package iface

import "github.com/PuerkitoBio/goquery"

type IParse interface {

	//传入待爬地址，返回文档对象
	GetDocument(url string) (*goquery.Document, error)

	//传入需要解析的文档对象
	ParseHtml(*goquery.Document) (interface{}, error)
}
