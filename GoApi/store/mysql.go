package store

import (
	"log"

	"my_go_api/g"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Orm *xorm.Engine

//连接mysql
func InitMysql() {
	cfg := g.Config

	var err error
	Orm, err = xorm.NewEngine("mysql", cfg.Mysql.Addr)
	if err != nil {
		log.Fatalln("fail to connect mysql", err)
	}
	Orm.SetMaxIdleConns(cfg.Mysql.Idle)
	Orm.SetMaxOpenConns(cfg.Mysql.Max)
}
