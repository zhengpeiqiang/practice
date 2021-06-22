package controller

import (
	"github.com/gin-gonic/gin"
	"origin/middleware"
	"origin/modules/i18n"
	"origin/router/controller/v1/interfaceStruct"
)

type Controller struct {
}

func RegisterController(router *gin.RouterGroup) {
	this := &Controller{}
	router.POST("/login", this.login)
}

// @Summary 登录
// @Description
// @Tags 微信
// @ID   /login
// @Accept  json
// @Produce  json
// @Param polygon body interfaceStruct.LoginInput true "body"
// @Success 200 {object} middleware.Response{data=interfaceStruct.LoginOutput} "success"
// @Router /login [post]
func (c *Controller) login(router *gin.Context) {
	params := &interfaceStruct.LoginInput{}
	if err := params.BindParams(router); err != nil {
		middleware.ResponseData(router, i18n.PARMAS_ERROR_CODE, nil, err, false)
		return
	}
}
