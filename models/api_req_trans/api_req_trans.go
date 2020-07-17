package api_req_trans

type ApiReqTrans struct {
	Id      string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	ApiFrId string `orm:"column(api_fr_id);type(char);size(36)" json:"apifrid"`
	Config  string `orm:"column(config);type(text)" json:"config"`
	Enable  int    `orm:"column(enable);size(2)" json:"enable"`
}
