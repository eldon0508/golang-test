package player

type Level struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level Level  `json:"level"`
}

type PlayerInput struct {
	Name    string `json:"name"`
	LevelID string `json:"level_id"`
}

var LevelsData = []Level{
	{ID: "1", Name: "Beginner"},
	{ID: "2", Name: "Intermediate"},
}
