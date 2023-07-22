package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"polemica_scrapper/database/adapter"
)

// StatsHandler is a simple handler
type StatsHandler struct {
	l *log.Logger
}

// NewStatsHandler creates new handler with given logger
func NewStatsHandler(l *log.Logger) *StatsHandler {
	return &StatsHandler{l: l}
}

func (h *StatsHandler) FetchRatings(rw http.ResponseWriter, r *http.Request) {
	h.l.Printf("FetchRatings router")

	//req to db (wtih params opt)
	resp, err := adapter.GetStats()
	if err != nil {
		h.l.Printf("client: could not read ratings: %s\n", err)
		http.Error(rw, "Unable to read current ratings", http.StatusInternalServerError)
	}

	h.l.Printf("res: %v", resp)

	res, err := json.Marshal(resp)
	if err != nil {
		h.l.Printf("client: could not read response body: %s\n", err)
		http.Error(rw, "Unable to read response", http.StatusInternalServerError)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(res)
}
