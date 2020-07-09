package main

import (
	"apigw/controllers/filter"
	"apigw/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/goharbor/harbor/src/core/config"


	_ "github.com/mvc/models"
	_ "github.com/mvc/routers"
	"os"
	"strconv"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = config.SessionCookieName
	logs.Info("initializing configurations...")
	//db init 使用
	database, err := config.Database()
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	//初始化route
	routers.Init()
	beego.Run()

}
