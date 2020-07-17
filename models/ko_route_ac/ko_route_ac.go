package ko_route_ac

type KoRouteAc struct {
	Id              string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoRouteId       string `orm:"column(ko_route_id);type(char);size(36)" json:"korouteid"`
	KoPluginUuid    string `orm:"column(ko_plugin_uuid);type(char);size(36)" json:"kopluginuuid"`
	AccessControlId string `orm:"column(access_control_id);type(char);size(36)" json:"accesscontrolid"`
}
