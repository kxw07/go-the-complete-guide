package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kxw07/REST-API/module/event"
	"github.com/kxw07/REST-API/module/user"
	"github.com/kxw07/REST-API/storage"
	"log/slog"
)

func main() {
	slog.Info("starting REST API server")

	server := gin.Default()
	server.SetTrustedProxies([]string{"127.0.0.1"})

	sto := storage.NewStorage()
	userRepository := user.NewRepository(sto)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(server)

	eventRepository := event.NewRepository(sto)
	eventService := event.NewService(eventRepository)
	eventHandler := event.NewHandler(eventService)
	eventHandler.RegisterRoutes(server)

	server.Run(":8080")
}
