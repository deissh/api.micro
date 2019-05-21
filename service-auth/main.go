package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-auth/common"
	service "github.com/deissh/api.micro/service-auth/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/deissh/api.micro/service-auth/docs"
)

// @title Service Auth API
// @version 1.0
// @description Auth, create tokens, and refresh old

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	conn := common.Init()
	common.Migrate()

	handlers := service.CreateHandlers(conn)

	g := r.Group("/")
	{
		g.GET("/token.create", handlers.TokenCreate)
		g.GET("/token.refresh", handlers.TokenRefresh)
		g.GET("/token.remove", handlers.TokenRemove)
		g.GET("/token.check", handlers.TokenCheck)

		g.GET("/_/health", handlers.HealthCheck)
		g.GET("/_/ping", handlers.PingCheck)
		g.GET("/_/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
