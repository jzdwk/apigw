/*
@Time : 20-7-16
@Author : jzd
@Project: apigw
*/
package dao

//api template dao , only get/list. if do CUD, try dao.WithTransaction
type ApiTplDao interface {
	Get(apiTplId string) (apiTplInfo *interface{}, err error)
	//paged list
	List(apiTplQuery *interface{}) (apiTplList *interface{}, err error)
}

type defaultApiTplDao struct {
}

func NewDefaultApiTplDao() ApiTplDao {
	return &defaultApiTplDao{}
}

func (*defaultApiTplDao) Get(apiId string) (apiInfo *interface{}, err error) {
	panic("implement me")
}
func (*defaultApiTplDao) List(apiQuery *interface{}) (apiList *interface{}, err error) {
	panic("implement me")
}
