/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package ecpinfo

import (
	ecpinfoDao "apigw/dao/ecpinfo"
	"apigw/models/ecpinfo"
	"apigw/util/logs"
)

//ecp interface define all ecp CRUDs
type EcpManager interface {
	//do Post
	Post(ecp *ecpinfo.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpinfo.EcpInfoResp, error)
	//todo implement
	Delete(uuid string) (interface{}, error)
	Update(uuid string, info interface{}) (interface{}, error)
}

//
func NewDefaultEcpManager() EcpManager {
	return &defaultEcpMg{dao: ecpinfoDao.NewDefalutDao()}
}

type defaultEcpMg struct {
	//dao reference
	dao ecpinfoDao.Dao
}

func (e *defaultEcpMg) Delete(uuid string) (interface{}, error) {
	panic("implement me")
}

func (e *defaultEcpMg) Update(uuid string, info interface{}) (interface{}, error) {
	panic("implement me")
}

// post example
func (e *defaultEcpMg) Post(ecp *ecpinfo.EcpInfo) (id string, err error) {
	if _, err := e.dao.Post(ecp); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// get example
func (e *defaultEcpMg) Get(uuid string) (ecp *ecpinfo.EcpInfoResp, err error) {
	ecp, err = e.dao.Get(uuid)
	if err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return nil, err
	}
	return ecp, nil
}
