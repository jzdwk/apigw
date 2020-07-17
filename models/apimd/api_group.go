package apimd

import "time"

type ApiGroup struct {
	ID          string    `orm:"column(id);pk;type(char);size(36)" json:"id,omitempty"`
	Name        string    `orm:"column(name);type(varchar);size(20)" json:"name"`
	Description string    `orm:"column(description);type(varchar);size(255)" json:"description"`
	User        string    `orm:"column(user_id);type(varchar);size(36)" json:"userid"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);auto_now_add" json:"createtime"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);auto_now" json:"updatetime"`
	AuthType    int       `orm:"column(auth_type);type(varchar);size(2)" json:"authtype"`
	DomainId    string    `orm:"column(domain_id);type(char);size(36)" json:"domain"`
	Status      int       `orm:"column(status);size(2)" json:"status"`
}
