package player

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerService struct {
	Players *[]Player // Pointer to the players slice in main.go
}

// NewPlayerService creates a new instance of PlayerService
func NewPlayerService(p *[]Player) *PlayerService {
	return &PlayerService{
		Players: p,
	}
}

func (ps *PlayerService) GetPlayers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, *ps.Players)
}

func (ps *PlayerService) CreatePlayer(c *gin.Context) {
	var input PlayerInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Error binding JSON for creating player: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newPlayer Player
	newPlayer.ID = fmt.Sprintf("%d", len(*ps.Players)+1)
	newPlayer.Name = input.Name

	foundLevel := false
	for _, l := range LevelsData {
		if l.ID == input.LevelID {
			newPlayer.Level = l
			foundLevel = true
			break
		}
	}

	if !foundLevel {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid LevelID provided"})
		return
	}

	*ps.Players = append(*ps.Players, newPlayer)
	c.IndentedJSON(http.StatusCreated, newPlayer.ID)
}

func (ps *PlayerService) GetPlayerByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range *ps.Players {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Player not found"})
}

func (ps *PlayerService) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")

	var input PlayerInput
	if err := c.BindJSON(&input); err != nil {
		fmt.Printf("Error binding JSON for update player: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedLevel Level
	foundLevel := false
	for _, l := range LevelsData {
		if l.ID == input.LevelID {
			updatedLevel = l
			foundLevel = true
			break
		}
	}

	if !foundLevel {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid LevelID provided"})
		return
	}

	playerFoundInSlice := false
	for i, p := range *ps.Players {
		if p.ID == id {
			playerFoundInSlice = true
			(*ps.Players)[i].Name = input.Name
			(*ps.Players)[i].Level = updatedLevel
			fmt.Println("Updated Player:", (*ps.Players)[i])
			c.IndentedJSON(http.StatusOK, (*ps.Players)[i])
			return
		}
	}

	if !playerFoundInSlice {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Player not found"})
	}
}

func (ps *PlayerService) DeletePlayer(c *gin.Context) {
	id := c.Param("id")

	for i, p := range *ps.Players {
		if p.ID == id {
			*ps.Players = append((*ps.Players)[:i], (*ps.Players)[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Player has been deleted"})
			return
		}
	}
}

func (ps *PlayerService) GetLevels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, LevelsData)
}

func (ps *PlayerService) CreateLevel(c *gin.Context) {
	var newLevel Level

	if err := c.BindJSON(&newLevel); err != nil {
		fmt.Printf("Error binding JSON for creating level: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newLevel.ID = fmt.Sprintf("%d", len(LevelsData)+1)

	LevelsData = append(LevelsData, newLevel)
	c.IndentedJSON(http.StatusCreated, newLevel.ID)
}
