package api_bk

type ApiBk struct {
	Id       string `orm:"column(id);type(char);size(36);pk" json:"id"`
	Protocol int    `orm:"column(protocol);size(2)" json:"protocol"`
	Method   int    `orm:"column(method);size(4)" json:"method"`
	Path     string `orm:"column(path);type(varchar);size(255)" json:"path"`
	Auth     int    `orm:"column(auth);size(2)" json:"auth"`
	ApiFrId  string `orm:"column(api_fr_id);type(char);size(36)" json:"apifrid"`
}
