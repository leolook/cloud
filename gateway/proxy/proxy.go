package proxy

import (
	"cloud/gateway/conf"
	"cloud/gateway/regexp"
	"fmt"
	"github.com/gitbubhwt/baseserver/util"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Handle struct{}

//web proxy
func (h *Handle) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	suc, addr := h.checkPath(path)
	if !suc {
		h.do404(w)
		return
	}

	//去掉gateway请求头
	req.URL.Path = regexp.GetReplaceReg().ReplaceAllString(path, "")

	ok := h.allow(req.URL.Path)
	if ok {
		h.doProxy(w, req, addr)
		return
	}

	ok = h.checkToken(w, req)
	if ok {
		h.doProxy(w, req, addr)
		return
	}

	h.doInvalidToken(w)
}

//执行代理
func (h *Handle) doProxy(w http.ResponseWriter, req *http.Request, addr string) {
	remote, err := url.Parse(fmt.Sprintf("http://%s", addr))
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	var pTransport http.RoundTripper = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	proxy.Transport = pTransport
	proxy.ServeHTTP(w, req)
}

//检查访问路径
func (h *Handle) checkPath(path string) (bool, string) {

	gatewayReg := regexp.GetGatewayReg()
	suc := gatewayReg.FindStringIndex(path)
	if suc == nil || len(suc) <= 0 {
		return false, ""
	}

	serverReg, addr := regexp.GetServerReg()
	for i, v := range serverReg {
		suc := v.FindStringIndex(path)
		if len(suc) > 0 {
			return true, addr[i]
		}
	}

	return false, ""
}

//返回404
func (h *Handle) do404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	io.Writer.Write(w, util.FailByte(http.StatusNotFound, "404"))
}

//无效token
func (h *Handle) doInvalidToken(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNonAuthoritativeInfo)
	io.Writer.Write(w, util.FailByte(-1, "invalid token"))
}

//允许通过的路径
func (h *Handle) allow(path string) bool {
	allowMp := conf.GetAllowPath()
	if _, ok := allowMp[path]; ok {
		return true
	}
	return false
}

//token 验证
func (h *Handle) checkToken(w http.ResponseWriter, req *http.Request) bool {

	return false
}
