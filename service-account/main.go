package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nekko-ru/api/helpers"
	"github.com/nekko-ru/api/service-account/common"
	service "github.com/nekko-ru/api/service-account/handlers"
	log "github.com/sirupsen/logrus"
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
