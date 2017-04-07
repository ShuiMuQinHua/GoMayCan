package models

import (
	"GoMayCan/candivide/helpers"
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := helpers.GetMyConfig("db.host")
	dbport := helpers.GetMyConfig("db.port")
	dbuser := helpers.GetMyConfig("db.user")
	dbpassword := helpers.GetMyConfig("db.password")
	dbname := helpers.GetMyConfig("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone) //url.QueryEscape,s进行转码使之可以安全的用在URL查询里
	}
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(RpotLog), new(Ter), new(Tendor))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
