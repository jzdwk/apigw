/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package service

type Auth interface {
	//Post
	Create(config *interface{}) (koUUID string, err error)
	//Delete
	Delete(koUUID string) (err error)
}
