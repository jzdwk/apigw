package api_args_fr

type ApiArgsfr struct {
	Id          string `orm:"column(id);type(char);size(36);pk" json:"id"`
	ApiFrId     string `orm:"column(api_fr_id);type(char);size(36)" json:"apifrid"`
	Name        string `orm:"column(name);type(char);size(20)" json:"name"`
	Location    int    `orm:"column(location);size(4)" json:"location"`
	Type        int    `orm:"column(type);size(2)" json:"type"`
	Required    int    `orm:"column(required);size(1)" json:"required"`
	Default     string `orm:"column(default);type(char);size(255)" json:"default"`
	Max         string `orm:"column(max);type(char);size(255)" json:"max"`
	Min         string `orm:"column(min);type(char);size(255)" json:"min"`
	Example     string `orm:"column(example);type(char);size(255)" json:"example"`
	Description string `orm:"column(description);type(char);size(255)" json:"description"`
}
