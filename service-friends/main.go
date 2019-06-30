package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-friends/common"
	service "github.com/deissh/api.micro/service-friends/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
)

func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/friends.add")
		g.GET("/friends.addList")
		g.GET("/friends.areFriends")
		g.GET("/friends.delete")
		g.GET("/friends.deleteAllRequests")
		g.GET("/friends.deleteList")
		g.GET("/friends.get")
		g.GET("/friends.getOnline")
		g.GET("/friends.getRequests")
		g.GET("/friends.getSuggestions")
		g.GET("/friends.search")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
