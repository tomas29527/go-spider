package models

import "time"

//t_moive 表
type Movie struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	Definition string `db:"definition"`
	Pic        string `db:"pic"`
	Introduce  string `db:"introduce"`
	Mold       int    `db:"mold"`
	DownPage   string `db:"down_page"`
	CreateTime string `db:"create_time"`
}

func (m *Movie) MovieInsert() error {
	sql := "insert into t_movie (name,definition,pic,introduce,mold,down_page,create_time) values (?,?,?,?,?,?,?)"
	_, e := mySqlDB.Exec(sql, m.Name, m.Definition, m.Pic, m.Introduce, m.Mold, m.DownPage, time.Now())
	return e
}

//创建Moive对象
func NewMovie(name string, definition string, picUrl string, introduce string, mold int, downPageUrl string) (m *Movie) {
	m = &Movie{
		Name:       name,
		Definition: definition,
		Pic:        picUrl,
		Introduce:  introduce,
		Mold:       mold,
		DownPage:   downPageUrl,
	}
	return
}
