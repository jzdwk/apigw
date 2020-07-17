/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package dao

const ApiGroup = "api_group"

type GroupDao interface {
	//group CRUD
	Create(groupInfo *interface{}) (uuid string, err error)
	Delete(uuid string) error
	Update(uuid string, groupInfo *interface{}) error
	Get(id string) (groupInfo *interface{}, err error)
	//paged list
	List(groupQuery *interface{}) (groupList interface{}, err error)
	All() (groupInfo *interface{}, err error)
}

type defaultGroupDao struct{}

func (*defaultGroupDao) Create(groupInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (*defaultGroupDao) Delete(uuid string) error {
	panic("implement me")
}

func (*defaultGroupDao) Update(uuid string, groupInfo *interface{}) error {
	panic("implement me")
}

func (*defaultGroupDao) Get(id string) (groupInfo *interface{}, err error) {
	panic("implement me")
}

func (*defaultGroupDao) List(groupQuery *interface{}) (groupList interface{}, err error) {
	panic("implement me")
}

func (*defaultGroupDao) All() (groupInfo *interface{}, err error) {
	panic("implement me")
}

func NewDefaultGroupDao() GroupDao {
	return &defaultGroupDao{}
}
