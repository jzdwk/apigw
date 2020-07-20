/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package service

type Acl interface {
	//Create
	Create(templateId string) (koUUID string, err error)
	//Delete
	Delete(koUUID string) (err error)
}
