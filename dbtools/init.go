package dbtools

import (
	"log"
	"os"
)

func Init(dbConf *DbConf){
	//初始化Mysql连接池
	mysql := GetMysqlInstance().InitMysqlPool(dbConf)
	if !mysql {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}

