package flag

import "flag"

var (
	ConfPath        string
	GatewayConfPath string
	MonitorConfPath string
)

func init() {
	flag.StringVar(&ConfPath, "conf", "config.yml", "config path")
	flag.StringVar(&GatewayConfPath, "gatewayConf", "gateway.yml", "gateway path")
	flag.StringVar(&MonitorConfPath, "monitorConf", "monitor.yml", "monitor path")
	flag.Parse()
}
