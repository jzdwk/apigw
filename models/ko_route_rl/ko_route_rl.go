package ko_route_rl

type KoRouteRl struct {
	Id           string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoPluginUuid string `orm:"column(ko_plugin_uuid);type(char);size(36)" json:"kopluginuuid"`
	KoRouteId    string `orm:"column(ko_route_id);type(char);size(36)" json:"korouteid"`
	RateLimitId  string `orm:"column(rate_limit_id);type(char);size(36)" json:"ratelimitid"`
}
