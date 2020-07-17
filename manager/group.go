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
	Create(groupId string) (uuid string, err error)
	Delete(groupId string) error
	Update(groupId string, groupInfo *interface{}) error
	Get(groupId string) (groupInfo *interface{}, err error)
	List(groupQuery *interface{}) (groupList *interface{}, err error)
	All() (groupList *interface{}, err error)
}

//default impl
type defaultGroupMg struct {
	groupDao dao.GroupDao
}

func (*defaultGroupMg) Create(groupId string) (uuid string, err error) {
	panic("implement me")
}

func (*defaultGroupMg) Delete(groupId string) error {
	//1. check
	//2. delete
	panic("implement me")
}

func (*defaultGroupMg) Update(groupId string, groupInfo *interface{}) error {
	//1. check
	//2. update
	panic("implement me")
}

func (*defaultGroupMg) Get(groupId string) (groupInfo *interface{}, err error) {
	panic("implement me")
}

func (*defaultGroupMg) List(groupQuery *interface{}) (groupList *interface{}, err error) {
	panic("implement me")
}

func (*defaultGroupMg) All() (groupList *interface{}, err error) {
	panic("implement me")
}

//to do check
func (*defaultGroupMg) GroupNumByEcp(ecpId string) (count int, err error) {
	panic("implement me")
}

func NewDefualtGroupMg() GroupManager {
	return &defaultGroupMg{groupDao: dao.NewDefaultGroupDao()}
}
