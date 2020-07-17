/*
@Time : 20-3-2
@Author : jzd
@Project: apigw
*/
package ecpmd

//ecpmd info table define
type EcpInfo struct {
	UUID   string `orm:"column(id);pk;type(char);size(36);default('undefine')" json:"id,omitempty"`
	Name   string `orm:"column(name);size(20);null" json:"name"`
	Region string `orm:"column(region);size(40);null" json:"region"`
	Info   string `orm:"column(info);type(text);null" json:"info"`
	IP     string `orm:"column(ip);size(40);null" json:"ip"`
	Port   string `orm:"column(port);size(20);null" json:"port"`
	Status int    `orm:"column(status);null" json:"status"`
}

//ecpmd info for resp
type EcpInfoResp struct {
	/*UUID		string	`json:"id,omitempty"`
	Name	string	`json:"name"`
	Region	string	`json:"region"`
	Info 	string	`json:"info"`
	IP 		string	`json:"ip"`
	Port 	string	`json:"port"`
	Status 	int 	`json:"status"`*/
	*EcpInfo
}
