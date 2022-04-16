package config

import (
	"net/url"
	"strings"

	"github.com/codeyifei/proxy-router/handler"
	"github.com/codeyifei/proxy-router/types"
	"github.com/creasty/defaults"
)

type ProxyRouterConfig struct {
	BaseRouterConfig
	Scheme  string `default:"http"`
	Host    string
	BaseUri string `default:"/"`
}

func (p ProxyRouterConfig) Handler() types.Handler {
	return handler.NewProxyHandler(p.BaseRouterConfig.Prefix, types.NewProxy(p.Scheme, p.Host, p.BaseUri))
}

var _ RouterConfig = (*ProxyRouterConfig)(nil)

func NewProxyRouterConfig(pathPrefix, scheme, host, baseUri string) *ProxyRouterConfig {
	if !strings.HasPrefix(baseUri, "/") {
		baseUri = "/" + baseUri
	}
	c := &ProxyRouterConfig{
		BaseRouterConfig: NewBaseRouterConfig(pathPrefix),
		Scheme:           scheme,
		Host:             host,
		BaseUri:          baseUri,
	}
	defaults.MustSet(c)
	return c
}

func NewProxyRouterConfigByUrl(pathPrefix, proxyPass string) (*ProxyRouterConfig, error) {
	u, err := url.Parse(proxyPass)
	if err != nil {
		return nil, err
	}
	return NewProxyRouterConfig(pathPrefix, u.Scheme, u.Host, u.Path), nil
}
