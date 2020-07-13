package dao

import (
	"apigw/util/logs"
	"github.com/astaxie/beego/orm"
	"sync"
)

var (
	globalOrm orm.Ormer
	once      sync.Once
)

// singleton init ormer ,only use for normal db operation
// if you begin transaction，please use WithTransaction
func Ormer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

// WithTransaction helper for transaction
func WithTransaction(handler func(o orm.Ormer) error) error {
	o := orm.NewOrm()

	if err := o.Begin(); err != nil {
		logs.Error("begin transaction failed: %v", err)
		return err
	}

	if err := handler(o); err != nil {
		if e := o.Rollback(); e != nil {
			logs.Error("rollback transaction failed: %v", e)
			return e
		}

		return err
	}

	if err := o.Commit(); err != nil {
		logs.Error("commit transaction failed: %v", err)
		return err
	}

	return nil
}
