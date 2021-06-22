package redis

import (
	"fmt"
	goRedis "github.com/go-redis/redis"
	"origin/cache/cacheMap"
	"strings"
)

var ClientRedis *goRedis.Client

func InitRedis() {
	proxyList := cacheMap.ConfRedis.ProxyList
	if len(proxyList) == 0 {
		return
	}
	proxy := strings.Split(proxyList[0], ":")
	options := goRedis.Options{
		Network:            "tcp",
		Addr:               fmt.Sprintf("%s:%s", proxy[0], proxy[1]),
		Dialer:             nil,
		OnConnect:          nil,
		Password:           cacheMap.ConfRedis.Password,
		DB:                 cacheMap.ConfRedis.DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	// 新建一个client
	ClientRedis = goRedis.NewClient(&options)
}
