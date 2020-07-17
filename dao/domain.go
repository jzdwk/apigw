/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package dao

type DomainDao interface {
	//create
	Create(domainInfo *interface{}) (uuid string, err error)
	//get
	Get(uuid string) (domainInfo *interface{}, err error)
	//todo more
	Update(uuid string, domainInfo *interface{}) error
	Delete(uuid string) error
	//paged list
	List(domainQuery *interface{}) (domainList *interface{}, err error)
	//all list
	All() (ecpList *interface{}, err error)
}

type defaultDomainDao struct {
}

func (*defaultDomainDao) Create(domainInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (*defaultDomainDao) Get(uuid string) (domainInfo *interface{}, err error) {
	panic("implement me")
}

func (*defaultDomainDao) Update(uuid string, domainInfo *interface{}) error {
	panic("implement me")
}

func (*defaultDomainDao) Delete(uuid string) error {
	panic("implement me")
}

func (*defaultDomainDao) List(domainQuery *interface{}) (domainList *interface{}, err error) {
	panic("implement me")
}

func (*defaultDomainDao) All() (ecpList *interface{}, err error) {
	panic("implement me")
}
