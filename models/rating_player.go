package models

type RatingPlayer struct {
	Name     string  `json:"username"`
	Games    int32   `json:"games_count"`
	Score    float32 `json:"score"`
	AvgScore float32 `json:"avg_score"`
}
