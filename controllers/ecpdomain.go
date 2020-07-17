/*
@Time : 20-3-2
@Author : jzd
@Project: apigw
*/
package controllers

import (
	"apigw/controllers/base"
	"apigw/manager"
)

type EcpDomainController struct {
	base.ApiController
	EcpDomainMg manager.DomainManager
}

// Post ...
// @Title Post
// @Description create test
// @Param	body	body	ecp_domain.EcpDomain	true	"body for test content"
// @Success 201 {string} success message
// @Failure 403 body is empty
// @router / [post]
func (c *EcpDomainController) Post() {
	//todo
	c.Success("create test success")
}
