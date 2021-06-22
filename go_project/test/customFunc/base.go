package customFunc

import (
	"github.com/e421083458/golang_common/log"
	"github.com/gogf/gf/util/gconv"
)

var Base *BaseBuf

type BaseBuf struct {

}

func init()  {
	Base = &BaseBuf{}
	log.Info(gconv.String(Base))
}