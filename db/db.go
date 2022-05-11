/**
 * Created by lock
 * Date: 2019-09-22
 * Time: 22:37
 */
package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gochat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var dbMap = map[string]*gorm.DB{}
var syncLock sync.Mutex

func init() {
	initDB("gochat")
}

func initDB(dbName string) {
	var e error
	// if prod env , you should change mysql driver for yourself !!!
	myConf := config.Conf.Common.CommonMySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		myConf.Username,
		myConf.Password,
		myConf.Host,
		myConf.Dbname,
		myConf.Timeout,
	)

	syncLock.Lock()
	_db, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbMap[dbName] = _db
	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(myConf.MaxConnections)     //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(myConf.MaxIdleConnections) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	syncLock.Unlock()
	if e != nil {
		logrus.Error("connect db fail:%s", e.Error())
	}
}

func GetDb(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok {
		return db
	} else {
		return nil
	}
}

type DbGoChat struct {
}

func (*DbGoChat) GetDbName() string {
	return "gochat"
}
