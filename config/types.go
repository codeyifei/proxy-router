package config

import (
	"sort"
	"strings"

	"github.com/codeyifei/proxy-router/handler"
	"github.com/codeyifei/proxy-router/types"
	"github.com/gorilla/mux"
)

type Config struct {
	Host    string
	Port    uint
	Routers RouterConfigSlice
}

func NewConfig(host string, port uint, routers ...RouterConfig) Config {
	sort.Sort(RouterConfigSlice(routers))
	return Config{host, port, routers}
}

type RouterConfig interface {
	PathPrefix() string
	EnableCore()
	IsNeedCore() bool
	Handler() types.Handler
}

type RouterConfigSlice []RouterConfig

func (r RouterConfigSlice) Len() int {
	return len(r)
}

func (r RouterConfigSlice) Less(i, j int) bool {
	return strings.HasPrefix(r[i].PathPrefix(), r[j].PathPrefix())
}

func (r RouterConfigSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

var _ sort.Interface = (*RouterConfigSlice)(nil)

func (r RouterConfigSlice) Register(router *mux.Router) {
	for _, item := range r {
		h := item.Handler()
		if item.IsNeedCore() {
			h = handler.NewCoreHandler(item.Handler())
		}
		router.PathPrefix(item.PathPrefix()).Handler(h)
	}
}
