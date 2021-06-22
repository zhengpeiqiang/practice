package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"origin/modules/i18n"
	"origin/modules/redis"
	"strings"
	"time"
)

type Response struct {
	ErrorCode int         `json:"errno"`
	ErrorMsg  string      `json:"errmsg"`
	Data      interface{} `json:"data"`
	TraceId   interface{} `json:"trace_id"`
	Stack     interface{} `json:"stack"`
}

func ResponseData(c *gin.Context, code int, data interface{}, err error, isRedis bool, redisKey ...string) {
	if err == nil {
		err = errors.New("")
	}
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	if code == 200 {
		trace, _ := c.Get("trace")
		traceContext, _ := trace.(*lib.TraceContext)
		traceId := ""
		if traceContext != nil {
			traceId = traceContext.TraceId
		}

		resp := &Response{ErrorCode: code, ErrorMsg: i18n.MsgFlags[code], Data: data, TraceId: traceId}
		c.JSON(200, resp)
		response, _ := json.Marshal(resp)
		c.Set("response", string(response))
		//存入缓存
		if isRedis {
			jsonContent, err := json.Marshal(data)
			if err != nil {
				ResponseData(c, i18n.INTERNAL_ERROR_CODE, nil, err, false)
				return
			}
			redis.ClientRedis.Set(redisKey[0], jsonContent, 60*time.Second)
		}
	} else {
		stack := ""
		if c.Query("is_debug") == "1" || lib.GetConfEnv() == "dev" {
			stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
		}

		data := make([]string, 0)
		resp := &Response{ErrorCode: code, ErrorMsg: i18n.MsgFlags[code], Data: data, TraceId: traceId, Stack: stack}

		c.JSON(200, resp)
		response, _ := json.Marshal(resp)
		c.Set("response", string(response))
		if err == nil {
			err = errors.New(i18n.MsgFlags[i18n.INTERNAL_ERROR_CODE])
		}
		_ = c.AbortWithError(200, err)
	}
}