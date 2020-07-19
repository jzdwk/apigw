/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

import "apigw/dao"

type ApiTplManager interface {
	//template
	CreateApiTemplate(templateInfo *interface{}) (uuid string, err error)
	DeleteApiTemplate(templateId string) error
	//update api or template
	UpdateApiTemplate(templateId string, templateInfo *interface{}) error
	GetApiTemplate(templateId string) (templateInfo *interface{}, err error)
	//paged list
	ListApiTemplate(templateQuery *interface{}) (templateList *interface{}, err error)
}

type defaultApiTplMg struct {
	apiTplDao dao.ApiTplDao
}

func (a *defaultApiTplMg) CreateApiTemplate(templateInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (a *defaultApiTplMg) DeleteApiTemplate(templateId string) error {
	panic("implement me")
}

func (a *defaultApiTplMg) UpdateApiTemplate(templateId string, templateInfo *interface{}) error {
	panic("implement me")
}

func (a *defaultApiTplMg) GetApiTemplate(templateId string) (templateInfo *interface{}, err error) {
	//a.apiTplDao.Get(templateId)
	panic("implement me")
}

func (a *defaultApiTplMg) ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error) {
	//a.apiTplDao.List(templateId)
	panic("implement me")
}
