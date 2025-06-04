package challenge

type Challenge struct {
	ID       string  `json:"id"`
	Fee      float64 `json:"fee"`
	PlayerID string  `json:"player_id"`
	Status   string  `json:"status"`
}

var ChallengesData = []Challenge{
	{ID: "1", Fee: 20.01, PlayerID: "1", Status: "Failed"},
	{ID: "2", Fee: 20.01, PlayerID: "1", Status: "Failed"},
}
