package handlers

import (
	"log"
	"net/http"
)

// GameHandler is a simple handler
type GameHandler struct {
	l *log.Logger
}

// NewGameSync creates new handler with given logger
func NewGameSync(l *log.Logger) *GameHandler {
	return &GameHandler{l: l}
}

func (h *GameHandler) RegisterGame(rw http.ResponseWriter, r *http.Request) {

}
