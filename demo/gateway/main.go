package main

import (
	"cloud/demo/gateway/proxy"
	etcd "cloud/lib/etcd3"

	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	log "cloud/lib/log"
)

func main() {

	name := "gateway"
	addr := "0.0.0.0:8001"
	target := "http://127.0.0.1:2379"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		//log.Error("Failed to net.Listen(%s,%s),err=%v\n", "tcp", addr, err)
		return
	}

	log.Infof("test:=%s","test")

	err = etcd.Register(name, addr, target, 10*time.Second, 15)
	if err != nil {
		fmt.Println(err)
	}

	option := proxy.Option()
	server := grpc.NewServer(
		option...,
	)
	server.Serve(listener)
}
