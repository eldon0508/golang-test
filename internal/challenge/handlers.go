package challenge

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChallenge(c *gin.Context) {
	var newChallenge Challenge

	if err := c.BindJSON(&newChallenge); err != nil {
		fmt.Printf("Error binding JSON for creating challenge: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newChallenge.ID = fmt.Sprintf("%d", len(ChallengesData)+1)
	ChallengesData = append(ChallengesData, newChallenge)
	c.IndentedJSON(http.StatusCreated, newChallenge)
}

func CheckResult(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ChallengesData)
}
