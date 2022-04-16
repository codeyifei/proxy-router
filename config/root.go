package config

import (
	"github.com/codeyifei/proxy-router/handler"
	"github.com/codeyifei/proxy-router/types"
	"github.com/creasty/defaults"
)

type RootRouterConfig struct {
	BaseRouterConfig
	StaticPath string
	IndexPath  []string `default:"[\"index.html\"]"`
}

func (c *RootRouterConfig) Handler() types.Handler {
	return handler.NewRootHandler(c.Prefix, c.StaticPath, c.IndexPath...)
}

var _ RouterConfig = (*RootRouterConfig)(nil)

func NewRootRouterConfig(pathPrefix, staticPath string, indexPath ...string) *RootRouterConfig {
	c := &RootRouterConfig{
		BaseRouterConfig: NewBaseRouterConfig(pathPrefix),
		StaticPath:       staticPath,
		IndexPath:        indexPath,
	}
	defaults.MustSet(c)
	return c
}
