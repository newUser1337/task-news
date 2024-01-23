package muxserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxServer struct {
	router *mux.Router
	port   int
}

func NewMuxServer(port int, router *mux.Router) *muxServer {
	return &muxServer{
		router: router,
		port:   port,
	}
}

func (m *muxServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", m.port), m.router)
}
