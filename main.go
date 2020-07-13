package main

import (
	_ "apigw/conf"
	"apigw/controller/filter"
	"apigw/initial"
	_ "apigw/model"
	_ "apigw/router"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("initializing configurations...")
	//init config
	//conf.InitConfig()
	//init db
	logs.Info("initializing database...")
	initial.InitDb()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	//init route
	logs.Info("initializing router...")
	//init logger
	//gwlogs.InitLog()
	beego.Run()
}
