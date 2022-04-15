package main

import (
	"net/http"

	"github.com/codeyifei/goproxy/handler"
	"github.com/codeyifei/goproxy/middleware"
	"github.com/codeyifei/goproxy/types"
	"github.com/gorilla/mux"
)

var proxyMap = map[string]types.Proxy{
	"/api":                                types.NewProxy("api.unified-authority.server.test", "/v1"),
	"/upload-operating":                   types.NewProxy("api.operating.server.test", "/"),
	"/unified-authority-api":              types.NewProxy("api.unified-authority.server.test", "/v1"),
	"/unified-authority-upload-operating": types.NewProxy("api.operating.server.test", "/"),
}

func main() {
	r := mux.NewRouter()
	for k, v := range proxyMap {
		r.PathPrefix(k).Handler(handler.NewProxyHandler(k, v))
	}
	r.PathPrefix("/").Handler(handler.NewRootHandler("dist", "index.html"))
	// 404
	r.NotFoundHandler = handler.NewNotFoundHandler()
	r.Use(middleware.AccessLog)

	s := &http.Server{Addr: ":80", Handler: r}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
