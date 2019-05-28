package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"go-spider/httputil"
)

const (
	MOLD0 = iota //0:动作片
	MOLD1        //1:喜剧片
	MOLD2        //2爱情片
	MOLD3        //3科幻片
	MOLD4        //4剧情片
	MOLD5        //5悬疑片
	MOLD6        //6战争片
	MOLD7        //7恐怖片
	MOLD8        //8灾难片
)

type PaohuaParse struct {
}

//传入待爬地址，返回文档对象
func (p *PaohuaParse) GetDocument(url string) (*goquery.Document, error) {
	doc, err := httputil.Get("https://www.piaohua.com/html/dongzuo/list_2.html")
	if err != nil {
		logs.Error("获取document出错:", err)
		return nil, err
	}
	return doc, nil
}

//传入需要解析的文档对象
func (p *PaohuaParse) ParseHtml(*goquery.Document) (interface{}, error) {

	return nil, nil
}
