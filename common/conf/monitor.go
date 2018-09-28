package conf

import (
	"cloud/common/flag"
	"cloud/lib/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MonitorConf struct {
	//Redis相关配置
	Monitor struct {
		Name string
		Addr string
	}
}

var monitorConf *MonitorConf

func Monitor() *MonitorConf {

	if monitorConf != nil {
		return monitorConf
	}

	bytes, err := ioutil.ReadFile(flag.MonitorConfPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = yaml.Unmarshal(bytes, &monitorConf)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return monitorConf
}
