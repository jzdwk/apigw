/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package dao

import (
	"apigw/models/ecpinfo"
	"apigw/util/logs"
)

type EcpDao interface {
	//create
	Create(ecp *ecpinfo.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpinfo.EcpInfo, error)
	//todo more
	Update(uuid string, ecpInfo *ecpinfo.EcpInfo) error
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

func (e *defaultEcpDao) Update(uuid string, ecpInfo *ecpinfo.EcpInfo) error {
	panic("implement me")
}

func (e *defaultEcpDao) Delete(uuid string) error {
	panic("implement me")
}

func NewDefalutEcpDao() EcpDao {
	return &defaultEcpDao{}
}

const EcpInfo = "ecp_info"

// post example
func (e *defaultEcpDao) Create(ecp *ecpinfo.EcpInfo) (id string, err error) {
	ecp.UUID = UUID()
	if _, err := Ormer().Insert(ecp); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// post example
func (e *defaultEcpDao) Get(uuid string) (rst *ecpinfo.EcpInfo, err error) {
	ecp := &ecpinfo.EcpInfo{}
	qs := Ormer().QueryTable(EcpInfo).Filter("id", uuid)
	if err := qs.One(ecp); err != nil {
		return nil, err
	}
	//wrap
	return ecp, nil
}
