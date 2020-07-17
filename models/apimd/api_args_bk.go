package apimd

type ApiArgsBk struct {
	Id       string `orm:"column(id);type(char);size(36);pk" json:"id"`
	ApiBkId  string `orm:"column(api_bk_id);type(char);size(36)" json:"apibkid"`
	Name     string `orm:"column(name);type(varchar);size(255)" json:"name"`
	Location int    `orm:"column(location);size(2)" json:"location"`
}
