package config

import (
	"strings"
)

type BaseRouterConfig struct {
	Prefix string `default:"/"`
	core   bool
}

func (c *BaseRouterConfig) PathPrefix() string {
	return c.Prefix
}

func (c *BaseRouterConfig) EnableCore() {
	c.core = true
}

func (c *BaseRouterConfig) IsNeedCore() bool {
	return c.core
}

func NewBaseRouterConfig(prefix string) BaseRouterConfig {
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	return BaseRouterConfig{Prefix: prefix}
}
