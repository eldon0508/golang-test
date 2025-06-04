package log

import "time"

type Log struct {
	ID        string    `json:"id"`
	PlayerID  string    `json:"player_id"`
	Action    string    `json:"action"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var LogsData = []Log{
	{ID: "log1", PlayerID: "1", Action: "login", StartTime: time.Date(2025, time.June, 4, 9, 0, 0, 0, time.UTC), EndTime: time.Date(2025, time.June, 4, 9, 30, 0, 0, time.UTC)},
	{ID: "log2", PlayerID: "2", Action: "login", StartTime: time.Date(2025, time.June, 4, 10, 15, 0, 0, time.UTC), EndTime: time.Date(2025, time.June, 4, 10, 45, 0, 0, time.UTC)},
	{ID: "log3", PlayerID: "1", Action: "enter-room", StartTime: time.Date(2025, time.June, 4, 11, 0, 0, 0, time.UTC), EndTime: time.Date(2025, time.June, 4, 11, 0, 0, 0, time.UTC)},
	{ID: "log4", PlayerID: "3", Action: "login", StartTime: time.Date(2025, time.June, 5, 9, 45, 0, 0, time.UTC), EndTime: time.Date(2025, time.June, 5, 10, 15, 0, 0, time.UTC)},
	{ID: "log5", PlayerID: "1", Action: "start-challenge", StartTime: time.Date(2025, time.June, 5, 10, 30, 0, 0, time.UTC), EndTime: time.Date(2025, time.June, 5, 10, 50, 0, 0, time.UTC)},
}
