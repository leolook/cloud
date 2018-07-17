package regexp

import (
	"cloud/gateway/conf"
	"fmt"
	"regexp"
)

var (
	gatewayReg, replaceReg *regexp.Regexp
	addr                   []string
	regs                   []*regexp.Regexp
)

func init() {

	gatewayConf := conf.GetGatewayConf()
	gatewayReg = regexp.MustCompile(fmt.Sprintf("/%s/([a-z])", gatewayConf.Gateway.Name))
	replaceReg = regexp.MustCompile(fmt.Sprintf("/%s/", gatewayConf.Gateway.Name))

	server := gatewayConf.Server
	gateway := gatewayConf.Gateway

	names := []string{server.Shelves.Name, server.File.Name, server.Center.Name}
	addr = []string{server.Shelves.Addr, server.File.Addr, server.Center.Addr}

	regs = make([]*regexp.Regexp, len(names))
	for i, v := range names {
		regs[i] = regexp.MustCompile(fmt.Sprintf("/%s/%s/([a-z])", gateway.Name, v))
	}
}

func GetGatewayReg() *regexp.Regexp {
	return gatewayReg
}

func GetServerReg() ([]*regexp.Regexp, []string) {
	return regs, addr
}

func GetReplaceReg() *regexp.Regexp {
	return replaceReg
}
