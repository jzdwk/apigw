/*
@Time : 20-7-17
@Author : jzd
@Project: apigw
*/
package dao

type ApiDao interface {
	Get(apiId string) (apiInfo *interface{}, err error)
	//paged list
	List(apiQuery *interface{}) (apiList *interface{}, err error)
}
