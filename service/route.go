/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package service

import (
	"apigw/kong"
	"apigw/kong/crud"
	"apigw/kong/solver"
	"apigw/util/logs"
	gokong "github.com/hbagdi/go-kong/kong"
)

type Route interface {
	//Post
	Create(config *interface{}) (koUUID string, err error)
	//Delete
	Delete(koUUID string) (err error)
}

type defaultRouteSvc struct{}

func (r *defaultRouteSvc) Create(config *interface{}) (koUUID string, err error) {
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
	/* str := "userGo1"
	e := crud.Event{
		Op:   crud.Create,
		Kind: "consumer",
		Obj:  &kong.Consumer{Consumer: gokong.Consumer{Username: &str}},
	}*/
	routeName := "testApi"
	method1, method2 := "http", "https"
	path := "/test"
	e := crud.Event{
		Op:   crud.Create,
		Kind: "route",
		Obj:  &kong.Route{Route: gokong.Route{Name: &routeName, Methods: []*string{&method1, &method2}, Paths: []*string{&path}}},
	}
	result, err := solver.Solve(e)
	if err != nil {
		logs.Error("solver err: ", err)
	}
	logs.Info("result: ", *(result.(*kong.Consumer).Consumer.Username))

	return *(result.(*kong.Consumer).Consumer.ID), nil
}

func (r *defaultRouteSvc) Delete(koUUID string) (err error) {
	panic("implement me")
}
