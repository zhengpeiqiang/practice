package i18n

var (
	Error1   = ""
	Success1 = ""
)

var (
	SUCCESS_CODE               = 200
	INTERNAL_ERROR_CODE        = 4001 //服务器内部错误
	PARMAS_ERROR_CODE          = 4002 //参数错误
	USER_NOT_EXIST             = 4003 //用户不存在
	PASSWORD_ERROR_CODE        = 4004 //密码错误
	GENERATE_TOKEN_ERROR_CODE  = 4005 //生成token失败
	USERNAME_OR_PWD_ERROR_CODE = 4006 //用户名或者密码错误
	DATA_EXIST                 = 4007 //数据已经存在,请勿重复创建
	TIME_OUT                   = 4008 //
	DATA_NOT_EXIST             = 4009 //
)

var MsgFlags = map[int]string{
	SUCCESS_CODE:               "成功",
	INTERNAL_ERROR_CODE:        "服务器内部错误",
	PARMAS_ERROR_CODE:          "参数错误",
	USER_NOT_EXIST:             "用户不存在",
	PASSWORD_ERROR_CODE:        "密码错误",
	GENERATE_TOKEN_ERROR_CODE:  "生成token失败",
	USERNAME_OR_PWD_ERROR_CODE: "用户名或者密码错误",
	DATA_EXIST:                 "数据已经存在,请勿重复创建",
	TIME_OUT:                   "请求超时",
	DATA_NOT_EXIST:             "数据不存在",
}
