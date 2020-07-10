package initial

import (
	"apigw/conf"
	"apigw/util/logs"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

const DbDriverName = "mysql"

func InitDb() {
	orm.RegisterDriver(DbDriverName, orm.DRMySQL)
	// ensure database exist
	err := ensureDatabase()
	if err != nil {
		logs.Error("database init err: %v", err)
		panic(err)
	}
	db, err := orm.GetDB()
	if err != nil {
		logs.Error("database init err: %v", err)
		panic(err)
	}
	ttl := beego.AppConfig.DefaultInt("DBConnTTL", 30)

	db.SetConnMaxLifetime(time.Duration(ttl) * time.Second)

	orm.Debug = beego.AppConfig.DefaultBool("ShowSql", false)
}

func ensureDatabase() error {
	dbName := conf.ConfStoreMgr.GetItem(conf.DBNameKey).GetString()
	tns := conf.ConfStoreMgr.GetItem(conf.DBTnsKey).GetString()
	usr := conf.ConfStoreMgr.GetItem(conf.DBUserKey).GetString()
	pwd := conf.ConfStoreMgr.GetItem(conf.DBPwdKey).GetString()
	dbURL := fmt.Sprintf("%s:%s@%s/", usr,
		pwd, tns)
	db, err := sql.Open(DbDriverName, fmt.Sprintf("%s%s", dbURL, dbName))
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		switch e := err.(type) {
		case *mysql.MySQLError:
			// MySQL error unkonw database;
			// refer https://dev.mysql.com/doc/refman/5.6/en/error-messages-server.html
			if e.Number == 1049 {
				//needInit = true
				dbForCreateDatabase, err := sql.Open(DbDriverName, addLocation(dbURL))
				if err != nil {
					return err
				}
				defer dbForCreateDatabase.Close()
				_, err = dbForCreateDatabase.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8 COLLATE utf8_general_ci;", dbName))
				if err != nil {
					return err
				}

			} else {
				return err
			}
		default:
			return err
		}
	}

	logs.Info("Initialize database connection: %s", strings.Replace(dbURL, pwd, "****", 1))
	err = orm.RegisterDataBase("default", "mysql", addLocation(fmt.Sprintf("%s%s", dbURL, dbName)))
	if err != nil {
		logs.Error("register database failed")
		return err
	}
	return nil
}

func addLocation(dbURL string) string {
	return fmt.Sprintf("%s?charset=utf8&loc=%s", dbURL, conf.ConfStoreMgr.GetItem(conf.DBLocKey).GetString())
}
