package main

import (
	_ "zjlife/docs"
	"zjlife/logger"
	_ "zjlife/models"
	_ "zjlife/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("dbuser")+":"+beego.AppConfig.String("dbpassword")+"@/"+beego.AppConfig.String("dbname")+"?charset=utf8")
}

func main() {
	logger.SetLevel(logger.DEBUG, logger.DEBUG)
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
