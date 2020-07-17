package apimd

type ApiFr struct {
	Id       string `orm:"column(id);type(char);pk;size(36)" json:"id"`
	Name     string `orm:"column(name);type(varchar);size(20)" json:"name"`
	GroupId  string `orm:"column(group_id);type(char);size(36)" json:"grouid"`
	User     string `orm:"column(user);type(varchar);size(20)" json:"user"`
	Path     string `orm:"column(path);type(varchar);size(255)" json:"path"`
	Method   int    `orm:"column(method);size(4)" json:"method"`
	Protocol int    `orm:"column(protocol);size(4)" json:"protocol"`
	Match    int    `orm:"column(match);size(4)" json:"match"`
	Status   int    `orm:"column(status);size(2)" json:"status"`
}
