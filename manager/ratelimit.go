/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package manager

import (
	"apigw/dao"
	"apigw/service"
)

type RateLimitManager interface {
	// rl template
	CreateRateLimit(rateLimitInfo *interface{}) (uuid string, err error)
	DeleteRateLimit(rateLimitId string) error
	UpdateRateLimit(rateLimitId string, rateLimitInfo *interface{}) error
	GetRateLimit(rateLimitId string) (rateLimitInfo *interface{}, err error)
	//paged list
	ListRateLimit(rateLimitQuery *interface{}) (rateLimitInfo *interface{}, err error)
	//bind api
	BindApi(rateLimitId, apiId string) error
	UnbindApi(rateLimitId, apiId string) error
}

type defaultRateLimtMg struct {
	//kong service
	koRateLimit service.RateLimit
	//dao
	rateLimitDao dao.RateLimitDao
}

func (r *defaultRateLimtMg) CreateRateLimit(rateLimitInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (r *defaultRateLimtMg) DeleteRateLimit(rateLimitId string) error {
	panic("implement me")
}

func (r *defaultRateLimtMg) UpdateRateLimit(rateLimitId string, rateLimitInfo *interface{}) error {
	panic("implement me")
}

func (r *defaultRateLimtMg) GetRateLimit(rateLimitId string) (rateLimitInfo *interface{}, err error) {
	panic("implement me")
}

func (r *defaultRateLimtMg) ListRateLimit(rateLimitQuery *interface{}) (rateLimitInfo *interface{}, err error) {
	panic("implement me")
}

func (r *defaultRateLimtMg) BindApi(rateLimitId, apiId string) error {
	panic("implement me")
}

func (r *defaultRateLimtMg) UnbindApi(rateLimitId, apiId string) error {
	panic("implement me")
}
