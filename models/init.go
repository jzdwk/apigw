package models

import (
	"apigw/models/apimd"
	"apigw/models/appmd"
	cap2 "apigw/models/capmd"
	"apigw/models/ecpmd"
	"apigw/models/komd"
	"github.com/astaxie/beego/orm"
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	//ecpmd
	orm.RegisterModel(new(ecpmd.EcpInfo), new(ecpmd.EcpDomain))
	//apimd
	orm.RegisterModel(
		new(apimd.ApiGroup),
		new(apimd.ApiAccessControl),
		new(apimd.ApiArgsfr),
		new(apimd.ApiArgsBk),
		new(apimd.ApiBk),
		new(apimd.ApiFr),
		new(apimd.ApiGroupDc),
		new(apimd.ApiRateLimit),
		new(apimd.ApiReqTrans))
	//kong
	orm.RegisterModel(
		new(komd.KoRoute),
		new(komd.KoRouteAc),
		new(komd.KoRoutePlugin),
		new(komd.KoRouteReqTrans),
		new(komd.KoRouteRl),
		new(komd.KoService),
		new(komd.KoServiceAcl),
		new(komd.KoServiceAuth),
		new(komd.KoServiceHmac),
		new(komd.KoServicePlugin))
	//appmd
	orm.RegisterModel(new(appmd.AppInfo), new(appmd.AppCredential), new(appmd.AppGroupDc))
	//capmd
	orm.RegisterModel(new(cap2.CapGroupDc), new(cap2.CapMy))

}
