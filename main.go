package main

import (
	"apigw/controllers/filter"
	"apigw/initial"
	_ "apigw/models"
	"apigw/routers"
	_ "apigw/routers"
	"apigw/util/logs"
	"github.com/astaxie/beego"
)

func main() {
	// log init
	logs.InitLogs()
	// router init
	routers.InitRouter()
	// db init
	initial.InitDb()
	// filter init
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	beego.Run()
}
