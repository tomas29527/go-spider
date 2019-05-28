package models

import "time"

//t_moive 表
type Moive struct {
	id         int    `db:"id"`
	Name       string `db:"name"`
	Definition string `db:"definition"`
	Pic        string `db:"pic"`
	Introduce  string `db:"introduce"`
	Mold       int    `db:"mold"`
	DownPage   string `db:"down_page"`
	CreateTime string `db:"create_time"`
}

func (m *Moive) insert() error {
	sql := "insert into t_moive (name,definition,pic,introduce,mold,down_page,create_time) values (?,?,?,?,?,?,?)"
	_, e := mySqlDB.Exec(sql, m.Name, m.Definition, m.Pic, m.Introduce, m.Mold, m.DownPage, time.Now())
	return e
}
