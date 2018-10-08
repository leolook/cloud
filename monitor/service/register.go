package service

import (
	"cloud/httpproto/monitor"
	"cloud/lib/log"
)

func (s *Server) Register(req *monitor.RegReq) error {

	log.Infof("register=%+v", req)

	return nil
}

func (s *Server) Register1(req *monitor.RegReq) (rsp *monitor.RefRsp, err error) {

	return nil, nil
}
