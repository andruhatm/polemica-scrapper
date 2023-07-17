package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	h.l.Printf("RegisterGame router")

	//req to id path var
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["gameId"], 10, 0)
	if err != nil {
		http.Error(rw, "Unable to Marshal Game Id", http.StatusBadRequest)
		return
	}

	//make request to
	//https://polemicagames.kz/game-statistics/1

	h.l.Printf("Retrieving object with id = %d", id)
}
