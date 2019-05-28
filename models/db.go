package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

type Db struct {
}

var mySqlDB *sqlx.DB

func (d *Db) InitDb() error {
	username := beego.AppConfig.String("prod::mysql.username")
	password := beego.AppConfig.String("prod::mysql.password")
	logs.Debug("===mysql username=:%s,password=:%s", username, password)
	if username == "" || password == "" {
		logs.Error("===================")
		logs.Error("数据库初始化失败")
		return errors.New("数据库初始化失败")
	}
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(106.13.60.183:3306)/star_child?charset=utf8", username, password)
	logs.Debug("===mysqlUrl is ==:%s", mysqlUrl)
	db, err := sqlx.Open("mysql", mysqlUrl)
	if err != nil {
		logs.Error("数据库初始化失败:%v", err)
		return errors.New("数据库初始化失败")
	}
	mySqlDB = db
	return nil
}

func (d *Db) Close() {
	logs.Info("数据库连接关闭============")
	mySqlDB.Close()
}
