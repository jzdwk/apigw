/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package dao

type CapDao interface {
	Create(capInfo *interface{}) (uuid string, err error)
	Get(capId string) (capInfo *interface{}, err error)
	//paged list
	List(capQuery *interface{}) (capList *interface{}, err error)
}

type defaultCapDao struct{}

func NewDefaultCapDao() CapDao {
	return &defaultCapDao{}
}

func (c *defaultCapDao) Create(capInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (c *defaultCapDao) Get(capId string) (capInfo *interface{}, err error) {
	panic("implement me")
}

func (c *defaultCapDao) List(capQuery *interface{}) (capList *interface{}, err error) {
	panic("implement me")
}
