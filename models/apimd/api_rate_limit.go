package apimd

type ApiRateLimit struct {
	Id     string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	User   string `orm:"column(user);type(varchar);size(20)" json:"user"`
	Config string `orm:"column(config);type(text)" json:"config"`
	Status int    `orm:"column(status);size(2)" json:"status"`
}
