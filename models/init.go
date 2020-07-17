package models

import (
	"apigw/models/api_access_control"
	"apigw/models/api_args_bk"
	"apigw/models/api_args_fr"
	"apigw/models/api_bk"
	"apigw/models/api_fr"
	"apigw/models/api_group"
	"apigw/models/api_group_dc"
	"apigw/models/api_rate_limit"
	"apigw/models/api_req_trans"
	"apigw/models/app_credential"
	"apigw/models/app_group_dc"
	"apigw/models/app_info"
	"apigw/models/cap_group_dc"
	"apigw/models/cap_my"
	"apigw/models/dc_domain"
	"apigw/models/ecpinfo"
	"apigw/models/ko_route"
	"apigw/models/ko_route_ac"
	"apigw/models/ko_route_plugin"
	"apigw/models/ko_route_req_trans"
	"apigw/models/ko_route_rl"
	"apigw/models/ko_service"
	"apigw/models/ko_service_acl"
	"apigw/models/ko_service_auth"
	"apigw/models/ko_service_hmac"
	"apigw/models/ko_service_plugin"
	"github.com/astaxie/beego/orm"
)

func init() {
	//print sql
	orm.Debug = true
	// init orm tables
	orm.RegisterModel(new(ecpinfo.EcpInfo))
	orm.RegisterModel(new(api_group.ApiGroup), new(dc_domain.DcDomain), new(api_access_control.ApiAccessControl),
		new(api_args_fr.ApiArgsfr), new(api_args_bk.ApiArgsBk),  new(api_bk.ApiBk),
		new(api_fr.ApiFr), new(api_group_dc.ApiGroupDc), new(api_rate_limit.ApiRateLimit),
		new(api_req_trans.ApiReqTrans), new(app_credential.AppCredential), new(app_group_dc.AppGroupDc), new(app_info.AppInfo),
		new(cap_group_dc.CapGroupDc), new(cap_my.CapMy), new(ko_route.KoRoute), new(ko_route_ac.KoRouteAc),
		new(ko_route_plugin.KoRoutePlugin), new(ko_route_req_trans.KoRouteReqTrans), new(ko_route_rl.KoRouteRl),
		new(ko_service.KoService), new(ko_service_acl.KoServiceAcl), new(ko_service_auth.KoServiceAuth), new(ko_service_hmac.KoServiceHmac),
		new(ko_service_plugin.KoServicePlugin))
}
