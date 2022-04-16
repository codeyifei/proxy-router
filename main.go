package main

import (
	"fmt"
	"net/http"

	_ "github.com/spf13/viper"

	"github.com/codeyifei/proxy-router/config"
	"github.com/codeyifei/proxy-router/middleware"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group{}
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	for _, server := range conf {
		g.Go(func() error {
			r := mux.NewRouter()
			server.Routers.Register(r)
			// 404
			r.Use(middleware.AccessLog)

			return http.ListenAndServe(fmt.Sprintf(":%d", server.Port), r)
		})
	}

	if err := g.Wait(); err != nil {
		panic(fmt.Errorf("服务启动失败 %w", err))
	}
}
