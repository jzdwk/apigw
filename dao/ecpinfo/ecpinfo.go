/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package ecpinfo

import (
	"apigw/dao"
	"apigw/models/ecpinfo"
	"apigw/util/logs"
)

type Dao interface {
	//post
	Post(ecp *ecpinfo.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpinfo.EcpInfoResp, error)
	//todo more
}

type defaultDao struct {
}

func NewDefalutDao() Dao {
	return &defaultDao{}
}

const EcpInfo = "ecp_info"

// post example
func (e *defaultDao) Post(ecp *ecpinfo.EcpInfo) (id string, err error) {
	ecp.UUID = dao.UUID()
	if _, err := dao.Ormer().Insert(ecp); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// post example
func (e *defaultDao) Get(uuid string) (rst *ecpinfo.EcpInfoResp, err error) {
	ecp := &ecpinfo.EcpInfo{}
	qs := dao.Ormer().QueryTable(EcpInfo).Filter("id", uuid)
	if err := qs.One(ecp); err != nil {
		return nil, err
	}
	//wrap
	return &ecpinfo.EcpInfoResp{EcpInfo: ecp}, nil
}
