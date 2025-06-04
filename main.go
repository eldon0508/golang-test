package main

import (
	"fmt"

	"example/interview_juntaoyeap_20250604/internal/challenge"
	"example/interview_juntaoyeap_20250604/internal/log"
	"example/interview_juntaoyeap_20250604/internal/payment"
	"example/interview_juntaoyeap_20250604/internal/player"
	"example/interview_juntaoyeap_20250604/internal/room"

	"github.com/gin-gonic/gin"
)

var players = []player.Player{
	{ID: "1", Name: "Alice", Level: player.Level{ID: "1", Name: "Beginner"}},
	{ID: "2", Name: "Bob", Level: player.Level{ID: "2", Name: "Intermediate"}},
	{ID: "3", Name: "Cassie", Level: player.Level{ID: "1", Name: "Beginner"}},
}

func main() {
	fmt.Println("Starting API server...")

	playerService := player.NewPlayerService(&players)

	router := gin.Default()

	router.GET("/players", playerService.GetPlayers)
	router.POST("/players", playerService.CreatePlayer)
	router.GET("/players/:id", playerService.GetPlayerByID)
	router.PUT("/players/:id", playerService.UpdatePlayer)
	router.DELETE("/players/:id", playerService.DeletePlayer)

	router.GET("/levels", playerService.GetLevels)
	router.POST("levels", playerService.CreateLevel)

	router.GET("/rooms", room.GetRooms)
	router.POST("/rooms", room.CreateRoom)
	router.GET("/rooms/:id", room.GetRoomByID)
	router.PUT("/rooms/:id", room.UpdateRoom)
	router.DELETE("/rooms/:id", room.DeleteRoom)

	router.GET("/reservations", room.GetReservation)
	router.POST("/reservations", room.CreateReservation)

	router.POST("/challenges", challenge.CreateChallenge)
	router.GET("/challenges/results", challenge.CheckResult)

	router.GET("/logs", log.CheckLog)
	router.POST("/logs", log.AddLog)

	router.POST("/payments", payment.CreatePayment)
	router.GET("/payments/:id", payment.GetPaymentByID)

	router.Run("localhost:3000")
}
