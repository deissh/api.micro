package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-news/common"
	service "github.com/deissh/api.micro/service-news/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/deissh/api.micro/service-news/docs"
)

// @title Service Users API
// @version 1.0
// @description Users methods

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
		g.GET("/news.create", handlers.CreateNews)
		g.GET("/news.get", handlers.GetNews)
		g.GET("/news.update", handlers.UpdateNews)
		g.GET("/news.remove", handlers.RemoveNews)
		g.GET("/news.search")

		g.GET("/_/health", handlers.HealthCheckHandler)
		g.GET("/_/ping", handlers.PingHandler)
		g.GET("/_/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Use(gin.Recovery())

	if err := r.Run(helpers.GetEnv("HTTP_HOST", ":8080")); err != nil {
		log.Error(err)
	}
}
