package main

import (
	"cloud/gateway/conf"
	"cloud/gateway/proxy"
	"cloud/gateway/regexp"
	log "github.com/alecthomas/log4go"
	"net/http"
)

func main() {

	gateway := conf.GetGatewayConf().Gateway

	h := &proxy.Handle{}
	err := http.ListenAndServe(gateway.Addr, h)
	if err != nil {
		log.Exit(err)
	}
}

func init() {
	conf.GetGatewayConf()
	regexp.GetServerReg()
}
