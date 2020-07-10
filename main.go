package main

import (
	"apigw/conf"
	"apigw/controller/filter"
	"apigw/initial"
	"apigw/router"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info("initializing configurations...")
	conf.Init()
	//db init 使用
	logs.Info("initializing database...")
	initial.InitDb()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	//初始化route
	logs.Info("initializing router...")
	router.Init()
	beego.Run()

}
