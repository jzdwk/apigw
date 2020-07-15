package initial

import (
	"apigw/conf"
	"apigw/util/logs"
	"database/sql"
	"fmt"
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
		panic(err)
	}
	db, err := orm.GetDB()
	if err != nil {
		panic(err)
	}
	//ttl := beego.AppConfig.DefaultInt("DBConnTTL", 30)
	ttl := conf.ConfStoreMgr.GetItem(conf.DBConTTLKey).GetInt()
	db.SetConnMaxLifetime(time.Duration(ttl) * time.Second)
	//orm.Debug = beego.AppConfig.DefaultBool("ShowSql", false)
	orm.Debug = conf.ConfStoreMgr.GetItem(conf.DBShowSQL).GetBool()
}

func ensureDatabase() error {

	dbName := conf.ConfStoreMgr.GetItem(conf.DBNameKey).GetString()
	tns := conf.ConfStoreMgr.GetItem(conf.DBTnsKey).GetString()
	usr := conf.ConfStoreMgr.GetItem(conf.DBUserKey).GetString()
	pwd := conf.ConfStoreMgr.GetItem(conf.DBPwdKey).GetString()

	//needInit, err := beego.AppConfig.Bool("InitDBFlag")
	//dbName := beego.AppConfig.String("DBName")
	/*dbURL := fmt.Sprintf("%s:%s@%s/", beego.AppConfig.String("DBUser"),
	beego.AppConfig.String("DBPasswd"), beego.AppConfig.String("DBTns"))*/

	dbURL := fmt.Sprintf("%s:%s@%s/", usr, pwd, tns)
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

	//fmt.Println("Initialize database connection: %s", strings.Replace(dbURL, beego.AppConfig.String("DBPasswd"), "****", 1))
	logs.Info("Initialize database connection: %s", strings.Replace(dbURL, pwd, "****", 1))
	err = orm.RegisterDataBase("default", "mysql", addLocation(fmt.Sprintf("%s%s", dbURL, dbName)))
	if err != nil {
		fmt.Println("register database failed")
		return err
	}

	if true {
		fmt.Println("need init start...  runsyncdb")
		err = orm.RunSyncdb("default", false, true)
		if err != nil {
			return err
		}
		fmt.Println("create tables ")
		for _, insertSql := range InitialData {
			//skip err
			/*	_, err = orm.NewOrm().Raw(insertSql).Exec()
				if err != nil {
					return err
				}*/
			orm.NewOrm().Raw(insertSql).Exec()
		}

	}
	return nil
}

func addLocation(dbURL string) string {
	//return fmt.Sprintf("%s?charset=utf8&loc=%s", dbURL, beego.AppConfig.DefaultString("DBLoc", "Asia%2FShanghai"))
	return fmt.Sprintf("%s?charset=utf8&loc=%s", dbURL, conf.ConfStoreMgr.GetItem(conf.DBLocKey).GetString())
}
