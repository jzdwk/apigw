/*
@Time : 20-7-19
@Author : jzd
@Project: apigw
*/
package dao

type RateLimitDao interface {
	// rl template
	CreateRateLimit(rateLimitInfo *interface{}) (uuid string, err error)
	DeleteRateLimit(rateLimitId string) error
	UpdateRateLimit(rateLimitId string, rateLimitInfo *interface{}) error
	GetRateLimit(rateLimitId string) (rateLimitInfo *interface{}, err error)
	//paged list
	ListRateLimit(rateLimitQuery *interface{}) (rateLimitInfo *interface{}, err error)
}

func NewDefaultRateLimitDao() RateLimitDao {
	return &defaultRateLimitDao{}
}

type defaultRateLimitDao struct{}

func (r *defaultRateLimitDao) CreateRateLimit(rateLimitInfo *interface{}) (uuid string, err error) {
	panic("implement me")
}

func (r *defaultRateLimitDao) DeleteRateLimit(rateLimitId string) error {
	panic("implement me")
}

func (r *defaultRateLimitDao) UpdateRateLimit(rateLimitId string, rateLimitInfo *interface{}) error {
	panic("implement me")
}

func (r *defaultRateLimitDao) GetRateLimit(rateLimitId string) (rateLimitInfo *interface{}, err error) {
	panic("implement me")
}

func (r *defaultRateLimitDao) ListRateLimit(rateLimitQuery *interface{}) (rateLimitInfo *interface{}, err error) {
	panic("implement me")
}
