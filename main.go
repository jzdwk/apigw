package main

import (
	"apigw/config"
	"apigw/controller/filter"
	"apigw/initial"
	"apigw/router"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = config.SessionCookieName
	logs.Info("initializing configurations...")
	//db init 使用
	//database, err := config.Database()
	logs.Info("initializing configurations...")
	initial.InitDb()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	//初始化route
	router.Init()
	beego.Run()

}
