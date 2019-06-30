package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-notifications/common"
	service "github.com/deissh/api.micro/service-notifications/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/notifications.get")
		g.GET("/notifications.markAsViewed")
		g.GET("/notifications.sendMessage")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
