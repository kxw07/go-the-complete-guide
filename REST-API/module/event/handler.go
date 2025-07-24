package event

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/utils"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (handler Handler) RegisterRoutes(server *gin.Engine) {
	server.GET("/event/all", handler.getEvents)
	server.GET("/event/:id", handler.getEvent)
	server.POST("/event", handler.createEvent)
	server.PUT("/event/:id", handler.updateEvent)
	server.DELETE("/event/:id", handler.deleteEvent)
}

func (handler Handler) getEvents(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No Authorization"})
		return
	}

	err := utils.ValidToken(token)
	if err != nil {
		slog.Info("getEvents: valid token failed", "error", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No Authorization"})
		return
	}

	events, err := handler.svc.getEvents()
	if err != nil {
		slog.Info("getEvents: bind json failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (handler Handler) createEvent(context *gin.Context) {
	var event Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		slog.Info("createEvent: bind json failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := handler.svc.createEvent(event)
	if err != nil {
		slog.Info("createEvent: create event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": result})
}

func (handler Handler) getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("getEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := handler.svc.getEvent(eventId)
	if err != nil {
		slog.Info("getEvent: get event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func (handler Handler) updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("updateEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	var event Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		slog.Info("updateEvent: bind json failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	event.ID = eventId

	result, err := handler.svc.updateEvent(event)
	if err != nil {
		slog.Info("updateEvent: update event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated.", "event": result})
}

func (handler Handler) deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("deleteEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	err = handler.svc.deleteEvent(eventId)
	if err != nil {
		slog.Info("deleteEvent: delete event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted."})
}
