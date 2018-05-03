package constants

const (
	keyPrefix = "video."
	//管理员会话列表
	ADMIN_LIST_SESSION_HASH_KEY = keyPrefix + "admin.list.session"

	//管理员会话表
	ADMIN_SESSION_HASH_KEY    = keyPrefix + "admin.session.%s"
	ADMIN_SESSION_USER_NAME   = "user_name"
	ADMIN_SESSION_TOKEN       = "token"
	ADMIN_SESSION_CREATE_TIME = "create_time"
)
