package appmd

type AppInfo struct {
	Id           string `orm:"column(id);pk;type(char);size(36)" json:"id"`
	User         string `orm:"column(user);type(varchar);size(20)" json:"user"`
	Name         string `orm:"column(name);type(varchar);size(20)" json:"name"`
	ConsumerName string `orm:"column(consumer_name);type(varchar);size(50)" json:"consumername"`
	AuthKey      string `orm:"column(auth_key);type(varchar);size(255)" json:"authkey"`
	HmacSecret   string `orm:"column(hmac_secret);type(varchar);size(255)" json:"hmacsecret"`
	HmacUsername string `orm:"column(hmac_username);type(varchar);size(255)" json:"hmacusername"`
}
