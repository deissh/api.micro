package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type health struct {
	Alive bool `json:"alive"`
}

// HealthResponse response value
type HealthResponse struct {
	// API version
	Version string `json:"v"`
	Health  health `json:"Health"`
}

// HealthCheck godoc
func (h Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Version: "1",
		Health: health{
			Alive: true,
		},
	})
}
