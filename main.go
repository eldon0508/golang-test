package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
}

var players = []player{
	{ID: "1", Name: "Alice", Level: 99},
	{ID: "2", Name: "Bob", Level: 199},
	{ID: "3", Name: "Cassie", Level: 299},
}

func main() {
	fmt.Println("Hello, World!")

	router := gin.Default()
	router.GET("/players", playersIndex)
	router.POST("/players", playerStore)
	router.GET("/players/:id", playersView)
	router.PUT("/players/:id", playerUpdate)
	router.DELETE("/players/:id", playerDelete)

	router.Run("localhost:3000")
}

func playersIndex(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, players)
}

func playerStore(c *gin.Context) {
	var newPlayer player

	if err := c.BindJSON(&newPlayer); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPlayer.ID = fmt.Sprintf("%d", len(players)+1)

	players = append(players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer)
}

func playersView(c *gin.Context) {
	id := c.Param("id")
	for _, a := range players {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Player not found"})
}

func playerUpdate(c *gin.Context) {
	id := c.Param("id")

	var updatedPlayer player
	if err := c.BindJSON(&updatedPlayer); err != nil {
		fmt.Printf("Error binding JSON for update: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range players {
		if p.ID == id {
			fmt.Println(players[i])
			players[i].Name = updatedPlayer.Name
			players[i].Level = updatedPlayer.Level
			c.IndentedJSON(http.StatusOK, players[i])
			return
		}
	}
}

func playerDelete(c *gin.Context) {
	id := c.Param("id")

	for i, p := range players {
		if p.ID == id {
			players = append(players[:i], players[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Player deleted"})
			return
		}
	}
}
