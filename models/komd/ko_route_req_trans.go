package komd

type KoRouteReqTrans struct {
	Id           string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	ReqTransId   string `orm:"column(req_trans_id);type(char);size(36)" json:"reqtransid"`
	KoRouteId    string `orm:"column(ko_route_id);type(char);size(36)" json:"korouteid"`
	KoPluginUuid string `orm:"column(ko_plugin_uuid);type(char);size(36)" json:"kopluginuuid"`
}
