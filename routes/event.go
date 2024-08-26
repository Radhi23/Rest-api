package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"userapp/rest-api/models"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the events by id."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events.."})
		return
	}
	context.JSON(http.StatusOK, events)

}

func createEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)
	fmt.Println("even", err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the events"})
		return
	}
	userId := context.GetInt64("userId")
	fmt.Println("hh", userId)
	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event.."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "Event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the events by id."})
		return
	}
	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not request the events"})
		return
	}
	updateEvent.ID = eventId

	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not Update event.."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	fmt.Println("1", err)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	fmt.Println("2", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the events by id."})
		return
	}
	err = event.Delete()
	fmt.Println("3", err)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the events by id."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted."})
}
