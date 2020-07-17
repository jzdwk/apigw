package komd

type KoServicePlugin struct {
	Id          string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	PluginName  string `orm:"column(plugin_name);type(varchar);size(255)" json:"pluginname"`
	PluginId    string `orm:"column(plugin_id);type(char);size(36)" json:"pluginid"`
	KoServiceId string `orm:"column(ko_service_id);type(char);size(36)" json:"koserviceid"`
}
