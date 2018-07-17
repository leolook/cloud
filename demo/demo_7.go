package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type handle struct {
}

func main() {
	h := &handle{}
	err := http.ListenAndServe(":9090", h)
	if err != nil {
		//log.Fatalln("ListenAndServe: ", err)
	}

}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://127.0.0.1:3030")
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

	proxy.ServeHTTP(w, r)
}
