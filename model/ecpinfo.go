/*
@Time : 20-7-13
@Author : jzd
@Project: apigw
*/
package model

type EcpInfo struct {
	ID     string `orm:"pk;column(id)" json:"id"`
	Name   string `orm:"column(name)" json:"name"`
	Region string `orm:"column(region)" json:"region"`
	Ip     string `orm:"column(ip)" json:"ip"`
	Port   string `orm:"column(port)" json:"port"`
	Info   string `orm:"column(info)" json:"info"`
}
