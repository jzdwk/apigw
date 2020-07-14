/*
@Time : 20-7-9
@Author : jzd
@Project: apigw
*/
package security

import (
	"apigw/auth"
	"apigw/models"
	"fmt"
	"net/http"
)

// ContextValueKey for content value
type ContextValueKey string

const (
	// SecCtxKey is context value key for security context
	SecCtxKey ContextValueKey = "apigw_security_context"
)

//security context do authN/Z
type Context interface {
	// IsAuthenticated returns whether the context has been authenticated or not
	IsAuthenticated() bool
	// GetUsername returns the username of user related to the context
	GetUsername() string
	// IsSysAdmin returns whether the user is system admin
	IsSysAdmin() bool
	// Permit returns whether the user can do action on resource
	Permit(action auth.Action, resource auth.Resource) bool
}

// Default security context
// DefaultSecContext implements security.Context interface based on token
type DefaultSecContext struct {
	user *models.User
}

// NewSecurityContext ...
func NewDefaultSecContext(user *models.User) *DefaultSecContext {
	return &DefaultSecContext{
		user: user,
	}
}

// IsAuthenticated returns true if the user has been authenticated
func (s *DefaultSecContext) IsAuthenticated() bool {
	return s.user != nil
}

// GetUsername returns the username of the authenticated user
// It returns null if the user has not been authenticated
func (s *DefaultSecContext) GetUsername() string {
	if !s.IsAuthenticated() {
		return ""
	}
	return s.user.Name
}

// IsSysAdmin returns whether the authenticated user is system admin
// It returns false if the user has not been authenticated
func (s *DefaultSecContext) IsSysAdmin() bool {
	if !s.IsAuthenticated() {
		return false
	}
	return s.user.Admin
}

// Permit returns whether the user can do action on resource
func (s *DefaultSecContext) Permit(action auth.Action, resource auth.Resource) bool {
	//todo return permission based on rbac
	return true
}

func GetSecurityContext(req *http.Request) (Context, error) {
	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}

	ctx := req.Context().Value(SecCtxKey)
	if ctx == nil {
		return nil, fmt.Errorf("the security context got from request is nil")
	}

	c, ok := ctx.(Context)
	if !ok {
		return nil, fmt.Errorf("the variable got from request is not security context type")
	}
	return c, nil
}
