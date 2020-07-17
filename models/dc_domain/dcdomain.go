/*
@Time : 20-7-15
@Author : gzj
@Project: apigw
*/

package dc_domain

type DcDomain struct {
	ID     string `orm:"column(id);pk;type(char);size(36)" json:"id,omitempty"`
	Domain string `orm:"column(domain);type(varchar);size(255)" json:"domain"`
	Status int    `orm:"column(status);size(2);null" json:"status"`
	Count  int    `orm:"column(count)" json:"count"`
	User   string `orm:"column(user);type(varchar);size(20)" json:"user"`
	Name   string `orm:"column(name);type(varchar);size(20)" json:"name"`
	//ApiGroup []*ApiGroup `orm:"reverse(many)"`
}

type Api_bk struct {
	Id        string `orm:"column(id);type(char);size(38) json:"id"`
	Protocol  int    `orm:"column(protocol);type(int);size(38) json:"protocol"`
	Method    int    `orm:"column(method);type(int);size(38) json:"method"`
	Path      string `orm:"column(path);type(char);size(38) json:"path"`
	Auth      int    `orm:"column(auth);type(int);size(38) json:"auth"`
	Api_fr_id string `orm:"column(api_fr_id);type(char);size(38) json:"apifrid"`
}
