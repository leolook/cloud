package controller

import (
	pb "cloud/tool/protocol"
	"fmt"
	"cloud/lib/log"
)

func (s *Srv) MobileLog(req pb.MobileLogReq) (rsp *pb.MobileLogRsp, err error) {
	log.Info(fmt.Sprintf("Mobile log:%v", req.Data))
	return nil, nil
}
