/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
	"apigw/models/ecpmd"
	"apigw/util/logs"
)

//ecpmd interface define all ecpmd CRUDs
type EcpManager interface {
	//do Create
	Create(ecp *ecpmd.EcpInfo) (string, error)
	//get
	Get(uuid string) (*ecpmd.EcpInfoResp, error)
	//todo implement
	Delete(uuid string) error
	Update(uuid string, ecpInfo *interface{}) error
	List(ecpQuery *interface{}) (ecpList *interface{}, err error)
	All() (ecpList *interface{}, err error)
}

//
func NewDefaultEcpManager() EcpManager {
	return &defaultEcpMg{ecpDao: dao.NewDefalutEcpDao()}
}

type defaultEcpMg struct {
	//domainDao reference
	ecpDao dao.EcpDao
}

func (e *defaultEcpMg) Delete(uuid string) error {
	// delete
	panic("implement me")
}

func (e *defaultEcpMg) Update(uuid string, ecpInfo *interface{}) error {
	panic("implement me")
}

func (e *defaultEcpMg) List(ecpQuery *interface{}) (ecpList *interface{}, err error) {
	panic("implement me")
}

func (e *defaultEcpMg) All() (ecpList *interface{}, err error) {
	panic("implement me")
}

// post example
func (e *defaultEcpMg) Create(ecp *ecpmd.EcpInfo) (id string, err error) {
	if _, err := e.ecpDao.Create(ecp); err != nil {
		logs.Error("add ecpmd info err %v.", err.Error())
		return "", err
	}
	return ecp.UUID, nil
}

// get example
func (e *defaultEcpMg) Get(uuid string) (ecp *ecpmd.EcpInfoResp, err error) {
	ecpInfo, err := e.ecpDao.Get(uuid)
	if err != nil {
		logs.Error("add ecpmd info err %v.", err.Error())
		return nil, err
	}
	return &ecpmd.EcpInfoResp{EcpInfo: ecpInfo}, nil
}
