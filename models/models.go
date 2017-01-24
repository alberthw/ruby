package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//initSqlite()
	initPostgresql()
	orm.RegisterModel(new(Rubyconfig), new(Command), new(Reposetting), new(Filerepo), new(Devicesystemconfig), new(Devicehardwareconfig))
	createTables()
}

func initSqlite() {
	beego.Info("sqlite")
	orm.Debug = false
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db", 30)
}

func initPostgresql() {
	beego.Info("Postgresql")
	orm.Debug = false
	orm.RegisterDriver("postgres", orm.DRPostgres)
	connstr := "user=postgres password=123456 dbname=ruby sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connstr)
}

func createTables() error {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	return err
}
