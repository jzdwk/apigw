/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

type ApiManager interface {
	//template CRUD
	CreateTemplate(templateInfo interface{}) (uuid string, err error)
	DeleteTemplate(tId string) error
	UpdateTemplate(tId string, templateInfo interface{}) error
	GetTemplate(tId string) (templateInfo interface{}, err error)
	ListTemplate(templateQuery interface{}) (templateInfos interface{}, err error)
}
