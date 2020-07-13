package model

import (
	"github.com/astaxie/beego/orm"
)

//init db tables
func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(EcpInfo))

}
