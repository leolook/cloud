package constants

const (
	TABLE_VIDEO_ADMIN = "video_admin"
	TABLE_VIDEO_FILE  = "video_file"
	TABLE_VIDEO_PATH  = "video_file_path"
)

const (
	ADMIN_LOGIN_CHECK_SQL = "select id from " + TABLE_VIDEO_ADMIN + " where user_name=? and password=?"

	GET_LAST_ID_BY_INSERT_SQL = "SELECT LAST_INSERT_ID()"

	ADMIN_VIDEO_IS_EXIST_BY_NAME_SQL = "select count(*) from " + TABLE_VIDEO_FILE + " where name =?"
	ADMIN_VIDEO_FILE_INSERT_SQL      = "insert into " + TABLE_VIDEO_FILE + "(name,info,classify_id,update_user,create_time,update_time) values(?,?,?,?,unix_timestamp(now()),unix_timestamp(now()))"
	ADMIN_VIDEO_PATH_INSERT_SQL      = "insert into " + TABLE_VIDEO_PATH + "(file_id,path,number,info,create_time) values(?,?,?,?,?)"
)
