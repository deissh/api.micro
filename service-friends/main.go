package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-friends/common"
	service "github.com/nekko-ru/api/service-friends/handlers"
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
