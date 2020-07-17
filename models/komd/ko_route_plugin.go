package komd

type KoRoutePlugin struct {
	Id         string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	PluginName string `orm:"column(plugin_name);type(varchar);size(255)" json:"pluginname"`
	PluginId   string `orm:"column(plugin_id);type(char);size(36)" json:"pluginid"`
	KoRouteId  string `orm:"column(ko_route_id);type(char);size(36)" json:"korouteid"`
}
