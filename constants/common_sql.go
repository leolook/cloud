package constants

const (
	TABLE_VIDEO_ADMIN = "video_admin"
	TABLE_VIDEO_FILE  = "video_file"
	TABLE_VIDEO_PATH  = "video_file_path"
	TABLE_CLASSIFY    = "video_classify"
	TABLE_REAL_FILE   = "real_file"
)

const (
	ADMIN_LOGIN_CHECK_SQL = "select id,full_name,img_path from " + TABLE_VIDEO_ADMIN + " where user_name=? and password=?"

	GET_LAST_ID_BY_INSERT_SQL = "SELECT LAST_INSERT_ID()"

	ADMIN_VIDEO_IS_EXIST_BY_NAME_SQL = "select id from " + TABLE_VIDEO_FILE + " where name =?"
	ADMIN_VIDEO_FILE_GET_SQL         = "select a.id,a.name,a.cover,a.info,a.classify_id,b.name as classify from " + TABLE_VIDEO_FILE + " as a left join " + TABLE_CLASSIFY + " as b on a.classify_id=b.id where a.id=?"
	ADMIN_VIDEO_FILE_INSERT_SQL      = "insert into " + TABLE_VIDEO_FILE + "(name,info,cover,classify_id,update_user,create_time,update_time) values(?,?,?,?,?,unix_timestamp(now()),unix_timestamp(now()))"
	ADMIN_VIDEO_FILE_UPDATE_SQL      = "update " + TABLE_VIDEO_FILE + " set name=?,info=?,cover=?,classify_id=?,update_user=?,update_time=unix_timestamp(now()) where id=?"
	ADMIN_VIDEO_FILE_LIKE_SQL        = "and a.name like '%s' "
	ADMIN_VIDEO_FILE_LIST_SQL        = "select a.*,b.name as classify from " + TABLE_VIDEO_FILE + " as a left join " + TABLE_CLASSIFY + " as b on a.classify_id=b.id  where 1=1  %s  order by id desc limit ?,?"

	VIDEO_FILE_COUNT_SQL = "select count(*) from " + TABLE_VIDEO_FILE + " as a where 1=1  %s "

	ADMIN_VIDEO_PATH_GET_SQL    = "select * from " + TABLE_VIDEO_PATH + " where file_id=?"
	ADMIN_VIDEO_PATH_INSERT_SQL = "insert into " + TABLE_VIDEO_PATH + "(file_id,path,number,info,create_time) values(?,?,?,?,?)"
	ADMIN_VIDEO_PATH_DELETE_SQL = "delete from " + TABLE_VIDEO_PATH + " where file_id=?"

	ADMIN_CLASSIFY_ALL_SQL = "select id,name from " + TABLE_CLASSIFY

	ADMIN_REAL_FILE_INSERT      = "insert into " + TABLE_REAL_FILE + "(path,full_name,file_type,create_time) values(?,?,?,unix_timestamp(now()))" //插入新文件
	ADMIN_REAL_FILE_UPDATE      = "update " + TABLE_REAL_FILE + "set video_id=? where id=?"                                                       //更新文件关联关系
	ADMIN_REAL_FILE_LIST        = "select * from " + TABLE_REAL_FILE + " where 1=1 %s  order by id desc limit ?,?"
	ADMIN_REAL_FILE_DEL         = "delete from " + TABLE_REAL_FILE + " where id in(?)"
	ADMIN_REAL_FILE_DEL_BY_PATH = "delete from " + TABLE_REAL_FILE + " where path=?"

	USER_VIDEO_FILE_LIST_SQL = "select id,name,cover from " + TABLE_VIDEO_FILE + " where 1=1  %s limit ?,?"
)
