package proxy

import (
	"cloud/gateway/regexp"
	"fmt"
	log "github.com/alecthomas/log4go"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Handle struct{}

func (h *Handle) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	log.Info(fmt.Sprintf("req=%v", path))

	suc, addr := h.checkPath(path)
	if !suc {
		h.do404(w)
		return
	}

	//去掉gateway请求头
	replaceReg := regexp.GetReplaceReg()
	req.URL.Path = replaceReg.ReplaceAllString(path, "")

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

func (h *Handle) do404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "404 page not found")
}
