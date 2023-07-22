package adapter

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"polemica_scrapper/models"
)

func GetStats() ([]models.RatingPlayer, error) {
	resp := make([]models.RatingPlayer, 0)

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:123@localhost:5433/games")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	rows, _ := conn.Query(context.Background(), "select username,games_count,score,avg_score from player;")

	defer rows.Close()

	for rows.Next() {
		var playerStat models.RatingPlayer
		err := rows.Scan(&playerStat.Name, &playerStat.Games, &playerStat.Score, &playerStat.AvgScore)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return nil, err
		}
		resp = append(resp, playerStat)
	}
	return resp, nil
}

func SaveRatings(err error) {

}

func SaveGameHistory(err error) {

}
