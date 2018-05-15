package constants

const (
	TABLE_VIDEO_ADMIN = "video_admin"
	TABLE_VIDEO_FILE  = "video_file"
)

const (
	ADMIN_LOGIN_CHECK_SQL = "select id from " + TABLE_VIDEO_ADMIN + " where user_name=? and password=?"

	ADMIN_VIDEO_IS_EXIST_BY_NAME = "select count(*) from " + TABLE_VIDEO_FILE + " where name =?"
)
