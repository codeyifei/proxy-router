package config

import (
	"github.com/codeyifei/proxy-router/handler"
	"github.com/codeyifei/proxy-router/types"
)

type CoreRouterConfig struct {
	BaseRouterConfig
	RouterConfig
	Origins []string
	Headers []string
	Methods []string
}

func (c *CoreRouterConfig) Handler() types.Handler {
	h := handler.NewCoreHandler(c.RouterConfig.Handler())
	h.SetOrigins(c.Origins...)
	h.SetHeaders(c.Headers...)
	h.SetMethods(c.Methods...)
	return h
}
