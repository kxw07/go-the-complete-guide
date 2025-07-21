package event

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	svc Service
}

func (handler Handler) RegisterRoutes(server *gin.Engine) {
	server.GET("/event/all", handler.getEvents)
	server.GET("/event/:id", handler.getEvent)
	server.POST("/event", handler.createEvent)
	server.PUT("/event/:id", handler.updateEvent)
}

func (handler Handler) getEvents(context *gin.Context) {
	events, err := handler.svc.getEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (handler Handler) createEvent(context *gin.Context) {
	var event Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.svc.createEvent(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": result})
}

func (handler Handler) getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	event, err := handler.svc.getEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func (handler Handler) updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = eventId

	result, err := handler.svc.updateEvent(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated.", "event": result})
}
