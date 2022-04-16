package types

import "net/http"

type Option func(http.Handler)
