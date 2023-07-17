package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"
	"os"
	"polemica_scrapper/models"
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
	//vars := mux.Vars(r)
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5433/games")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	resp := make([]models.RatingPlayer, 0)

	rows, _ := conn.Query(context.Background(), "select username,games_count,score,avg_score from player;")

	defer rows.Close()

	for rows.Next() {
		var playerStat models.RatingPlayer
		err := rows.Scan(&playerStat.Name, &playerStat.Games, &playerStat.Score, &playerStat.AvgScore)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, playerStat)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	h.l.Printf("res: %v", resp)

	rw.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(resp)
	rw.Write(res)
}
