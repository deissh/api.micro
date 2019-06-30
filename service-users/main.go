package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-users/common"
	service "github.com/deissh/api.micro/service-users/handlers"
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
		g.GET("/users.get", handlers.GetUser)
		g.GET("/users.getFollowers")
		g.GET("/users.getSubscriptions")
		g.GET("/users.report")
		g.GET("/users.search")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
