/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package dao

import (
	"apigw/models/ecpmd"
	"apigw/util/logs"
)

type EcpDao interface {
	//create
	Create(ecp *ecpmd.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpmd.EcpInfo, error)
	//todo more
	Update(uuid string, ecpInfo *ecpmd.EcpInfo) error
	Delete(uuid string) error
	//paged list
	List(ecpQuery *interface{}) (ecpList *interface{}, err error)
	//all list
	All() (ecpInfo *interface{}, err error)
}

type defaultEcpDao struct {
}

func (e *defaultEcpDao) List(ecpQuery *interface{}) (ecpList *interface{}, err error) {
	panic("implement me")
}

func (e *defaultEcpDao) All() (ecpInfo *interface{}, err error) {
	panic("implement me")
}

func (e *defaultEcpDao) Update(uuid string, ecpInfo *ecpmd.EcpInfo) error {
	panic("implement me")
}

func (e *defaultEcpDao) Delete(uuid string) error {
	panic("implement me")
}

func NewDefalutEcpDao() EcpDao {
	return &defaultEcpDao{}
}

// post example
func (e *defaultEcpDao) Create(ecp *ecpmd.EcpInfo) (id string, err error) {
	ecp.UUID = UUID()
	if _, err := Ormer().Insert(ecp); err != nil {
		logs.Error("add ecpmd info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// post example
func (e *defaultEcpDao) Get(uuid string) (rst *ecpmd.EcpInfo, err error) {
	ecp := &ecpmd.EcpInfo{}
	qs := Ormer().QueryTable(&ecpmd.EcpInfo{}).Filter("id", uuid)
	if err := qs.One(ecp); err != nil {
		return nil, err
	}
	//wrap
	return ecp, nil
}
