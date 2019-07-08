package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-news/common"
	service "github.com/nekko-ru/api/service-news/handlers"
)

func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.POST("/news.create", handlers.CreateNews)
		g.GET("/news.get", handlers.GetNews)
		g.POST("/news.update", handlers.UpdateNews)
		g.GET("/news.remove", handlers.RemoveNews)
		g.GET("/news.search")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
