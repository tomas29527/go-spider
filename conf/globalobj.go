package conf

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type GlobalObj struct {
	PiaohuaIndexUrl string
}

var Global *GlobalObj

func init() {
	indexPageurl := beego.AppConfig.String("prod::spider.url")
	if indexPageurl == "" {
		logs.Error("==================爬取地址不能为空!")
		return
	}
	Global = &GlobalObj{
		indexPageurl,
	}
}
