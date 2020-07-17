/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package manager

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
