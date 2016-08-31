package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	initSqlite()
	orm.RegisterModel(new(Rubyconfig), new(Request), new(Response))
	createTables()

}

func initSqlite() {
	beego.Info("sqlite")
	orm.Debug = false
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db", 30)
}

func createTables() error {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	return err
}
