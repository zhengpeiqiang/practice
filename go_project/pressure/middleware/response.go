package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"strings"
)

type ResponseCode int

//1000以下为通用码，1000以上为用户自定义码
const (
	SUCCESS_CODE ResponseCode = iota

	INTERNAL_ERROR_CODE        = 4001 //服务器内部错误
	PARMAS_ERROR_CODE          = 4002 //参数错误
	USER_NOT_EXIST             = 4003 //用户不存在
	PASSWORD_ERROR_CODE        = 4004 //密码错误
	GENERATE_TOKEN_ERROR_CODE  = 4005 //生成token失败
	USERNAME_OR_PWD_ERROR_CODE = 4006 //用户名或者密码错误
	DATA_EXIST                 = 4007 //数据已经存在,请勿重复创建

)

var MsgFlags = map[ResponseCode]string{
	SUCCESS_CODE:               "success",
	INTERNAL_ERROR_CODE:        "服务器内部错误",
	PARMAS_ERROR_CODE:          "参数错误",
	USER_NOT_EXIST:             "用户不存在",
	PASSWORD_ERROR_CODE:        "密码错误",
	GENERATE_TOKEN_ERROR_CODE:  "生成token失败",
	USERNAME_OR_PWD_ERROR_CODE: "用户名或者密码错误",
	DATA_EXIST:                 "数据已经存在,请勿重复创建",
}

type Response struct {
	ErrorCode ResponseCode `json:"errno"`
	ErrorMsg  string       `json:"errmsg"`
	Data      interface{}  `json:"data"`
	TraceId   interface{}  `json:"trace_id"`
	Stack     interface{}  `json:"stack"`
}

func getErrorMsg(code ResponseCode, err error) string {

	errorMsg, ok := MsgFlags[code]
	if ok {
		if err != nil {
			return errorMsg + " " + err.Error()
		}

		return errorMsg
	}
	return ""
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	if err == nil {
		err = errors.New("")
	}
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || lib.GetConfEnv() == "dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	data := make([]string, 0)
	resp := &Response{ErrorCode: code, ErrorMsg: getErrorMsg(code, err), Data: data, TraceId: traceId, Stack: stack}

	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	if err == nil {
		err = errors.New("服务器内部错误")
	}
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{ErrorCode: SUCCESS_CODE, ErrorMsg: getErrorMsg(SUCCESS_CODE, nil), Data: data, TraceId: traceId}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
