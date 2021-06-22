package middleware

import (
	"github.com/gin-gonic/gin"
	"origin/modules/ginFunc"
	"origin/modules/i18n/i18nInit"
)

//跨域中间件
func CrossDomainMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginFunc.CrossDomain(c)
		c.Next()
	}
}

//翻译中间件；用以确定返回内容语言
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		i18nInit.LangCheck(c)
		c.Next()
	}
}

//恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//记录日志
			}
		}()
		c.Next()
	}
}
