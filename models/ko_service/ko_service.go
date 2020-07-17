package ko_service

type KoService struct {
	Id        string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	KoSvcUuid string `orm:"column(ko_svc_uuid);type(char);size(36)" json:"kosvcuuid"`
	GroupDcId string `orm:"column(group_dc_id);type(char);size(36)" json:"groupdcid"`
	Url       string `orm:"column(url);type(varchar);size(255)" json:"url"`
}
