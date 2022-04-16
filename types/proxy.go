package types

type Proxy struct {
	Scheme  string
	Host    string
	BaseUri string
}

func NewProxy(scheme, host, baseUri string) Proxy {
	return Proxy{scheme, host, baseUri}
}
