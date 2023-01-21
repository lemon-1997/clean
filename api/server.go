package api

import (
	"github.com/lemon-1997/clean/api/handler"
	"github.com/lemon-1997/clean/config"
	"net/http"
)

type Handle interface {
	Handler() map[string]http.HandlerFunc
}

type Server struct {
	addr  string
	order *handler.Order
}

func NewServer(c *config.Server, order *handler.Order) *Server {
	return &Server{
		addr:  c.Addr,
		order: order,
	}
}

func (s *Server) Run() error {
	s.register()
	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) register() {
	registerOrder(s.order)
}

func registerOrder(h Handle) {
	for path, handle := range h.Handler() {
		http.HandleFunc(path, handle)
	}
}
