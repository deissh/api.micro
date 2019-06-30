package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-anime/common"
	service "github.com/deissh/api.micro/service-anime/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	r := SetupRouter()

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}

// SetupRouter create Gin router and return one
func SetupRouter() *gin.Engine {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.POST("/anime.create", handlers.CreateAnime)
		g.GET("/anime.get", handlers.GetAnime)
		g.POST("/anime.update", handlers.UpdateAnime)
		g.GET("/anime.remove", handlers.RemoveAnime)
		g.GET("/anime.search")

		g.GET("/_/ping", handlers.PingCheck)
	}

	return r
}
