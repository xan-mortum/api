package router

import (
	"net/http"

	"github.com/xan-mortum/api/internal/hello/handler"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Get(prefix string) http.Handler {
	mux := http.NewServeMux()

	h := handler.NewHelloHandler()
	mux.Handle("/", h.World())

	return http.StripPrefix(prefix, mux)
}
