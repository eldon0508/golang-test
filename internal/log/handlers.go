package log

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func timeToMinutes(t time.Time) int {
	return t.Hour()*60 + t.Minute()
}

func CheckLog(c *gin.Context) {
	playerID := c.Query("player_id")
	action := c.Query("action")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")
	limitStr := c.Query("limit")

	var filterStartTime time.Time
	var filterEndTime time.Time
	var err error

	if startTimeStr != "" {
		filterStartTime, err = time.Parse("15:04", startTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Please use HH:MM (e.g., 10:30)."})
			return
		}
	}

	if endTimeStr != "" {
		filterEndTime, err = time.Parse("15:04", endTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format. Please use HH:MM (e.g., 10:30)."})
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

	var filteredLogs []Log
	count := 0

	for _, l := range LogsData {
		if playerID != "" && playerID != l.PlayerID {
			continue
		}

		if action != "" && action != l.Action {
			continue
		}

		if !filterStartTime.IsZero() {
			logStartMinutes := timeToMinutes(l.StartTime)
			filterStartMinutes := timeToMinutes(filterStartTime)
			if logStartMinutes < filterStartMinutes {
				continue // Skip if log starts before the filter's start time
			}
		}

		// Apply end_time filter (log's EndTime must be <= filterEndTime's time of day)
		if !filterEndTime.IsZero() {
			logEndMinutes := timeToMinutes(l.EndTime)
			filterEndMinutes := timeToMinutes(filterEndTime)
			if logEndMinutes > filterEndMinutes {
				continue // Skip if log ends after the filter's end time
			}
		}

		// If all filters pass, add to results
		filteredLogs = append(filteredLogs, l)
		count++

		// Apply limit if specified
		if limit != -1 && count >= limit {
			break // Stop if limit is reached
		}
	}

	c.IndentedJSON(http.StatusOK, filteredLogs)
}

func AddLog(c *gin.Context) {
	var newLog Log
	if err := c.BindJSON(&newLog); err != nil {
		fmt.Printf("Error binding JSON for creating reservation: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newLog.ID = fmt.Sprintf("log%d", len(LogsData)+1)

	LogsData = append(LogsData, newLog)
	c.IndentedJSON(http.StatusCreated, newLog.ID)
}
