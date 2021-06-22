package router

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"use_granphviz/middleware"
	"use_granphviz/pressure"
	"use_granphviz/public"
	"use_granphviz/task"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

var goNum = 0

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	defer func() {
		if x := recover(); x != nil {
			fmt.Print(x)
		}
	}()
	//[GIN-debug] GET    /getfunc                  --> rock/customer/test/router.InitRouter.func1 (3 handlers)
	//[GIN-debug] POST   /postfunc                 --> rock/customer/test/router.InitRouter.func2 (3 handlers)
	//[GIN-debug] POST   /groupfunc/funcone        --> rock/customer/test/router.(*groupFunc).funcOne-fm (6 handlers)

	router := gin.Default()
	pprof.Register(router)
	router.Use(Cors())
	//_ = http.ListenAndServeTLS(":443", "cert/www.domain.cn.crt", "cert/www.domain.cn.key", router)
	router.GET("/getfunc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	router.GET("/timeout", func(c *gin.Context) {
		pressure.TimeOut(
			goNum,
			0, "oR52qv4HG8K1MyFdGhXaH6Qjqz7A", 2,
			"PQ", "https://thirdwx.qlogo.cn/mmopen/vi_32/j7SnZnaxybGgsYNl4YNIpJibAH6JZdxImoUiaAWz2paK33BGYu9EEMsBvBxuZ0cvZop7YCZfhpamzN6ZCURfuwow/132", "127317",
			"127317")
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	router.GET("/pressure", func(c *gin.Context) {
		// 一共并发多少次
		runNum := gconv.Int(c.DefaultQuery("run_num", "10"))
		// 一次并发执行多少次程序
		concurrentNum := gconv.Int(c.DefaultQuery("concurrent_num", "10"))
		//
		task.TaskList.Set(public.RandomString(10), task.TaskListContent{
			RunNum:        runNum,
			ConcurrentNum: concurrentNum,
			Type:          "pressure",
		})
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	router.GET("/channel", func(c *gin.Context) {
		// 一共并发多少次
		runNum := gconv.Int(c.DefaultQuery("run_num", "10"))
		// 一次并发执行多少次程序
		concurrentNum := gconv.Int(c.DefaultQuery("concurrent_num", "10"))
		//
		channelType := gconv.Int(c.DefaultQuery("channel_type", "1"))
		//
		task.TaskList.Set(public.RandomString(10), task.TaskListContent{
			RunNum:        runNum,
			ConcurrentNum: concurrentNum,
			Type:          "channel",
			Maps: map[string]interface{}{
				"channel_type": channelType,
			},
		})
		c.JSON(200, gin.H{
			"message": "ok!!!",
		})
	})

	router.POST("/postfunc", func(c *gin.Context) {
		defer func() {
			if x := recover(); x != nil {
				fmt.Print(x)
			}
		}()
		postFunc(c)
	})

	//非登陆接口
	apiNormalGroup := router.Group("/groupfunc")
	apiNormalGroup.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		RegisterGroupFunc(apiNormalGroup)
	}

	return router
}

type groupFunc struct {
}

func RegisterGroupFunc(router *gin.RouterGroup) {
	this := &groupFunc{}
	router.POST("/funcone", this.funcOne)
}

func (a *groupFunc) funcOne(router *gin.Context) {
	middleware.ResponseSuccess(router, gin.H{
		"msg": "ok!",
	})
	panic("报错了！！！")
}

func postFunc(router *gin.Context) {
	middleware.ResponseSuccess(router, gin.H{
		"msg": "ok!",
	})
	panic("报错了！！！")
}
