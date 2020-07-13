/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
	"apigw/model"
	"apigw/util/logs"
)

//ecp interface define all ecp CRUDs
type EcpManager interface {
	//do Post
	Post(ecp *model.EcpInfo) (string, error)
}

type defaultEcpMg struct {
}

func NewDefaultEcpManager() EcpManager {
	return &defaultEcpMg{}
}

// post example
func (e *defaultEcpMg) Post(ecp *model.EcpInfo) (id string, err error) {
	ecp.ID = uuid()
	if _, err := dao.Ormer().Insert(ecp); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return "", err
	}
	return ecp.ID, nil
}
