package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-anime/common"
	service "github.com/nekko-ru/api/service-anime/handlers"

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
		g.GET("/anime.search", handlers.FindAnime)

		g.GET("/_/ping", handlers.PingCheck)
	}

	return r
}
