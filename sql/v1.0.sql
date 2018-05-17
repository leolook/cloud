CREATE database if NOT EXISTS `video` DEFAULT char set 'UTF8';

use `video`;

SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `video_file`;
CREATE TABLE `video_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE utf8_unicode_ci NOT NULL COMMENT '视频名称',
  `info` varchar(1000) COLLATE utf8_unicode_ci  COMMENT '视频描述',
  `cover` varchar(1000) COLLATE utf8_unicode_ci  COMMENT '视频封面',
  `classify_id` int(11) NOT NULL COMMENT '视频分类',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_user` varchar(200) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作人',
  `update_time` int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `video_classify`;
CREATE TABLE `video_classify` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(200) COLLATE utf8_unicode_ci NOT NULL COMMENT '分类名称',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `video_file_path`;
CREATE TABLE `video_file_path` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `file_id` int(11) NOT NULL COMMENT '视频id',
  `path` varchar(200) COLLATE utf8_unicode_ci NOT NULL COMMENT '视频路径',
  `number` int(2) NOT NULL COMMENT '视频文件顺序',
  `info` varchar(1000) COLLATE utf8_unicode_ci NOT NULL COMMENT '简要描述',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `video_admin`;
CREATE TABLE `video_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT '密钥',
  `is_login` int(1) NOT NULL DEFAULT 0 COMMENT '是否登录 默认 0 登录 1',
  `last_login_ip` varchar(20) COLLATE utf8_unicode_ci COMMENT '最后登陆IP',
  `last_login_time` int(11)  DEFAULT '0' COMMENT '最后登陆时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=137 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT video_admin set user_name='admin',password=sha1(md5('admin')),create_time=unix_timestamp(now());

INSERT INTO video_classify(name,create_time) VALUES('电影',unix_timestamp(now()));
INSERT INTO video_classify(name,create_time) VALUES('短视频',unix_timestamp(now()));
INSERT INTO video_classify(name,create_time) VALUES('电视',unix_timestamp(now()));
INSERT INTO video_classify(name,create_time) VALUES('其它',unix_timestamp(now()));
