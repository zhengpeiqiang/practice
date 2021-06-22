package i18nInit

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/text/gstr"
	"origin/modules/ginFunc"
	"origin/modules/i18n/en"
	"origin/modules/i18n/zh"
)

var LangList = []string{
	"zh", "en",
}

func LangCheck(c *gin.Context) {
	lang := ginFunc.GetGinQuery(c, "lang")
	if lang == "" {
		zh.InitZh()
		return
	}
	if gstr.InArray(LangList, lang) {
		switch lang {
		case "zh":
			zh.InitZh()
			break
		case "en":
			en.InitEn()
			break
		default:
			zh.InitZh()
			break
		}
		return
	}
	zh.InitZh()
	return
}
