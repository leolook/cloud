package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var engine *xorm.Engine

func main() {
	var err error
	driverUrl := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "icsoc", "127.0.0.1:3306", "video")
	engine, err = xorm.NewEngine("mysql", driverUrl)
	if err != nil {
		log.Println("connect err:", err)
		return
	}
	log.Println("engine", engine)
	err = engine.Ping()
	if err != nil {
		log.Println("ping err:", err)
	}
}
