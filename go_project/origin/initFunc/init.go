package initFunc

import (
	"github.com/e421083458/golang_common/lib"
	"origin/cache/cacheMap"
	_ "origin/customFunc"
)

func init()  {
	configPath := "./conf/"
	_ = lib.InitModule(configPath, []string{"base"})
	defer lib.Destroy()

	cacheMap.InitConf()
}