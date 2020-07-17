package main

import (
	"apigw/controllers/filter"
	"apigw/initial"
	"github.com/astaxie/beego"
)

func main() {
	//Db init
	initial.InitDb()
	//filter init
	beego.InsertFilter("/*", beego.BeforeRouter, filter.SecurityFilter)
	beego.Run()
}

