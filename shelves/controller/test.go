package controller

import (
	pb "cloud/shelves/protocol"
	"fmt"
	log "github.com/alecthomas/log4go"
)

func (s *Srv) Test(common pb.CommonReq, req pb.TestReq) (pb.TestRsp, error) {
	log.Info(fmt.Sprintf("[common=%v][req=%v]", common, req))
	return pb.TestRsp{
		Name: "test",
		Id:   common.UserId,
	}, nil
}

func (s *Srv) Test1(req pb.TestReq) (pb.TestRsp, error) {
	return pb.TestRsp{
		Name: "test1",
	}, nil
}
