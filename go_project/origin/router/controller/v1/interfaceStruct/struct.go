package interfaceStruct

import (
	"github.com/gin-gonic/gin"
	"origin/modules/ginFunc"
)

type LoginInput struct {
	ID      int    `json:"id" example:"1"`       // 用户id
	UserId  int    `json:"user_id" example:"1"`  // 用户userId
	UnionId string `json:"union_id" example:"1"` // 用户唯一id
	Openid  string `json:"openid" example:"1"`   // 用户应用id
}

type LoginOutput struct {
	ID      int    `json:"id" example:"1"`       // 用户id
	UserId  int    `json:"user_id" example:"1"`  // 用户userId
	UnionId string `json:"union_id" example:"1"` // 用户唯一id
	Openid  string `json:"openid" example:"1"`   // 用户应用id
	Name    string `json:"name" example:"1"`     // 用户昵称
	Age     int    `json:"age" example:"1"`      // 用户年龄
	Job     string `json:"job" example:"1"`      // 用户职业
	Remark  string `json:"remark" example:"1"`   // 用户备注
}

func (params *LoginInput) BindParams(c *gin.Context) error {
	return ginFunc.BindParams(c, params)
}
