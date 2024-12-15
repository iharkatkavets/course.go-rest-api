package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event id because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	err = event.Register(userId)
	if err != nil {
		message := fmt.Sprintf("Couldn't register user for event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registered"})
}

func cancelRegistration(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event id because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		message := fmt.Sprintf("Couldn't cancel event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}
