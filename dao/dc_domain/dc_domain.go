package dc_domain

import (
	"apigw/dao"
	"apigw/models"
	"apigw/util/logs"
)

/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
type Dao interface {
	//post
	Post(domain *models.DcDomain) (error)
	//get
	//Get(uuid string) (*models.DcDomain, error)
	////todo more
}

type defaultDao struct {
}

func NewDefalutDao() Dao {
	return &defaultDao{}
}

const DomainInfo = "dc_domain"

// post example
func (e *defaultDao) Post(domain *models.DcDomain) (err error) {
	if _, err := dao.Ormer().Insert(domain); err != nil {
		logs.Error("add ecp info err %v.", err.Error())
		return err
	}
	return nil
}

//// post example
//func (e *defaultDao) Get(uuid string) (rst *models.DcDomain, err error) {
//	ecp := &ecpinfo.EcpInfo{}
//	qs := dao.Ormer().QueryTable(DomainInfo).Filter("id", uuid)
//	if err := qs.One(ecp); err != nil {
//		return nil, err
//	}
//	//wrap
//	return &*models.DcDomain{EcpInfo: ecp}, nil
//}
