/*
@Time : 20-7-9
@Author : jzd
@Project: apigw
*/
package filter

import (
	"apigw/auth"
	"apigw/security"
	"apigw/util/logs"
	beectx "github.com/astaxie/beego/context"
)

// configCtxModifier populates to the configuration values to context, which are to be read by subsequent
// filters.
type reqTokenCtxModifier struct {
}

//check token, then add to req context
func (c *reqTokenCtxModifier) Modify(ctx *beectx.Context) bool {
	//base auth token check
	authString := ctx.Input.Header("Authorization")
	if authString == "" {
		logs.Error("empty token from http header.")
		return false
	}
	login, usr := auth.SSOLogin(authString)
	if !login {
		logs.Error("token authenticates failed.")
		return false
	}
	sc := security.NewDefaultSecContext(usr)
	addToReqContext(ctx.Request, SecCtxKey, sc)
	return true
}
