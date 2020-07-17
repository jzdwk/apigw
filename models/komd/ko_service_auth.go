package komd

type KoServiceAuth struct {
	Id           string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoPluginUuid string `orm:"column(ko_plugin_uuid);type(char);size(36)" json:"kopluginuuid"`
	ServiceId    string `orm:"column(service_id);type(char);size(36)" json:"serviceid"`
	Config       string `orm:"column(config);type(text)" json:"config"`
}
