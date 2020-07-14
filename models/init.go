package models

import (
	"apigw/models/ecpinfo"
	"github.com/astaxie/beego/orm"
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(ecpinfo.EcpInfo))
}
