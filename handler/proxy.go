package handler

import (
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/codeyifei/goproxy/types"
)

type ProxyHandler struct {
	PathPrefix string
	Proxy      types.Proxy
}

func NewProxyHandler(pathPrefix string, proxy types.Proxy) *ProxyHandler {
	return &ProxyHandler{
		PathPrefix: pathPrefix,
		Proxy:      proxy,
	}
}

func (h ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = h.Proxy.Host
			req.Host = h.Proxy.Host
			req.URL.Path = h.Proxy.BaseUri + strings.TrimPrefix(r.URL.Path, h.PathPrefix)
		},
	}
	reverseProxy.ServeHTTP(w, r)
}

var _ http.Handler = (*ProxyHandler)(nil)
