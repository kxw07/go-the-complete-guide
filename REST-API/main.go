package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/module/event"
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

var eventHandler = event.Handler{}

func main() {
	slog.Info("starting REST API server")

	storage.InitDB()

	server := gin.Default()
	server.SetTrustedProxies([]string{"127.0.0.1"})

	server.GET("/event/all", eventHandler.GetEvents)
	server.POST("/event", eventHandler.CreateEvent)

	server.Run(":8080")
}
