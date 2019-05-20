package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type health struct {
	Alive bool `json:"alive"`
}

// HealthResponse response struct
type HealthResponse struct {
	// API version
	Version string `json:"v"`
	Health  health `json:"Health"`
}

// HealthCheck godoc
// @Summary Show health service
// @ID get-service-health
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.HealthResponse
// @Router /_/health [get]
func (h Handler) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Version: "1",
		Health: health{
			Alive: true,
		},
	})
}
