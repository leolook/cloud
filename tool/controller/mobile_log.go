package controller

import (
	pb "cloud/tool/protocol"
	"fmt"
	log "github.com/alecthomas/log4go"
)

func (s *Srv) MobileLog(req pb.MobileLogReq) (rsp *pb.MobileLogRsp, err error) {
	log.Info(fmt.Sprintf("Mobile log:%v", req.Data))
	return nil, nil
}
