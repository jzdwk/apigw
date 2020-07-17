package capmd

type CapGroupDc struct {
	Id        string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	GroupDcId string `orm:"column(group_dc_id);type(char);size(36)" json:"groupdcid"`
	CapId     string `orm:"column(cap_id);type(char);size(36)" json:"capid"`
}
