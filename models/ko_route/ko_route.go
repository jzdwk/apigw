package ko_route

type KoRoute struct {
	Id          string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	GroupDcId   string `orm:"column(group_dc_id);type(char);size(36)" json:"groupdcid"`
	KoRouteUuid string `orm:"column(ko_route_uuid);type(char);size(36)" json:"korouteuuid"`
	KoServiceId string `orm:"column(ko_service_id);type(char);size(36)" json:"koserviceid"`
	ApiFrId     string `orm:"column(api_fr_id);type(char);size(36)" json:"apifrid"`
	Name        string `orm:"column(name);type(varchar);size(255)" json:"name"`
	Status      int    `orm:"column(status);size(4)" json:"status"`
}
