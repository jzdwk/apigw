package model

import (
	"apigw/util/encode"
)

type UserType int

const PWD_SALT = "apigw#1.&168"

type userModel struct{}

//user infos from sso login
type User struct {
	Name   string
	Admin  bool
	Detail string
}
