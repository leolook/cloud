package middle

import (
	pb "cloud/tool/protocol"
	"encoding/base64"
	"fmt"
	log "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	. "github.com/gitbubhwt/baseserver/util"
)

func Common(c *gin.Context) (interface{}, error) {

	url := c.GetHeader("Url")
	if url == "" {
		return nil, Error(pb.ERR_URL_EMPTY, "url empty")
	}

	bytes, err := base64.StdEncoding.DecodeString(url)
	if err != nil {
		log.Error(fmt.Sprintf("Base64 decode fail,[err=%v] [url=%v]", err, url))
		return nil, Error(pb.ERR_URL_PARSE, "url parse error")
	}

	if len(bytes) <= 0 {
		return nil, Error(pb.ERR_URL_EMPTY, "url empty")
	}

	common := pb.CommonReq{
		Url: string(bytes),
	}
	return common, nil
}
