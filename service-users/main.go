package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-users/common"
	service "github.com/nekko-ru/api/service-users/handlers"

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
