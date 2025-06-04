package room

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RoomService struct {
	Rooms *[]Room
}

func GetRooms(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, RoomsData)
}

func CreateRoom(c *gin.Context) {
	var newRoom Room
	if err := c.BindJSON(&newRoom); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRoom.ID = fmt.Sprintf("%d", len(RoomsData)+1)

	RoomsData = append(RoomsData, newRoom)
	c.IndentedJSON(http.StatusCreated, newRoom.ID)
}

func GetRoomByID(c *gin.Context) {
	id := c.Param("id")
	for _, r := range RoomsData {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Room not found"})
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")

	var updatedRoom Room
	if err := c.BindJSON(&updatedRoom); err != nil {
		fmt.Printf("Error binding JSON for updating room: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, r := range RoomsData {
		if r.ID == id {
			RoomsData[i].Name = updatedRoom.Name
			RoomsData[i].Description = updatedRoom.Description
			c.IndentedJSON(http.StatusOK, RoomsData[i])
			return
		}
	}
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")

	for i, r := range RoomsData {
		if r.ID == id {
			RoomsData = append(RoomsData[:i], RoomsData[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Room has been deleted"})
			return
		}
	}
}

func GetReservation(c *gin.Context) {
	roomID := c.Query("room_id")
	dateStr := c.Query("date")
	limitStr := c.Query("limit")

	var filterDate time.Time
	var err error

	if dateStr != "" {
		filterDate, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Printf("Error parsing date '%s': %v\n", dateStr, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Expected YYYY-MM-DD."})
			return
		}
	}

	limit := -1 // -1 means no limit
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit. Must be a non-negative integer."})
			return
		}
	}

	var filteredReservations []Reservation
	count := 0
	for _, res := range ReservationsData {
		if roomID != "" && roomID != res.RoomID {
			continue
		}

		if !filterDate.IsZero() && !(res.Date.Year() == filterDate.Year() &&
			res.Date.Month() == filterDate.Month() &&
			res.Date.Day() == filterDate.Day()) {
			continue
		}

		// Add to filtered results
		filteredReservations = append(filteredReservations, res)
		count++

		// Apply limit if specified
		if limit != -1 && count >= limit {
			break // Stop if limit is reached
		}
	}

	c.IndentedJSON(http.StatusOK, filteredReservations)
}

func CreateReservation(c *gin.Context) {
	var newRes Reservation

	if err := c.BindJSON(&newRes); err != nil {
		fmt.Printf("Error binding JSON for creating reservation: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRes.ID = fmt.Sprintf("res%d", len(ReservationsData)+1)

	ReservationsData = append(ReservationsData, newRes)
	c.IndentedJSON(http.StatusCreated, newRes.ID)
}
