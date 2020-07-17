package appmd

type AppCredential struct {
	Id              string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoConsumerName  string `orm:"column(ko_consumer_name);type(varchar);size(30)" json:"koconsumername"`
	KoConsumerKey   string `orm:"column(ko_consumer_key);type(varchar);size(255)" json:"koconsumerkey"`
	KoConsumerUuid  string `orm:"column(ko_consumer_uuid);type(char);size(36)" json:"koconsumeruuid"`
	AppGroupDcId    string `orm:"column(app_group_dc_id);type(char);size(36)" json:"appgroupdcid"`
	KoConsumerGroup string `orm:"column(ko_consumer_group);type(varchar);size(255)" json:"koconsumergroup"`
}
