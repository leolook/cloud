package fileserver

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"net/http"
	"os"
)

func Remove(w http.ResponseWriter, req *http.Request) {

	//解析表单
	err := req.ParseForm()
	if err != nil {
		log.Error(fmt.Sprintf("Parse form fail,[err=%v]", err))
		ReplyJson(w, FailRsp(ERR_SERVER_ERR, "server err"))
		return
	}

	path := req.FormValue("path")

	if path == "" {
		ReplyJson(w, FailRsp(ERR_IS_EMPTY, "path is empty"))
		return
	}

	err = os.Remove(path)
	if err != nil {
		ReplyJson(w, FailRsp(ERR_SERVER_ERR, "remove fail"))
		return
	}

	ReplyJson(w, SucRsp(http.StatusOK, "remove success"))
}
