package models

import (
	"time"
)
type DcDomain struct {
	ID   string                   `orm:"column(id);pk;type(char);size(36)" json:"id,omitempty"`
	Domain   string               `orm:"column(domain);type(varchar);size(255)" json:"domain"`
	Status int                    `orm:"column(status);size(2);null" json:"status"`
	Count  int                    `orm:"column(count)" json:"count"`
	User  string                  `orm:"column(user);type(varchar);size(20)" json:"user"`
	Name   string                 `orm:"column(name);type(varchar);size(20)" json:"name"`
	ApiGroup []*ApiGroup `orm:"reverse(many)"`
}

type ApiGroup struct {
	ID   string                 `orm:"column(id);pk;type(char);size(36)" json:"id,omitempty"`
	Name   string               `orm:"column(name);type(varchar);size(20)" json:"name"`
	Description string          `orm:"column(description);type(varchar);size(255)" json:"description"`
	UserId  string              `orm:"column(user_id);type(char);size(36)" json:"userid"`
	CreateTime time.Time        `orm:"column(create_time);type(datetime);auto_now_add" json:"createtime"`
	UpdateTime time.Time        `orm:"column(update_time);type(datetime);auto_now" json:"updatetime"`
	AuthType  int               `orm:"column(auth_type);type(varchar);size(2)" json:"authtype"`
	DomainId *DcDomain `orm:"rel(fk)"`
}

