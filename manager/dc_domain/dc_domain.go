/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package dc_domain

import (
	DomainDao "apigw/dao/dc_domain"
	"apigw/models"
	"apigw/util/logs"
)

//ecp interface define all ecp CRUDs
type DomainManager interface {
	//do Post
	Post(domain *models.DcDomain) (error)
	////get
	//Get(uuid string) (*models.DcDomain, error)
}

type defaultDomainMg struct {
	//dao reference
	dao DomainDao.Dao
}

//
func NewDefaultDomainManager() DomainManager {
	return &defaultDomainMg{dao: DomainDao.NewDefalutDao()}
}

// post example
func (d *defaultDomainMg) Post(domain *models.DcDomain) ( err error) {
	if  err := d.dao.Post(domain); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return err
	}
	return nil
}

//// get example
//func (d *defaultDomainMg) Get(uuid string) (domain *models.DcDomain, err error) {
//	ecp, err = d.dao.Get(uuid)
//	if err != nil {
//		logs.Error("add ecp info err %v.", err.Error())
//		return nil, err
//	}
//	return ecp, nil
//}
