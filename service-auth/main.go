package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-auth/common"
	service "github.com/nekko-ru/api/service-auth/handlers"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	r := gin.New()
	r.Use(helpers.Logger(), gin.Recovery())

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/token.create", handlers.TokenCreate)
		g.GET("/token.refresh", handlers.TokenRefresh)
		g.GET("/token.remove", handlers.TokenRemove)
		g.GET("/token.check", handlers.TokenCheck)

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
