/*
@Time : 20-7-17
@Author : jzd
@Project: apigw
*/
package manager

type ApiManager interface {
	//future mode
	CreateKoResource(work chan struct{}, apiTplId, ecpId string, backServiceInfo *interface{}) error
	Create

	DeleteTemplate(templateId string) error
	UpdateTemplate(templateId string, templateInfo *interface{}) error
	GetTemplate(templateId string) (templateInfo *interface{}, err error)
	//paged list
	ListTemplate(templateQuery *interface{}) (templateList *interface{}, err error)
	//args
}
