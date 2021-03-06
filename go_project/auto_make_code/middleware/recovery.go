package middleware

import (
	"errors"
	"fmt"
	"github.com/e421083458/gin_scaffold/public"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//先做一下日志记录
				fmt.Println(string(debug.Stack()))
				public.ComLogNotice(c, "_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				ResponseError(c, 500, errors.New(fmt.Sprint(err)))
				return
			}
		}()
		c.Next()
	}
}