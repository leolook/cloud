package conf

import (
	"cloud/common/flag"
	"fmt"
	log "github.com/alecthomas/log4go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	gatewayConf GatewayConf
)

type GatewayConf struct {
	Gateway struct {
		Name string
		Addr string
	}

	Server struct {
		Center  ServerInfo
		Shelves ServerInfo
		File    ServerInfo
	}

	Allow struct {
		Path []string
	}
}

type ServerInfo struct {
	Addr string
	Name string
}

func LoadGatewayConf(path string) error {
	gatewayConf = GatewayConf{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, &gatewayConf)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	path := flag.GatewayConfPath
	err := LoadGatewayConf(path)
	if err != nil {
		log.Exit(fmt.Sprintf("Load gatewag conf fail,[err=%v]", err))
	}
	log.Info(fmt.Sprintf("gatewayConf=%v", gatewayConf))
}

//获取gateway 配置文件
func GetGatewayConf() GatewayConf {
	return gatewayConf
}

//获取allow 通过的路径
func GetAllowPath() map[string]string {
	mp := make(map[string]string)
	for _, v := range gatewayConf.Allow.Path {
		mp[v] = v
	}
	return mp
}
