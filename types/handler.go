package types

import "net/http"

type Handler interface {
	http.Handler
}
