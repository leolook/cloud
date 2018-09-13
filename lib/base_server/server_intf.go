package base_server

import "cloud/lib/base_server/impl"

type ServerIntf interface {
	StartServer() error
}

func NewGrpcServer(addr, name string, srv interface{}) ServerIntf {
	return &impl.GrpcImpl{
		Name: name,
		Addr: addr,
		Srv:  srv,
	}
}
