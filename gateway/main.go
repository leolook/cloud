package main

import (
	"cloud/gateway/conf"
	"cloud/gateway/proxy"
	"cloud/gateway/regexp"
	"cloud/lib/log"
	"net/http"
)

func main() {

	gateway := conf.GetGatewayConf().Gateway

	h := &proxy.Handle{}
	err := http.ListenAndServe(gateway.Addr, h)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	conf.GetGatewayConf()
	regexp.GetServerReg()
}
