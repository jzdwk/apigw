/*
@Time : 20-7-15
@Author : jzd
@Project: apigw
*/
package service

type Acl interface {
	//Post
	Create(templateId string) (koUUID string, err error)
	//Delete
	Delete(koUUID string) (err error)
}
