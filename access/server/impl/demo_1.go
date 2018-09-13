package impl

import (
	"cloud/pb"
	"context"
	"fmt"
)

func (s *Srv)Demo(ctx context.Context, req *pb.DemoReq) (*pb.DemoRsp, error) {

	fmt.Println("req", req.String())

	return &pb.DemoRsp{
		Id: req.Id,
	}, nil
}
