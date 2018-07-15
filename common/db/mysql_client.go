package db

import (
	"cloud/common/config"
	logger "github.com/alecthomas/log4go"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func GetEngineClient() error {
	if engine == nil {
		var err error
		mysqlConf := config.GetConf().Mysql
		driverUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysqlConf.User, mysqlConf.Password, mysqlConf.Addr, mysqlConf.Database)
		engine, err = xorm.NewEngine("mysql", driverUrl)
		if err != nil {
			logger.Error(fmt.Sprintf("New engine err:%v", err))
			return err
		}
		err = engine.Ping()
		if err != nil {
			logger.Error(fmt.Sprintf("Ping mysql server err:%v", err))
			return err
		}
		//engine.ShowSQL(true)
	}
	return nil
}

func GetEngine() *xorm.Engine {
	if engine == nil {
		GetEngineClient()
	}
	return engine
}
