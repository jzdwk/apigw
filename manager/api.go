/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

import "apigw/dao"

type ApiManager interface {
	//template
	CreateTemplate(templateInfo *interface{}) (uuid string, err error)
	DeleteTemplate(templateId string) error
	UpdateTemplate(templateId string, templateInfo *interface{}) error
	GetTemplate(templateId string) (templateInfo *interface{}, err error)
	//paged list
	ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error)
	//args
}

type defaultApiMg struct {
	apiDao dao.ApiFrDao
}

func (a *defaultApiMg) CreateTemplate(templateInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (a *defaultApiMg) DeleteTemplate(templateId string) error {
	panic("implement me")
}

func (a *defaultApiMg) UpdateTemplate(templateId string, templateInfo *interface{}) error {
	panic("implement me")
}

func (a *defaultApiMg) GetTemplate(templateId string) (templateInfo *interface{}, err error) {
	//a.apiDao.Get(templateId)
	panic("implement me")
}

func (a *defaultApiMg) ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error) {
	//a.apiDao.List(templateId)
	panic("implement me")
}
