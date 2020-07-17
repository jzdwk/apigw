/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

import "apigw/dao"

type ApiTplManager interface {
	//template
	CreateTemplate(templateInfo *interface{}) (uuid string, err error)
	DeleteTemplate(templateId string) error
	UpdateTemplate(templateId string, templateInfo *interface{}) error
	GetTemplate(templateId string) (templateInfo *interface{}, err error)
	//paged list
	ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error)
	//args
}

type defaultApiTplMg struct {
	apiTplDao dao.ApiTplDao
}

func (a *defaultApiTplMg) CreateTemplate(templateInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (a *defaultApiTplMg) DeleteTemplate(templateId string) error {
	panic("implement me")
}

func (a *defaultApiTplMg) UpdateTemplate(templateId string, templateInfo *interface{}) error {
	panic("implement me")
}

func (a *defaultApiTplMg) GetTemplate(templateId string) (templateInfo *interface{}, err error) {
	//a.apiTplDao.Get(templateId)
	panic("implement me")
}

func (a *defaultApiTplMg) ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error) {
	//a.apiTplDao.List(templateId)
	panic("implement me")
}
