package cap_my

import "time"

type CapMy struct {
	Id            string    `orm:"column(id);pk;type(char);size(36)" json:"id"`
	ProductId     string    `orm:"column(product_id);type(char);size(36)" json:"productid"`
	User          string    `orm:"column(user);type(varchar);size(36)" json:"user"`
	ProductInfo   string    `orm:"column(product_info);type(varchar);size(255)" json:"productinfo"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime)" json:"createtime"`
	ProductDetail string    `orm:"column(product_detail);type(varchar);size(255)" json:"productdetail"`
}
