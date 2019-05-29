package parse

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/logs"
	"go-spider/conf"
	"go-spider/httputil"
	"go-spider/models"
	"time"
)

//电影类型
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

var (
	mold_dongzuo   = "动作片"
	mold_xiju      = "喜剧片"
	mold_aiqing    = "爱情片"
	mold_kehuan    = "科幻片"
	mold_juqing    = "剧情片"
	mold_xuanyi    = "悬疑片"
	mold_zhanzheng = "战争片"
	mold_kongbu    = "恐怖片"
	mold_zhainan   = "灾难片"
)

type PaohuaParse struct {
	//动作类型页管道
	Dongzuo chan bool
}

func NewPaohuaParse() *PaohuaParse {
	piaohua := &PaohuaParse{
		Dongzuo: make(chan bool, 1),
	}
	return piaohua
}

//传入待爬地址，返回文档对象
func (p *PaohuaParse) GetDocument(url string) (*goquery.Document, error) {
	doc, err := httputil.Get(url)
	if err != nil {
		logs.Error("请求地址:%s,获取document出错:%v", url, err)
		return nil, err
	}
	return doc, nil
}

//传入需要解析的文档对象
func (p *PaohuaParse) ParseHtml(doc *goquery.Document) {
	//首页类容,找到导航栏
	nav := doc.Find("div.nav .wp ul li")
	nav.Each(func(i int, selection *goquery.Selection) {
		var moldDownPage string
		var exists bool
		if moldDownPage, exists = selection.Find("a").Attr("href"); exists {
			logs.Debug("类型下载地址:%s", moldDownPage)
		}
		logs.Debug("类型:%s", selection.Find("a").Text())
		if i > 0 && i <= 9 {
			mold := selection.Find("a").Text()
			switch mold {
			case mold_dongzuo:
				dongzuoUrl := conf.Global.PiaohuaIndexUrl + moldDownPage
				logs.Info("动作类型地址dongzuoUrl:%s", dongzuoUrl)
				go p.parseDongzuo(dongzuoUrl)
			}
		}
	})
}

//爬取动作类型电影
func (p *PaohuaParse) parseDongzuo(url string) {
	//解析动作页面
	document, e := p.GetDocument(url)
	if e != nil {
		logs.Error("请求地址:%s,获取document出错:%v", url, e)
		return
	}

	//解析当前页数据,存入数据库
	p.parseCurentPage(document)

	//判断是否还有下一页
	pageNext := document.Find("div.pages ul li.pages-next")
	if pageNext != nil {
		//下一页地址
		nextPageUrl, exists := pageNext.Find("a").Attr("href")
		if exists {

			nextPage := conf.Global.PiaohuaIndexUrl + "/html/dongzuo/" + nextPageUrl
			logs.Debug("nextPage:%s", nextPageUrl)
			time.Sleep(500 * time.Millisecond)
			p.parseDongzuo(nextPage)
		}
	}

}

func (p *PaohuaParse) parseCurentPage(doc *goquery.Document) {

	find := doc.Find("div.m-film .ul-imgtxt2 .col-md-6")

	find.Each(func(i int, s *goquery.Selection) {
		//下载页面地址
		downPageUrl, _ := s.Find("div.pic a").Attr("href")
		//小图片地址
		picUrl, _ := s.Find("div.pic").Find("a").Find("img").Attr("src")
		//电影名
		name := s.Find("div.txt h3 a b font").Text()
		//清晰渡
		definition := s.Find("div.txt h3 a em").Text()
		//简介
		introduce := s.Find("div.txt p").Text()
		logs.Debug("下载类容:%s", downPageUrl)
		logs.Debug("电影小图地址:%s", picUrl)
		logs.Debug("电影名:%s", name)
		logs.Debug("清晰度:%s", definition)
		logs.Debug("简介:%s", introduce)

		m := models.NewMovie(name, definition, picUrl, introduce, MOLD0, downPageUrl)
		err := m.MovieInsert()
		if err != nil {
			logs.Error("保存数据库失败:%v", err)
			logs.Error("保存失败对象Movie=:%v", m)
		}
	})
}
