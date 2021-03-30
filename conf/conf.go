package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

type Config struct {
	Database map[string]database
}

type database struct {
	Host string
	Port string // 对应toml配置文件中的Cats数组
	User string
	Pwd  string
}

func GetConf(path string,conf *Config) error {
	// 通过toml.DecodeFile将toml配置文件的内容，解析到struct对象
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " 格式转换出错:"+err.Error())
		return err
	}
	fmt.Println(conf)
	return nil
}