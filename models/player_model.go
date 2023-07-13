package models

type Player struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Role     struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	} `json:"role"`
	TablePosition   int    `json:"tablePosition"`
	WL              string `json:"w_l"`
	Points          int    `json:"points"`
	AchievementsSum struct {
		Points       float64 `json:"points"`
		Achievements struct {
			Victory struct {
				Title  string `json:"title"`
				Points int    `json:"points"`
			} `json:"victory"`
			Voting struct {
				Title  string  `json:"title"`
				Points float64 `json:"points"`
			} `json:"voting"`
		} `json:"achievements"`
	} `json:"achievementsSum"`
}
