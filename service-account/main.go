package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nekko-ru/api/service-account/common"
	service "github.com/nekko-ru/api/service-account/handlers"
	"github.com/nekko-ru/api/service-account/helpers"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})

	r := gin.New()
	r.Use(helpers.Logger(log), gin.Recovery())

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.POST("/account.create", handlers.AccountCreate)
		g.GET("/account.activate", handlers.Activate)
		g.GET("/account.getProfileInfo", handlers.GetProfile)
		g.POST("/account.setProfileInfo", handlers.UpdateProfile)
		g.POST("/account.passwordRestore", handlers.PasswordRestore)
		g.POST("/account.passwordChange", handlers.PasswordChange)
		g.GET("/account.getSettings")
		g.POST("/account.setSettings")
		g.GET("/account.getPushSettings")
		g.POST("/account.setPushSettings")

		g.GET("/_/ping", handlers.PingCheck)
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
