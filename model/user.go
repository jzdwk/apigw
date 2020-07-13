/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package model

//user info from token
type User struct {
	Name  string
	Admin bool
	//todo more infos
	Detail string
}
