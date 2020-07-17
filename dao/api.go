/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package dao

type ApiFrDao interface {
	Get(apiId string) (apiInfo *interface{}, err error)
	//paged list
	List(apiQuery *interface{}) (apiList *interface{}, err error)
}

type defaultApiFrDao struct {
}

func NewDefaultApiFrDao() ApiFrDao {
	return &defaultApiFrDao{}
}

func (*defaultApiFrDao) Get(apiId string) (apiInfo *interface{}, err error) {
	panic("implement me")
}
func (*defaultApiFrDao) List(apiQuery *interface{}) (apiList *interface{}, err error) {
	panic("implement me")
}
