package handler

import (
	"net/http"
	"strings"

	"github.com/codeyifei/proxy-router/types"
)

var (
	defaultOrigins = []string{"*"}
	defaultHeaders = []string{"Content-Type", "AccessToken", "X-CSRF-Token", "Authorization", "Token"}
	defaultMethods = []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"}
)

type CoreHandler struct {
	http.Handler
	origins          []string
	headers          []string
	allowCredentials bool
	methods          []string
}

func NewCoreHandler(handler http.Handler) *CoreHandler {
	return &CoreHandler{
		Handler: handler,
		origins: defaultOrigins,
		headers: defaultHeaders,
		methods: defaultMethods,
	}
}

func (h *CoreHandler) SetOrigins(origins ...string) *CoreHandler {
	if len(origins) == 0 {
		origins = defaultHeaders
	}

	h.origins = origins
	return h
}

func (h *CoreHandler) SetHeaders(headers ...string) *CoreHandler {
	if len(headers) == 0 {
		headers = defaultHeaders
	}

	h.headers = headers
	return h
}

func (h *CoreHandler) AllowCredentials() *CoreHandler {
	h.allowCredentials = true
	return h
}

func (h *CoreHandler) SetMethods(methods ...string) *CoreHandler {
	if len(methods) == 0 {
		methods = defaultMethods
	}

	h.methods = methods
	return h
}

func (h *CoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", strings.Join(h.origins, ","))
	w.Header().Add("Access-Control-Allow-Headers", strings.Join(h.headers, ","))
	if h.allowCredentials {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
	}
	w.Header().Add("Access-Control-Allow-Methods", strings.Join(h.methods, ","))
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	h.Handler.ServeHTTP(w, r)
}

var _ types.Handler = (*CoreHandler)(nil)
