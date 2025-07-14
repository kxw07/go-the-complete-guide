package event

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	svc Service
}

func (handler Handler) GetEvents(context *gin.Context) {
	context.JSON(http.StatusOK, handler.svc.getEvents())
}

func (handler Handler) CreateEvent(context *gin.Context) {
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

func (handler Handler) GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID."})
		return
	}

	context.JSON(http.StatusOK, handler.svc.getEvent(eventId))
}
