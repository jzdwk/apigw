package app_group_dc

type AppGroupDc struct {
	Id        string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	AppId     string `orm:"column(app_id);type(char);size(36)" json:"appid"`
	GroupDcId string `orm:"column(group_dc_id);type(char);size(36)" json:"groupdcid"`
}
