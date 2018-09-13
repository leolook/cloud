package impl

import (
	"cloud/pb"
	"google.golang.org/grpc"
)

type Srv struct {
	pb.DemoServiceServer
}

func (s *Srv) RegServer(grpcServer *grpc.Server) {
	pb.RegisterDemoServiceServer(grpcServer, s)
}
