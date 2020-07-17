package ko_service_acl

type KoServiceAcl struct {
	Id         string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoPluginId string `orm:"column(ko_plugin_id);type(char);size(36)" json:"kopluginid"`
	ServiceId  string `orm:"column(service_id);type(char);size(36)" json:"serviceid"`
	Config     string `orm:"column(config);type(text)" json:"config"`
	Whitelist  string `orm:"column(whitelist);type(varchar);size(50)" json:"whitelist"`
	Status     string `orm:"column(status);type(varchar);size(50)" json:"status"`
}
