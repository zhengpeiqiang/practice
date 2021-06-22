package controller

import (
	"auto_make_code/middleware"
	"github.com/gin-gonic/gin"
)

var DemandToFunc = map[string]string{
	"要一个登录功能":"login",
	"要一个支付功能":"pay",
}

type groupFunc struct {
}

func RegisterGroupFunc(router *gin.RouterGroup) {
	this := &groupFunc{}
	router.POST("/wordtocode", this.wordToCode)
}

/*
	wordToCode 文字转换成代码
*/
func (a *groupFunc) wordToCode(router *gin.Context) {
	middleware.ResponseSuccess(router, gin.H{
		"msg": "ok!",
	})
}
