package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"polemica_scrapper/parser"
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
	h.l.Printf("Retrieving object with id = %d", id)

	//make request to Polemica
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://polemicagame.com/game-statistics/%d", id), nil)
	if err != nil {
		h.l.Printf("client: could not create request: %s\n", err)
		http.Error(rw, "Unable to create request", http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		h.l.Printf("client: error making http request: %s\n", err)
		http.Error(rw, "Unable to make request", http.StatusInternalServerError)
		return
	}

	h.l.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		h.l.Printf("client: could not read response body: %s\n", err)
		http.Error(rw, "Unable to read response", http.StatusInternalServerError)
		return
	}

	//log.Println(string(resBody))

	gameData, err := parser.ParseGameData(resBody)
	if err != nil {
		h.l.Printf("client: could not parse game data: %s\n", err)
		http.Error(rw, "Unable to parse response", http.StatusInternalServerError)
		return
	}

	//store incoming data to db
	h.l.Printf("Retrieving object with id = %d", id, gameData)
}
