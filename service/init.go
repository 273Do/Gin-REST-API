package service

import (
	"errors"
	"fmt"
	"gin-rest-api/model"
	"log"

	"github.com/go-xorm/xorm"
)

var DbEngine *xorm.Engine

// データベースへの接続とテーブルの初期化
func init() {
	driverName := "mysql"
	DsName := "root:root@(192.168.99.100:3306)/gin?charset=utf8"
	err := errors.New("")
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	DbEngine.Sync2(new(model.Book))
	fmt.Println("init data base ok")
}
