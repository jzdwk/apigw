/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
	"github.com/astaxie/beego/orm"
)

type DomainManager interface {
	Create(domainId string) (uuid string, err error)
	Delete(domainId string) error
	Update(domainId string, domainInfo *interface{}) error
	Get(domainId string) (domainInfo *interface{}, err error)
	List(domainQuery *interface{}) (domainList *interface{}, err error)
	All() (domainList *interface{}, err error)
}

type defaultDomainMg struct {
	domainDao dao.DomainDao
}

func NewDefaultDomainMg() DomainManager {
	return &defaultDomainMg{domainDao: dao.NewDefaultDomainDao()}
}

func (*defaultDomainMg) Create(domainId string) (uuid string, err error) {
	panic("implement me")
}

func (*defaultDomainMg) Delete(domainId string) error {
	//check delete like
	dao.IsExist(dao.Ormer(), nil, nil)
	panic("implement me")
}

func (d *defaultDomainMg) Update(domainId string, domainInfo *interface{}) error {
	//check update like
	dao.IsExist(dao.Ormer(), nil, nil)
	// or with transaction
	_ = dao.WithTransaction(func(o orm.Ormer) error {
		dao.IsExist(o, nil, nil)
		//todo more
		//update
		err := d.domainDao.Update("example", nil)
		return err
	})
	panic("implement me")
}

func (*defaultDomainMg) Get(domainId string) (domainInfo *interface{}, err error) {
	panic("implement me")
}

func (*defaultDomainMg) List(domainQuery *interface{}) (domainList *interface{}, err error) {
	panic("implement me")
}

func (*defaultDomainMg) All() (domainList *interface{}, err error) {
	panic("implement me")
}
