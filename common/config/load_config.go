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
}

var conf AppConfig

func LoadConfig(file string) {
	conf = AppConfig{}
	err := configor.Load(&conf, file)
	if err != nil {
		logger.Error("Failed to find configuration ", file)
		os.Exit(1)
	}
}

func init() {
	confPath := flag.ConfPath
	LoadConfig(confPath)
	for k, v := range os.Args {
		logger.Info(fmt.Sprintf("k:%d, v:%s", k, v))
	}
	logger.Info(fmt.Sprintf("load config path %s", confPath))
	logger.Info("redis addr:", conf.Db.Addr)
}

func GetConf() AppConfig {
	return conf
}
