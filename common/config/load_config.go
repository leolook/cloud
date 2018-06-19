package config

import (
	"cloud/common/flag"
	"cloud/common/logger"
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

type AppConfig struct {
	//Redis相关配置
	Db struct {
		Addr     string
		Password string
		User     string
		Database int
	}

	Http struct {
		Addr string
	}

	Mysql struct {
		Addr     string
		User     string
		Password string
		Database string
	}

	UploadFile struct {
		Path string
	}
}

var conf AppConfig

func LoadConfig(file string) {
	conf = AppConfig{}
	err := configor.Load(&conf, file)
	if err != nil {
		logger.Error("Failed to find configuration ", file)
		os.Exit(1)
	}
	logger.Info("conf", conf)
}

func init() {
	confPath := flag.ConfPath
	LoadConfig(confPath)
	for k, v := range os.Args {
		logger.Info(fmt.Sprintf("k:%d, v:%s", k, v))
	}
	logger.Info(fmt.Sprintf("load config path: %s", confPath))
	logger.Info(fmt.Sprintf("redis addr:%s", conf.Db.Addr))
	logger.Info(fmt.Sprintf("httpServer addr:%s", conf.Http.Addr))
	logger.Info(fmt.Sprintf("mysql:%v", conf.Mysql))
	logger.Info(fmt.Sprintf("upload file path:%v", conf.UploadFile))
}

func GetConf() AppConfig {
	return conf
}
