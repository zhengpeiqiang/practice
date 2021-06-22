package en

import "origin/modules/i18n"

func InitEn()  {
	i18n.Error1   = "fail！"
	i18n.Success1 = "success！"
	i18n.MsgFlags = map[int]string{
		i18n.SUCCESS_CODE:               "success",
		i18n.INTERNAL_ERROR_CODE:        "server internal error",
		i18n.PARMAS_ERROR_CODE:          "parameter error",
		i18n.USER_NOT_EXIST:             "user does not exist",
		i18n.PASSWORD_ERROR_CODE:        "password error",
		i18n.GENERATE_TOKEN_ERROR_CODE:  "failed to generate token",
		i18n.USERNAME_OR_PWD_ERROR_CODE: "wrong user name or password",
		i18n.DATA_EXIST:                 "the data already exists. do not create it repeatedly",
	}
}