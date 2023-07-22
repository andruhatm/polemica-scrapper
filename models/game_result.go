package models

type GameResults struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	DaysNumber  float64     `json:"daysNumber"`
	WinnerCode  int         `json:"winnerCode"`
	Judge       interface{} `json:"judge"`
	Players     []Player    `json:"players"`
	FirstKilled string      `json:"firstKilled"`
}

type Player struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Image    string `json:"image"`
	Role     struct {
		Type  string `json:"type"`
		Title string `json:"title"`
	} `json:"role"`
	TablePosition int    `json:"tablePosition"`
	WL            string `json:"w_l"`
	Points        int    `json:"points"`
	Coins         int    `json:"coins"`
	Achievements  []struct {
		Sum   float32       `json:"sum"`
		Array []interface{} `json:"array"`
	} `json:"achievements"`
	AchievementsSum struct {
		Points float32 `json:"points"`
		//TODO parse types of achievements
		Achievements Achievements `json:"achievements"`
	} `json:"achievementsSum"`
}

type Achievements struct {
	Achievements struct {
		BestMove struct {
			Points float32 `json:"points"`
		} `json:"best_move"`
		SheriffDetection struct {
			Points float32 `json:"points"`
		} `json:"sheriff_detection"`
		Voting struct {
			Points float32 `json:"points"`
		} `json:"voting"`
		Victory struct {
			Points int `json:"points"`
		} `json:"victory"`
	} `json:"achievements"`
}
