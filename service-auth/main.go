package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nekko-ru/api/service-auth/common"
	"github.com/nekko-ru/api/service-auth/handlers"
	"github.com/nekko-ru/api/service-auth/helpers"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})

	r := gin.New()
	r.Use(helpers.Logger(log), gin.Recovery())

	conn := common.Init()
	h := handlers.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/token.create", h.TokenCreate)
		g.GET("/token.refresh", h.TokenRefresh)
		g.GET("/token.remove", h.TokenRemove)
		g.GET("/token.check", h.TokenCheck)

		g.GET("/_/ping", h.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
