package dbtools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

/*
* MysqlConnectiPool
* 数据库连接操作库
* 基于gorm封装开发
 */
type MysqlConnectPool struct {
}

var mysqlInstance *MysqlConnectPool
var mysqlOnce sync.Once

var db *gorm.DB
var errDb error

func GetMysqlInstance() *MysqlConnectPool {
	mysqlOnce.Do(func() {
		mysqlInstance = &MysqlConnectPool{}
	})
	return mysqlInstance
}

type DbConf struct {
	Host   string
	Port   string
	User   string
	Pwd    string
	DbName string
}

/*
* @fuc 初始化数据库连接(可在mail()适当位置调用)
 */
func (m *MysqlConnectPool) InitMysqlPool(dbConf *DbConf) (issucc bool) {
	port := dbConf.Port
	db, errDb = gorm.Open("mysql", dbConf.User+":"+dbConf.Pwd+"@tcp("+dbConf.Host+":"+port+")/"+dbConf.DbName+"?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	fmt.Println(errDb)
	if errDb != nil {
		log.Fatal(errDb)
		return false
	}
	//关闭数据库，db会被多个goroutine共享，可以不调用
	// defer db.Close()
	return true
}

/*
* @fuc  对外获取数据库连接对象db
 */
func (m *MysqlConnectPool) GetMysqlPool() *gorm.DB {
	//db.LogMode(true)
	return db
}

func GetMysqlDb() (db *gorm.DB) {
	return GetMysqlInstance().GetMysqlPool()
}
