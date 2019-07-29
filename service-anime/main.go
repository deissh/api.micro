package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nekko-ru/api/service-anime/common"
	"github.com/nekko-ru/api/service-anime/handlers"
	"github.com/nekko-ru/api/service-anime/helpers"
	"github.com/nekko-ru/api/service-anime/services"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})

	r := SetupRouter()

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}

// SetupRouter create Gin router and return one
func SetupRouter() *gin.Engine {
	conn := common.Init()
	common.Migrate()

	r := gin.New()
	r.Use(helpers.Logger(log), gin.Recovery())

	srv := services.Services{Db: conn, Log: log}

	h := handlers.Handler{srv}

	g := r.Group("/")
	{
		g.POST("/anime.create", h.CreateAnime)
		g.GET("/anime.get", h.GetAnime)
		g.POST("/anime.update", h.UpdateAnime)
		g.GET("/anime.search", h.FindAnime)
		g.POST("/anime.vote")

		g.GET("/_/ping", h.PingCheck)
	}

	return r
}
