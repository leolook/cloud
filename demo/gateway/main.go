package main

import (
	"cloud/demo/gateway/proxy"
	etcd "cloud/lib/etcd3"

	"fmt"
	"net"
	"time"

	log "cloud/lib/log"

	"google.golang.org/grpc"
)

func main() {

	name := "gateway"
	addr := "0.0.0.0:8001"
	target := "http://127.0.0.1:2379"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("Failed to net.Listen(%s,%s),err=%v\n", "tcp", addr, err)
		return
	}

	log.Infof("test:=%s", "test")

	log.Infof("hu=%s", "hu")

	log.Infof("teshhhhhhhhhhh")

	log.Infof("test=%+v", "test")

	err = etcd.Register(name, addr, target, 10*time.Second, 15)
	if err != nil {
		fmt.Println(err)
	}

	log.Debugf("debug=%+v", listener)

	option := proxy.Option()
	server := grpc.NewServer(
		option...,
	)
	server.Serve(listener)
}
