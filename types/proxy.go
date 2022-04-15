package types

type Proxy struct {
	Host    string
	BaseUri string
}

func NewProxy(host, baseUri string) Proxy {
	return Proxy{host, baseUri}
}
