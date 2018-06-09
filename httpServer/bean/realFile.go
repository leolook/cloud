package bean

type RealFile struct {
	Id           int64  `json:"key"`                               //文件id
	VideoId      string `json:"videoId" xorm:"video_id"`           //视频id
	Path         string `json:"path"`                              //文件路径
	FullName     string `json:"fullName" xorm:"full_name"`         //操作人
	FileType     int    `json:"fileType" xorm:"file_type"`         //文件类型  0-图片 1-视频
	CreateTime   int64  `json:"createTime" xorm:"create_time"`     //创建时间 时间戳
	RelativeTime int64  `json:"relativeTime" xorm:"relative_time"` //关联时间 时间戳
}

type RealFilePage struct {
	BasePageReq
	FileType   int `json:"fileType"`
	IsRelative int `json:"isRelative"`
}
