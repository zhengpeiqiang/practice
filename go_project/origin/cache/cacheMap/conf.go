package cacheMap

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/gogf/gf/util/gconv"
)

type Http struct {
	Protocol       string
	Host           string
	Port           string
	Http           string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
	AllowIp        []string
}

var ConfHttp *Http

func initConfHttp() {
	conf := Http{}
	conf.Protocol = lib.GetStringConf("base.http.protocol")
	conf.Host = lib.GetStringConf("base.http.host")
	conf.Port = lib.GetStringConf("base.http.port")
	conf.Http = lib.GetStringConf("base.http.http")
	conf.ReadTimeout = lib.GetIntConf("base.http.read_timeout")
	conf.WriteTimeout = lib.GetIntConf("base.http.write_timeout")
	conf.MaxHeaderBytes = lib.GetIntConf("base.http.max_header_bytes")
	conf.AllowIp = lib.GetStringSliceConf("base.http.allow_ip")
	ConfHttp = &conf
}

type File struct {
	ViewPath  string
	MusicFile string
}

var ConfFile *File

func initConfFile() {
	conf := File{}
	conf.ViewPath = lib.GetStringConf("base.file.view_path")
	conf.MusicFile = lib.GetStringConf("base.file.music_file")
	ConfFile = &conf
}

type DB struct {
	DriverName      string
	DataSourceName  string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifeTime int
}

var ConfDB *map[string]DB

func initConfDB() {
	conf := make(map[string]DB, 0)
	list := lib.GetStringMapConf("base.mysql")
	for k, v := range list {
		if value, ok := v.(map[string]interface{}); ok {
			driverName := ""
			dataSourceName := ""
			maxOpenConn := 0
			maxIdleConn := 0
			maxConnLifeTime := 0
			for kk, vv := range value {
				if kk == "driver_name" {
					driverName = vv.(string)
				}
				if kk == "data_source_name" {
					dataSourceName = vv.(string)
				}
				if kk == "max_open_conn" {
					maxOpenConn = gconv.Int(vv)
				}
				if kk == "max_idle_conn" {
					maxIdleConn = gconv.Int(vv)
				}
				if kk == "max_conn_life_time" {
					maxConnLifeTime = gconv.Int(vv)
				}
			}
			conf[k] = DB{
				DriverName:      driverName,
				DataSourceName:  dataSourceName,
				MaxOpenConn:     maxOpenConn,
				MaxIdleConn:     maxIdleConn,
				MaxConnLifeTime: maxConnLifeTime,
			}
		}
	}
	ConfDB = &conf
}

type Redis struct {
	ProxyList []string
	Password  string
	DB        int
	MaxActive int
	MaxIdle   int
}

var ConfRedis *Redis

func initConfRedis() {
	conf := Redis{}
	conf.ProxyList = lib.GetStringSliceConf("base.redis.proxy_list")
	conf.Password = lib.GetStringConf("base.redis.password")
	conf.DB = lib.GetIntConf("base.redis.db")
	conf.MaxActive = lib.GetIntConf("base.redis.max_active")
	conf.MaxIdle = lib.GetIntConf("base.redis.max_idle")
	ConfRedis = &conf
}

type Swagger struct {
	Title    string
	Desc     string
	Host     string
	BasePath string
}

var SwaggerConf Swagger

func initConfSwagger() {
	conf := &SwaggerConf
	conf.Title = lib.GetStringConf("base.swagger.title")
	conf.Desc = lib.GetStringConf("base.swagger.desc")
	conf.Host = lib.GetStringConf("base.swagger.host")
	conf.BasePath = lib.GetStringConf("base.swagger.base_path")
}
