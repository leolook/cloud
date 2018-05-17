package constants

const (
	TABLE_VIDEO_ADMIN = "video_admin"
	TABLE_VIDEO_FILE  = "video_file"
	TABLE_VIDEO_PATH  = "video_file_path"
	TABLE_CLASSIFY    = "video_classify"
)

const (
	ADMIN_LOGIN_CHECK_SQL = "select id from " + TABLE_VIDEO_ADMIN + " where user_name=? and password=?"

	GET_LAST_ID_BY_INSERT_SQL = "SELECT LAST_INSERT_ID()"

	ADMIN_VIDEO_IS_EXIST_BY_NAME_SQL = "select id from " + TABLE_VIDEO_FILE + " where name =?"
	ADMIN_VIDEO_FILE_GET_SQL         = "select id,name,info,classify_id from " + TABLE_VIDEO_FILE + " where id=?"
	ADMIN_VIDEO_FILE_INSERT_SQL      = "insert into " + TABLE_VIDEO_FILE + "(name,info,cover,classify_id,update_user,create_time,update_time) values(?,?,?,?,?,unix_timestamp(now()),unix_timestamp(now()))"
	ADMIN_VIDEO_FILE_UPDATE_SQL      = "update " + TABLE_VIDEO_FILE + " set name=?,info=?,cover=?,classify_id=?,update_user=?,update_time=unix_timestamp(now()) where id=?"
	ADMIN_VIDEO_FILE_LIKE_SQL        = "and name like '%s' "
	ADMIN_VIDEO_FILE_LIST_SQL        = "select * from " + TABLE_VIDEO_FILE + " where 1=1  %s limit ?,?"
	ADMIN_VIDEO_FILE_COUNT_SQL       = "select count(*) from " + TABLE_VIDEO_FILE + " where 1=1  %s "

	ADMIN_VIDEO_PATH_GET_SQL    = "select * from " + TABLE_VIDEO_PATH + " where file_id=?"
	ADMIN_VIDEO_PATH_INSERT_SQL = "insert into " + TABLE_VIDEO_PATH + "(file_id,path,number,info,create_time) values(?,?,?,?,?)"
	ADMIN_VIDEO_PATH_DELETE_SQL = "delete from " + TABLE_VIDEO_PATH + " where file_id=?"

	ADMIN_CLASSIFY_ALL_SQL = "select id,name from " + TABLE_CLASSIFY
)
