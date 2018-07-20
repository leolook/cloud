package controller

import (
	"cloud/tool/dao"
	pb "cloud/tool/protocol"
)

//创建model
func (s *Srv) CreateModel(com pb.CommonReq, req pb.ModelReq) (string, error) {
	return dao.Model.CreateModel(com.Url, req.Name)
}
