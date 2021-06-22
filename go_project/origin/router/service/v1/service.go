package v1

import (
	"github.com/e421083458/golang_common/lib"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"origin/router/model"
)

// 从gin的Context中获取数据
func GetGinTraceContext(c *gin.Context) *lib.TraceContext {
	// 防御
	if c == nil {
		return lib.NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*lib.TraceContext); ok {
			return tc
		}
	}
	return lib.NewTrace()
}

type UserService struct {
}

//创建用户
func (w *UserService) CreateUserDetail(c *gin.Context, tx *gorm.DB, userDetail model.User) error {
	return tx.SetCtx(GetGinTraceContext(c)).Create(&userDetail).Error
}