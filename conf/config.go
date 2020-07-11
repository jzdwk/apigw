package conf

import (
	"github.com/astaxie/beego"
	"os"
	"strconv"
)

//some const config
const (
	AppAddrKey = iota
	AppPortKey
	//db
	DBNameKey
	DBTnsKey
	DBUserKey
	DBPwdKey
	DBLocKey
	DBConTTLKey
	DBShowSQL
	//log
	LogLevelKey
	LogFileKey
	SentryEnableKey
	SentryDSNKey
	SentryLogLevelKey
)

var (
	//init self-define config
	configMap = map[int]*item{
		//bee http listen
		AppAddrKey: {envKey: "APIGW_HTTP_ADDR", defaultValue: "localhost"},
		//bee http port
		AppPortKey: {envKey: "APIGW_HTTP_PORT", defaultValue: "8090"},
		//bee db
		DBNameKey:   {envKey: "APIGW_DB_NAME", defaultValue: "apigw"},
		DBTnsKey:    {envKey: "APIGW_DB_NAME", defaultValue: "tcp(127.0.0.1:3306)"},
		DBUserKey:   {envKey: "APIGW_DB_USER", defaultValue: "tcp(127.0.0.1:3306)"},
		DBPwdKey:    {envKey: "APIGW_DB_PWD", defaultValue: "123456"},
		DBLocKey:    {envKey: "APIGW_DB_LOC", defaultValue: "Asia%2FShanghai"},
		DBConTTLKey: {envKey: "APIGW_DB_TTL", defaultValue: "30"},
		DBShowSQL:   {envKey: "APIGW_SHOW_SQL", defaultValue: "false"},
		//bee log
		LogLevelKey:       {envKey: "APIGW_LOG_LVL", defaultValue: "6"},
		LogFileKey:        {envKey: "APIGW_FILE", defaultValue: ""},
		SentryEnableKey:   {envKey: "APIGW_LOG_SENTRY", defaultValue: "false"},
		SentryDSNKey:      {envKey: "APIGW_LOG_DSN", defaultValue: ""},
		SentryLogLevelKey: {envKey: "APIGW_LOG_SENTRY_LVL", defaultValue: "4"},
		//todo more self-define config
	}

	ConfStoreMgr StoreMgr
)

//config item const
type item struct {
	//env key, if no env key, use default
	envKey       string
	defaultValue string
}

func (i *item) getValue() string {
	connStr, ex := os.LookupEnv(i.envKey)
	if !ex {
		return i.defaultValue
	}
	return connStr
}
func (i *item) GetString() string {
	return i.getValue()
}
func (i *item) GetInt() int {
	rstStr := i.getValue()
	rst, _ := strconv.Atoi(rstStr)
	return rst
}
func (i *item) GetBool() bool {
	rstStr := i.getValue()
	rst, _ := strconv.ParseBool(rstStr)
	return rst
}

//config store manager
type StoreMgr interface {
	GetItem(int) *item
}

type storeMgr struct {
	config map[int]*item
}

func NewStoreMgr() StoreMgr {
	return &storeMgr{config: configMap}
}

func (s *storeMgr) GetItem(key int) *item {
	return s.config[key]
}

func InitConfig() {
	//beego init config
	beego.BConfig.AppName = "apigw"
	beego.BConfig.Listen = beego.Listen{HTTPAddr: configMap[AppAddrKey].GetString(), HTTPPort: configMap[AppPortKey].GetInt()}
	beego.BConfig.RunMode = "dev"
	beego.BConfig.CopyRequestBody = true

	ConfStoreMgr = NewStoreMgr()

}
