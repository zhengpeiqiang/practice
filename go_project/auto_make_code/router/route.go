package router

import (
	"auto_make_code/controller"
	"auto_make_code/middleware"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
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

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	defer func() {
		if x := recover(); x != nil {
			fmt.Print(x)
		}
	}()

	router := gin.Default()
	pprof.Register(router)
	router.Use(Cors())
	//_ = http.ListenAndServeTLS(":443", "cert/www.domain.cn.crt", "cert/www.domain.cn.key", router)

	//非登陆接口
	apiNormalGroup := router.Group("/func")
	apiNormalGroup.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.RegisterGroupFunc(apiNormalGroup)
	}

	return router
}
