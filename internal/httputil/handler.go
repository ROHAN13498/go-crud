package httputil

import (
	"net/http"
)

type HttpHandler interface {
	Get(pattern string, handler func(http.ResponseWriter, *http.Request))
	Post(pattern string, handler func(http.ResponseWriter, *http.Request))
	Put(pattern string, handler func(http.ResponseWriter, *http.Request))
	Delete(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type Handler struct{}

func (h *Handler) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("GET "+pattern, handler)
}

func (h *Handler) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("POST "+pattern, handler)
}

func (h *Handler) Put(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("PUT "+pattern, handler)
}

func (h *Handler) Delete(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("DELETE "+pattern, handler)
}
