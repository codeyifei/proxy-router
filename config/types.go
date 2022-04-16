package config

import (
	"sort"
	"strings"

	"github.com/codeyifei/goproxy/handler"
	"github.com/codeyifei/goproxy/types"
	"github.com/gorilla/mux"
)

type Config struct {
	Port    uint
	Routers RouterConfigSlice
}

func NewConfig(port uint, routers ...RouterConfig) Config {
	sort.Sort(RouterConfigSlice(routers))
	return Config{port, routers}
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
