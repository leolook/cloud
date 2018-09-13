package main

import (
	"cloud/access/server/impl"
	base "cloud/lib/base_server"
	"fmt"
	log "github.com/alecthomas/log4go"
)

func main() {
	flag := make(chan int, 1)
	addr := fmt.Sprintf("0.0.0.0:8002")
	log.Info("addr=%v", addr)
	grpc := base.NewGrpcServer(addr, "access", &impl.Srv{})
	err := grpc.StartServer()
	if err != nil {
		log.Exit(err)
	}
	<-flag
}
