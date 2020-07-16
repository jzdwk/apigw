/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
)

type GroupManager interface {
	Create(dId string) (uuid string, err error)
	Delete(gId string) error
	Update(gId string, gInfo interface{}) error
	Get(gId string) (gInfo interface{}, err error)
	List(gId string, gQuery interface{}) (gInfo interface{}, err error)
}

//default impl
type defaultGroupMg struct {
	groupDao dao.GroupDao
}

func (*defaultGroupMg) Create(dId string) (uuid string, err error) {
	panic("implement me")
}

func (*defaultGroupMg) Delete(gId string) error {
	panic("implement me")
}

func (*defaultGroupMg) Update(gId string, gInfo interface{}) error {
	panic("implement me")
}

func (*defaultGroupMg) Get(gId string) (gInfo interface{}, err error) {
	panic("implement me")
}

func (*defaultGroupMg) List(gId string, gQuery interface{}) (gInfo interface{}, err error) {
	panic("implement me")
}

func NewDefualtGroupMg() GroupManager {
	return &defaultGroupMg{groupDao: dao.NewDefaultGroupDao()}
}
