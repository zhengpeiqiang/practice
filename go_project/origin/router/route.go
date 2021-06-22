package router

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io/ioutil"
	"net/http"
	"origin/cache/cacheMap"
	"origin/docs"
	"origin/middleware"
	"origin/router/controller/v1"
	"time"
)

const (
	ErJiYuMing      = "/origin"
	ErJiYuMingDebug = "origin-debug"
	MusicUrlPath    = "/origin/music"
)

func Router() *gin.Engine {

	// programatically set swagger info
	docs.SwaggerInfo.Title = cacheMap.SwaggerConf.Title
	docs.SwaggerInfo.Description = cacheMap.SwaggerConf.Desc
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = cacheMap.SwaggerConf.Host
	docs.SwaggerInfo.BasePath = cacheMap.SwaggerConf.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	pprof.Register(router, ErJiYuMingDebug)

	router.GET("/customer/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		middleware.CrossDomainMiddleware(),
		middleware.TranslationMiddleware(),
		middleware.RecoveryMiddleware(),
		middleware.TimeoutMiddleware(10*time.Millisecond),
	)

	InitRouter(router)

	return router
}

func InitRouter(e *gin.Engine) {

	group := e.Group("/customer")
	group.Use()
	{
		controller.RegisterController(group)
	}

	e.GET(fmt.Sprintf("%v/welcome", ErJiYuMing), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok!!!",
			"data": map[string]interface{}{
				"http.Http":           cacheMap.ConfHttp.Http,
				"http.ReadTimeout":    cacheMap.ConfHttp.ReadTimeout,
				"http.WriteTimeout":   cacheMap.ConfHttp.WriteTimeout,
				"http.MaxHeaderBytes": cacheMap.ConfHttp.MaxHeaderBytes,
				"http.AllowIp":        cacheMap.ConfHttp.AllowIp,
			},
		})
	})

	e.Static(fmt.Sprintf("%v/music", ErJiYuMing), "file/music")

	e.LoadHTMLGlob(cacheMap.ConfFile.ViewPath)

	e.GET(fmt.Sprintf("%v/musiclist", ErJiYuMing), func(c *gin.Context) {

		fileMusicBasePath := cacheMap.ConfFile.MusicFile
		httpBaseUrl := cacheMap.ConfHttp.Http + MusicUrlPath
		fileDir, err := ioutil.ReadDir(fileMusicBasePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		musicVal := make([]map[string]string, 0)
		for _, v := range fileDir {
			musicVal = append(musicVal, map[string]string{
				"key":   v.Name(),
				"value": httpBaseUrl + "/" + v.Name(),
			})
		}

		c.HTML(http.StatusOK, "musiclist.html", gin.H{
			"title":    "文件列表",
			"musicVal": musicVal,
		})
	})
}
