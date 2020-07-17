package api_group_dc

type ApiGroupDc struct {
	Id           string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	GroupId      string `orm:"column(group_id);type(char);size(36)" json:"groupid"`
	DatacenterId string `orm:"column(datacenter_id);type(char);size(36)" json:"datacenterid"`
}
