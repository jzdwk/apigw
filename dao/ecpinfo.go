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

type Dao interface {
	//post
	Create(ecp *ecpinfo.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpinfo.EcpInfo, error)
	//todo more
}

type defaultDao struct {
}

func NewDefalutDao() Dao {
	return &defaultDao{}
}

const EcpInfo = "ecp_info"

// post example
func (e *defaultDao) Create(ecp *ecpinfo.EcpInfo) (id string, err error) {
	ecp.UUID = UUID()
	if _, err := Ormer().Insert(ecp); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// post example
func (e *defaultDao) Get(uuid string) (rst *ecpinfo.EcpInfo, err error) {
	ecp := &ecpinfo.EcpInfo{}
	qs := Ormer().QueryTable(EcpInfo).Filter("id", uuid)
	if err := qs.One(ecp); err != nil {
		return nil, err
	}
	//wrap
	return ecp, nil
}
