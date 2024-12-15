package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	events := []models.Event{}
	storedEvents, err := models.GetAllEvents()
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch events because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}
	if storedEvents != nil {
		events = storedEvents
	}
	ctx.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event id because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event json because %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	event.UserID = ctx.GetInt64("userId")
	err = event.Save()
	if err != nil {
		message := fmt.Sprintf("Couldn't create event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event id because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	userId := ctx.GetInt64("userId")
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	var updatedEvent models.Event
	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event json because %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		message := fmt.Sprintf("Couldn't update event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		message := fmt.Sprintf("Couldn't parse event id because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		message := fmt.Sprintf("Couldn't fetch event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	userId := ctx.GetInt64("userId")
	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = event.Delete()
	if err != nil {
		message := fmt.Sprintf("Couldn't delete event because %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
