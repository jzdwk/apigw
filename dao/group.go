/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package dao

const ApiGroup = "api_group"

type GroupDao interface {
	//group CRUD
	Create() (uuid string, err error)
	Delete(id string) error
	Update(id string, gInfo interface{}) error
	Get(id string) (gInfo interface{}, err error)
	List(gQuery interface{}) (gInfo interface{}, err error)
}

type defaultGroupDao struct{}

func NewDefaultGroupDao() GroupDao {
	return &defaultGroupDao{}
}

func (defaultGroupDao) Create() (uuid string, err error) {
	panic("implement me")
}

func (defaultGroupDao) Delete(id string) error {
	panic("implement me")
}

func (defaultGroupDao) Update(id string, gInfo interface{}) error {
	panic("implement me")
}

func (defaultGroupDao) Get(id string) (gInfo interface{}, err error) {
	panic("implement me")
}

func (defaultGroupDao) List(gQuery interface{}) (gInfo interface{}, err error) {
	panic("implement me")
}
