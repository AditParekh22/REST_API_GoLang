package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {

	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.Try again later"})
		return
	}
	context.JSON(200, events)
}
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	events, err := models.GetEventID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	}
	context.JSON(200, events)
}
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBind(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events.Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	_, err = models.GetEventID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
		return
	}
	fmt.Println(updatedEvent)
	context.JSON(http.StatusOK, gin.H{"message": "Event updated."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event, err := models.GetEventID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
