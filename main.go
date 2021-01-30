package main

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "sogo/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	orm.Debug = true
	beego.Run()
}

func init() {

	sqlconn, _ := beego.AppConfig.String("sqlconn")


	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", sqlconn)

	// create table
	//orm.RunSyncdb("default", false, true)

	orm.RunCommand()
}
