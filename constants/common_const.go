package constants

const (
	OTHER_FILE_DIRECTORY      = "other" //未知文件目录
	ADMIN_SESSION_EXPIRE_TIME = 10      //管理员会话保存时间
)

const (
	TIME_FORMAT_Y_M_D = "2006-01-02" //年月日时间格式化
)

const (
	STR_IS_EMPTY = "" //字符串为空
)

const (
	CODE_TOKEN_INVALID  = -1  //token失效
	CODE_PARAM_IS_NULL  = 401 //参数为空
	CODE_PARAM_IS_WRONG = 402 //参数不合法
	CODE_SYSTEM_ERROR   = 500 //系统报错
)

const (
	ERR_SYSTEM_ERROR                          = "系统报错,请稍后重试"
	ERR_FILE_PATH_CAN_NOT_BE_EMPTY            = "文件路径不能为空"
	ERR_FILE_PATH_IS_NOT_EXIST                = "文件路径不存在"
	ERR_USERNAME_OR_PASSWORD_CAN_NOT_BE_EMPTY = "用户名或者密码不能为空"
	ERR_USERNAME_OR_PASSWORD_IS_WRONG         = "用户名或者密码错误"
	ERR_LOGIN_OUT                             = "退出登录失败"
	ERR_VIDEO_NAME_CAN_NOT_BE_EMPTY           = "视频名称不能为空"
	ERR_VIDEO_DESCRIBE_CAN_NOT_BE_EMPTY       = "视频描述不能为空"
	ERR_VIDEO_CLASSIFY_CAN_NOT_BE_EMPTY       = "视频分类不能为空"
	ERR_ADD_VIDEO_FAIL                        = "添加视频失败"
)

const (
	SUCCESS_DEL_FILE = "文件删除成功"
)
