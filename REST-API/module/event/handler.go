package event

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/module/middlewares"
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
	authenticated := server.Group("/event")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/all", handler.getEvents)
	authenticated.GET("/:id", handler.getEvent)
	authenticated.POST("/", handler.createEvent)
	authenticated.PUT("/:id", handler.updateEvent)
	authenticated.DELETE("/:id", handler.deleteEvent)
	authenticated.POST("/:id/registry", handler.registerEvent)
	authenticated.DELETE("/:id/registry", handler.unregisterEvent)
}

func (handler Handler) getEvents(context *gin.Context) {
	userId := context.GetInt64("userId")

	events, err := handler.svc.getEvents(userId)
	if err != nil {
		slog.Info("getEvents: bind json failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (handler Handler) createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	var event Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		slog.Info("createEvent: bind json failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := handler.svc.createEvent(event, userId)
	if err != nil {
		slog.Info("createEvent: create event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": result})
}

func (handler Handler) getEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("getEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	event, err := handler.svc.getEvent(eventId, userId)
	if err != nil {
		slog.Info("getEvent: get event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func (handler Handler) updateEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

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

	result, err := handler.svc.updateEvent(event, userId)
	if err != nil {
		slog.Info("updateEvent: update event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event updated.", "event": result})
}

func (handler Handler) deleteEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("deleteEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	err = handler.svc.deleteEvent(eventId, userId)
	if err != nil {
		slog.Info("deleteEvent: delete event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted."})
}

func (handler Handler) registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("deleteEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	err = handler.svc.registerEvent(eventId, userId)
	if err != nil {
		slog.Info("registerEvent: register event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event."})
}

func (handler Handler) unregisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		slog.Info("deleteEvent: get eventId failed", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	err = handler.svc.unregisterEvent(eventId, userId)
	if err != nil {
		slog.Info("unregisterEvent: unregister event failed", "error", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unregister from event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Unregistered from event."})
}
