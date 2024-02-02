package muxserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxServer struct {
	server *http.Server
}

func NewMuxServer(port int, router *mux.Router) *muxServer {
	return &muxServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
	}
}

func (m *muxServer) Run() error {
	return m.server.ListenAndServe()
}

func (m *muxServer) Stop() error {
	return m.server.Close()
}
