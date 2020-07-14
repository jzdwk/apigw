package dao

import (
	"apigw/common"
	"apigw/util/logs"
	"errors"
	"github.com/astaxie/beego/orm"
	"sync"
)

const SqlErrorCode int64 = -1

var (
	globalOrm orm.Ormer
	once      sync.Once
)

var UUID = func() string {
	return "uuid-test"
}

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

// 所有实体表的单表查询统一封装
// @Description get entity by id
// @Param	table   table name for query
// @Param	rst		record
// @Param	id		primary key
// @Param	relates relate table name
func GetOne(table interface{}, record interface{}, id int64, relates ...string) (err error) {
	qs := Ormer().QueryTable(table).Filter("Deleted", 0).Filter("UUID", id)
	//级联参数
	if relates != nil {
		for _, relate := range relates {
			qs = qs.RelatedSel(relate)
		}
	}
	//qs = qs.RelatedSel("user")
	if err := qs.One(record); err != nil {
		return err
	}
	return nil
}

// 所有实体表的单表更新统一封装
// @Description update entity
// @Param	table   table name for update
// @Param	id      primary key
// @Param	update  update value
// @Param	cols	columns allowed to update
func Update(table interface{}, id int64, update interface{}, cols []string) (num int64, err error) {
	qs := Ormer().QueryTable(table).Filter("UUID", id).Filter("Deleted", 0)
	if !qs.Exist() {
		return SqlErrorCode, errors.New("record isn't exist")
	}
	if num, err = Ormer().Update(update, cols...); err != nil {
		return SqlErrorCode, err
	}
	return num, nil
}

// 所有实体表的单表增加统一封装
// @Description add entity
// @Param	table   table name for insert
// @Param	record	 record for insert
func Add(table interface{}, record interface{}) (num int64, err error) {
	qs := Ormer().QueryTable(table)
	i, err := qs.PrepareInsert()
	if err != nil {
		return SqlErrorCode, err
	}
	if num, err = i.Insert(record); err != nil {
		return SqlErrorCode, err
	}
	return num, nil
}

// 所有实体表的单表软删除统一封装
// @Description add entity
// @Param	table   table name
// @Param	id    primary key
func SoftDelete(table interface{}, id int64) (num int64, err error) {
	qs := Ormer().QueryTable(table).Filter("UUID", id).Filter("Deleted", 0)
	if !qs.Exist() {
		return SqlErrorCode, errors.New("record isn't exist")
	}
	param := make(map[string]interface{})
	// 删除位
	param["Deleted"] = 1
	if num, err = qs.Update(param); err != nil {
		return num, err
	}
	return num, nil

}

// 所有实体表的单表删除统一封装,慎用,考虑特殊情况的数据直接删除
// @Description add entity
// @Param	table   table name
// @Param	id    primary key
func Delete(table interface{}, id int64) (num int64, err error) {
	qs := Ormer().QueryTable(table)
	if num, err = qs.Filter("Deleted", 0).Filter("UUID", id).Delete(); err != nil {
		return SqlErrorCode, err
	}
	return num, nil

}

// 所有实体的唯一约束鉴别表
// @Description exist key check
func IsExist(table interface{}, param map[string]interface{}) (rst bool) {
	qs := Ormer().QueryTable(table).Filter("Deleted", 0)
	// query k=v
	for k, v := range param {
		if v != nil && v != "" {
			qs = qs.Filter(k, v)
		}
	}
	var nums int64 = 0
	nums, _ = qs.Count()
	return nums > 0
}

func GetTotal(queryTable interface{}, q *common.QueryParam) (int64, error) {
	qs := Ormer().QueryTable(queryTable)
	qs = BuildFilter(qs, q.Query)
	if len(q.GroupBy) != 0 {
		qs = qs.GroupBy(q.GroupBy...)
	}
	return qs.Count()
}

func GetAll(queryTable interface{}, list interface{}, q *common.QueryParam) error {
	qs := Ormer().QueryTable(queryTable)
	qs = BuildFilter(qs, q.Query)
	if len(q.GroupBy) != 0 {
		qs = qs.GroupBy(q.GroupBy...)
	}
	if q.Order != "" {
		qs = qs.OrderBy(q.Order)
	}
	if _, err := qs.Limit(q.Limit(), q.Offset()).All(list); err != nil {
		return err
	}
	return nil
}

func BuildFilter(qs orm.QuerySeter, query map[string]interface{}) orm.QuerySeter {
	// query k=v
	for k, v := range query {
		if v != nil && v != "" {
			qs = qs.Filter(k, v)
		}
	}
	return qs
}
