package main

import (
	"modeltools/conf"
	"modeltools/dbtools"
)

func main() {
	// 首先初始化struct对象，用于保存配置文件信息
	var confS conf.Config
	// 通过toml.DecodeFile将toml配置文件的内容，解析到struct对象
	err := conf.GetConf("./conf/mysql_map.toml",&confS)
	if err != nil {
		return
	}
	for _,v := range confS.Database {
		var MasterDbConfig = &dbtools.DbConf{
			Host:   v.Host,
			Port:   v.Port,
			User:   v.User,
			Pwd:    v.Pwd,
			DbName: "wkb_v2",
		}
		dbtools.Init(MasterDbConfig)//初始化数据库
		dbtools.Generate("./models/", "wkb_v2", "wkb_im_robot",) //生成指定表信息，可变参数可传入多个表名
	}

}
