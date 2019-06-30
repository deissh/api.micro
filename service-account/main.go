package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-account/common"
	service "github.com/deissh/api.micro/service-account/handlers"
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
		g.POST("/account.create", handlers.AccountCreate)
		g.GET("/account.activate", handlers.Activate)
		g.GET("/account.getProfileInfo", handlers.GetProfile)
		g.POST("/account.setProfileInfo")
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
