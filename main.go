package main

import (
	"github.com/astaxie/beego/logs"
	"go-spider/conf"
	"go-spider/models"
	"go-spider/parse"
)

func main() {
	//日志配置
	conf.LogSetting()
	//初始化数据库
	db := &models.Db{}
	initDbErr := db.InitDb()
	if initDbErr != nil {
		logs.Error("数据库初始化失败..")
		return
	}

	piaohua := &parse.PaohuaParse{}
	if conf.Global == nil {
		logs.Error("全局配置初始化失败")
		return
	}
	document, e := piaohua.GetDocument(conf.Global.PiaohuaIndexUrl)
	if e != nil {
		logs.Error("获取文档对象失败:%v", e)
		return
	}
	piaohua.ParseHtml(document)

	select {}
}
