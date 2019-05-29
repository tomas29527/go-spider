package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Db struct {
}

var mySqlDB *sqlx.DB

func (d *Db) InitDb() error {
	username := beego.AppConfig.String("prod::mysql.username")
	password := beego.AppConfig.String("prod::mysql.password")
	ip := beego.AppConfig.String("prod::mysql.ip")
	port := beego.AppConfig.String("prod::mysql.port")
	logs.Debug("===mysql ip =: %s port =: %s username=:%s,password=:%s", ip, port, username, password)
	if username == "" || password == "" {
		logs.Error("===================")
		logs.Error("数据库初始化失败")
		return errors.New("数据库初始化失败")
	}
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/go_spider?charset=utf8", username, password, ip, port)
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
