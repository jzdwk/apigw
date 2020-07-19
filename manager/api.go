/*
@Time : 20-7-17
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/kong/solver"
	"apigw/util/logs"
	gokong "github.com/hbagdi/go-kong/kong"
)

type ApiManager interface {
	CreateApi(apiTplId, ecpId string) error
	DeleteApi(apiId string) error
	GetApi(apiId string) (apiInfo *interface{}, err error)
	//paged list
	ListApi(apiQuery *interface{}) (apiList *interface{}, err error)
}

type defaultApiMg struct {
	//dao
	apiDao dao.ApiDao
	//kong

}

func (a *defaultApiMg) CreateApi(apiTplId, ecpId string) error {
	/* do it with transaction
	1. add group_dc or escape
	2. add ko_service or escape, add kong service in future mode
	3. add ko_service_plugin & ko_service_auth or ko_service_hmac, add kong auth-key/hmac in future mode
	4. add ko_service_plugin & ko_service_acl, add kong acl in future mode with whitelist is self
	5. add ko_route & ko_req_trans, ko_route_plugin, add kong route in future mode
	6. add kong trans plugin to kon
	*/

	/*
		kong example
	*/
	opt := kong.KongClientConfig{
		Address: "http://192.168.182.135:8001",
	}
	client, err := kong.GetKongClient(opt)
	if err != nil {
		logs.Error("get kong client err: ", err)
	}

	solver.BuildRegistry(client)
	/**/ str := "userGo1"
	e := crud.Event{
		Op:   crud.Create,
		Kind: "consumer",
		Obj:  &kong.Consumer{Consumer: gokong.Consumer{Username: &str}},
	}
	/*routeName :="testApi"
	e := crud.Event{
		Op:crud.Create,
		Kind:"route",
		Obj:&kong.Route{Route:gokong.Route{Name:"test",Methods:&["http","https"],Path:""}},
	}*/

	result, err := solver.Solve(e)
	if err != nil {
		logs.Error("solver err: ", err)
	}
	logs.Info("result: ", *(result.(*kong.Consumer).Consumer.Username))
	//调试代码

	panic("implement me")
}

func (a *defaultApiMg) DeleteApi(apiId string) error {
	panic("implement me")
}

func (a *defaultApiMg) ListApi(apiQuery *interface{}) (apiList *interface{}, err error) {
	panic("implement me")
}

//return apimd template info
func (a *defaultApiMg) GetApi(apiId string) (apiInfo *interface{}, err error) {
	panic("implement me")
}
