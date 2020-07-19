/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package manager

type AppManager interface {
	//add my cao
	CreateApp(appInfo *interface{}) (uuid string, err error)
	DeleteApp(appId string) error
	//could update?
	//UpdateApp(appId string,capInfo *interface{})error
	GetApp(appId string) (appInfo *interface{}, err error)
	ListApp(appQuery *interface{}) (appList *interface{}, err error)
	// bind
	BindGroupEcp(groupId, ecpId, appId string) error
	UnbindGroupEcp(groupId, ecpId, appId string) error
}

type defaultAppMg struct {
	//kong

	//dao
}

func (a *defaultAppMg) CreateApp(appInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (a *defaultAppMg) DeleteApp(appId string) error {
	panic("implement me")
}

func (a *defaultAppMg) GetApp(appId string) (appInfo *interface{}, err error) {
	panic("implement me")
}

func (a *defaultAppMg) ListApp(appQuery *interface{}) (appList *interface{}, err error) {
	panic("implement me")
}

func (a *defaultAppMg) BindGroupEcp(groupId, ecpId, appId string) error {
	panic("implement me")
}

func (a *defaultAppMg) UnbindGroupEcp(groupId, ecpId, appId string) error {
	panic("implement me")
}
