package api_access_control

type ApiAccessControl struct {
	Id     string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	User   string `orm:"column(user);type(varchar);size(20)" json:"user"`
	Config string `orm:"column(config);type(text)" json:"config"`
	Enable int    `orm:"column(enable);size(2)" json:"enable"`
}
