package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gtimer"
	"origin/modules/i18n"
	"origin/modules/redis"
	"net/http"
	"time"
)

// timeout middleware wraps the request context with a timeout
func TimeoutMiddleware(timeout time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {

		// wrap the request context with a timeout
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)

		defer func() {
			// check if context timeout was reached
			if ctx.Err() == context.DeadlineExceeded {
				// write response and abort the request
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				ResponseData(c, i18n.TIME_OUT, nil, errors.New("接口超时！！"), false)
				c.Abort()
			}

			//cancel to cleaurces after finished
			cancel()
		}()

		// replace request with context wrapped request
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func TimedHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// get the underlying request context
		ctx := c.Request.Context()
		// create the response data type to use as a channel type
		type responseData struct {
			status int
			body   map[string]interface{}
		}
		// create a done channel to tell the request it's done
		doneChan := make(chan responseData)
		// here you put the actual work needed for the request
		// and then send the doneChan with the status and body
		// to finish the request by writing the response
		go func() {
			doneChan <- responseData{
				status: 200,
				body:   gin.H{"hello": "world"},
			}
		}()
		// non-blocking select on two channels see if the request
		// times out or finishes
		select {
		// if the context is done it timed out or was cancelled
		// so don't return anything
		case <-ctx.Done():
			return
			// if the request finished then finish the request by
			// writing the response
		case res := <-doneChan:
			c.JSON(res.status, res.body)
		}
	}
}

// timeout middleware wraps the request context with a timeout
func Timeout(c *gin.Context, duration time.Duration, traceId string, cacheOverdue time.Duration) {
	redis.ClientRedis.Set(traceId, 0, cacheOverdue)
	gtimer.AddOnce(duration, func() {
		redisCache := redis.ClientRedis.Get(traceId)
		val, err := redisCache.Result()
		result := 0
		if err == nil {
			err = json.Unmarshal([]byte(val), &result)
			if err == nil {
				if result == 0 {
					redis.ClientRedis.Set(traceId, 1, 30*time.Second)
					ResponseData(c, i18n.TIME_OUT, nil, errors.New("请求超时"), false)
					return
				}
			}
		}
	})
}
