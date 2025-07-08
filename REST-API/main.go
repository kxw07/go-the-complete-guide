package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/model"
	"github.com/kxw07/REST-API/storage"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("starting REST API server")
	
	storage.InitDB()

	server := gin.Default()
	server.SetTrustedProxies([]string{"127.0.0.1"})

	server.GET("/event/all", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, model.GetAll())
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	model.AddEvent(event)

	context.JSON(http.StatusCreated, gin.H{"message": "Event created.", "event": event})
}
