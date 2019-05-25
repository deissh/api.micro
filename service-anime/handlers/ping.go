package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ping struct {
	ID          string    `json:"id,omitempty"`
	ServiceName string    `json:"service,omitempty"`
	Time        time.Time `json:"time,omitempty"`
}

// PingResponse response value
type PingResponse struct {
	// API version
	Version string `json:"v"`
	Ping    ping   `json:"ping"`
}

// PingCheck godoc
// @Summary Ping service
// @ID ping-service
// @Accept  json
// @Produce  json
// @Success 200 {object} handlers.PingResponse
// @Router /_/ping [get]
func (h Handler) PingCheck(c *gin.Context) {
	ping := ping{
		ID:          uuid.New().String(),
		ServiceName: "service-auth",
		Time:        time.Now().Local(),
	}

	c.JSON(http.StatusOK, PingResponse{
		Version: "1",
		Ping:    ping,
	})
}
