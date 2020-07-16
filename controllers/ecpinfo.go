/*
@Time : 20-3-2
@Author : jzd
@Project: apigw
*/
package controllers

import (
	"apigw/controllers/base"
	"apigw/manager"
	"apigw/models/ecpinfo"
	"apigw/util/logs"
	"encoding/json"
)

type EcpInfoController struct {
	base.ApiController
	EcpManager manager.EcpManager
}

// Post ...
// @Title Post
// @Description create test
// @Param	body	body	models.EcpInfo	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router / [post]
func (c *EcpInfoController) Post() {
	var v ecpinfo.EcpInfo
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		logs.Error("get body error. %v", err)
		c.AbortBadRequestFormat("test")
		return
	}
	if _, err := c.EcpManager.Post(&v); err != nil {
		logs.Error("create test error. %v", err)
		c.HandleError(err)
		return
	}
	c.Success("create test success")
}

// GetOne EcpInfo
// @Title Get One
// @Description get method test
// @Param	id		path 	string	 true	"test id"
// @Success 200 {object} models.EcpInfo
// @Failure 403
// @router /:id  [get]
func (c *EcpInfoController) Get() {
	uuid := c.GetStrIDFromURL()
	rst, err := c.EcpManager.Get(uuid)
	if err != nil {
		logs.Error("get ecp info error. %v", err)
	}
	c.Success(rst)
}
