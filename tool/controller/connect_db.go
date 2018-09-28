package controller

import (
	"cloud/tool/dao"
	"cloud/tool/mysql"
	pb "cloud/tool/protocol"
	"encoding/base64"
	"fmt"
	"cloud/lib/log"
	. "github.com/gitbubhwt/baseserver/util"
)

//连接数据库
func (s *Srv) ConnectDb(req pb.ConnectInfoReq) (*pb.ConnectInfoRsp, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		req.User, req.Pwd, req.Addr, req.Db)
	db := mysql.GetClient(url)
	if db != nil {
		db.Close()
		mysql.RemoveClient(url)
	}
	var err error
	db, err = mysql.NewClient(url)
	if err != nil {
		log.Error(err)
		return nil, Error(pb.ERR_FAIL_CONNECT_DB, fmt.Sprintf("Fail connect db,[err=%v]", err))
	}
	if db == nil {
		return nil, Error(pb.ERR_FAIL_CONNECT_DB, "Fail connect db")
	}
	//获取所有表
	list, err := dao.Model.ShowTables(url)
	if err != nil {
		log.Error(err)
		return nil, Error(pb.ERR_FAIL_CONNECT_DB, fmt.Sprintf("Fail query db,[err=%v]", err))
	}
	url = base64.StdEncoding.EncodeToString([]byte(url))
	return &pb.ConnectInfoRsp{Url: url,
		List: list}, nil
}
