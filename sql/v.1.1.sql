CREATE database if NOT EXISTS `video` DEFAULT char set 'UTF8';

use `video`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `real_file`;
CREATE TABLE `real_file` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`video_id` int(11) DEFAULT 0 COMMENT 'video_id',
`path` varchar(200) COLLATE utf8_unicode_ci NOT NULL COMMENT '文件路径',
`full_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人',
`file_type` int(1) NOT NULL DEFAULT 0 COMMENT '文件类型 默认 0-图片 1-视频',
`create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
`relative_time` int(11) NOT NULL DEFAULT '0' COMMENT '关联时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
