package handler

import (
	"net/http"
)

type NotFoundHandler struct {
}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (n NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}

var _ http.Handler = (*NotFoundHandler)(nil)
