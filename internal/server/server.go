package server

import (
	"github.com/newUser1337/task-news/internal/server/muxserver"

	"github.com/gorilla/mux"
)

type Server interface {
	Run() error
	Stop() error
}

func NewServer(port int, router *mux.Router) Server {
	return muxserver.NewMuxServer(port, router)
}
