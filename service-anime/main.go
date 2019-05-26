package main

import (
	"github.com/deissh/api.micro/helpers"
	"github.com/deissh/api.micro/service-anime/common"
	service "github.com/deissh/api.micro/service-anime/handlers"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/deissh/api.micro/service-anime/docs"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Service Anime API
// @version 1.0
// @description Anims methods

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

		g.GET("/_/health", handlers.HealthCheck)
		g.GET("/_/ping", handlers.PingCheck)
		g.GET("/_/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
