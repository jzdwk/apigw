/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package ecpinfo

import (
	"apigw/controller/base"
	"apigw/manager"
	"apigw/model"
	"apigw/util/logs"
	"encoding/json"
)

type Controller struct {
	base.ApiController
	EcpManager manager.EcpManager
}

// Post ...
// @Title Post
// @Description create ecp test
// @Param	body	body	model.EcpInfo	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router / [post]
func (c *Controller) Post() {
	var v model.EcpInfo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}
	uuid, err := c.EcpManager.Post(&v)
	if err != nil {
		logs.Error("create ecp info error. %v", err)
		c.HandleError(err)
		return
	}
	c.Success(uuid)
}
