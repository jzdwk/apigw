package base

import (
	"apigw/security"
	"apigw/util/logs"
)

type ApiController struct {
	ParamBuilderController
	SecurityCtx security.Context
}

func (c *ApiController) Prepare() {
	ctx, err := security.GetSecurityContext(c.Ctx.Request)
	if err != nil {
		logs.Error("failed to get security context: %v", err)
		c.AbortInternalServerError("get security context error.")
		return
	}
	c.SecurityCtx = ctx
}
