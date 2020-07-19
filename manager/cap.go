/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package manager

import "apigw/dao"

type CapManager interface {
	//add my cao
	AddMyCap(capInfo *interface{}) (uuid string, err error)
	GetMyCap(capId string) (capInfo *interface{}, err error)
	ListMyCap(capQuery *interface{}) (capList *interface{}, err error)
	// could delete?
	//DeleteMyCap(capId string) error
}

type defaultCapMg struct {
	//dao
	capDao dao.CapDao
}

func (c *defaultCapMg) AddMyCap(capInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (c *defaultCapMg) GetMyCap(capId string) (capInfo *interface{}, err error) {
	panic("implement me")
}

func (c *defaultCapMg) ListMyCap(capQuery *interface{}) (capList *interface{}, err error) {
	panic("implement me")
}
