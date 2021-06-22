package zh

import "origin/modules/i18n"

func InitZh() {
	i18n.Error1 = "失败！"
	i18n.Success1 = "成功！"
	i18n.MsgFlags = map[int]string{
		i18n.SUCCESS_CODE:               "成功",
		i18n.INTERNAL_ERROR_CODE:        "服务器内部错误",
		i18n.PARMAS_ERROR_CODE:          "参数错误",
		i18n.USER_NOT_EXIST:             "用户不存在",
		i18n.PASSWORD_ERROR_CODE:        "密码错误",
		i18n.GENERATE_TOKEN_ERROR_CODE:  "生成token失败",
		i18n.USERNAME_OR_PWD_ERROR_CODE: "用户名或者密码错误",
		i18n.DATA_EXIST:                 "数据已经存在,请勿重复创建",
	}
}
