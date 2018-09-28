package proxy

import (
	"fmt"
	"cloud/lib/log"
	"google.golang.org/grpc"
	//"strings"
)

func Proxy(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

	log.Info(fmt.Sprintf("stream=%+v", stream))
	return nil
}
