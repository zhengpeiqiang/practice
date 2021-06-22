package test

import (
	"github.com/gin-gonic/gin"
	"origin/middleware"
	"origin/modules/i18n"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.Use(
		middleware.CrossDomainMiddleware(),
		middleware.TranslationMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	InitRouter(router)

	return router
}

func InitRouter(e *gin.Engine) {
	e.GET("testlang", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": i18n.Success1,
		})
	})
}
