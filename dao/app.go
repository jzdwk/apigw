/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package dao

type AppDao interface {
	Create(appInfo *interface{}) (uuid string, err error)
	Delete(appId string) error
	Get(appId string) (appInfo *interface{}, err error)
	List(appQuery *interface{}) (appList *interface{}, err error)
}

type defaultAppDao struct{}

func NewDefaultAppDao() AppDao {
	return &defaultAppDao{}
}

func (a *defaultAppDao) Create(appInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (a *defaultAppDao) Delete(appId string) error {
	panic("implement me")
}

func (a *defaultAppDao) Get(appId string) (appInfo *interface{}, err error) {
	panic("implement me")
}

func (a *defaultAppDao) List(appQuery *interface{}) (appList *interface{}, err error) {
	panic("implement me")
}
