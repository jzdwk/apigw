package filter

import (
	"context"
	beectx "github.com/astaxie/beego/context"
	"net/http"
)

// ContextValueKey for content value
type ContextValueKey string

const (
	// SecCtxKey is context value key for security context
	SecCtxKey ContextValueKey = "apigw_security_context"
)

var reqCtxModifiers []ReqCtxModifier

func Init() {
	// reqContext adds info to context
	reqCtxModifiers = []ReqCtxModifier{
		&baseCtxModifier{},
		&reqTokenCtxModifier{},
	}
}

func SecurityFilter(ctx *beectx.Context) {
	if ctx == nil {
		return
	}
	req := ctx.Request
	if req == nil {
		return
	}
	//do filter
	for _, modifier := range reqCtxModifiers {
		if modifier.Modify(ctx) {
			break
		}
	}
}

func addToReqContext(req *http.Request, key, value interface{}) {
	*req = *(req.WithContext(context.WithValue(req.Context(), key, value)))
}

// ReqCtxModifier modifies the context of request
type ReqCtxModifier interface {
	Modify(*beectx.Context) bool
}

// configCtxModifier populates to the configuration values to context
type baseCtxModifier struct {
}

func (c *baseCtxModifier) Modify(ctx *beectx.Context) bool {
	//todo add some config info
	return false
}
