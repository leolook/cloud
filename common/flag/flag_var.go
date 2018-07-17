package flag

import "flag"

var (
	ConfPath        string
	GatewayConfPath string
)

func init() {
	flag.StringVar(&ConfPath, "conf", "config.yml", "config path")
	flag.StringVar(&GatewayConfPath, "gateway", "gateway.yml", "gateway path")
	flag.Parse()
}
