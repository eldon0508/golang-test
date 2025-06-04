package room

import "time"

type Room struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Reservation struct {
	ID       string    `json:"id"`
	RoomID   string    `json:"room_id"`
	Date     time.Time `json:"date"`
	PlayerID string    `json:"player_id"`
}

var RoomsData = []Room{
	{ID: "1", Name: "Single", Description: "Single room"},
	{ID: "2", Name: "Double", Description: "Double room"},
}

var ReservationsData = []Reservation{
	{ID: "res1", RoomID: "1", Date: time.Date(2025, time.June, 5, 0, 0, 0, 0, time.UTC), PlayerID: "1"},
	{ID: "res2", RoomID: "2", Date: time.Date(2025, time.June, 5, 0, 0, 0, 0, time.UTC), PlayerID: "2"},
}
