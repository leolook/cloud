package main

import (
	"cloud/demo/gateway/proxy"
	etcd "cloud/lib/etcd3"

	"net"
	"time"

	log "cloud/lib/log"

	"google.golang.org/grpc"
)

func main() {

	name := "main.go"
	addr := "0.0.0.0:8001"
	target := "http://127.0.0.1:2379"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("failed to net.Listen(%s,%s),err=%v", "tcp", addr, err)
		return
	}

	err = etcd.Register(name, addr, target, 10*time.Second, 15)
	if err != nil {
		log.Errorf("failed to register etcd,err=%+v,name=%v,addr=%v,target=%v", err, name, addr, target)
		return
	}

	option := proxy.Option()
	server := grpc.NewServer(
		option...,
	)
	server.Serve(listener)
}
